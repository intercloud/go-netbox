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

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamFhrpGroupsBulkDeleteReader is a Reader for the IpamFhrpGroupsBulkDelete structure.
type IpamFhrpGroupsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamFhrpGroupsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamFhrpGroupsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamFhrpGroupsBulkDeleteNoContent creates a IpamFhrpGroupsBulkDeleteNoContent with default headers values
func NewIpamFhrpGroupsBulkDeleteNoContent() *IpamFhrpGroupsBulkDeleteNoContent {
	return &IpamFhrpGroupsBulkDeleteNoContent{}
}

/*
IpamFhrpGroupsBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamFhrpGroupsBulkDeleteNoContent ipam fhrp groups bulk delete no content
*/
type IpamFhrpGroupsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this ipam fhrp groups bulk delete no content response has a 2xx status code
func (o *IpamFhrpGroupsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam fhrp groups bulk delete no content response has a 3xx status code
func (o *IpamFhrpGroupsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam fhrp groups bulk delete no content response has a 4xx status code
func (o *IpamFhrpGroupsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam fhrp groups bulk delete no content response has a 5xx status code
func (o *IpamFhrpGroupsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam fhrp groups bulk delete no content response a status code equal to that given
func (o *IpamFhrpGroupsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *IpamFhrpGroupsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/fhrp-groups/][%d] ipamFhrpGroupsBulkDeleteNoContent ", 204)
}

func (o *IpamFhrpGroupsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ipam/fhrp-groups/][%d] ipamFhrpGroupsBulkDeleteNoContent ", 204)
}

func (o *IpamFhrpGroupsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
