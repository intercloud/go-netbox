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

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WirelessWirelessLansDeleteReader is a Reader for the WirelessWirelessLansDelete structure.
type WirelessWirelessLansDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLansDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewWirelessWirelessLansDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWirelessWirelessLansDeleteNoContent creates a WirelessWirelessLansDeleteNoContent with default headers values
func NewWirelessWirelessLansDeleteNoContent() *WirelessWirelessLansDeleteNoContent {
	return &WirelessWirelessLansDeleteNoContent{}
}

/*
WirelessWirelessLansDeleteNoContent describes a response with status code 204, with default header values.

WirelessWirelessLansDeleteNoContent wireless wireless lans delete no content
*/
type WirelessWirelessLansDeleteNoContent struct {
}

// IsSuccess returns true when this wireless wireless lans delete no content response has a 2xx status code
func (o *WirelessWirelessLansDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless lans delete no content response has a 3xx status code
func (o *WirelessWirelessLansDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless lans delete no content response has a 4xx status code
func (o *WirelessWirelessLansDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless lans delete no content response has a 5xx status code
func (o *WirelessWirelessLansDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless lans delete no content response a status code equal to that given
func (o *WirelessWirelessLansDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *WirelessWirelessLansDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /wireless/wireless-lans/{id}/][%d] wirelessWirelessLansDeleteNoContent ", 204)
}

func (o *WirelessWirelessLansDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /wireless/wireless-lans/{id}/][%d] wirelessWirelessLansDeleteNoContent ", 204)
}

func (o *WirelessWirelessLansDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
