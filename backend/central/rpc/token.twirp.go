// Code generated by protoc-gen-twirp v5.12.0, DO NOT EDIT.
// source: token.proto

package rpc

import bytes "bytes"
import strings "strings"
import context "context"
import fmt "fmt"
import ioutil "io/ioutil"
import http "net/http"
import strconv "strconv"

import jsonpb "github.com/golang/protobuf/jsonpb"
import proto "github.com/golang/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

// ======================
// TokenService Interface
// ======================

type TokenService interface {
	Auth(context.Context, *Empty) (*TokenAuthResponse, error)

	PostMessage(context.Context, *Empty) (*TokenPostMessageResponse, error)

	GetMessages(context.Context, *Empty) (*TokenGetMessagesResponse, error)
}

// ============================
// TokenService Protobuf Client
// ============================

type tokenServiceProtobufClient struct {
	client HTTPClient
	urls   [3]string
	opts   twirp.ClientOptions
}

// NewTokenServiceProtobufClient creates a Protobuf client that implements the TokenService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewTokenServiceProtobufClient(addr string, client HTTPClient, opts ...twirp.ClientOption) TokenService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	prefix := urlBase(addr) + TokenServicePathPrefix
	urls := [3]string{
		prefix + "Auth",
		prefix + "PostMessage",
		prefix + "GetMessages",
	}

	return &tokenServiceProtobufClient{
		client: client,
		urls:   urls,
		opts:   clientOpts,
	}
}

func (c *tokenServiceProtobufClient) Auth(ctx context.Context, in *Empty) (*TokenAuthResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "TokenService")
	ctx = ctxsetters.WithMethodName(ctx, "Auth")
	out := new(TokenAuthResponse)
	err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *tokenServiceProtobufClient) PostMessage(ctx context.Context, in *Empty) (*TokenPostMessageResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "TokenService")
	ctx = ctxsetters.WithMethodName(ctx, "PostMessage")
	out := new(TokenPostMessageResponse)
	err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *tokenServiceProtobufClient) GetMessages(ctx context.Context, in *Empty) (*TokenGetMessagesResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "TokenService")
	ctx = ctxsetters.WithMethodName(ctx, "GetMessages")
	out := new(TokenGetMessagesResponse)
	err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[2], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ========================
// TokenService JSON Client
// ========================

type tokenServiceJSONClient struct {
	client HTTPClient
	urls   [3]string
	opts   twirp.ClientOptions
}

// NewTokenServiceJSONClient creates a JSON client that implements the TokenService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewTokenServiceJSONClient(addr string, client HTTPClient, opts ...twirp.ClientOption) TokenService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	prefix := urlBase(addr) + TokenServicePathPrefix
	urls := [3]string{
		prefix + "Auth",
		prefix + "PostMessage",
		prefix + "GetMessages",
	}

	return &tokenServiceJSONClient{
		client: client,
		urls:   urls,
		opts:   clientOpts,
	}
}

func (c *tokenServiceJSONClient) Auth(ctx context.Context, in *Empty) (*TokenAuthResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "TokenService")
	ctx = ctxsetters.WithMethodName(ctx, "Auth")
	out := new(TokenAuthResponse)
	err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *tokenServiceJSONClient) PostMessage(ctx context.Context, in *Empty) (*TokenPostMessageResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "TokenService")
	ctx = ctxsetters.WithMethodName(ctx, "PostMessage")
	out := new(TokenPostMessageResponse)
	err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *tokenServiceJSONClient) GetMessages(ctx context.Context, in *Empty) (*TokenGetMessagesResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "TokenService")
	ctx = ctxsetters.WithMethodName(ctx, "GetMessages")
	out := new(TokenGetMessagesResponse)
	err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[2], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ===========================
// TokenService Server Handler
// ===========================

type tokenServiceServer struct {
	TokenService
	hooks *twirp.ServerHooks
}

