// Package task provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package task

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

// Task defines model for task.
type Task struct {
	CreatedAt    time.Time  `json:"created_at" mapstructure:"created_at"`
	Description  *string    `json:"description,omitempty" mapstructure:"description"`
	EndDate      *time.Time `json:"end_date,omitempty" mapstructure:"end_date"`
	Priority     *string    `json:"priority,omitempty" mapstructure:"priority"`
	StartDate    *time.Time `json:"start_date,omitempty" mapstructure:"start_date"`
	Status       *string    `json:"status,omitempty" mapstructure:"status"`
	TaskId       int64      `json:"task_id" mapstructure:"task_id"`
	TaskName     string     `json:"task_name" mapstructure:"task_name"`
	TasksGroupId int64      `json:"tasks_group_id" mapstructure:"tasks_group_id"`
	UpdatedAt    time.Time  `json:"updated_at" mapstructure:"updated_at"`
}

// TaskCreateRequest defines model for task_create_request.
type TaskCreateRequest struct {
	Description  *string    `json:"description,omitempty" mapstructure:"description"`
	EndDate      *time.Time `json:"end_date,omitempty" mapstructure:"end_date"`
	Priority     *string    `json:"priority,omitempty" mapstructure:"priority"`
	StartDate    *time.Time `json:"start_date,omitempty" mapstructure:"start_date"`
	Status       *string    `json:"status,omitempty" mapstructure:"status"`
	TaskName     string     `json:"task_name" mapstructure:"task_name"`
	TasksGroupId int64      `json:"tasks_group_id" mapstructure:"tasks_group_id"`
}

// TaskCreateResponse defines model for task_create_response.
type TaskCreateResponse struct {
	Task Task `json:"task"`
}

// TaskDeleteRequest defines model for task_delete_request.
type TaskDeleteRequest struct {
	TaskId int64 `json:"task_id" mapstructure:"task_id"`
}

// TaskDeleteResponse defines model for task_delete_response.
type TaskDeleteResponse struct {
	Message string `json:"message" mapstructure:"message"`
}

// TaskRequest defines model for task_request.
type TaskRequest struct {
	TaskId int64 `json:"task_id" mapstructure:"task_id"`
}

// TaskResponse defines model for task_response.
type TaskResponse struct {
	Task Task `json:"task"`
}

// TaskUpdateRequest defines model for task_update_request.
type TaskUpdateRequest struct {
	Description  *string    `json:"description,omitempty" mapstructure:"description"`
	EndDate      *time.Time `json:"end_date,omitempty" mapstructure:"end_date"`
	Priority     *string    `json:"priority,omitempty" mapstructure:"priority"`
	StartDate    *time.Time `json:"start_date,omitempty" mapstructure:"start_date"`
	Status       *string    `json:"status,omitempty" mapstructure:"status"`
	TaskId       int64      `json:"task_id" mapstructure:"task_id"`
	TaskName     *string    `json:"task_name,omitempty" mapstructure:"task_name"`
	TasksGroupId *int64     `json:"tasks_group_id,omitempty" mapstructure:"tasks_group_id"`
}

// TaskUpdateResponse defines model for task_update_response.
type TaskUpdateResponse struct {
	Task Task `json:"task"`
}

// TaskUser defines model for task_user.
type TaskUser struct {
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at"`
	Role      string    `json:"role" mapstructure:"role"`
	TaskId    int64     `json:"task_id" mapstructure:"task_id"`
	UpdatedAt time.Time `json:"updated_at" mapstructure:"updated_at"`
	UserId    int64     `json:"user_id" mapstructure:"user_id"`
}

// TaskUserCreateRequest defines model for task_user_create_request.
type TaskUserCreateRequest struct {
	Role   string `json:"role" mapstructure:"role"`
	TaskId int64  `json:"task_id" mapstructure:"task_id"`
	UserId int64  `json:"user_id" mapstructure:"user_id"`
}

