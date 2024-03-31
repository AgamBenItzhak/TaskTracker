// Package project provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package project

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
)

// Project defines model for project.
type Project struct {
	CreatedAt   time.Time  `json:"created_at" mapstructure:"created_at"`
	Description *string    `json:"description,omitempty" mapstructure:"description"`
	EndDate     *time.Time `json:"end_date,omitempty" mapstructure:"end_date"`
	ProjectId   int64      `json:"project_id" mapstructure:"project_id"`
	ProjectName string     `json:"project_name" mapstructure:"project_name"`
	StartDate   *time.Time `json:"start_date,omitempty" mapstructure:"start_date"`
	Status      *string    `json:"status,omitempty" mapstructure:"status"`
	UpdatedAt   time.Time  `json:"updated_at" mapstructure:"updated_at"`
}

// ProjectCreateRequest defines model for project_create_request.
type ProjectCreateRequest struct {
	Description *string    `json:"description,omitempty" mapstructure:"description"`
	EndDate     *time.Time `json:"end_date,omitempty" mapstructure:"end_date"`
	ProjectName string     `json:"project_name" mapstructure:"project_name"`
	StartDate   *time.Time `json:"start_date,omitempty" mapstructure:"start_date"`
	Status      *string    `json:"status,omitempty" mapstructure:"status"`
}

// ProjectCreateResponse defines model for project_create_response.
type ProjectCreateResponse struct {
	Project Project `json:"project"`
}

// ProjectDeleteRequest defines model for project_delete_request.
type ProjectDeleteRequest struct {
	ProjectId int64 `json:"project_id" mapstructure:"project_id"`
}

// ProjectDeleteResponse defines model for project_delete_response.
type ProjectDeleteResponse struct {
	Message string `json:"message" mapstructure:"message"`
}

// ProjectRequest defines model for project_request.
type ProjectRequest struct {
	ProjectId int64 `json:"project_id" mapstructure:"project_id"`
}

// ProjectResponse defines model for project_response.
type ProjectResponse struct {
	Project Project `json:"project"`
}

// ProjectUpdateRequest defines model for project_update_request.
type ProjectUpdateRequest struct {
	Description *string    `json:"description,omitempty" mapstructure:"description"`
	EndDate     *time.Time `json:"end_date,omitempty" mapstructure:"end_date"`
	ProjectId   int64      `json:"project_id" mapstructure:"project_id"`
	ProjectName *string    `json:"project_name,omitempty" mapstructure:"project_name"`
	StartDate   *time.Time `json:"start_date,omitempty" mapstructure:"start_date"`
	Status      *string    `json:"status,omitempty" mapstructure:"status"`
}

// ProjectUpdateResponse defines model for project_update_response.
type ProjectUpdateResponse struct {
	Project Project `json:"project"`
}

// ProjectUser defines model for project_user.
type ProjectUser struct {
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at"`
	ProjectId int64     `json:"project_id" mapstructure:"project_id"`
	Role      string    `json:"role" mapstructure:"role"`
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at"`
	UserId    int64     `json:"user_id" mapstructure:"user_id"`
}

// ProjectUserCreateRequest defines model for project_user_create_request.
type ProjectUserCreateRequest struct {
	ProjectId int64  `json:"project_id" mapstructure:"project_id"`
	Role      string `json:"role" mapstructure:"role"`
	UserId    int64  `json:"user_id" mapstructure:"user_id"`
}

// ProjectUserCreateResponse defines model for project_user_create_response.
type ProjectUserCreateResponse struct {
	ProjectUser ProjectUser `json:"project_user"`
}

// ProjectUserDeleteRequest defines model for project_user_delete_request.
type ProjectUserDeleteRequest struct {
	ProjectId int64 `json:"project_id" mapstructure:"project_id"`
	UserId    int64 `json:"user_id" mapstructure:"user_id"`
}

// ProjectUserDeleteResponse defines model for project_user_delete_response.
type ProjectUserDeleteResponse struct {
	Message string `json:"message" mapstructure:"message"`
}

// ProjectUserRequest defines model for project_user_request.
type ProjectUserRequest struct {
	ProjectId int64 `json:"project_id" mapstructure:"project_id"`
	UserId    int64 `json:"user_id" mapstructure:"user_id"`
}

// ProjectUserResponse defines model for project_user_response.
type ProjectUserResponse struct {
	ProjectUser ProjectUser `json:"project_user"`
}

// ProjectUserUpdateRequest defines model for project_user_update_request.
type ProjectUserUpdateRequest struct {
	ProjectId int64   `json:"project_id" mapstructure:"project_id"`
	Role      *string `json:"role,omitempty" mapstructure:"role"`
	UserId    int64   `json:"user_id" mapstructure:"user_id"`
}

// ProjectUserUpdateResponse defines model for project_user_update_response.
type ProjectUserUpdateResponse struct {
	ProjectUser ProjectUser `json:"project_user"`
}

// ProjectUsersRequest defines model for project_users_request.
type ProjectUsersRequest struct {
	ProjectId int64 `json:"project_id" mapstructure:"project_id"`
}

// ProjectUsersResponse defines model for project_users_response.
type ProjectUsersResponse struct {
	ProjectUsers []ProjectUser `json:"project_users" mapstructure:"project_users"`
}

// ProjectsRequest defines model for projects_request.
type ProjectsRequest struct {
	ProjectIds []int64 `json:"project_ids" mapstructure:"project_ids"`
}

// ProjectsResponse defines model for projects_response.
type ProjectsResponse struct {
	Projects []Project `json:"projects" mapstructure:"projects"`
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xXTW/bMAz9L9yONnIZdvBfGYZAsxlXQyxpFA20KPzfB8kfVdraVuo6cbbcjEh5fOR7",
	"pOlnyHVltELFFrJnsPkDVsI/GtK/Mefu0SCxRH+QEwrGYi/82UFT5Z6gEIwpywohAX4yCBlYJqlKSOAx",
	"1cLINNcFlqhSfGQSKYvS41XCWKY655rcnwL4pkmgQJuTNCy1cpeXIIdQDhpVsXesV0ljAHeRulruZXES",
	"Syr+/u0ljlSMJdKZgQLsMJQSFS4t2AmWA7csiNerWQDfRePaLk2iQ3GAtSnWdG4A37hwhH9qSVhA9iNU",
	"6ZVGSej4E44/Bzr6l2/FQN72P3sXAu07TfpPtc3dy6+9POIun1yUbazRyuJb3wRT/yvhATL4snt5Rey6",
	"98OuvzZCZJJDgUecsu7FhuV4i0bRHythhdaKcrFhe5g3PPuDKZK3WtyrGrMdvf/JTL2vItsb3zPNMfjz",
	"mj1ika79SXBBC5M+Lraux7j0+pmAU2qVEvXA0ztuf6ur4ce2XA8yt+rerB02JVC8DDPTZxgRESOovTtG",
	"0x/O8trMOrkZReNLtuUV1jO9qxqv6uY6c26fvg/udQ0Ruy9e2Bf2Zr9Oe/Yx9fQ/SMbKnlfZIbwgEk8f",
	"zLEl0EzJZKcyjZLoNMM5sT4nMRd1Qr3ZpGaUO1u0z0lrPKf3EnJXpTpoyFR9PCagDSphJGQACRjBD7Y9",
	"af4GAAD//1E5BPTdGQAA",
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