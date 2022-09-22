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

// DcimPowerOutletsReadReader is a Reader for the DcimPowerOutletsRead structure.
type DcimPowerOutletsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerOutletsReadOK creates a DcimPowerOutletsReadOK with default headers values
func NewDcimPowerOutletsReadOK() *DcimPowerOutletsReadOK {
	return &DcimPowerOutletsReadOK{}
}

/*
DcimPowerOutletsReadOK describes a response with status code 200, with default header values.

DcimPowerOutletsReadOK dcim power outlets read o k
*/
type DcimPowerOutletsReadOK struct {
	Payload *models.PowerOutlet
}

// IsSuccess returns true when this dcim power outlets read o k response has a 2xx status code
func (o *DcimPowerOutletsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power outlets read o k response has a 3xx status code
func (o *DcimPowerOutletsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power outlets read o k response has a 4xx status code
func (o *DcimPowerOutletsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power outlets read o k response has a 5xx status code
func (o *DcimPowerOutletsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power outlets read o k response a status code equal to that given
func (o *DcimPowerOutletsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimPowerOutletsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/power-outlets/{id}/][%d] dcimPowerOutletsReadOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletsReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/power-outlets/{id}/][%d] dcimPowerOutletsReadOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletsReadOK) GetPayload() *models.PowerOutlet {
	return o.Payload
}

func (o *DcimPowerOutletsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutlet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
