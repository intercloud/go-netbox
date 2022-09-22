// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/intercloud/go-netbox/netbox/models"
)

// ExtrasExportTemplatesBulkPartialUpdateReader is a Reader for the ExtrasExportTemplatesBulkPartialUpdate structure.
type ExtrasExportTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasExportTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasExportTemplatesBulkPartialUpdateOK creates a ExtrasExportTemplatesBulkPartialUpdateOK with default headers values
func NewExtrasExportTemplatesBulkPartialUpdateOK() *ExtrasExportTemplatesBulkPartialUpdateOK {
	return &ExtrasExportTemplatesBulkPartialUpdateOK{}
}

/*
ExtrasExportTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasExportTemplatesBulkPartialUpdateOK extras export templates bulk partial update o k
*/
type ExtrasExportTemplatesBulkPartialUpdateOK struct {
	Payload *models.ExportTemplate
}

// IsSuccess returns true when this extras export templates bulk partial update o k response has a 2xx status code
func (o *ExtrasExportTemplatesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras export templates bulk partial update o k response has a 3xx status code
func (o *ExtrasExportTemplatesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras export templates bulk partial update o k response has a 4xx status code
func (o *ExtrasExportTemplatesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras export templates bulk partial update o k response has a 5xx status code
func (o *ExtrasExportTemplatesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras export templates bulk partial update o k response a status code equal to that given
func (o *ExtrasExportTemplatesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasExportTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/export-templates/][%d] extrasExportTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasExportTemplatesBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /extras/export-templates/][%d] extrasExportTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasExportTemplatesBulkPartialUpdateOK) GetPayload() *models.ExportTemplate {
	return o.Payload
}

func (o *ExtrasExportTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
