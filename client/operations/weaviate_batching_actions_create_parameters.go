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

package operations

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

// NewWeaviateBatchingActionsCreateParams creates a new WeaviateBatchingActionsCreateParams object
// with the default values initialized.
func NewWeaviateBatchingActionsCreateParams() *WeaviateBatchingActionsCreateParams {
	var ()
	return &WeaviateBatchingActionsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateBatchingActionsCreateParamsWithTimeout creates a new WeaviateBatchingActionsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateBatchingActionsCreateParamsWithTimeout(timeout time.Duration) *WeaviateBatchingActionsCreateParams {
	var ()
	return &WeaviateBatchingActionsCreateParams{

		timeout: timeout,
	}
}

// NewWeaviateBatchingActionsCreateParamsWithContext creates a new WeaviateBatchingActionsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateBatchingActionsCreateParamsWithContext(ctx context.Context) *WeaviateBatchingActionsCreateParams {
	var ()
	return &WeaviateBatchingActionsCreateParams{

		Context: ctx,
	}
}

// NewWeaviateBatchingActionsCreateParamsWithHTTPClient creates a new WeaviateBatchingActionsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateBatchingActionsCreateParamsWithHTTPClient(client *http.Client) *WeaviateBatchingActionsCreateParams {
	var ()
	return &WeaviateBatchingActionsCreateParams{
		HTTPClient: client,
	}
}

/*WeaviateBatchingActionsCreateParams contains all the parameters to send to the API endpoint
for the weaviate batching actions create operation typically these are written to a http.Request
*/
type WeaviateBatchingActionsCreateParams struct {

	/*Body*/
	Body WeaviateBatchingActionsCreateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate batching actions create params
func (o *WeaviateBatchingActionsCreateParams) WithTimeout(timeout time.Duration) *WeaviateBatchingActionsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate batching actions create params
func (o *WeaviateBatchingActionsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate batching actions create params
func (o *WeaviateBatchingActionsCreateParams) WithContext(ctx context.Context) *WeaviateBatchingActionsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate batching actions create params
func (o *WeaviateBatchingActionsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate batching actions create params
func (o *WeaviateBatchingActionsCreateParams) WithHTTPClient(client *http.Client) *WeaviateBatchingActionsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate batching actions create params
func (o *WeaviateBatchingActionsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the weaviate batching actions create params
func (o *WeaviateBatchingActionsCreateParams) WithBody(body WeaviateBatchingActionsCreateBody) *WeaviateBatchingActionsCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the weaviate batching actions create params
func (o *WeaviateBatchingActionsCreateParams) SetBody(body WeaviateBatchingActionsCreateBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateBatchingActionsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