// TaskUserCreateResponse defines model for task_user_create_response.
type TaskUserCreateResponse struct {
	TaskUser TaskUser `json:"task_user"`
}

// TaskUserDeleteRequest defines model for task_user_delete_request.
type TaskUserDeleteRequest struct {
	TaskId int64 `json:"task_id" mapstructure:"task_id"`
	UserId int64 `json:"user_id" mapstructure:"user_id"`
}

// TaskUserDeleteResponse defines model for task_user_delete_response.
type TaskUserDeleteResponse struct {
	Message string `json:"message" mapstructure:"message"`
}

// TaskUserRequest defines model for task_user_request.
type TaskUserRequest struct {
	TaskId int64 `json:"task_id" mapstructure:"task_id"`
	UserId int64 `json:"user_id" mapstructure:"user_id"`
}

// TaskUserResponse defines model for task_user_response.
type TaskUserResponse struct {
	TaskUser TaskUser `json:"task_user"`
}

// TaskUserUpdateRequest defines model for task_user_update_request.
type TaskUserUpdateRequest struct {
	Role   *string `json:"role,omitempty" mapstructure:"role"`
	TaskId int64   `json:"task_id" mapstructure:"task_id"`
	UserId int64   `json:"user_id" mapstructure:"user_id"`
}

// TaskUserUpdateResponse defines model for task_user_update_response.
type TaskUserUpdateResponse struct {
	TaskUser TaskUser `json:"task_user"`
}

// TaskUsersRequest defines model for task_users_request.
type TaskUsersRequest struct {
	TaskId int64 `json:"task_id" mapstructure:"task_id"`
}

// TaskUsersResponse defines model for task_users_response.
type TaskUsersResponse struct {
	TaskUsers []TaskUser `json:"task_users" mapstructure:"task_users"`
}

// TasksGroup defines model for tasks_group.
type TasksGroup struct {
	CreatedAt    time.Time  `json:"created_at" mapstructure:"created_at"`
	Description  *string    `json:"description,omitempty" mapstructure:"description"`
	EndDate      *time.Time `json:"end_date,omitempty" mapstructure:"end_date"`
	GroupName    string     `json:"group_name" mapstructure:"group_name"`
	Priority     *string    `json:"priority,omitempty" mapstructure:"priority"`
	ProjectId    int64      `json:"project_id" mapstructure:"project_id"`
	StartDate    *time.Time `json:"start_date,omitempty" mapstructure:"start_date"`
	Status       *string    `json:"status,omitempty" mapstructure:"status"`
	TasksGroupId int64      `json:"tasks_group_id" mapstructure:"tasks_group_id"`
	UpdatedAt    time.Time  `json:"updated_at" mapstructure:"updated_at"`
}

// TasksGroupCreateRequest defines model for tasks_group_create_request.
type TasksGroupCreateRequest struct {
	Description *string    `json:"description,omitempty" mapstructure:"description"`
	EndDate     *time.Time `json:"end_date,omitempty" mapstructure:"end_date"`
	GroupName   string     `json:"group_name" mapstructure:"group_name"`
	Priority    *string    `json:"priority,omitempty" mapstructure:"priority"`
	ProjectId   int64      `json:"project_id" mapstructure:"project_id"`
	StartDate   *time.Time `json:"start_date,omitempty" mapstructure:"start_date"`
	Status      *string    `json:"status,omitempty" mapstructure:"status"`
}

// TasksGroupCreateResponse defines model for tasks_group_create_response.
type TasksGroupCreateResponse struct {
	TasksGroup TasksGroup `json:"tasks_group"`
}

// TasksGroupDeleteRequest defines model for tasks_group_delete_request.
type TasksGroupDeleteRequest struct {
	TasksGroupId int64 `json:"tasks_group_id" mapstructure:"tasks_group_id"`
}

// TasksGroupDeleteResponse defines model for tasks_group_delete_response.
type TasksGroupDeleteResponse struct {
	Message string `json:"message" mapstructure:"message"`
}

