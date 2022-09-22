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

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimDeviceBaysBulkDeleteReader is a Reader for the DcimDeviceBaysBulkDelete structure.
type DcimDeviceBaysBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBaysBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDeviceBaysBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceBaysBulkDeleteNoContent creates a DcimDeviceBaysBulkDeleteNoContent with default headers values
func NewDcimDeviceBaysBulkDeleteNoContent() *DcimDeviceBaysBulkDeleteNoContent {
	return &DcimDeviceBaysBulkDeleteNoContent{}
}

/*
DcimDeviceBaysBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimDeviceBaysBulkDeleteNoContent dcim device bays bulk delete no content
*/
type DcimDeviceBaysBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim device bays bulk delete no content response has a 2xx status code
func (o *DcimDeviceBaysBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device bays bulk delete no content response has a 3xx status code
func (o *DcimDeviceBaysBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device bays bulk delete no content response has a 4xx status code
func (o *DcimDeviceBaysBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device bays bulk delete no content response has a 5xx status code
func (o *DcimDeviceBaysBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device bays bulk delete no content response a status code equal to that given
func (o *DcimDeviceBaysBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimDeviceBaysBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-bays/][%d] dcimDeviceBaysBulkDeleteNoContent ", 204)
}

func (o *DcimDeviceBaysBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/device-bays/][%d] dcimDeviceBaysBulkDeleteNoContent ", 204)
}

func (o *DcimDeviceBaysBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
