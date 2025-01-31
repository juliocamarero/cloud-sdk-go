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

package clusters_elasticsearch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// MoveEsClusterInstancesAdvancedReader is a Reader for the MoveEsClusterInstancesAdvanced structure.
type MoveEsClusterInstancesAdvancedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MoveEsClusterInstancesAdvancedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewMoveEsClusterInstancesAdvancedAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMoveEsClusterInstancesAdvancedBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMoveEsClusterInstancesAdvancedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMoveEsClusterInstancesAdvancedNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewMoveEsClusterInstancesAdvancedRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMoveEsClusterInstancesAdvancedAccepted creates a MoveEsClusterInstancesAdvancedAccepted with default headers values
func NewMoveEsClusterInstancesAdvancedAccepted() *MoveEsClusterInstancesAdvancedAccepted {
	return &MoveEsClusterInstancesAdvancedAccepted{}
}

/* MoveEsClusterInstancesAdvancedAccepted describes a response with status code 202, with default header values.

The move command was issued successfully. Use the "GET" command on the /{deployment_id} resource to monitor progress
*/
type MoveEsClusterInstancesAdvancedAccepted struct {
	Payload *models.ClusterCommandResponse
}

func (o *MoveEsClusterInstancesAdvancedAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/instances/_move][%d] moveEsClusterInstancesAdvancedAccepted  %+v", 202, o.Payload)
}
func (o *MoveEsClusterInstancesAdvancedAccepted) GetPayload() *models.ClusterCommandResponse {
	return o.Payload
}

func (o *MoveEsClusterInstancesAdvancedAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMoveEsClusterInstancesAdvancedBadRequest creates a MoveEsClusterInstancesAdvancedBadRequest with default headers values
func NewMoveEsClusterInstancesAdvancedBadRequest() *MoveEsClusterInstancesAdvancedBadRequest {
	return &MoveEsClusterInstancesAdvancedBadRequest{}
}

/* MoveEsClusterInstancesAdvancedBadRequest describes a response with status code 400, with default header values.

 * The cluster definition contained errors. (code: `clusters.cluster_invalid_plan`)
* The cluster definition contained errors. (code: `clusters.plan_feature_not_implemented`)
*/
type MoveEsClusterInstancesAdvancedBadRequest struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *MoveEsClusterInstancesAdvancedBadRequest) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/instances/_move][%d] moveEsClusterInstancesAdvancedBadRequest  %+v", 400, o.Payload)
}
func (o *MoveEsClusterInstancesAdvancedBadRequest) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *MoveEsClusterInstancesAdvancedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-error-codes
	hdrXCloudErrorCodes := response.GetHeader("x-cloud-error-codes")

	if hdrXCloudErrorCodes != "" {
		o.XCloudErrorCodes = hdrXCloudErrorCodes
	}

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMoveEsClusterInstancesAdvancedForbidden creates a MoveEsClusterInstancesAdvancedForbidden with default headers values
func NewMoveEsClusterInstancesAdvancedForbidden() *MoveEsClusterInstancesAdvancedForbidden {
	return &MoveEsClusterInstancesAdvancedForbidden{}
}

/* MoveEsClusterInstancesAdvancedForbidden describes a response with status code 403, with default header values.

The move command was prohibited for the given cluster. (code: `clusters.command_prohibited`)
*/
type MoveEsClusterInstancesAdvancedForbidden struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *MoveEsClusterInstancesAdvancedForbidden) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/instances/_move][%d] moveEsClusterInstancesAdvancedForbidden  %+v", 403, o.Payload)
}
func (o *MoveEsClusterInstancesAdvancedForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *MoveEsClusterInstancesAdvancedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-error-codes
	hdrXCloudErrorCodes := response.GetHeader("x-cloud-error-codes")

	if hdrXCloudErrorCodes != "" {
		o.XCloudErrorCodes = hdrXCloudErrorCodes
	}

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMoveEsClusterInstancesAdvancedNotFound creates a MoveEsClusterInstancesAdvancedNotFound with default headers values
func NewMoveEsClusterInstancesAdvancedNotFound() *MoveEsClusterInstancesAdvancedNotFound {
	return &MoveEsClusterInstancesAdvancedNotFound{}
}

/* MoveEsClusterInstancesAdvancedNotFound describes a response with status code 404, with default header values.

The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
*/
type MoveEsClusterInstancesAdvancedNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *MoveEsClusterInstancesAdvancedNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/instances/_move][%d] moveEsClusterInstancesAdvancedNotFound  %+v", 404, o.Payload)
}
func (o *MoveEsClusterInstancesAdvancedNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *MoveEsClusterInstancesAdvancedNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-error-codes
	hdrXCloudErrorCodes := response.GetHeader("x-cloud-error-codes")

	if hdrXCloudErrorCodes != "" {
		o.XCloudErrorCodes = hdrXCloudErrorCodes
	}

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMoveEsClusterInstancesAdvancedRetryWith creates a MoveEsClusterInstancesAdvancedRetryWith with default headers values
func NewMoveEsClusterInstancesAdvancedRetryWith() *MoveEsClusterInstancesAdvancedRetryWith {
	return &MoveEsClusterInstancesAdvancedRetryWith{}
}

/* MoveEsClusterInstancesAdvancedRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type MoveEsClusterInstancesAdvancedRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *MoveEsClusterInstancesAdvancedRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/instances/_move][%d] moveEsClusterInstancesAdvancedRetryWith  %+v", 449, o.Payload)
}
func (o *MoveEsClusterInstancesAdvancedRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *MoveEsClusterInstancesAdvancedRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header x-cloud-error-codes
	hdrXCloudErrorCodes := response.GetHeader("x-cloud-error-codes")

	if hdrXCloudErrorCodes != "" {
		o.XCloudErrorCodes = hdrXCloudErrorCodes
	}

	o.Payload = new(models.BasicFailedReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
