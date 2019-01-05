/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN: Bob van Luijt (bob@k10y.co)
 */
// Code generated by go-swagger; DO NOT EDIT.

package graphql

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// WeaviateGraphqlBatchReader is a Reader for the WeaviateGraphqlBatch structure.
type WeaviateGraphqlBatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateGraphqlBatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateGraphqlBatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateGraphqlBatchUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateGraphqlBatchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateGraphqlBatchUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateGraphqlBatchOK creates a WeaviateGraphqlBatchOK with default headers values
func NewWeaviateGraphqlBatchOK() *WeaviateGraphqlBatchOK {
	return &WeaviateGraphqlBatchOK{}
}

/*WeaviateGraphqlBatchOK handles this case with default header values.

Successful query (with select).
*/
type WeaviateGraphqlBatchOK struct {
	Payload models.GraphQLResponses
}

func (o *WeaviateGraphqlBatchOK) Error() string {
	return fmt.Sprintf("[POST /graphql/batch][%d] weaviateGraphqlBatchOK  %+v", 200, o.Payload)
}

func (o *WeaviateGraphqlBatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateGraphqlBatchUnauthorized creates a WeaviateGraphqlBatchUnauthorized with default headers values
func NewWeaviateGraphqlBatchUnauthorized() *WeaviateGraphqlBatchUnauthorized {
	return &WeaviateGraphqlBatchUnauthorized{}
}

/*WeaviateGraphqlBatchUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateGraphqlBatchUnauthorized struct {
}

func (o *WeaviateGraphqlBatchUnauthorized) Error() string {
	return fmt.Sprintf("[POST /graphql/batch][%d] weaviateGraphqlBatchUnauthorized ", 401)
}

func (o *WeaviateGraphqlBatchUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateGraphqlBatchForbidden creates a WeaviateGraphqlBatchForbidden with default headers values
func NewWeaviateGraphqlBatchForbidden() *WeaviateGraphqlBatchForbidden {
	return &WeaviateGraphqlBatchForbidden{}
}

/*WeaviateGraphqlBatchForbidden handles this case with default header values.

The used API-key has insufficient permissions.
*/
type WeaviateGraphqlBatchForbidden struct {
}

func (o *WeaviateGraphqlBatchForbidden) Error() string {
	return fmt.Sprintf("[POST /graphql/batch][%d] weaviateGraphqlBatchForbidden ", 403)
}

func (o *WeaviateGraphqlBatchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateGraphqlBatchUnprocessableEntity creates a WeaviateGraphqlBatchUnprocessableEntity with default headers values
func NewWeaviateGraphqlBatchUnprocessableEntity() *WeaviateGraphqlBatchUnprocessableEntity {
	return &WeaviateGraphqlBatchUnprocessableEntity{}
}

/*WeaviateGraphqlBatchUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the class is defined in the configuration file?
*/
type WeaviateGraphqlBatchUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateGraphqlBatchUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /graphql/batch][%d] weaviateGraphqlBatchUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateGraphqlBatchUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
