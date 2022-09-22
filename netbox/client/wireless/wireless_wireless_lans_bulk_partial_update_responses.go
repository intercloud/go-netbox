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
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/intercloud/go-netbox/netbox/models"
)

// WirelessWirelessLansBulkPartialUpdateReader is a Reader for the WirelessWirelessLansBulkPartialUpdate structure.
type WirelessWirelessLansBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLansBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWirelessWirelessLansBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewWirelessWirelessLansBulkPartialUpdateOK creates a WirelessWirelessLansBulkPartialUpdateOK with default headers values
func NewWirelessWirelessLansBulkPartialUpdateOK() *WirelessWirelessLansBulkPartialUpdateOK {
	return &WirelessWirelessLansBulkPartialUpdateOK{}
}

/*
WirelessWirelessLansBulkPartialUpdateOK describes a response with status code 200, with default header values.

WirelessWirelessLansBulkPartialUpdateOK wireless wireless lans bulk partial update o k
*/
type WirelessWirelessLansBulkPartialUpdateOK struct {
	Payload *models.WirelessLAN
}

// IsSuccess returns true when this wireless wireless lans bulk partial update o k response has a 2xx status code
func (o *WirelessWirelessLansBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless lans bulk partial update o k response has a 3xx status code
func (o *WirelessWirelessLansBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless lans bulk partial update o k response has a 4xx status code
func (o *WirelessWirelessLansBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless lans bulk partial update o k response has a 5xx status code
func (o *WirelessWirelessLansBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless lans bulk partial update o k response a status code equal to that given
func (o *WirelessWirelessLansBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *WirelessWirelessLansBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /wireless/wireless-lans/][%d] wirelessWirelessLansBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *WirelessWirelessLansBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /wireless/wireless-lans/][%d] wirelessWirelessLansBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *WirelessWirelessLansBulkPartialUpdateOK) GetPayload() *models.WirelessLAN {
	return o.Payload
}

func (o *WirelessWirelessLansBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WirelessLAN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
