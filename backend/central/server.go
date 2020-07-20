package central

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/gorilla/sessions"
	"github.com/twitchtv/twirp"

	"github.com/mreider/koto/backend/central/bcrypt"
	"github.com/mreider/koto/backend/central/config"
	"github.com/mreider/koto/backend/central/repo"
	"github.com/mreider/koto/backend/central/rpc"
	"github.com/mreider/koto/backend/central/services"
	"github.com/mreider/koto/backend/common"
	"github.com/mreider/koto/backend/token"
)

const (
	cookieAuthenticationKey = "oSKDA9fDNa6jIHArw8MHGBPe0XZm4hnY"
	sessionName             = "auth-session"
	sessionUserKey          = "user-id"
)

type Server struct {
	cfg            config.Config
	pubKeyPEM      string
	repos          repo.Repos
	tokenGenerator token.Generator
	s3Storage      *common.S3Storage
	sessionStore   *sessions.CookieStore
}

func NewServer(cfg config.Config, pubKeyPEM string, repos repo.Repos, tokenGenerator token.Generator, s3Storage *common.S3Storage) *Server {
	sessionStore := sessions.NewCookieStore([]byte(cookieAuthenticationKey))
	sessionStore.Options.HttpOnly = true
	sessionStore.Options.MaxAge = 0

	return &Server{
		cfg:            cfg,
		pubKeyPEM:      pubKeyPEM,
		repos:          repos,
		tokenGenerator: tokenGenerator,
		s3Storage:      s3Storage,
		sessionStore:   sessionStore,
	}
}

func (s *Server) Run() error {
	r := chi.NewRouter()
	s.setupMiddlewares(r)

	rpcHooks := &twirp.ServerHooks{}
	baseService := services.NewBase(s.repos, s.s3Storage)

	passwordHash := bcrypt.NewPasswordHash()

	authService := services.NewAuth(baseService, sessionUserKey, passwordHash)
	authServiceHandler := rpc.NewAuthServiceServer(authService, rpcHooks)
	r.Handle(authServiceHandler.PathPrefix()+"*", s.authSessionProvider(authServiceHandler))

	infoService := services.NewInfo(baseService, s.pubKeyPEM)
	infoServiceHandler := rpc.NewInfoServiceServer(infoService, rpcHooks)
	r.Handle(infoServiceHandler.PathPrefix()+"*", infoServiceHandler)

	tokenService := services.NewToken(baseService, s.tokenGenerator, s.cfg.TokenDuration())
	tokenServiceHandler := rpc.NewTokenServiceServer(tokenService, rpcHooks)
	r.Handle(tokenServiceHandler.PathPrefix()+"*", s.checkAuth(tokenServiceHandler))

	userService := services.NewUser(baseService)
	userServiceHandler := rpc.NewUserServiceServer(userService, rpcHooks)
	r.Handle(userServiceHandler.PathPrefix()+"*", s.checkAuth(userServiceHandler))

	nodeService := services.NewNode(baseService)
	nodeServiceHandler := rpc.NewNodeServiceServer(nodeService, rpcHooks)
	r.Handle(nodeServiceHandler.PathPrefix()+"*", s.checkAuth(nodeServiceHandler))

	inviteService := services.NewInvite(baseService)
	inviteServiceHandler := rpc.NewInviteServiceServer(inviteService, rpcHooks)
	r.Handle(inviteServiceHandler.PathPrefix()+"*", s.checkAuth(inviteServiceHandler))

	blobService := services.NewBlob(baseService)
	blobServiceHandler := rpc.NewBlobServiceServer(blobService, rpcHooks)
	r.Handle(blobServiceHandler.PathPrefix()+"*", s.checkAuth(blobServiceHandler))

	log.Println("started on " + s.cfg.ListenAddress)
	return http.ListenAndServe(s.cfg.ListenAddress, r)
}

func (s *Server) setupMiddlewares(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	corsOptions := cors.Options{
		AllowOriginFunc: func(r *http.Request, origin string) bool {
			return true
		},
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
	}
	r.Use(cors.New(corsOptions).Handler)
}

func (s *Server) checkAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := s.sessionStore.Get(r, sessionName)
		userID, ok := session.Values[sessionUserKey].(string)
		if !ok {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}
		user, err := s.repos.User.FindUserByID(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if user == nil {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), services.ContextUserKey, *user)
		ctx = context.WithValue(ctx, services.ContextIsAdminKey, s.cfg.IsAdmin(user.Name) || s.cfg.IsAdmin(user.Email))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *Server) authSessionProvider(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := s.sessionStore.Get(r, sessionName)
		var sessionWrapper services.Session = &sessionWrapper{
			session: session,
			w:       w,
			r:       r,
		}
		ctx := context.WithValue(r.Context(), services.ContextSession, sessionWrapper)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type sessionWrapper struct {
	session *sessions.Session
	w       http.ResponseWriter
	r       *http.Request
}

func (s *sessionWrapper) SetValue(key, value interface{}) {
	s.session.Values[key] = value
}

func (s *sessionWrapper) Clear() {
	s.session.Values = nil
}

func (s *sessionWrapper) Save() error {
	return s.session.Save(s.r, s.w)
}
