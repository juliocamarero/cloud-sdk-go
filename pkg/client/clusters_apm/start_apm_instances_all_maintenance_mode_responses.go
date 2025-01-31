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

// StartApmInstancesAllMaintenanceModeReader is a Reader for the StartApmInstancesAllMaintenanceMode structure.
type StartApmInstancesAllMaintenanceModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartApmInstancesAllMaintenanceModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewStartApmInstancesAllMaintenanceModeAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewStartApmInstancesAllMaintenanceModeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStartApmInstancesAllMaintenanceModeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewStartApmInstancesAllMaintenanceModeRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStartApmInstancesAllMaintenanceModeAccepted creates a StartApmInstancesAllMaintenanceModeAccepted with default headers values
func NewStartApmInstancesAllMaintenanceModeAccepted() *StartApmInstancesAllMaintenanceModeAccepted {
	return &StartApmInstancesAllMaintenanceModeAccepted{}
}

/* StartApmInstancesAllMaintenanceModeAccepted describes a response with status code 202, with default header values.

The start maintenance mode command was issued successfully. Use the "GET" command on the /{deployment_id} resource to monitor progress
*/
type StartApmInstancesAllMaintenanceModeAccepted struct {
	Payload *models.ClusterCommandResponse
}

func (o *StartApmInstancesAllMaintenanceModeAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/instances/maintenance-mode/_start][%d] startApmInstancesAllMaintenanceModeAccepted  %+v", 202, o.Payload)
}
func (o *StartApmInstancesAllMaintenanceModeAccepted) GetPayload() *models.ClusterCommandResponse {
	return o.Payload
}

func (o *StartApmInstancesAllMaintenanceModeAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartApmInstancesAllMaintenanceModeForbidden creates a StartApmInstancesAllMaintenanceModeForbidden with default headers values
func NewStartApmInstancesAllMaintenanceModeForbidden() *StartApmInstancesAllMaintenanceModeForbidden {
	return &StartApmInstancesAllMaintenanceModeForbidden{}
}

/* StartApmInstancesAllMaintenanceModeForbidden describes a response with status code 403, with default header values.

The start maintenance mode command was prohibited for the given cluster. (code: `clusters.command_prohibited`)
*/
type StartApmInstancesAllMaintenanceModeForbidden struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartApmInstancesAllMaintenanceModeForbidden) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/instances/maintenance-mode/_start][%d] startApmInstancesAllMaintenanceModeForbidden  %+v", 403, o.Payload)
}
func (o *StartApmInstancesAllMaintenanceModeForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartApmInstancesAllMaintenanceModeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStartApmInstancesAllMaintenanceModeNotFound creates a StartApmInstancesAllMaintenanceModeNotFound with default headers values
func NewStartApmInstancesAllMaintenanceModeNotFound() *StartApmInstancesAllMaintenanceModeNotFound {
	return &StartApmInstancesAllMaintenanceModeNotFound{}
}

/* StartApmInstancesAllMaintenanceModeNotFound describes a response with status code 404, with default header values.

The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
*/
type StartApmInstancesAllMaintenanceModeNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartApmInstancesAllMaintenanceModeNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/instances/maintenance-mode/_start][%d] startApmInstancesAllMaintenanceModeNotFound  %+v", 404, o.Payload)
}
func (o *StartApmInstancesAllMaintenanceModeNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartApmInstancesAllMaintenanceModeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStartApmInstancesAllMaintenanceModeRetryWith creates a StartApmInstancesAllMaintenanceModeRetryWith with default headers values
func NewStartApmInstancesAllMaintenanceModeRetryWith() *StartApmInstancesAllMaintenanceModeRetryWith {
	return &StartApmInstancesAllMaintenanceModeRetryWith{}
}

/* StartApmInstancesAllMaintenanceModeRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type StartApmInstancesAllMaintenanceModeRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartApmInstancesAllMaintenanceModeRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/apm/{cluster_id}/instances/maintenance-mode/_start][%d] startApmInstancesAllMaintenanceModeRetryWith  %+v", 449, o.Payload)
}
func (o *StartApmInstancesAllMaintenanceModeRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartApmInstancesAllMaintenanceModeRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
