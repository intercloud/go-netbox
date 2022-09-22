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
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/intercloud/go-netbox/netbox/models"
)

// DcimRegionsListReader is a Reader for the DcimRegionsList structure.
type DcimRegionsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRegionsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRegionsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRegionsListOK creates a DcimRegionsListOK with default headers values
func NewDcimRegionsListOK() *DcimRegionsListOK {
	return &DcimRegionsListOK{}
}

/*
DcimRegionsListOK describes a response with status code 200, with default header values.

DcimRegionsListOK dcim regions list o k
*/
type DcimRegionsListOK struct {
	Payload *DcimRegionsListOKBody
}

// IsSuccess returns true when this dcim regions list o k response has a 2xx status code
func (o *DcimRegionsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim regions list o k response has a 3xx status code
func (o *DcimRegionsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim regions list o k response has a 4xx status code
func (o *DcimRegionsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim regions list o k response has a 5xx status code
func (o *DcimRegionsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim regions list o k response a status code equal to that given
func (o *DcimRegionsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimRegionsListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/regions/][%d] dcimRegionsListOK  %+v", 200, o.Payload)
}

func (o *DcimRegionsListOK) String() string {
	return fmt.Sprintf("[GET /dcim/regions/][%d] dcimRegionsListOK  %+v", 200, o.Payload)
}

func (o *DcimRegionsListOK) GetPayload() *DcimRegionsListOKBody {
	return o.Payload
}

func (o *DcimRegionsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimRegionsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DcimRegionsListOKBody dcim regions list o k body
swagger:model DcimRegionsListOKBody
*/
type DcimRegionsListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.Region `json:"results"`
}

// Validate validates this dcim regions list o k body
func (o *DcimRegionsListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimRegionsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimRegionsListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimRegionsListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimRegionsListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimRegionsListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimRegionsListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimRegionsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimRegionsListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimRegionsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimRegionsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim regions list o k body based on the context it is used
func (o *DcimRegionsListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimRegionsListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimRegionsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimRegionsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimRegionsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimRegionsListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimRegionsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
