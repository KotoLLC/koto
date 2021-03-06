// Code generated by protoc-gen-twirp v8.0.0, DO NOT EDIT.
// source: info.proto

package rpc

import context "context"
import fmt "fmt"
import http "net/http"
import ioutil "io/ioutil"
import json "encoding/json"
import strconv "strconv"
import strings "strings"

import protojson "google.golang.org/protobuf/encoding/protojson"
import proto "google.golang.org/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the twirp package used in your project.
// A compilation error at this line likely means your copy of the
// twirp package needs to be updated.
const _ = twirp.TwirpPackageIsVersion7

// =====================
// InfoService Interface
// =====================

type InfoService interface {
	PublicKey(context.Context, *Empty) (*InfoPublicKeyResponse, error)

	Version(context.Context, *Empty) (*InfoVersionResponse, error)
}

// ===========================
// InfoService Protobuf Client
// ===========================

type infoServiceProtobufClient struct {
	client      HTTPClient
	urls        [2]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewInfoServiceProtobufClient creates a Protobuf client that implements the InfoService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewInfoServiceProtobufClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) InfoService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(clientOpts.PathPrefix(), "rpc", "InfoService")
	urls := [2]string{
		serviceURL + "PublicKey",
		serviceURL + "Version",
	}

	return &infoServiceProtobufClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *infoServiceProtobufClient) PublicKey(ctx context.Context, in *Empty) (*InfoPublicKeyResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "InfoService")
	ctx = ctxsetters.WithMethodName(ctx, "PublicKey")
	caller := c.callPublicKey
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *Empty) (*InfoPublicKeyResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*Empty)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*Empty) when calling interceptor")
					}
					return c.callPublicKey(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*InfoPublicKeyResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*InfoPublicKeyResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *infoServiceProtobufClient) callPublicKey(ctx context.Context, in *Empty) (*InfoPublicKeyResponse, error) {
	out := new(InfoPublicKeyResponse)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
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

func (c *infoServiceProtobufClient) Version(ctx context.Context, in *Empty) (*InfoVersionResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "InfoService")
	ctx = ctxsetters.WithMethodName(ctx, "Version")
	caller := c.callVersion
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *Empty) (*InfoVersionResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*Empty)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*Empty) when calling interceptor")
					}
					return c.callVersion(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*InfoVersionResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*InfoVersionResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *infoServiceProtobufClient) callVersion(ctx context.Context, in *Empty) (*InfoVersionResponse, error) {
	out := new(InfoVersionResponse)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
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
// InfoService JSON Client
// =======================

type infoServiceJSONClient struct {
	client      HTTPClient
	urls        [2]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewInfoServiceJSONClient creates a JSON client that implements the InfoService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewInfoServiceJSONClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) InfoService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(clientOpts.PathPrefix(), "rpc", "InfoService")
	urls := [2]string{
		serviceURL + "PublicKey",
		serviceURL + "Version",
	}

	return &infoServiceJSONClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *infoServiceJSONClient) PublicKey(ctx context.Context, in *Empty) (*InfoPublicKeyResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "InfoService")
	ctx = ctxsetters.WithMethodName(ctx, "PublicKey")
	caller := c.callPublicKey
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *Empty) (*InfoPublicKeyResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*Empty)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*Empty) when calling interceptor")
					}
					return c.callPublicKey(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*InfoPublicKeyResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*InfoPublicKeyResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *infoServiceJSONClient) callPublicKey(ctx context.Context, in *Empty) (*InfoPublicKeyResponse, error) {
	out := new(InfoPublicKeyResponse)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
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

func (c *infoServiceJSONClient) Version(ctx context.Context, in *Empty) (*InfoVersionResponse, error) {
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "InfoService")
	ctx = ctxsetters.WithMethodName(ctx, "Version")
	caller := c.callVersion
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *Empty) (*InfoVersionResponse, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*Empty)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*Empty) when calling interceptor")
					}
					return c.callVersion(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*InfoVersionResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*InfoVersionResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *infoServiceJSONClient) callVersion(ctx context.Context, in *Empty) (*InfoVersionResponse, error) {
	out := new(InfoVersionResponse)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
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
// InfoService Server Handler
// ==========================

type infoServiceServer struct {
	InfoService
	interceptor      twirp.Interceptor
	hooks            *twirp.ServerHooks
	pathPrefix       string // prefix for routing
	jsonSkipDefaults bool   // do not include unpopulated fields (default values) in the response
}

// NewInfoServiceServer builds a TwirpServer that can be used as an http.Handler to handle
// HTTP requests that are routed to the right method in the provided svc implementation.
// The opts are twirp.ServerOption modifiers, for example twirp.WithServerHooks(hooks).
func NewInfoServiceServer(svc InfoService, opts ...interface{}) TwirpServer {
	serverOpts := twirp.ServerOptions{}
	for _, opt := range opts {
		switch o := opt.(type) {
		case twirp.ServerOption:
			o(&serverOpts)
		case *twirp.ServerHooks: // backwards compatibility, allow to specify hooks as an argument
			twirp.WithServerHooks(o)(&serverOpts)
		case nil: // backwards compatibility, allow nil value for the argument
			continue
		default:
			panic(fmt.Sprintf("Invalid option type %T on NewInfoServiceServer", o))
		}
	}

	return &infoServiceServer{
		InfoService:      svc,
		pathPrefix:       serverOpts.PathPrefix(),
		interceptor:      twirp.ChainInterceptors(serverOpts.Interceptors...),
		hooks:            serverOpts.Hooks,
		jsonSkipDefaults: serverOpts.JSONSkipDefaults,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *infoServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// handleRequestBodyError is used to handle error when the twirp server cannot read request
func (s *infoServiceServer) handleRequestBodyError(ctx context.Context, resp http.ResponseWriter, msg string, err error) {
	if context.Canceled == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.Canceled, "failed to read request: context canceled"))
		return
	}
	if context.DeadlineExceeded == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.DeadlineExceeded, "failed to read request: deadline exceeded"))
		return
	}
	s.writeError(ctx, resp, twirp.WrapError(malformedRequestError(msg), err))
}

