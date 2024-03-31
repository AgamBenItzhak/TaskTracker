// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package api

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
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Project defines model for project.
type Project struct {
	CreatedAt   *time.Time          `json:"created_at,omitempty" mapstructure:"created_at"`
	Description *string             `json:"description,omitempty" mapstructure:"description"`
	EndDate     *openapi_types.Date `json:"end_date,omitempty" mapstructure:"end_date"`
	ProjectId   *int64              `json:"project_id,omitempty" mapstructure:"project_id"`
	ProjectName *string             `json:"project_name,omitempty" mapstructure:"project_name"`
	StartDate   *openapi_types.Date `json:"start_date,omitempty" mapstructure:"start_date"`
	Status      *string             `json:"status,omitempty" mapstructure:"status"`
	UpdatedAt   *time.Time          `json:"updated_at,omitempty" mapstructure:"updated_at"`
}

// ProjectUser defines model for project_user.
type ProjectUser struct {
	CreatedAt *time.Time `json:"created_at,omitempty" mapstructure:"created_at"`
	ProjectId *int64     `json:"project_id,omitempty" mapstructure:"project_id"`
	Role      *string    `json:"role,omitempty" mapstructure:"role"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" mapstructure:"updated_at"`
	UserId    *int64     `json:"user_id,omitempty" mapstructure:"user_id"`
}

// Task defines model for task.
type Task struct {
	CreatedAt   *time.Time          `json:"created_at,omitempty" mapstructure:"created_at"`
	Description *string             `json:"description,omitempty" mapstructure:"description"`
	EndDate     *openapi_types.Date `json:"end_date,omitempty" mapstructure:"end_date"`
	Priority    *string             `json:"priority,omitempty" mapstructure:"priority"`
	StartDate   *openapi_types.Date `json:"start_date,omitempty" mapstructure:"start_date"`
	Status      *string             `json:"status,omitempty" mapstructure:"status"`
	TaskGroupId *int64              `json:"task_group_id,omitempty" mapstructure:"task_group_id"`
	TaskId      *int64              `json:"task_id,omitempty" mapstructure:"task_id"`
	TaskName    *string             `json:"task_name,omitempty" mapstructure:"task_name"`
	UpdatedAt   *time.Time          `json:"updated_at,omitempty" mapstructure:"updated_at"`
}

// TaskGroup defines model for task_group.
type TaskGroup struct {
	CreatedAt   *time.Time `json:"created_at,omitempty" mapstructure:"created_at"`
	Description *string    `json:"description,omitempty" mapstructure:"description"`
	GroupName   *string    `json:"group_name,omitempty" mapstructure:"group_name"`
	ProjectId   *int64     `json:"project_id,omitempty" mapstructure:"project_id"`
	TaskGroupId *int64     `json:"task_group_id,omitempty" mapstructure:"task_group_id"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" mapstructure:"updated_at"`
}

// User defines model for user.
type User struct {
	CreatedAt *time.Time           `json:"created_at,omitempty" mapstructure:"created_at"`
	Email     *openapi_types.Email `json:"email,omitempty" mapstructure:"email"`
	FirstName *string              `json:"first_name,omitempty" mapstructure:"first_name"`
	LastName  *string              `json:"last_name,omitempty" mapstructure:"last_name"`
	LastSeen  *time.Time           `json:"last_seen,omitempty" mapstructure:"last_seen"`
	UpdatedAt *time.Time           `json:"updated_at,omitempty" mapstructure:"updated_at"`
	UserId    *int64               `json:"user_id,omitempty" mapstructure:"user_id"`
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+yWO2/bMBDHvwpxsyS4D3TwljFbhuzGVTorrM0HjqciRqDvXpysWhTapZDlpG02Qo/f",
	"nzz+7/ECdXAxePKSYPsCqX4ih8MycvhGtYzLSCyWhhc1Ewo1Oxze7QM7XUGDQqVYR1CAnCLBFpKw9S0U",
	"8FwGjLasQ0Mt+ZKehbEUbAeew5iEu1o61p8yfN8X0FCq2UaxwevHS8g5StHkm53u+pdjLDzBhasiYxh3",
	"tpnJWC9fPk861gu1xH8olLFzKY+OlsZqxlJ4EmRZJVwZeRSSLi3d/0hRYBebNf2a4XuVG1nh65A92bV0",
	"ifi1s+mGbuRwXOzCgXHrOyxAb2qVEP0E/94ogunwXm6Xllsb2Mppef0bOX937VNL7VoOXVzFz3P8RXA1",
	"qVzkGl1uAr2FRjFF8x+uAme3XOPyMtKNW9vNs+q1nfkWRhdyaI8zofOThc1jYCh/bzldZ3bOSAo+4pW4",
	"E+iCTUR+leBP9P9k/NJH1u+D/jYrf3D3cG/2gc0jpsMjY30gLgyasSYYhx5bcuTFoG+MJq8R/cz61qRT",
	"EnKVbtyKzsOQUczdwz0U8J04nZU21Ydqo/kWInmMFrbwqdpUH6GAiPKkR+r7HwEAAP//3XkY8bMPAAA=",
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
