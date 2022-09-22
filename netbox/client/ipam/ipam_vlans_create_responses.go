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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/intercloud/go-netbox/netbox/models"
)

// IpamVlansCreateReader is a Reader for the IpamVlansCreate structure.
type IpamVlansCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlansCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamVlansCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVlansCreateCreated creates a IpamVlansCreateCreated with default headers values
func NewIpamVlansCreateCreated() *IpamVlansCreateCreated {
	return &IpamVlansCreateCreated{}
}

/*
IpamVlansCreateCreated describes a response with status code 201, with default header values.

IpamVlansCreateCreated ipam vlans create created
*/
type IpamVlansCreateCreated struct {
	Payload *models.VLAN
}

// IsSuccess returns true when this ipam vlans create created response has a 2xx status code
func (o *IpamVlansCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vlans create created response has a 3xx status code
func (o *IpamVlansCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vlans create created response has a 4xx status code
func (o *IpamVlansCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vlans create created response has a 5xx status code
func (o *IpamVlansCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vlans create created response a status code equal to that given
func (o *IpamVlansCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *IpamVlansCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/vlans/][%d] ipamVlansCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamVlansCreateCreated) String() string {
	return fmt.Sprintf("[POST /ipam/vlans/][%d] ipamVlansCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamVlansCreateCreated) GetPayload() *models.VLAN {
	return o.Payload
}

func (o *IpamVlansCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLAN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
