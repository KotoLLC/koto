// Code generated by protoc-gen-twirp v5.12.0, DO NOT EDIT.
// source: user.proto

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

// =====================
// UserService Interface
// =====================

type UserService interface {
	Friends(context.Context, *Empty) (*UserFriendsResponse, error)

	FriendsOfFriends(context.Context, *Empty) (*UserFriendsOfFriendsResponse, error)

	Me(context.Context, *Empty) (*UserMeResponse, error)
}

// ===========================
// UserService Protobuf Client
// ===========================

type userServiceProtobufClient struct {
	client HTTPClient
	urls   [3]string
	opts   twirp.ClientOptions
}

// NewUserServiceProtobufClient creates a Protobuf client that implements the UserService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewUserServiceProtobufClient(addr string, client HTTPClient, opts ...twirp.ClientOption) UserService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	prefix := urlBase(addr) + UserServicePathPrefix
	urls := [3]string{
		prefix + "Friends",
		prefix + "FriendsOfFriends",
		prefix + "Me",
	}

	return &userServiceProtobufClient{
		client: client,
		urls:   urls,
		opts:   clientOpts,
	}
}

func (c *userServiceProtobufClient) Friends(ctx context.Context, in *Empty) (*UserFriendsResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "Friends")
	out := new(UserFriendsResponse)
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

func (c *userServiceProtobufClient) FriendsOfFriends(ctx context.Context, in *Empty) (*UserFriendsOfFriendsResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "FriendsOfFriends")
	out := new(UserFriendsOfFriendsResponse)
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

func (c *userServiceProtobufClient) Me(ctx context.Context, in *Empty) (*UserMeResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "Me")
	out := new(UserMeResponse)
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

// =======================
// UserService JSON Client
// =======================

type userServiceJSONClient struct {
	client HTTPClient
	urls   [3]string
	opts   twirp.ClientOptions
}

// NewUserServiceJSONClient creates a JSON client that implements the UserService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewUserServiceJSONClient(addr string, client HTTPClient, opts ...twirp.ClientOption) UserService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	prefix := urlBase(addr) + UserServicePathPrefix
	urls := [3]string{
		prefix + "Friends",
		prefix + "FriendsOfFriends",
		prefix + "Me",
	}

	return &userServiceJSONClient{
		client: client,
		urls:   urls,
		opts:   clientOpts,
	}
}

func (c *userServiceJSONClient) Friends(ctx context.Context, in *Empty) (*UserFriendsResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "Friends")
	out := new(UserFriendsResponse)
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

func (c *userServiceJSONClient) FriendsOfFriends(ctx context.Context, in *Empty) (*UserFriendsOfFriendsResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "FriendsOfFriends")
	out := new(UserFriendsOfFriendsResponse)
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