// TasksGroupRequest defines model for tasks_group_request.
type TasksGroupRequest struct {
	TasksGroupId int64 `json:"tasks_group_id" mapstructure:"tasks_group_id"`
}

// TasksGroupResponse defines model for tasks_group_response.
type TasksGroupResponse struct {
	TasksGroup TasksGroup `json:"tasks_group"`
}

// TasksGroupUpdateRequest defines model for tasks_group_update_request.
type TasksGroupUpdateRequest struct {
	Description  *string    `json:"description,omitempty" mapstructure:"description"`
	EndDate      *time.Time `json:"end_date,omitempty" mapstructure:"end_date"`
	GroupName    *string    `json:"group_name,omitempty" mapstructure:"group_name"`
	Priority     *string    `json:"priority,omitempty" mapstructure:"priority"`
	ProjectId    *int64     `json:"project_id,omitempty" mapstructure:"project_id"`
	StartDate    *time.Time `json:"start_date,omitempty" mapstructure:"start_date"`
	Status       *string    `json:"status,omitempty" mapstructure:"status"`
	TasksGroupId int64      `json:"tasks_group_id" mapstructure:"tasks_group_id"`
}

// TasksGroupUpdateResponse defines model for tasks_group_update_response.
type TasksGroupUpdateResponse struct {
	TasksGroup TasksGroup `json:"tasks_group"`
}

// TasksGroupUser defines model for tasks_group_user.
type TasksGroupUser struct {
	CreatedAt    time.Time `json:"created_at" mapstructure:"created_at"`
	Role         string    `json:"role" mapstructure:"role"`
	TasksGroupId int64     `json:"tasks_group_id" mapstructure:"tasks_group_id"`
	UpdatedAt    time.Time `json:"updated_at" mapstructure:"updated_at"`
	UserId       int64     `json:"user_id" mapstructure:"user_id"`
}

// TasksGroupUserCreateRequest defines model for tasks_group_user_create_request.
type TasksGroupUserCreateRequest struct {
	Role         string `json:"role" mapstructure:"role"`
	TasksGroupId int64  `json:"tasks_group_id" mapstructure:"tasks_group_id"`
	UserId       int64  `json:"user_id" mapstructure:"user_id"`
}

// TasksGroupUserCreateResponse defines model for tasks_group_user_create_response.
type TasksGroupUserCreateResponse struct {
	TasksGroupUser TasksGroupUser `json:"tasks_group_user"`
}

// TasksGroupUserDeleteRequest defines model for tasks_group_user_delete_request.
type TasksGroupUserDeleteRequest struct {
	TasksGroupId int64 `json:"tasks_group_id" mapstructure:"tasks_group_id"`
	UserId       int64 `json:"user_id" mapstructure:"user_id"`
}

// TasksGroupUserDeleteResponse defines model for tasks_group_user_delete_response.
type TasksGroupUserDeleteResponse struct {
	Message string `json:"message" mapstructure:"message"`
}

// TasksGroupUserRequest defines model for tasks_group_user_request.
type TasksGroupUserRequest struct {
	TasksGroupId int64 `json:"tasks_group_id" mapstructure:"tasks_group_id"`
	UserId       int64 `json:"user_id" mapstructure:"user_id"`
}

// TasksGroupUserResponse defines model for tasks_group_user_response.
type TasksGroupUserResponse struct {
	TasksGroupUser TasksGroupUser `json:"tasks_group_user"`
}

// TasksGroupUserUpdateRequest defines model for tasks_group_user_update_request.
type TasksGroupUserUpdateRequest struct {
	Role         *string `json:"role,omitempty" mapstructure:"role"`
	TasksGroupId int64   `json:"tasks_group_id" mapstructure:"tasks_group_id"`
	UserId       int64   `json:"user_id" mapstructure:"user_id"`
}

// TasksGroupUserUpdateResponse defines model for tasks_group_user_update_response.
type TasksGroupUserUpdateResponse struct {
	TasksGroupUser TasksGroupUser `json:"tasks_group_user"`
}

