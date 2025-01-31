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

package clusters_apm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/elastic/cloud-sdk-go/pkg/models"
)

// StartApmMaintenanceModeReader is a Reader for the StartApmMaintenanceMode structure.
type StartApmMaintenanceModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartApmMaintenanceModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewStartApmMaintenanceModeAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewStartApmMaintenanceModeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStartApmMaintenanceModeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewStartApmMaintenanceModeRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStartApmMaintenanceModeAccepted creates a StartApmMaintenanceModeAccepted with default headers values
func NewStartApmMaintenanceModeAccepted() *StartApmMaintenanceModeAccepted {
	return &StartApmMaintenanceModeAccepted{}
}

/* StartApmMaintenanceModeAccepted describes a response with status code 202, with default header values.

The start maintenance mode command was issued successfully. Use the "GET" command on the /{deployment_id} resource to monitor progress
*/
type StartApmMaintenanceModeAccepted struct {
	Payload *models.ClusterCommandResponse
}

func (o *StartApmMaintenanceModeAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/instances/{instance_ids}/maintenance-mode/_start][%d] startApmMaintenanceModeAccepted  %+v", 202, o.Payload)
}
func (o *StartApmMaintenanceModeAccepted) GetPayload() *models.ClusterCommandResponse {
	return o.Payload
}

func (o *StartApmMaintenanceModeAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartApmMaintenanceModeForbidden creates a StartApmMaintenanceModeForbidden with default headers values
func NewStartApmMaintenanceModeForbidden() *StartApmMaintenanceModeForbidden {
	return &StartApmMaintenanceModeForbidden{}
}

/* StartApmMaintenanceModeForbidden describes a response with status code 403, with default header values.

The start maintenance mode command was prohibited for the given cluster. (code: `clusters.command_prohibited`)
*/
type StartApmMaintenanceModeForbidden struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartApmMaintenanceModeForbidden) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/instances/{instance_ids}/maintenance-mode/_start][%d] startApmMaintenanceModeForbidden  %+v", 403, o.Payload)
}
func (o *StartApmMaintenanceModeForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartApmMaintenanceModeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStartApmMaintenanceModeNotFound creates a StartApmMaintenanceModeNotFound with default headers values
func NewStartApmMaintenanceModeNotFound() *StartApmMaintenanceModeNotFound {
	return &StartApmMaintenanceModeNotFound{}
}

/* StartApmMaintenanceModeNotFound describes a response with status code 404, with default header values.

 * The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
* One or more of the instances specified at {instance_ids} could not be found. (code: `clusters.instances_not_found`)
*/
type StartApmMaintenanceModeNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartApmMaintenanceModeNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/instances/{instance_ids}/maintenance-mode/_start][%d] startApmMaintenanceModeNotFound  %+v", 404, o.Payload)
}
func (o *StartApmMaintenanceModeNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartApmMaintenanceModeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStartApmMaintenanceModeRetryWith creates a StartApmMaintenanceModeRetryWith with default headers values
func NewStartApmMaintenanceModeRetryWith() *StartApmMaintenanceModeRetryWith {
	return &StartApmMaintenanceModeRetryWith{}
}

/* StartApmMaintenanceModeRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type StartApmMaintenanceModeRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartApmMaintenanceModeRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/instances/{instance_ids}/maintenance-mode/_start][%d] startApmMaintenanceModeRetryWith  %+v", 449, o.Payload)
}
func (o *StartApmMaintenanceModeRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartApmMaintenanceModeRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