func (c *userServiceJSONClient) Me(ctx context.Context, in *Empty) (*UserMeResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
	ctx = ctxsetters.WithMethodName(ctx, "Me")
	out := new(UserMeResponse)
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

// ==========================
// UserService Server Handler
// ==========================

type userServiceServer struct {
	UserService
	hooks *twirp.ServerHooks
}

func NewUserServiceServer(svc UserService, hooks *twirp.ServerHooks) TwirpServer {
	return &userServiceServer{
		UserService: svc,
		hooks:       hooks,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *userServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// UserServicePathPrefix is used for all URL paths on a twirp UserService server.
// Requests are always: POST UserServicePathPrefix/method
// It can be used in an HTTP mux to route twirp requests along with non-twirp requests on other routes.
const UserServicePathPrefix = "/rpc.UserService/"

func (s *userServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "UserService")
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
	case "/rpc.UserService/Friends":
		s.serveFriends(ctx, resp, req)
		return
	case "/rpc.UserService/FriendsOfFriends":
		s.serveFriendsOfFriends(ctx, resp, req)
		return
	case "/rpc.UserService/Me":
		s.serveMe(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}
}

func (s *userServiceServer) serveFriends(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveFriendsJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveFriendsProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *userServiceServer) serveFriendsJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Friends")
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
	var respContent *UserFriendsResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.UserService.Friends(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *UserFriendsResponse and nil error while calling Friends. nil responses are not supported"))
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

func (s *userServiceServer) serveFriendsProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Friends")
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
	var respContent *UserFriendsResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.UserService.Friends(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *UserFriendsResponse and nil error while calling Friends. nil responses are not supported"))
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

func (s *userServiceServer) serveFriendsOfFriends(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveFriendsOfFriendsJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveFriendsOfFriendsProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *userServiceServer) serveFriendsOfFriendsJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "FriendsOfFriends")
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
	var respContent *UserFriendsOfFriendsResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.UserService.FriendsOfFriends(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *UserFriendsOfFriendsResponse and nil error while calling FriendsOfFriends. nil responses are not supported"))
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

func (s *userServiceServer) serveFriendsOfFriendsProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "FriendsOfFriends")
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
	var respContent *UserFriendsOfFriendsResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.UserService.FriendsOfFriends(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *UserFriendsOfFriendsResponse and nil error while calling FriendsOfFriends. nil responses are not supported"))
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

func (s *userServiceServer) serveMe(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveMeJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveMeProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *userServiceServer) serveMeJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Me")
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
	var respContent *UserMeResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.UserService.Me(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *UserMeResponse and nil error while calling Me. nil responses are not supported"))
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

func (s *userServiceServer) serveMeProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Me")
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
	var respContent *UserMeResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.UserService.Me(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *UserMeResponse and nil error while calling Me. nil responses are not supported"))
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

func (s *userServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor5, 0
}

func (s *userServiceServer) ProtocGenTwirpVersion() string {
	return "v5.12.0"
}

func (s *userServiceServer) PathPrefix() string {
	return UserServicePathPrefix
}

var twirpFileDescriptor5 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xd9, 0x46, 0x9a, 0x74, 0xa2, 0x22, 0xdb, 0x4b, 0x2c, 0x0a, 0x31, 0x45, 0xec, 0x69,
	0x0b, 0xf5, 0xe6, 0xad, 0x82, 0x1e, 0x84, 0x22, 0xa4, 0x78, 0xf1, 0x52, 0x62, 0x32, 0x85, 0x05,
	0x93, 0x2c, 0xbb, 0xdb, 0x82, 0x3f, 0xc2, 0xff, 0xe1, 0xcf, 0x94, 0xdd, 0x68, 0x9a, 0x04, 0xda,
	0x1e, 0xe7, 0xed, 0xbc, 0x37, 0xef, 0x83, 0x05, 0xd8, 0x28, 0x94, 0x4c, 0xc8, 0x52, 0x97, 0xd4,
	0x91, 0x22, 0x1d, 0xf9, 0x79, 0x99, 0xe1, 0x67, 0xa5, 0x44, 0x0f, 0x30, 0x7c, 0x53, 0x28, 0x9f,
	0x25, 0xc7, 0x22, 0x53, 0x31, 0x2a, 0x51, 0x16, 0x0a, 0xe9, 0x18, 0xdc, 0x75, 0x25, 0x05, 0x24,
	0x74, 0x26, 0xfe, 0x6c, 0xc0, 0xa4, 0x48, 0x99, 0x59, 0x8d, 0xff, 0x5f, 0xa2, 0x6f, 0x02, 0x51,
	0xc3, 0xfc, 0xba, 0xee, 0xa4, 0x54, 0x23, 0xbd, 0x86, 0x13, 0x53, 0x21, 0x20, 0x21, 0x69, 0x07,
	0x59, 0x99, 0x8e, 0xe1, 0x8c, 0x17, 0x5b, 0xae, 0x71, 0xa5, 0x74, 0xa2, 0x37, 0x2a, 0xe8, 0x85,
	0x64, 0x32, 0x88, 0x4f, 0x2b, 0x71, 0x69, 0xb5, 0x66, 0x1f, 0x67, 0x6f, 0x9f, 0x04, 0xae, 0x0e,
	0xd5, 0xa1, 0xf3, 0x2e, 0xd4, 0x5d, 0x1d, 0x72, 0x18, 0x61, 0x77, 0xe2, 0x05, 0xce, 0xcd, 0xfa,
	0x02, 0xeb, 0xd0, 0x23, 0x74, 0x97, 0xe0, 0x71, 0xb5, 0x4a, 0xb2, 0x9c, 0x17, 0x16, 0xcc, 0x8b,
	0x5d, 0xae, 0xe6, 0x66, 0x9c, 0xfd, 0x10, 0xf0, 0xcd, 0xe6, 0x12, 0xe5, 0x96, 0xa7, 0x48, 0xa7,
	0xe0, 0xfe, 0x5d, 0xa7, 0x60, 0x63, 0x9e, 0x72, 0xa1, 0xbf, 0x46, 0x41, 0xb7, 0x64, 0x83, 0xe7,
	0xa2, 0xdb, 0xbb, 0xe5, 0xbc, 0x39, 0x8a, 0x47, 0x6f, 0xa1, 0xb7, 0xc0, 0x96, 0x69, 0x58, 0x9b,
	0x76, 0x90, 0x8f, 0xde, 0x7b, 0x9f, 0xb1, 0xa9, 0x14, 0xe9, 0x47, 0xdf, 0x7e, 0x9b, 0xfb, 0xdf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0x1d, 0x41, 0x0d, 0x56, 0x02, 0x00, 0x00,
}
