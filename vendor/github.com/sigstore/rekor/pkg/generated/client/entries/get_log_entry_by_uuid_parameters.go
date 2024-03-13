// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright 2021 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package entries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetLogEntryByUUIDParams creates a new GetLogEntryByUUIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLogEntryByUUIDParams() *GetLogEntryByUUIDParams {
	return &GetLogEntryByUUIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLogEntryByUUIDParamsWithTimeout creates a new GetLogEntryByUUIDParams object
// with the ability to set a timeout on a request.
func NewGetLogEntryByUUIDParamsWithTimeout(timeout time.Duration) *GetLogEntryByUUIDParams {
	return &GetLogEntryByUUIDParams{
		timeout: timeout,
	}
}

// NewGetLogEntryByUUIDParamsWithContext creates a new GetLogEntryByUUIDParams object
// with the ability to set a context for a request.
func NewGetLogEntryByUUIDParamsWithContext(ctx context.Context) *GetLogEntryByUUIDParams {
	return &GetLogEntryByUUIDParams{
		Context: ctx,
	}
}

// NewGetLogEntryByUUIDParamsWithHTTPClient creates a new GetLogEntryByUUIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLogEntryByUUIDParamsWithHTTPClient(client *http.Client) *GetLogEntryByUUIDParams {
	return &GetLogEntryByUUIDParams{
		HTTPClient: client,
	}
}

/*
GetLogEntryByUUIDParams contains all the parameters to send to the API endpoint

	for the get log entry by UUID operation.

	Typically these are written to a http.Request.
*/
type GetLogEntryByUUIDParams struct {

	/* EntryUUID.

	   the UUID of the entry for which the inclusion proof information should be returned
	*/
	EntryUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get log entry by UUID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogEntryByUUIDParams) WithDefaults() *GetLogEntryByUUIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get log entry by UUID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogEntryByUUIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get log entry by UUID params
func (o *GetLogEntryByUUIDParams) WithTimeout(timeout time.Duration) *GetLogEntryByUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get log entry by UUID params
func (o *GetLogEntryByUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get log entry by UUID params
func (o *GetLogEntryByUUIDParams) WithContext(ctx context.Context) *GetLogEntryByUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get log entry by UUID params
func (o *GetLogEntryByUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get log entry by UUID params
func (o *GetLogEntryByUUIDParams) WithHTTPClient(client *http.Client) *GetLogEntryByUUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get log entry by UUID params
func (o *GetLogEntryByUUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEntryUUID adds the entryUUID to the get log entry by UUID params
func (o *GetLogEntryByUUIDParams) WithEntryUUID(entryUUID string) *GetLogEntryByUUIDParams {
	o.SetEntryUUID(entryUUID)
	return o
}

// SetEntryUUID adds the entryUuid to the get log entry by UUID params
func (o *GetLogEntryByUUIDParams) SetEntryUUID(entryUUID string) {
	o.EntryUUID = entryUUID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLogEntryByUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param entryUUID
	if err := r.SetPathParam("entryUUID", o.EntryUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