func NewTokenServiceServer(svc TokenService, hooks *twirp.ServerHooks) TwirpServer {
	return &tokenServiceServer{
		TokenService: svc,
		hooks:        hooks,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *tokenServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// TokenServicePathPrefix is used for all URL paths on a twirp TokenService server.
// Requests are always: POST TokenServicePathPrefix/method
// It can be used in an HTTP mux to route twirp requests along with non-twirp requests on other routes.
const TokenServicePathPrefix = "/rpc.TokenService/"

func (s *tokenServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "TokenService")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}

	switch req.URL.Path {
	case "/rpc.TokenService/Auth":
		s.serveAuth(ctx, resp, req)
		return
	case "/rpc.TokenService/PostMessage":
		s.servePostMessage(ctx, resp, req)
		return
	case "/rpc.TokenService/GetMessages":
		s.serveGetMessages(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}
}

func (s *tokenServiceServer) serveAuth(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveAuthJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveAuthProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *tokenServiceServer) serveAuthJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Auth")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(Empty)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *TokenAuthResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.TokenService.Auth(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *TokenAuthResponse and nil error while calling Auth. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *tokenServiceServer) serveAuthProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Auth")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(Empty)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *TokenAuthResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.TokenService.Auth(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *TokenAuthResponse and nil error while calling Auth. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *tokenServiceServer) servePostMessage(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.servePostMessageJSON(ctx, resp, req)
	case "application/protobuf":
		s.servePostMessageProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *tokenServiceServer) servePostMessageJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "PostMessage")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(Empty)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *TokenPostMessageResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.TokenService.PostMessage(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *TokenPostMessageResponse and nil error while calling PostMessage. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *tokenServiceServer) servePostMessageProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "PostMessage")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(Empty)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *TokenPostMessageResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.TokenService.PostMessage(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *TokenPostMessageResponse and nil error while calling PostMessage. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *tokenServiceServer) serveGetMessages(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveGetMessagesJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveGetMessagesProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *tokenServiceServer) serveGetMessagesJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "GetMessages")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(Empty)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *TokenGetMessagesResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.TokenService.GetMessages(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *TokenGetMessagesResponse and nil error while calling GetMessages. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	var buf bytes.Buffer
	marshaler := &jsonpb.Marshaler{OrigName: true}
	if err = marshaler.Marshal(&buf, respContent); err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	respBytes := buf.Bytes()
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *tokenServiceServer) serveGetMessagesProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "GetMessages")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to read request body"))
		return
	}
	reqContent := new(Empty)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *TokenGetMessagesResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.TokenService.GetMessages(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *TokenGetMessagesResponse and nil error while calling GetMessages. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *tokenServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor4, 0
}

func (s *tokenServiceServer) ProtocGenTwirpVersion() string {
	return "v5.12.0"
}

func (s *tokenServiceServer) PathPrefix() string {
	return TokenServicePathPrefix
}

var twirpFileDescriptor4 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xc9, 0xcf, 0x4e,
	0xcd, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x96, 0xe2, 0xce, 0xcd,
	0x4f, 0x49, 0xcd, 0x81, 0x88, 0x28, 0x69, 0x72, 0x09, 0x86, 0x80, 0x14, 0x38, 0x96, 0x96, 0x64,
	0x04, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x89, 0x70, 0xb1, 0x82, 0x75, 0x49, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x4a, 0x33, 0x18, 0xb9, 0x24, 0xc0, 0x6a, 0x03, 0xf2,
	0x8b, 0x4b, 0x7c, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x53, 0xe1, 0x5a, 0x1c, 0xb9, 0xd8, 0xc0, 0xaa,
	0x8a, 0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x34, 0xf5, 0x8a, 0x0a, 0x92, 0xf5, 0x70, 0x29,
	0x87, 0x48, 0x14, 0xbb, 0xe6, 0x95, 0x14, 0x55, 0x06, 0x41, 0x35, 0x4a, 0x59, 0x72, 0x71, 0x23,
	0x09, 0x0b, 0x09, 0x70, 0x31, 0x67, 0xa7, 0x56, 0x42, 0x9d, 0x00, 0x62, 0x82, 0x9c, 0x55, 0x96,
	0x98, 0x53, 0x9a, 0x2a, 0xc1, 0x04, 0x71, 0x16, 0x98, 0x63, 0xc5, 0x64, 0xc1, 0x88, 0x70, 0x9a,
	0x7b, 0x2a, 0xcc, 0xaa, 0x62, 0x62, 0x9c, 0x86, 0x45, 0x39, 0x95, 0x9d, 0x66, 0xb4, 0x8e, 0x91,
	0x8b, 0x07, 0xac, 0x37, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x48, 0x8b, 0x8b, 0x05, 0x14,
	0xd8, 0x42, 0x5c, 0x60, 0x67, 0xb8, 0xe6, 0x16, 0x94, 0x54, 0x4a, 0x89, 0x21, 0x9c, 0x84, 0x12,
	0x11, 0x16, 0x5c, 0xdc, 0x48, 0xa1, 0x87, 0xa2, 0x45, 0x16, 0x6f, 0x00, 0x83, 0x74, 0x22, 0x79,
	0x0e, 0x97, 0x4e, 0x2c, 0xfe, 0x77, 0xe2, 0x88, 0x62, 0xd3, 0xd3, 0xd3, 0x2f, 0x2a, 0x48, 0x4e,
	0x62, 0x03, 0x27, 0x11, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xac, 0xc8, 0x5a, 0xee, 0x43,
	0x02, 0x00, 0x00,
}
