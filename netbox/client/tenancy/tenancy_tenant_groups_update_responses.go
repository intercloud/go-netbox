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

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/intercloud/go-netbox/netbox/models"
)

// TenancyTenantGroupsUpdateReader is a Reader for the TenancyTenantGroupsUpdate structure.
type TenancyTenantGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyTenantGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTenancyTenantGroupsUpdateOK creates a TenancyTenantGroupsUpdateOK with default headers values
func NewTenancyTenantGroupsUpdateOK() *TenancyTenantGroupsUpdateOK {
	return &TenancyTenantGroupsUpdateOK{}
}

/*
TenancyTenantGroupsUpdateOK describes a response with status code 200, with default header values.

TenancyTenantGroupsUpdateOK tenancy tenant groups update o k
*/
type TenancyTenantGroupsUpdateOK struct {
	Payload *models.TenantGroup
}

// IsSuccess returns true when this tenancy tenant groups update o k response has a 2xx status code
func (o *TenancyTenantGroupsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy tenant groups update o k response has a 3xx status code
func (o *TenancyTenantGroupsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy tenant groups update o k response has a 4xx status code
func (o *TenancyTenantGroupsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy tenant groups update o k response has a 5xx status code
func (o *TenancyTenantGroupsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy tenant groups update o k response a status code equal to that given
func (o *TenancyTenantGroupsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *TenancyTenantGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /tenancy/tenant-groups/{id}/][%d] tenancyTenantGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyTenantGroupsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /tenancy/tenant-groups/{id}/][%d] tenancyTenantGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyTenantGroupsUpdateOK) GetPayload() *models.TenantGroup {
	return o.Payload
}

func (o *TenancyTenantGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TenantGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