// TasksGroupUsersRequest defines model for tasks_group_users_request.
type TasksGroupUsersRequest struct {
	TasksGroupId int64 `json:"tasks_group_id" mapstructure:"tasks_group_id"`
}

// TasksGroupUsersResponse defines model for tasks_group_users_response.
type TasksGroupUsersResponse struct {
	TasksGroupUsers []TasksGroupUser `json:"tasks_group_users" mapstructure:"tasks_group_users"`
}

// TasksGroupsRequest defines model for tasks_groups_request.
type TasksGroupsRequest struct {
	TasksGroupIds []int64 `json:"tasks_group_ids" mapstructure:"tasks_group_ids"`
}

// TasksGroupsResponse defines model for tasks_groups_response.
type TasksGroupsResponse struct {
	TasksGroups []TasksGroup `json:"tasks_groups" mapstructure:"tasks_groups"`
}

// TasksRequest defines model for tasks_request.
type TasksRequest struct {
	TaskIds []int64 `json:"task_ids" mapstructure:"task_ids"`
}

// TasksResponse defines model for tasks_response.
type TasksResponse struct {
	Tasks []Task `json:"tasks" mapstructure:"tasks"`
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xZzXLbIBB+F9qjPLl0evCrdDoeKm0cWgsorGaSyfjdO0hGQbUlUPQTInNqXOOP3f32",
	"g132leSilIIDR032r0TnT1DS+k+k+o/5VyohQSGD+n9zBRShOFA0nx6FKs1fpKAIO2QlkIzgiwSyJxoV",
	"40eSkeedoJLtclHAEfgOnlHRHdJjjVdSqVFVOVbK/MiBP58zUoDOFZPIBK9NmoDsQhlo4MXBWL2IGy24",
	"2UkqJhTDl6ketDgGVCNVuJwDDvxlN6z0VAcuKAbQJNeBFR3bGcfv397sZhzhCGrkHha43YTTEqYa/gZk",
	"YfXhqEQlF3PBwTdbVrJYUnUO/Nlsp+BvxRQUZP+jDeiV3258M1e4HXN/tpaJX78hR2J5aX5wMFuBxuuD",
	"Jkl/y9LfgCpv6KRHHX4NaCm4hmsR2Dv4q4JHsidfHt5u64fLVf1Qr7llTf++BZxgSHsrnM49p4zf5L5Q",
	"laA1PU5OKgtzZaH9otfCzxTN1TOuuRHSaZ8KvVTojbpSBmXcqmptNWtQH90aKnGaTHeNsVqWrlnEZ8Rw",
	"tIhLFnigU7BLLiy9ozuoEXwtwidMgXhICQz90MnSngO+46VZeNO0+pthWyKoVuNgLjBM0VbItZmJRl98",
	"IpCcr1xPJ+/8xAcVc+vwrz9VM2st9sat/sQQSj0igu2uVCn68h6nmq3PvVzoXtcuzUGawszQnDdt1hxN",
	"ooO0WNcvlTCJsIjSHOyNPDDc1zSm88rskNnJ8fE9l4W9i8FMOg22eBr8J5Y+bQQrYKiicO5mXyVhlw6I",
	"2WtTSBMc0fgr2J1Ym1Vr6FYCHlUy38U0KF0xqeCMQPhBzwpr63+L86MNNyEf9XzW6XQmD5Xc7Ft/trR2",
	"ekRH2DhaAk6r4LfQzvohswdfRjt2RleOR8X4uBDG3gL4x1aJaU/sIlfz6uOulDBhfMSdN3ozrwMhI7TO",
	"2lGTtG7kJw/UuoacfVR6x2vBNHa99hE6n6NmZw/DQU4GsPsuYudzddjPISd9A+u1uesnzcOWh6ZR/MxB",
	"TA8jNzww6xh/FGTPq9MpI0ICp5KRPSEZkRSfdPPN+V8AAAD//50Kich1OAAA",
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
