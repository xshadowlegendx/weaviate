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

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// NewWeaviateActionUpdateParams creates a new WeaviateActionUpdateParams object
// no default values defined in spec.
func NewWeaviateActionUpdateParams() WeaviateActionUpdateParams {

	return WeaviateActionUpdateParams{}
}

// WeaviateActionUpdateParams contains all the bound params for the weaviate action update operation
// typically these are obtained from a http.Request
//
// swagger:parameters weaviate.action.update
type WeaviateActionUpdateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Unique ID of the Action.
	  Required: true
	  In: path
	*/
	ActionID strfmt.UUID
	/*
	  Required: true
	  In: body
	*/
	Body *models.ActionUpdate
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewWeaviateActionUpdateParams() beforehand.
func (o *WeaviateActionUpdateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rActionID, rhkActionID, _ := route.Params.GetOK("actionId")
	if err := o.bindActionID(rActionID, rhkActionID, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ActionUpdate
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("body", "body"))
			} else {
				res = append(res, errors.NewParseError("body", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	} else {
		res = append(res, errors.Required("body", "body"))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindActionID binds and validates parameter ActionID from path.
func (o *WeaviateActionUpdateParams) bindActionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("actionId", "path", "strfmt.UUID", raw)
	}
	o.ActionID = *(value.(*strfmt.UUID))

	if err := o.validateActionID(formats); err != nil {
		return err
	}

	return nil
}

// validateActionID carries on validations for parameter ActionID
func (o *WeaviateActionUpdateParams) validateActionID(formats strfmt.Registry) error {

	if err := validate.FormatOf("actionId", "path", "uuid", o.ActionID.String(), formats); err != nil {
		return err
	}
	return nil
}
