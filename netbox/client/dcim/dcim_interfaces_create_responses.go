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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/intercloud/go-netbox/netbox/models"
)

// DcimInterfacesCreateReader is a Reader for the DcimInterfacesCreate structure.
type DcimInterfacesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfacesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimInterfacesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInterfacesCreateCreated creates a DcimInterfacesCreateCreated with default headers values
func NewDcimInterfacesCreateCreated() *DcimInterfacesCreateCreated {
	return &DcimInterfacesCreateCreated{}
}

/*
DcimInterfacesCreateCreated describes a response with status code 201, with default header values.

DcimInterfacesCreateCreated dcim interfaces create created
*/
type DcimInterfacesCreateCreated struct {
	Payload *models.Interface
}

// IsSuccess returns true when this dcim interfaces create created response has a 2xx status code
func (o *DcimInterfacesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim interfaces create created response has a 3xx status code
func (o *DcimInterfacesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim interfaces create created response has a 4xx status code
func (o *DcimInterfacesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim interfaces create created response has a 5xx status code
func (o *DcimInterfacesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim interfaces create created response a status code equal to that given
func (o *DcimInterfacesCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *DcimInterfacesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/interfaces/][%d] dcimInterfacesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimInterfacesCreateCreated) String() string {
	return fmt.Sprintf("[POST /dcim/interfaces/][%d] dcimInterfacesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimInterfacesCreateCreated) GetPayload() *models.Interface {
	return o.Payload
}

func (o *DcimInterfacesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Interface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
