// Code generated by protoc-gen-twirp v5.12.0, DO NOT EDIT.
// source: invite.proto

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

// =======================
// InviteService Interface
// =======================

type InviteService interface {
	Create(context.Context, *InviteCreateRequest) (*Empty, error)

	Accept(context.Context, *InviteAcceptRequest) (*Empty, error)

	FromMe(context.Context, *Empty) (*InviteFromMeResponse, error)

	ForMe(context.Context, *Empty) (*InviteForMeResponse, error)
}

// =============================
// InviteService Protobuf Client
// =============================

type inviteServiceProtobufClient struct {
	client HTTPClient
	urls   [4]string
	opts   twirp.ClientOptions
}

// NewInviteServiceProtobufClient creates a Protobuf client that implements the InviteService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewInviteServiceProtobufClient(addr string, client HTTPClient, opts ...twirp.ClientOption) InviteService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	prefix := urlBase(addr) + InviteServicePathPrefix
	urls := [4]string{
		prefix + "Create",
		prefix + "Accept",
		prefix + "FromMe",
		prefix + "ForMe",
	}

	return &inviteServiceProtobufClient{
		client: client,
		urls:   urls,
		opts:   clientOpts,
	}
}

func (c *inviteServiceProtobufClient) Create(ctx context.Context, in *InviteCreateRequest) (*Empty, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "InviteService")
	ctx = ctxsetters.WithMethodName(ctx, "Create")
	out := new(Empty)
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

func (c *inviteServiceProtobufClient) Accept(ctx context.Context, in *InviteAcceptRequest) (*Empty, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "InviteService")
	ctx = ctxsetters.WithMethodName(ctx, "Accept")
	out := new(Empty)
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

func (c *inviteServiceProtobufClient) FromMe(ctx context.Context, in *Empty) (*InviteFromMeResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "InviteService")
	ctx = ctxsetters.WithMethodName(ctx, "FromMe")
	out := new(InviteFromMeResponse)
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

func (c *inviteServiceProtobufClient) ForMe(ctx context.Context, in *Empty) (*InviteForMeResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "InviteService")
	ctx = ctxsetters.WithMethodName(ctx, "ForMe")
	out := new(InviteForMeResponse)
	err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[3], in, out)
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

// =========================
// InviteService JSON Client
// =========================

type inviteServiceJSONClient struct {
	client HTTPClient
	urls   [4]string
	opts   twirp.ClientOptions
}

// NewInviteServiceJSONClient creates a JSON client that implements the InviteService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewInviteServiceJSONClient(addr string, client HTTPClient, opts ...twirp.ClientOption) InviteService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	prefix := urlBase(addr) + InviteServicePathPrefix
	urls := [4]string{
		prefix + "Create",
		prefix + "Accept",
		prefix + "FromMe",
		prefix + "ForMe",
	}

	return &inviteServiceJSONClient{
		client: client,
		urls:   urls,
		opts:   clientOpts,
	}
}

func (c *inviteServiceJSONClient) Create(ctx context.Context, in *InviteCreateRequest) (*Empty, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "InviteService")
	ctx = ctxsetters.WithMethodName(ctx, "Create")
	out := new(Empty)
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

func (c *inviteServiceJSONClient) Accept(ctx context.Context, in *InviteAcceptRequest) (*Empty, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "InviteService")
	ctx = ctxsetters.WithMethodName(ctx, "Accept")
	out := new(Empty)
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

func (c *inviteServiceJSONClient) FromMe(ctx context.Context, in *Empty) (*InviteFromMeResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "InviteService")
	ctx = ctxsetters.WithMethodName(ctx, "FromMe")
	out := new(InviteFromMeResponse)
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

func (c *inviteServiceJSONClient) ForMe(ctx context.Context, in *Empty) (*InviteForMeResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "InviteService")
	ctx = ctxsetters.WithMethodName(ctx, "ForMe")
	out := new(InviteForMeResponse)
	err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[3], in, out)
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

// ============================
// InviteService Server Handler
// ============================

