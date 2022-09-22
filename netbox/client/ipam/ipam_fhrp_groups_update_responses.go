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

// IpamFhrpGroupsUpdateReader is a Reader for the IpamFhrpGroupsUpdate structure.
type IpamFhrpGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamFhrpGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamFhrpGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamFhrpGroupsUpdateOK creates a IpamFhrpGroupsUpdateOK with default headers values
func NewIpamFhrpGroupsUpdateOK() *IpamFhrpGroupsUpdateOK {
	return &IpamFhrpGroupsUpdateOK{}
}

/*
IpamFhrpGroupsUpdateOK describes a response with status code 200, with default header values.

IpamFhrpGroupsUpdateOK ipam fhrp groups update o k
*/
type IpamFhrpGroupsUpdateOK struct {
	Payload *models.FHRPGroup
}

// IsSuccess returns true when this ipam fhrp groups update o k response has a 2xx status code
func (o *IpamFhrpGroupsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam fhrp groups update o k response has a 3xx status code
func (o *IpamFhrpGroupsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam fhrp groups update o k response has a 4xx status code
func (o *IpamFhrpGroupsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam fhrp groups update o k response has a 5xx status code
func (o *IpamFhrpGroupsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam fhrp groups update o k response a status code equal to that given
func (o *IpamFhrpGroupsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamFhrpGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/fhrp-groups/{id}/][%d] ipamFhrpGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamFhrpGroupsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/fhrp-groups/{id}/][%d] ipamFhrpGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamFhrpGroupsUpdateOK) GetPayload() *models.FHRPGroup {
	return o.Payload
}

func (o *IpamFhrpGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FHRPGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
