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

// DcimVirtualChassisCreateReader is a Reader for the DcimVirtualChassisCreate structure.
type DcimVirtualChassisCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualChassisCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimVirtualChassisCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimVirtualChassisCreateCreated creates a DcimVirtualChassisCreateCreated with default headers values
func NewDcimVirtualChassisCreateCreated() *DcimVirtualChassisCreateCreated {
	return &DcimVirtualChassisCreateCreated{}
}

/*
DcimVirtualChassisCreateCreated describes a response with status code 201, with default header values.

DcimVirtualChassisCreateCreated dcim virtual chassis create created
*/
type DcimVirtualChassisCreateCreated struct {
	Payload *models.VirtualChassis
}

// IsSuccess returns true when this dcim virtual chassis create created response has a 2xx status code
func (o *DcimVirtualChassisCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim virtual chassis create created response has a 3xx status code
func (o *DcimVirtualChassisCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim virtual chassis create created response has a 4xx status code
func (o *DcimVirtualChassisCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim virtual chassis create created response has a 5xx status code
func (o *DcimVirtualChassisCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim virtual chassis create created response a status code equal to that given
func (o *DcimVirtualChassisCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *DcimVirtualChassisCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/virtual-chassis/][%d] dcimVirtualChassisCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimVirtualChassisCreateCreated) String() string {
	return fmt.Sprintf("[POST /dcim/virtual-chassis/][%d] dcimVirtualChassisCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimVirtualChassisCreateCreated) GetPayload() *models.VirtualChassis {
	return o.Payload
}

func (o *DcimVirtualChassisCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualChassis)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
