// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package job_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewEnableJobParams creates a new EnableJobParams object
// with the default values initialized.
func NewEnableJobParams() *EnableJobParams {
	var ()
	return &EnableJobParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEnableJobParamsWithTimeout creates a new EnableJobParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEnableJobParamsWithTimeout(timeout time.Duration) *EnableJobParams {
	var ()
	return &EnableJobParams{

		timeout: timeout,
	}
}

// NewEnableJobParamsWithContext creates a new EnableJobParams object
// with the default values initialized, and the ability to set a context for a request
func NewEnableJobParamsWithContext(ctx context.Context) *EnableJobParams {
	var ()
	return &EnableJobParams{

		Context: ctx,
	}
}

// NewEnableJobParamsWithHTTPClient creates a new EnableJobParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEnableJobParamsWithHTTPClient(client *http.Client) *EnableJobParams {
	var ()
	return &EnableJobParams{
		HTTPClient: client,
	}
}

/*EnableJobParams contains all the parameters to send to the API endpoint
for the enable job operation typically these are written to a http.Request
*/
type EnableJobParams struct {

	/*ID
	  The ID of the job to be enabled

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the enable job params
func (o *EnableJobParams) WithTimeout(timeout time.Duration) *EnableJobParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enable job params
func (o *EnableJobParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enable job params
func (o *EnableJobParams) WithContext(ctx context.Context) *EnableJobParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enable job params
func (o *EnableJobParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enable job params
func (o *EnableJobParams) WithHTTPClient(client *http.Client) *EnableJobParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enable job params
func (o *EnableJobParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the enable job params
func (o *EnableJobParams) WithID(id string) *EnableJobParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the enable job params
func (o *EnableJobParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EnableJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}