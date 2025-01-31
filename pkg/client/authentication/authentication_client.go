// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by go-swagger; DO NOT EDIT.

package authentication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new authentication API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for authentication API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateAPIKey(params *CreateAPIKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAPIKeyCreated, error)

	DeleteAPIKey(params *DeleteAPIKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAPIKeyOK, error)

	DeleteAPIKeys(params *DeleteAPIKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAPIKeysOK, error)

	DeleteUserAPIKey(params *DeleteUserAPIKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserAPIKeyOK, error)

	DeleteUserAPIKeys(params *DeleteUserAPIKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserAPIKeysOK, error)

	DeleteUsersAPIKeys(params *DeleteUsersAPIKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUsersAPIKeysOK, error)

	DisableElevatedPermissions(params *DisableElevatedPermissionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DisableElevatedPermissionsOK, error)

	EnableElevatedPermissions(params *EnableElevatedPermissionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EnableElevatedPermissionsOK, error)

	GetAPIKey(params *GetAPIKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPIKeyOK, error)

	GetAPIKeys(params *GetAPIKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPIKeysOK, error)

	GetAuthenticationInfo(params *GetAuthenticationInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAuthenticationInfoOK, error)

	GetUserAPIKey(params *GetUserAPIKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserAPIKeyOK, error)

	GetUserAPIKeys(params *GetUserAPIKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserAPIKeysOK, error)

	GetUsersAPIKeys(params *GetUsersAPIKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUsersAPIKeysOK, error)

	Login(params *LoginParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LoginOK, error)

	Logout(params *LogoutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LogoutOK, error)

	Methods(params *MethodsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MethodsOK, error)

	ReAuthenticate(params *ReAuthenticateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReAuthenticateOK, error)

	RefreshToken(params *RefreshTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RefreshTokenOK, error)

	SamlCallback(params *SamlCallbackParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) error

	SamlInit(params *SamlInitParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) error

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateAPIKey creates API key

  Creates a new API key.
