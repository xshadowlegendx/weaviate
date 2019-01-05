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
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewWeaviateSchemaActionsUpdateParams creates a new WeaviateSchemaActionsUpdateParams object
// with the default values initialized.
func NewWeaviateSchemaActionsUpdateParams() *WeaviateSchemaActionsUpdateParams {
	var ()
	return &WeaviateSchemaActionsUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateSchemaActionsUpdateParamsWithTimeout creates a new WeaviateSchemaActionsUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateSchemaActionsUpdateParamsWithTimeout(timeout time.Duration) *WeaviateSchemaActionsUpdateParams {
	var ()
	return &WeaviateSchemaActionsUpdateParams{

		timeout: timeout,
	}
}

// NewWeaviateSchemaActionsUpdateParamsWithContext creates a new WeaviateSchemaActionsUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateSchemaActionsUpdateParamsWithContext(ctx context.Context) *WeaviateSchemaActionsUpdateParams {
	var ()
	return &WeaviateSchemaActionsUpdateParams{

		Context: ctx,
	}
}

// NewWeaviateSchemaActionsUpdateParamsWithHTTPClient creates a new WeaviateSchemaActionsUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateSchemaActionsUpdateParamsWithHTTPClient(client *http.Client) *WeaviateSchemaActionsUpdateParams {
	var ()
	return &WeaviateSchemaActionsUpdateParams{
		HTTPClient: client,
	}
}

/*WeaviateSchemaActionsUpdateParams contains all the parameters to send to the API endpoint
for the weaviate schema actions update operation typically these are written to a http.Request
*/
type WeaviateSchemaActionsUpdateParams struct {

	/*Body*/
	Body WeaviateSchemaActionsUpdateBody
	/*ClassName*/
	ClassName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate schema actions update params
func (o *WeaviateSchemaActionsUpdateParams) WithTimeout(timeout time.Duration) *WeaviateSchemaActionsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate schema actions update params
func (o *WeaviateSchemaActionsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate schema actions update params
func (o *WeaviateSchemaActionsUpdateParams) WithContext(ctx context.Context) *WeaviateSchemaActionsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate schema actions update params
func (o *WeaviateSchemaActionsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate schema actions update params
func (o *WeaviateSchemaActionsUpdateParams) WithHTTPClient(client *http.Client) *WeaviateSchemaActionsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate schema actions update params
func (o *WeaviateSchemaActionsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the weaviate schema actions update params
func (o *WeaviateSchemaActionsUpdateParams) WithBody(body WeaviateSchemaActionsUpdateBody) *WeaviateSchemaActionsUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the weaviate schema actions update params
func (o *WeaviateSchemaActionsUpdateParams) SetBody(body WeaviateSchemaActionsUpdateBody) {
	o.Body = body
}

// WithClassName adds the className to the weaviate schema actions update params
func (o *WeaviateSchemaActionsUpdateParams) WithClassName(className string) *WeaviateSchemaActionsUpdateParams {
	o.SetClassName(className)
	return o
}

// SetClassName adds the className to the weaviate schema actions update params
func (o *WeaviateSchemaActionsUpdateParams) SetClassName(className string) {
	o.ClassName = className
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateSchemaActionsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param className
	if err := r.SetPathParam("className", o.ClassName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
