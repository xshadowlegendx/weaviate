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

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// WeaviateSchemaThingsPropertiesDeleteReader is a Reader for the WeaviateSchemaThingsPropertiesDelete structure.
type WeaviateSchemaThingsPropertiesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateSchemaThingsPropertiesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateSchemaThingsPropertiesDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateSchemaThingsPropertiesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateSchemaThingsPropertiesDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateSchemaThingsPropertiesDeleteOK creates a WeaviateSchemaThingsPropertiesDeleteOK with default headers values
func NewWeaviateSchemaThingsPropertiesDeleteOK() *WeaviateSchemaThingsPropertiesDeleteOK {
	return &WeaviateSchemaThingsPropertiesDeleteOK{}
}

/*WeaviateSchemaThingsPropertiesDeleteOK handles this case with default header values.

Removed the property from the ontology.
*/
type WeaviateSchemaThingsPropertiesDeleteOK struct {
}

func (o *WeaviateSchemaThingsPropertiesDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /schema/things/{className}/properties/{propertyName}][%d] weaviateSchemaThingsPropertiesDeleteOK ", 200)
}

func (o *WeaviateSchemaThingsPropertiesDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaThingsPropertiesDeleteUnauthorized creates a WeaviateSchemaThingsPropertiesDeleteUnauthorized with default headers values
func NewWeaviateSchemaThingsPropertiesDeleteUnauthorized() *WeaviateSchemaThingsPropertiesDeleteUnauthorized {
	return &WeaviateSchemaThingsPropertiesDeleteUnauthorized{}
}

/*WeaviateSchemaThingsPropertiesDeleteUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateSchemaThingsPropertiesDeleteUnauthorized struct {
}

func (o *WeaviateSchemaThingsPropertiesDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /schema/things/{className}/properties/{propertyName}][%d] weaviateSchemaThingsPropertiesDeleteUnauthorized ", 401)
}

func (o *WeaviateSchemaThingsPropertiesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaThingsPropertiesDeleteForbidden creates a WeaviateSchemaThingsPropertiesDeleteForbidden with default headers values
func NewWeaviateSchemaThingsPropertiesDeleteForbidden() *WeaviateSchemaThingsPropertiesDeleteForbidden {
	return &WeaviateSchemaThingsPropertiesDeleteForbidden{}
}

/*WeaviateSchemaThingsPropertiesDeleteForbidden handles this case with default header values.

Could not find the Thing class or property.
*/
type WeaviateSchemaThingsPropertiesDeleteForbidden struct {
}

func (o *WeaviateSchemaThingsPropertiesDeleteForbidden) Error() string {
	return fmt.Sprintf("[DELETE /schema/things/{className}/properties/{propertyName}][%d] weaviateSchemaThingsPropertiesDeleteForbidden ", 403)
}

func (o *WeaviateSchemaThingsPropertiesDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