*/
func (a *Client) CreateAPIKey(params *CreateAPIKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateAPIKeyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAPIKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "create-api-key",
		Method:             "POST",
		PathPattern:        "/users/auth/keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateAPIKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateAPIKeyCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for create-api-key: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteAPIKey deletes API key

  Delete or invalidate the API key.
*/
func (a *Client) DeleteAPIKey(params *DeleteAPIKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAPIKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPIKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-api-key",
		Method:             "DELETE",
		PathPattern:        "/users/auth/keys/{api_key_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAPIKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAPIKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-api-key: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteAPIKeys deletes API keys

  Delete or invalidate API keys.
*/
func (a *Client) DeleteAPIKeys(params *DeleteAPIKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteAPIKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPIKeysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-api-keys",
		Method:             "DELETE",
		PathPattern:        "/users/auth/keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAPIKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteAPIKeysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-api-keys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteUserAPIKey deletes an API key for a user

  Delete or invalidate an API key for a user.
*/
func (a *Client) DeleteUserAPIKey(params *DeleteUserAPIKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserAPIKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserAPIKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-user-api-key",
		Method:             "DELETE",
		PathPattern:        "/users/{user_id}/auth/keys/{api_key_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserAPIKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteUserAPIKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-user-api-key: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteUserAPIKeys deletes API keys for a user

  Delete or invalidate all of the API keys for a user.
*/
func (a *Client) DeleteUserAPIKeys(params *DeleteUserAPIKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUserAPIKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserAPIKeysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-user-api-keys",
		Method:             "DELETE",
		PathPattern:        "/users/{user_id}/auth/keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserAPIKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteUserAPIKeysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-user-api-keys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteUsersAPIKeys deletes API keys of multiple users

  Delete or invalidate the API keys for multiple users.
*/
func (a *Client) DeleteUsersAPIKeys(params *DeleteUsersAPIKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteUsersAPIKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUsersAPIKeysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete-users-api-keys",
		Method:             "DELETE",
		PathPattern:        "/users/auth/keys/_all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUsersAPIKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteUsersAPIKeysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete-users-api-keys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DisableElevatedPermissions disables elevated permissions

  Disables elevated permissions for the user.
*/
func (a *Client) DisableElevatedPermissions(params *DisableElevatedPermissionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DisableElevatedPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisableElevatedPermissionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "disable-elevated-permissions",
		Method:             "DELETE",
		PathPattern:        "/users/auth/_elevate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DisableElevatedPermissionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DisableElevatedPermissionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for disable-elevated-permissions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EnableElevatedPermissions enables elevated permissions

  Enables the elevated permissions for the current user. Elevated permissions allow the user to complete potentially destructive operations on clusters. Elevated permissions are available for a limited period of time and automatically expire if you do not renew them.
*/
func (a *Client) EnableElevatedPermissions(params *EnableElevatedPermissionsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*EnableElevatedPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnableElevatedPermissionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "enable-elevated-permissions",
		Method:             "POST",
		PathPattern:        "/users/auth/_elevate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &EnableElevatedPermissionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EnableElevatedPermissionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for enable-elevated-permissions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIKey gets API key

  Retrieves the metadata for an API key.
*/
func (a *Client) GetAPIKey(params *GetAPIKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPIKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-api-key",
		Method:             "GET",
		PathPattern:        "/users/auth/keys/{api_key_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAPIKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-api-key: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPIKeys gets all API keys

  Retrieves the metadata for all of the API keys that the user generated.
*/
func (a *Client) GetAPIKeys(params *GetAPIKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAPIKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIKeysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-api-keys",
		Method:             "GET",
		PathPattern:        "/users/auth/keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAPIKeysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-api-keys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAuthenticationInfo users authentication information

  Provides authentication information about a user, including elevated permission status and TOTP device availability.
*/
func (a *Client) GetAuthenticationInfo(params *GetAuthenticationInfoParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetAuthenticationInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAuthenticationInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-authentication-info",
		Method:             "GET",
		PathPattern:        "/users/auth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAuthenticationInfoReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAuthenticationInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-authentication-info: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUserAPIKey gets a user API key

  Retrieves the API key metadata for a user.
*/
func (a *Client) GetUserAPIKey(params *GetUserAPIKeyParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserAPIKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserAPIKeyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-user-api-key",
		Method:             "GET",
		PathPattern:        "/users/{user_id}/auth/keys/{api_key_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserAPIKeyReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUserAPIKeyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-user-api-key: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUserAPIKeys gets API key metadata for all keys created by the user

  Retrieves metadata for all API keys created by the given user.
*/
func (a *Client) GetUserAPIKeys(params *GetUserAPIKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserAPIKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserAPIKeysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-user-api-keys",
		Method:             "GET",
		PathPattern:        "/users/{user_id}/auth/keys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserAPIKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUserAPIKeysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-user-api-keys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUsersAPIKeys gets all API keys for all users

  Retrieves the metadata for all of the API keys for all users.
*/
func (a *Client) GetUsersAPIKeys(params *GetUsersAPIKeysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUsersAPIKeysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersAPIKeysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get-users-api-keys",
		Method:             "GET",
		PathPattern:        "/users/auth/keys/_all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersAPIKeysReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetUsersAPIKeysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get-users-api-keys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Login logins to e c e

  Authenticates against available users.
*/
func (a *Client) Login(params *LoginParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LoginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoginParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "login",
		Method:             "POST",
		PathPattern:        "/users/auth/_login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LoginReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LoginOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for login: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Logout logouts from e c e

  Destroys the current session.
*/
func (a *Client) Logout(params *LogoutParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*LogoutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLogoutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "logout",
		Method:             "POST",
		PathPattern:        "/users/auth/_logout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &LogoutReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LogoutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for logout: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Methods availables authentication methods

  Provides information about available authentication methods.
*/
func (a *Client) Methods(params *MethodsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*MethodsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMethodsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "methods",
		Method:             "GET",
		PathPattern:        "/users/auth/methods",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MethodsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MethodsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for methods: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReAuthenticate res authenticate to generate a token

  > WARNING
> This endpoint is deprecated and scheduled to be removed in the next major version.

Re-authenticate.
*/
func (a *Client) ReAuthenticate(params *ReAuthenticateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReAuthenticateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReAuthenticateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "re-authenticate",
		Method:             "POST",
		PathPattern:        "/users/auth/reauthenticate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ReAuthenticateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReAuthenticateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for re-authenticate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RefreshToken refreshes authentication token

  Issues a new authentication token.
*/
func (a *Client) RefreshToken(params *RefreshTokenParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RefreshTokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRefreshTokenParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "refresh-token",
		Method:             "POST",
		PathPattern:        "/users/auth/_refresh",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RefreshTokenReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RefreshTokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for refresh-token: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SamlCallback s a m l callback

  Accepts a callback request from an identity provider and authenticates the user.
*/
func (a *Client) SamlCallback(params *SamlCallbackParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSamlCallbackParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "saml-callback",
		Method:             "POST",
		PathPattern:        "/users/auth/saml/_callback",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SamlCallbackReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

/*
  SamlInit initiates s a m l protocol

  Calls the authentication cluster to initiate SAML Single Sign-on (Web Browser SSO profile) protocol and redirects the user to the identity provider for authentication. The authentication cluster must be configured prior to initiation.
*/
func (a *Client) SamlInit(params *SamlInitParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSamlInitParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "saml-init",
		Method:             "GET",
		PathPattern:        "/users/auth/saml/_init",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SamlInitReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	_, err := a.transport.Submit(op)
	if err != nil {
		return err
	}
	return nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
