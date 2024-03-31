// Package auth provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package auth

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	externalRef0 "github.com/AgamBenItzhak/TaskTracker/api/schemas/auth"
	externalRef1 "github.com/AgamBenItzhak/TaskTracker/api/schemas/errors"
	externalRef2 "github.com/AgamBenItzhak/TaskTracker/api/schemas/project"
	externalRef3 "github.com/AgamBenItzhak/TaskTracker/api/schemas/task"
	externalRef4 "github.com/AgamBenItzhak/TaskTracker/api/schemas/user"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
)

// PostAuthChangePasswordJSONRequestBody defines body for PostAuthChangePassword for application/json ContentType.
type PostAuthChangePasswordJSONRequestBody = externalRef0.PasswordChangeRequest

// PostAuthLoginJSONRequestBody defines body for PostAuthLogin for application/json ContentType.
type PostAuthLoginJSONRequestBody = externalRef0.LoginRequest

// PostAuthRefreshtokenJSONRequestBody defines body for PostAuthRefreshtoken for application/json ContentType.
type PostAuthRefreshtokenJSONRequestBody = externalRef0.RefreshTokenRequest

// PostAuthRegisterJSONRequestBody defines body for PostAuthRegister for application/json ContentType.
type PostAuthRegisterJSONRequestBody = externalRef4.User

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Change the user's password
	// (POST /auth/change-password)
	PostAuthChangePassword(w http.ResponseWriter, r *http.Request)
	// Log in to the system
	// (POST /auth/login)
	PostAuthLogin(w http.ResponseWriter, r *http.Request)
	// Log out of the system
	// (POST /auth/logout)
	PostAuthLogout(w http.ResponseWriter, r *http.Request)
	// Refresh the user's access token
	// (POST /auth/refreshtoken)
	PostAuthRefreshtoken(w http.ResponseWriter, r *http.Request)
	// Register a new user
	// (POST /auth/register)
	PostAuthRegister(w http.ResponseWriter, r *http.Request)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Change the user's password
// (POST /auth/change-password)
func (_ Unimplemented) PostAuthChangePassword(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Log in to the system
// (POST /auth/login)
func (_ Unimplemented) PostAuthLogin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Log out of the system
// (POST /auth/logout)
func (_ Unimplemented) PostAuthLogout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Refresh the user's access token
// (POST /auth/refreshtoken)
func (_ Unimplemented) PostAuthRefreshtoken(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Register a new user
// (POST /auth/register)
func (_ Unimplemented) PostAuthRegister(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// PostAuthChangePassword operation middleware
func (siw *ServerInterfaceWrapper) PostAuthChangePassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostAuthChangePassword(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostAuthLogin operation middleware
func (siw *ServerInterfaceWrapper) PostAuthLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostAuthLogin(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostAuthLogout operation middleware
func (siw *ServerInterfaceWrapper) PostAuthLogout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostAuthLogout(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostAuthRefreshtoken operation middleware
func (siw *ServerInterfaceWrapper) PostAuthRefreshtoken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostAuthRefreshtoken(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostAuthRegister operation middleware
func (siw *ServerInterfaceWrapper) PostAuthRegister(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostAuthRegister(w, r)
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
		r.Post(options.BaseURL+"/auth/change-password", wrapper.PostAuthChangePassword)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/auth/login", wrapper.PostAuthLogin)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/auth/logout", wrapper.PostAuthLogout)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/auth/refreshtoken", wrapper.PostAuthRefreshtoken)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/auth/register", wrapper.PostAuthRegister)
	})

	return r
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xX32/jNgz+VwRtwF6cOu3dhoPf2mEPxQ5Ycbs9HQqDsxlbV1vSKLo/VuR/HyQ7iR03",
	"vQypuwDrmyBTH6mPH2nxUWamtkajZieTR+myEmsISyQy5BeWjEVihWG7RuegQL/kB4sykY5J6UJG8n5m",
	"wKpZZnIsUM/wnglmDEV7DqxjajJuyB9awSyXy0gS/tUowlwmX9YfrqMVvvnzK2Ysl5GsTKF06q3R8Tg0",
	"rEFVfrEwVAPLpNuJDoq0xfBhWnDuzlA+cLHePMzLGmZEyOoOa4vnmHHWaIdjaiDL0LmUzQ3qQ1M3wPLR",
	"4r1VhC5VesCM0vzubEOL0owF0r9lf4PdErMgdOXLXGQI5uHDKm0xH2WOC2gqf5ULBAqhH+KuBz5K8oDT",
	"QRwDenek3jT8TO7/y4pdiTbNStAF7q7drCFCzemUNTby4W+i8W5SpwP8EXejkLYC2o/TY8z8oLh2533S",
	"gt4Kd/h5n6Df2un/rp02Dp949WSEwJinwIOc5MA4Y1XjoZ1pA++vslDkONVQH1y6PSQPXMEL4W6APGxj",
	"8ynZ6cEvuwylKt8ujp/eH1ocK+CRnFYfBpnpsxn1MzjgYywxD670wrQl4TJSlpXRMpHnV5diYUh8Bnfz",
	"mSC7QYoECEvGnxQ1aCiwRs0CdC4Y3I1gb6Z0IdyDY6xPPAWKK++vhyLOry5lJG+RXOtpfnJ6MvdqNxY1",
	"WCUT+e5kfnIWnplcBnJiaLiM23/MrP+HtKZt5L4+wId+mctEXhnH5w2XP4cDV5v/Wdf6L0z+ECrJaEYd",
	"AMDaSmUBIv7qjN4MIH71PeFCJvK7eDOhxN14Eu96VmxljqnBsNH28XCts/l8wjC6H0aIY5jc3371fL9/",
	"Qe/tgPaEqwvIxacVI97n6fQ+/9BeLobU35h7pz++xkUvNSNpqMTvSLdI4pfOMJKuqWugB5nIVo+CSxS+",
	"jH9wov+wC43gi/Sxy2t/sFV9GKe+rfWPwWwaiQ9n3VcW9tY4+Sbno5HzR1MIpQWboOi26z8vZNPwXkr2",
	"dtNKajCmvmnqqDRlGhZmsZ+ouof+emh4Xlqf+tbT9MqnZ81X7pk7Zsc3nR+Nzjsl9t8C7aAo1oPiTsUX",
	"ynE3GH5D7Z3lNEoP0+lewj6dwOdWnh2S6OaenqC383ILlcqF0rbhI1NDmykBQuNdEMQTCvBHAobffZQN",
	"VTKRJbNN4rgyGVSlcZx8mH+Yx36OWl4v/wkAAP//gXHBL1MZAAA=",
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

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(pathPrefix, "../schemas/auth.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	for rawPath, rawFunc := range externalRef1.PathToRawSpec(path.Join(pathPrefix, "../schemas/errors.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	for rawPath, rawFunc := range externalRef2.PathToRawSpec(path.Join(pathPrefix, "../schemas/project.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	for rawPath, rawFunc := range externalRef3.PathToRawSpec(path.Join(pathPrefix, "../schemas/task.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	for rawPath, rawFunc := range externalRef4.PathToRawSpec(path.Join(pathPrefix, "../schemas/user.yaml")) {
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