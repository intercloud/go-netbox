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

// DcimPlatformsDeleteReader is a Reader for the DcimPlatformsDelete structure.
type DcimPlatformsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPlatformsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPlatformsDeleteNoContent creates a DcimPlatformsDeleteNoContent with default headers values
func NewDcimPlatformsDeleteNoContent() *DcimPlatformsDeleteNoContent {
	return &DcimPlatformsDeleteNoContent{}
}

/*
DcimPlatformsDeleteNoContent describes a response with status code 204, with default header values.

DcimPlatformsDeleteNoContent dcim platforms delete no content
*/
type DcimPlatformsDeleteNoContent struct {
}

// IsSuccess returns true when this dcim platforms delete no content response has a 2xx status code
func (o *DcimPlatformsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim platforms delete no content response has a 3xx status code
func (o *DcimPlatformsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim platforms delete no content response has a 4xx status code
func (o *DcimPlatformsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim platforms delete no content response has a 5xx status code
func (o *DcimPlatformsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim platforms delete no content response a status code equal to that given
func (o *DcimPlatformsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimPlatformsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/platforms/{id}/][%d] dcimPlatformsDeleteNoContent ", 204)
}

func (o *DcimPlatformsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/platforms/{id}/][%d] dcimPlatformsDeleteNoContent ", 204)
}

func (o *DcimPlatformsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
