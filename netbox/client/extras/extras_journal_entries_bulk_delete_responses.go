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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasJournalEntriesBulkDeleteReader is a Reader for the ExtrasJournalEntriesBulkDelete structure.
type ExtrasJournalEntriesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasJournalEntriesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasJournalEntriesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasJournalEntriesBulkDeleteNoContent creates a ExtrasJournalEntriesBulkDeleteNoContent with default headers values
func NewExtrasJournalEntriesBulkDeleteNoContent() *ExtrasJournalEntriesBulkDeleteNoContent {
	return &ExtrasJournalEntriesBulkDeleteNoContent{}
}

/*
ExtrasJournalEntriesBulkDeleteNoContent describes a response with status code 204, with default header values.

ExtrasJournalEntriesBulkDeleteNoContent extras journal entries bulk delete no content
*/
type ExtrasJournalEntriesBulkDeleteNoContent struct {
}

// IsSuccess returns true when this extras journal entries bulk delete no content response has a 2xx status code
func (o *ExtrasJournalEntriesBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras journal entries bulk delete no content response has a 3xx status code
func (o *ExtrasJournalEntriesBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras journal entries bulk delete no content response has a 4xx status code
func (o *ExtrasJournalEntriesBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras journal entries bulk delete no content response has a 5xx status code
func (o *ExtrasJournalEntriesBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this extras journal entries bulk delete no content response a status code equal to that given
func (o *ExtrasJournalEntriesBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ExtrasJournalEntriesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/journal-entries/][%d] extrasJournalEntriesBulkDeleteNoContent ", 204)
}

func (o *ExtrasJournalEntriesBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /extras/journal-entries/][%d] extrasJournalEntriesBulkDeleteNoContent ", 204)
}

func (o *ExtrasJournalEntriesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
