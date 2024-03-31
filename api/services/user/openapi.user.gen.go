// Package user provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package user

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	externalRef0 "github.com/AgamBenItzhak/TaskTracker/api"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
)

// PutUsersUserIdJSONRequestBody defines body for PutUsersUserId for application/json ContentType.
type PutUsersUserIdJSONRequestBody = externalRef0.User

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Delete user by ID
	// (DELETE /users/{user_id})
	DeleteUsersUserId(w http.ResponseWriter, r *http.Request, userId int64)
	// Get user by ID
	// (GET /users/{user_id})
	GetUsersUserId(w http.ResponseWriter, r *http.Request, userId int64)
	// Update user by ID
	// (PUT /users/{user_id})
	PutUsersUserId(w http.ResponseWriter, r *http.Request, userId int64)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Delete user by ID
// (DELETE /users/{user_id})
func (_ Unimplemented) DeleteUsersUserId(w http.ResponseWriter, r *http.Request, userId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get user by ID
// (GET /users/{user_id})
func (_ Unimplemented) GetUsersUserId(w http.ResponseWriter, r *http.Request, userId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update user by ID
// (PUT /users/{user_id})
func (_ Unimplemented) PutUsersUserId(w http.ResponseWriter, r *http.Request, userId int64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// DeleteUsersUserId operation middleware
func (siw *ServerInterfaceWrapper) DeleteUsersUserId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId int64

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", chi.URLParam(r, "user_id"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteUsersUserId(w, r, userId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetUsersUserId operation middleware
func (siw *ServerInterfaceWrapper) GetUsersUserId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId int64

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", chi.URLParam(r, "user_id"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUsersUserId(w, r, userId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutUsersUserId operation middleware
func (siw *ServerInterfaceWrapper) PutUsersUserId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "user_id" -------------
	var userId int64

	err = runtime.BindStyledParameterWithOptions("simple", "user_id", chi.URLParam(r, "user_id"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "user_id", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutUsersUserId(w, r, userId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/users/{user_id}", wrapper.DeleteUsersUserId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/users/{user_id}", wrapper.GetUsersUserId)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/users/{user_id}", wrapper.PutUsersUserId)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/8SVz07kOBDGX8Wq3WO608uiFcqNVUsop+EAJ4SQSSppQ2KbcgVoobz7qJzQf6aZGaCn",
	"Zy7gtuzfV1VfpfwChWu9s2g5QPYCoVhgq+OyC0jy35PzSGww7haEmrG80Sy/KketrKDUjBM2LUICvPQI",
	"GQQmY2tI4HnitDeTwpVYo53gM5OesK4jr9U+MHUFdySXNvB9nwC22jRbQsPOfiIDQ/iVocA3Vrcox/Zh",
	"bpAE3OhfxF2DVtiAaA9S/DVdtLwO4clRuSW12txVeiLD+MU2S8iYOvyY9Ioryp0vD9ljG/ioFpBuzHaa",
	"xvJ/x2sZYxlrpI/qjOBeVEaSu73DgoctYysn10oMBRnPxlnI4PQ8V5UjdaHD/QXp4h4pUVp5cnJTtdrq",
	"Glu0rLQtFetwr1iOGVursAyM7VQCN9yI3gZFnZ7nkMAjUhiUZtN/pjPoE3AerfYGMvh3OpsegXjPi5hS",
	"KkmE9GXMpR/ibZBxN/J53FdayWF1u1S8QEMqn0tAMkO0HMzL1dFLYcufvIyapFtkpADZ1bfsfK5cJcAB",
	"zk6NUUgZY//wAhIYvrhV5RMgfOgMYfnak8N8e4/XfX8t14N3Ngyj72h2vJu0hD/GUkotj797yDpWlevs",
	"0OOha1tNy3XZXouWzyWW2FFXMZEA130CNfIu9wz5XeU+Q96r1iL+Wws9i0+Ns4w2pq29b0wRE0rvgrPr",
	"p0pWfxNWkMFf6fotS8eHLPbv8Lm9Yclox2dck9L/xDLfvWHZZRw+73LtvNvPtWHOHdy4hw4D/+/K5QE8",
	"2w6u/1N9Mr4Yn+uU0fEfNku8gfT4am1HDWSwYPZZmjau0M3CBc5OZiezVCZ1f91/DQAA//91rrahugkA",
	"AA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	pathPrefix := path.Dir(pathToFile)

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(pathPrefix, "../schema.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