// InfoServicePathPrefix is a convenience constant that could used to identify URL paths.
// Should be used with caution, it only matches routes generated by Twirp Go clients,
// that add a "/twirp" prefix by default, and use CamelCase service and method names.
// More info: https://twitchtv.github.io/twirp/docs/routing.html
const InfoServicePathPrefix = "/twirp/rpc.InfoService/"

func (s *infoServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "rpc")
	ctx = ctxsetters.WithServiceName(ctx, "InfoService")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	// Verify path format: [<prefix>]/<package>.<Service>/<Method>
	prefix, pkgService, method := parseTwirpPath(req.URL.Path)
	if pkgService != "rpc.InfoService" {
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
	if prefix != s.pathPrefix {
		msg := fmt.Sprintf("invalid path prefix %q, expected %q, on path %q", prefix, s.pathPrefix, req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	switch method {
	case "PublicKey":
		s.servePublicKey(ctx, resp, req)
		return
	case "Version":
		s.serveVersion(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
}

func (s *infoServiceServer) servePublicKey(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.servePublicKeyJSON(ctx, resp, req)
	case "application/protobuf":
		s.servePublicKeyProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *infoServiceServer) servePublicKeyJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "PublicKey")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(Empty)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.InfoService.PublicKey
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *Empty) (*InfoPublicKeyResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*Empty)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*Empty) when calling interceptor")
					}
					return s.InfoService.PublicKey(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*InfoPublicKeyResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*InfoPublicKeyResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *InfoPublicKeyResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *InfoPublicKeyResponse and nil error while calling PublicKey. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: true, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *infoServiceServer) servePublicKeyProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "PublicKey")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(Empty)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.InfoService.PublicKey
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *Empty) (*InfoPublicKeyResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*Empty)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*Empty) when calling interceptor")
					}
					return s.InfoService.PublicKey(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*InfoPublicKeyResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*InfoPublicKeyResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *InfoPublicKeyResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *InfoPublicKeyResponse and nil error while calling PublicKey. nil responses are not supported"))
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
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *infoServiceServer) serveVersion(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveVersionJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveVersionProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *infoServiceServer) serveVersionJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Version")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(Empty)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.InfoService.Version
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *Empty) (*InfoVersionResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*Empty)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*Empty) when calling interceptor")
					}
					return s.InfoService.Version(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*InfoVersionResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*InfoVersionResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *InfoVersionResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *InfoVersionResponse and nil error while calling Version. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: true, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *infoServiceServer) serveVersionProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "Version")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(Empty)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.InfoService.Version
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *Empty) (*InfoVersionResponse, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*Empty)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*Empty) when calling interceptor")
					}
					return s.InfoService.Version(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*InfoVersionResponse)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*InfoVersionResponse) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *InfoVersionResponse
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *InfoVersionResponse and nil error while calling Version. nil responses are not supported"))
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
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *infoServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor2, 0
}