type inviteServiceServer struct {
	InviteService
	hooks *twirp.ServerHooks
}

func NewInviteServiceServer(svc InviteService, hooks *twirp.ServerHooks) TwirpServer {
	return &inviteServiceServer{
		InviteService: svc,
		hooks:         hooks,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *inviteServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// InviteServicePathPrefix is used for all URL paths on a twirp InviteService server.
// Requests are always: POST InviteServicePathPrefix/method
// It can be used in an HTTP mux to route twirp requests along with non-twirp requests on other routes.
const InviteServicePathPrefix = "/rpc.InviteService/"

func (s *inviteServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "InviteService")
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
	case "/rpc.InviteService/Create":
		s.serveCreate(ctx, resp, req)
		return
	case "/rpc.InviteService/Accept":
		s.serveAccept(ctx, resp, req)
		return
	case "/rpc.InviteService/FromMe":
		s.serveFromMe(ctx, resp, req)
		return
	case "/rpc.InviteService/ForMe":
		s.serveForMe(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		err = badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, err)
		return
	}
}

func (s *inviteServiceServer) serveCreate(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveCreateJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveCreateProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *inviteServiceServer) serveCreateJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Create")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(InviteCreateRequest)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *Empty
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.InviteService.Create(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Empty and nil error while calling Create. nil responses are not supported"))
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

func (s *inviteServiceServer) serveCreateProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Create")
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
	reqContent := new(InviteCreateRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *Empty
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.InviteService.Create(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Empty and nil error while calling Create. nil responses are not supported"))
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

func (s *inviteServiceServer) serveAccept(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveAcceptJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveAcceptProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *inviteServiceServer) serveAcceptJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Accept")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	reqContent := new(InviteAcceptRequest)
	unmarshaler := jsonpb.Unmarshaler{AllowUnknownFields: true}
	if err = unmarshaler.Unmarshal(req.Body, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the json request could not be decoded"))
		return
	}

	// Call service method
	var respContent *Empty
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.InviteService.Accept(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Empty and nil error while calling Accept. nil responses are not supported"))
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

func (s *inviteServiceServer) serveAcceptProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Accept")
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
	reqContent := new(InviteAcceptRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	// Call service method
	var respContent *Empty
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.InviteService.Accept(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *Empty and nil error while calling Accept. nil responses are not supported"))
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

func (s *inviteServiceServer) serveFromMe(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveFromMeJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveFromMeProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *inviteServiceServer) serveFromMeJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "FromMe")
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
	var respContent *InviteFromMeResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.InviteService.FromMe(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *InviteFromMeResponse and nil error while calling FromMe. nil responses are not supported"))
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

func (s *inviteServiceServer) serveFromMeProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "FromMe")
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
	var respContent *InviteFromMeResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.InviteService.FromMe(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *InviteFromMeResponse and nil error while calling FromMe. nil responses are not supported"))
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

func (s *inviteServiceServer) serveForMe(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveForMeJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveForMeProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *inviteServiceServer) serveForMeJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ForMe")
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
	var respContent *InviteForMeResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.InviteService.ForMe(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *InviteForMeResponse and nil error while calling ForMe. nil responses are not supported"))
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

func (s *inviteServiceServer) serveForMeProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ForMe")
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
	var respContent *InviteForMeResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = s.InviteService.ForMe(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *InviteForMeResponse and nil error while calling ForMe. nil responses are not supported"))
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

func (s *inviteServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor2, 0
}

func (s *inviteServiceServer) ProtocGenTwirpVersion() string {
	return "v5.12.0"
}

func (s *inviteServiceServer) PathPrefix() string {
	return InviteServicePathPrefix
}

var twirpFileDescriptor2 = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x4e, 0xbb, 0x40,
	0x10, 0xc6, 0xc3, 0xbf, 0x7f, 0xb1, 0x9d, 0xea, 0x65, 0x35, 0x8a, 0x18, 0x63, 0xc3, 0xa9, 0x97,
	0x6e, 0x63, 0xf5, 0x05, 0xd0, 0xd8, 0xc8, 0x41, 0x0f, 0x78, 0xf3, 0xd2, 0xe0, 0x32, 0x26, 0x24,
	0xc2, 0xae, 0xcb, 0xda, 0xc4, 0xd7, 0xf0, 0xc5, 0x7c, 0x25, 0xc3, 0x0e, 0x1b, 0xd9, 0x78, 0xf4,
	0x06, 0xdf, 0xfc, 0xbe, 0x99, 0x8f, 0x19, 0x60, 0xaf, 0x6a, 0xb6, 0x95, 0x41, 0xae, 0xb4, 0x34,
	0x92, 0x8d, 0xb4, 0x12, 0xf1, 0xb4, 0x96, 0x25, 0xbe, 0x92, 0x92, 0x2c, 0xe0, 0x20, 0xb3, 0xc4,
	0x8d, 0xc6, 0xc2, 0x60, 0x8e, 0x6f, 0xef, 0xd8, 0x1a, 0x76, 0x04, 0xe1, 0x8b, 0xae, 0xb0, 0x29,
	0xa3, 0x60, 0x16, 0xcc, 0x27, 0x79, 0xff, 0x96, 0x5c, 0x39, 0x3c, 0x15, 0x02, 0x95, 0x71, 0xf8,
	0x19, 0x00, 0xcd, 0xd1, 0x9b, 0xca, 0x59, 0x26, 0xbd, 0x92, 0x95, 0xc9, 0x67, 0x00, 0x8c, 0x6c,
	0x6b, 0xdb, 0x86, 0x9e, 0xd9, 0x29, 0x4c, 0xa8, 0xed, 0x8f, 0x69, 0x4c, 0x42, 0x56, 0xb2, 0x73,
	0x98, 0xf6, 0xc5, 0xa6, 0xa8, 0x31, 0xfa, 0x67, 0xcb, 0x40, 0xd2, 0x43, 0x51, 0x63, 0x37, 0x53,
	0xd8, 0xcc, 0xe5, 0xa6, 0x30, 0xd1, 0x88, 0x66, 0xf6, 0x4a, 0x6a, 0x3a, 0x7f, 0x61, 0x33, 0x52,
	0xfd, 0x3f, 0xf9, 0x9d, 0x94, 0x9a, 0x24, 0x83, 0x43, 0x97, 0x49, 0xd6, 0xf7, 0x98, 0x63, 0xab,
	0x64, 0xd3, 0x22, 0xbb, 0x80, 0x5d, 0x4a, 0xde, 0x46, 0xc1, 0x6c, 0x34, 0x9f, 0xae, 0x8e, 0xb9,
	0x56, 0x82, 0xff, 0xce, 0x9f, 0x3b, 0x2e, 0xb9, 0x73, 0x5b, 0x59, 0x4b, 0xfd, 0xa7, 0x4e, 0xab,
	0xaf, 0x00, 0xf6, 0x49, 0x7b, 0x44, 0xbd, 0xad, 0x04, 0x32, 0x0e, 0x21, 0x9d, 0x86, 0x45, 0x03,
	0xb7, 0x77, 0xad, 0x18, 0x6c, 0xe5, 0xb6, 0x56, 0xe6, 0xa3, 0xe3, 0xe9, 0x36, 0x1e, 0xef, 0x9d,
	0xcb, 0xe3, 0x97, 0x10, 0xd2, 0x02, 0xd8, 0x40, 0x8d, 0x4f, 0xbc, 0xa4, 0xde, 0x7e, 0x16, 0xb0,
	0x63, 0x3f, 0xd3, 0xe3, 0x87, 0xb3, 0xbc, 0x25, 0x5c, 0x8f, 0x9f, 0x42, 0xce, 0x97, 0x5a, 0x89,
	0xe7, 0xd0, 0xfe, 0x71, 0x97, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xd7, 0x87, 0x32, 0x93,
	0x02, 0x00, 0x00,
}