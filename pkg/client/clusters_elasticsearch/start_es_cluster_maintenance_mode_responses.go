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

// StartEsClusterMaintenanceModeReader is a Reader for the StartEsClusterMaintenanceMode structure.
type StartEsClusterMaintenanceModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartEsClusterMaintenanceModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewStartEsClusterMaintenanceModeAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewStartEsClusterMaintenanceModeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStartEsClusterMaintenanceModeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 449:
		result := NewStartEsClusterMaintenanceModeRetryWith()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewStartEsClusterMaintenanceModeAccepted creates a StartEsClusterMaintenanceModeAccepted with default headers values
func NewStartEsClusterMaintenanceModeAccepted() *StartEsClusterMaintenanceModeAccepted {
	return &StartEsClusterMaintenanceModeAccepted{}
}

/* StartEsClusterMaintenanceModeAccepted describes a response with status code 202, with default header values.

The start maintenance mode command was issued successfully. Use the "GET" command on the /{deployment_id} resource to monitor progress
*/
type StartEsClusterMaintenanceModeAccepted struct {
	Payload *models.ClusterCommandResponse
}

func (o *StartEsClusterMaintenanceModeAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/instances/{instance_ids}/maintenance-mode/_start][%d] startEsClusterMaintenanceModeAccepted  %+v", 202, o.Payload)
}
func (o *StartEsClusterMaintenanceModeAccepted) GetPayload() *models.ClusterCommandResponse {
	return o.Payload
}

func (o *StartEsClusterMaintenanceModeAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterCommandResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartEsClusterMaintenanceModeForbidden creates a StartEsClusterMaintenanceModeForbidden with default headers values
func NewStartEsClusterMaintenanceModeForbidden() *StartEsClusterMaintenanceModeForbidden {
	return &StartEsClusterMaintenanceModeForbidden{}
}

/* StartEsClusterMaintenanceModeForbidden describes a response with status code 403, with default header values.

The start maintenance mode command was prohibited for the given cluster. (code: `clusters.command_prohibited`)
*/
type StartEsClusterMaintenanceModeForbidden struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartEsClusterMaintenanceModeForbidden) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/instances/{instance_ids}/maintenance-mode/_start][%d] startEsClusterMaintenanceModeForbidden  %+v", 403, o.Payload)
}
func (o *StartEsClusterMaintenanceModeForbidden) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartEsClusterMaintenanceModeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStartEsClusterMaintenanceModeNotFound creates a StartEsClusterMaintenanceModeNotFound with default headers values
func NewStartEsClusterMaintenanceModeNotFound() *StartEsClusterMaintenanceModeNotFound {
	return &StartEsClusterMaintenanceModeNotFound{}
}

/* StartEsClusterMaintenanceModeNotFound describes a response with status code 404, with default header values.

 * The cluster specified by {cluster_id} cannot be found. (code: `clusters.cluster_not_found`)
* One or more of the instances specified at {instance_ids} could not be found. (code: `clusters.instances_not_found`)
*/
type StartEsClusterMaintenanceModeNotFound struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartEsClusterMaintenanceModeNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/instances/{instance_ids}/maintenance-mode/_start][%d] startEsClusterMaintenanceModeNotFound  %+v", 404, o.Payload)
}
func (o *StartEsClusterMaintenanceModeNotFound) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartEsClusterMaintenanceModeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewStartEsClusterMaintenanceModeRetryWith creates a StartEsClusterMaintenanceModeRetryWith with default headers values
func NewStartEsClusterMaintenanceModeRetryWith() *StartEsClusterMaintenanceModeRetryWith {
	return &StartEsClusterMaintenanceModeRetryWith{}
}

/* StartEsClusterMaintenanceModeRetryWith describes a response with status code 449, with default header values.

Elevated permissions are required. (code: `root.unauthorized.rbac.elevated_permissions_required`)
*/
type StartEsClusterMaintenanceModeRetryWith struct {

	/* The error codes associated with the response
	 */
	XCloudErrorCodes string

	Payload *models.BasicFailedReply
}

func (o *StartEsClusterMaintenanceModeRetryWith) Error() string {
	return fmt.Sprintf("[POST /clusters/elasticsearch/{cluster_id}/instances/{instance_ids}/maintenance-mode/_start][%d] startEsClusterMaintenanceModeRetryWith  %+v", 449, o.Payload)
}
func (o *StartEsClusterMaintenanceModeRetryWith) GetPayload() *models.BasicFailedReply {
	return o.Payload
}

func (o *StartEsClusterMaintenanceModeRetryWith) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