func (s *infoServiceServer) ProtocGenTwirpVersion() string {
	return "v8.0.0"
}

// PathPrefix returns the base service path, in the form: "/<prefix>/<package>.<Service>/"
// that is everything in a Twirp route except for the <Method>. This can be used for routing,
// for example to identify the requests that are targeted to this service in a mux.
func (s *infoServiceServer) PathPrefix() string {
	return baseServicePath(s.pathPrefix, "rpc", "InfoService")
}

var twirpFileDescriptor2 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xca, 0xcc, 0x4b, 0xcb,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x96, 0xe2, 0xce, 0xcd, 0x4f,
	0x49, 0xcd, 0x81, 0x88, 0x28, 0x99, 0x71, 0x89, 0x7a, 0xe6, 0xa5, 0xe5, 0x07, 0x94, 0x26, 0xe5,
	0x64, 0x26, 0x7b, 0xa7, 0x56, 0x06, 0xa5, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0xc9, 0x72,
	0x71, 0x15, 0x80, 0x05, 0xe3, 0xb3, 0x53, 0x2b, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x38,
	0x0b, 0x60, 0xca, 0x94, 0x6c, 0xb8, 0x84, 0x41, 0xfa, 0xc2, 0x52, 0x8b, 0x8a, 0x33, 0xf3, 0xf3,
	0xe0, 0xba, 0x54, 0xb9, 0xf8, 0x52, 0xf2, 0x93, 0xb3, 0x53, 0x8b, 0xe2, 0x4b, 0x0b, 0x52, 0x12,
	0x4b, 0x52, 0x53, 0xa0, 0x3a, 0x79, 0x21, 0xa2, 0xa1, 0x10, 0x41, 0xa3, 0x62, 0x2e, 0x6e, 0x90,
	0xee, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x63, 0x2e, 0x4e, 0xb8, 0x03, 0x84, 0xb8,
	0xf4, 0x8a, 0x0a, 0x92, 0xf5, 0x5c, 0x73, 0x0b, 0x4a, 0x2a, 0xa5, 0xa4, 0xc0, 0x6c, 0xec, 0x0e,
	0xd4, 0xe7, 0x62, 0x87, 0xda, 0x8e, 0xa2, 0x45, 0x02, 0xae, 0x05, 0xcd, 0x6d, 0x4e, 0x1c, 0x51,
	0x6c, 0x7a, 0x7a, 0xfa, 0x45, 0x05, 0xc9, 0x49, 0x6c, 0x60, 0xbf, 0x1b, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x88, 0x9d, 0x05, 0x9e, 0x1b, 0x01, 0x00, 0x00,
}
