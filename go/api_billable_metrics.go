/*
Lago API documentation

Lago API allows your application to push customer information and metrics (events) from your application to the billing application.

API version: 0.52.0-beta
Contact: tech@getlago.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lagoapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// BillableMetricsAPIService BillableMetricsAPI service
type BillableMetricsAPIService service

type BillableMetricsAPICreateBillableMetricRequest struct {
	ctx context.Context
	ApiService *BillableMetricsAPIService
	billableMetricCreateInput *BillableMetricCreateInput
}

// Billable metric payload
func (r BillableMetricsAPICreateBillableMetricRequest) BillableMetricCreateInput(billableMetricCreateInput BillableMetricCreateInput) BillableMetricsAPICreateBillableMetricRequest {
	r.billableMetricCreateInput = &billableMetricCreateInput
	return r
}

func (r BillableMetricsAPICreateBillableMetricRequest) Execute() (*BillableMetric, *http.Response, error) {
	return r.ApiService.CreateBillableMetricExecute(r)
}

/*
CreateBillableMetric Create a billable metric

This endpoint creates a new billable metric representing a pricing component of your application.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BillableMetricsAPICreateBillableMetricRequest
*/
func (a *BillableMetricsAPIService) CreateBillableMetric(ctx context.Context) BillableMetricsAPICreateBillableMetricRequest {
	return BillableMetricsAPICreateBillableMetricRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BillableMetric
func (a *BillableMetricsAPIService) CreateBillableMetricExecute(r BillableMetricsAPICreateBillableMetricRequest) (*BillableMetric, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BillableMetric
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillableMetricsAPIService.CreateBillableMetric")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/billable_metrics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.billableMetricCreateInput == nil {
		return localVarReturnValue, nil, reportError("billableMetricCreateInput is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.billableMetricCreateInput
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiErrorBadRequest
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiErrorUnauthorized
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ApiErrorUnprocessableEntity
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type BillableMetricsAPIDestroyBillableMetricRequest struct {
	ctx context.Context
	ApiService *BillableMetricsAPIService
	code string
}

func (r BillableMetricsAPIDestroyBillableMetricRequest) Execute() (*BillableMetric, *http.Response, error) {
	return r.ApiService.DestroyBillableMetricExecute(r)
}

/*
DestroyBillableMetric Delete a billable metric

This endpoint deletes an existing billable metric representing a pricing component of your application.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param code Code of the existing billable metric.
 @return BillableMetricsAPIDestroyBillableMetricRequest
*/
func (a *BillableMetricsAPIService) DestroyBillableMetric(ctx context.Context, code string) BillableMetricsAPIDestroyBillableMetricRequest {
	return BillableMetricsAPIDestroyBillableMetricRequest{
		ApiService: a,
		ctx: ctx,
		code: code,
	}
}

// Execute executes the request
//  @return BillableMetric
func (a *BillableMetricsAPIService) DestroyBillableMetricExecute(r BillableMetricsAPIDestroyBillableMetricRequest) (*BillableMetric, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BillableMetric
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillableMetricsAPIService.DestroyBillableMetric")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/billable_metrics/{code}"
	localVarPath = strings.Replace(localVarPath, "{"+"code"+"}", url.PathEscape(parameterValueToString(r.code, "code")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiErrorUnauthorized
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiErrorNotFound
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type BillableMetricsAPIFindAllBillableMetricGroupsRequest struct {
	ctx context.Context
	ApiService *BillableMetricsAPIService
	code string
	page *int32
	perPage *int32
}

// Page number.
func (r BillableMetricsAPIFindAllBillableMetricGroupsRequest) Page(page int32) BillableMetricsAPIFindAllBillableMetricGroupsRequest {
	r.page = &page
	return r
}

// Number of records per page.
func (r BillableMetricsAPIFindAllBillableMetricGroupsRequest) PerPage(perPage int32) BillableMetricsAPIFindAllBillableMetricGroupsRequest {
	r.perPage = &perPage
	return r
}

func (r BillableMetricsAPIFindAllBillableMetricGroupsRequest) Execute() (*GroupsPaginated, *http.Response, error) {
	return r.ApiService.FindAllBillableMetricGroupsExecute(r)
}

/*
FindAllBillableMetricGroups Find a billable metric's groups

This endpoint retrieves all groups for a billable metric.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param code Code of the existing billable metric.
 @return BillableMetricsAPIFindAllBillableMetricGroupsRequest
*/
func (a *BillableMetricsAPIService) FindAllBillableMetricGroups(ctx context.Context, code string) BillableMetricsAPIFindAllBillableMetricGroupsRequest {
	return BillableMetricsAPIFindAllBillableMetricGroupsRequest{
		ApiService: a,
		ctx: ctx,
		code: code,
	}
}

// Execute executes the request
//  @return GroupsPaginated
func (a *BillableMetricsAPIService) FindAllBillableMetricGroupsExecute(r BillableMetricsAPIFindAllBillableMetricGroupsRequest) (*GroupsPaginated, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GroupsPaginated
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillableMetricsAPIService.FindAllBillableMetricGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/billable_metrics/{code}/groups"
	localVarPath = strings.Replace(localVarPath, "{"+"code"+"}", url.PathEscape(parameterValueToString(r.code, "code")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiErrorUnauthorized
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type BillableMetricsAPIFindAllBillableMetricsRequest struct {
	ctx context.Context
	ApiService *BillableMetricsAPIService
	page *int32
	perPage *int32
}

// Page number.
func (r BillableMetricsAPIFindAllBillableMetricsRequest) Page(page int32) BillableMetricsAPIFindAllBillableMetricsRequest {
	r.page = &page
	return r
}

// Number of records per page.
func (r BillableMetricsAPIFindAllBillableMetricsRequest) PerPage(perPage int32) BillableMetricsAPIFindAllBillableMetricsRequest {
	r.perPage = &perPage
	return r
}

func (r BillableMetricsAPIFindAllBillableMetricsRequest) Execute() (*BillableMetricsPaginated, *http.Response, error) {
	return r.ApiService.FindAllBillableMetricsExecute(r)
}

/*
FindAllBillableMetrics List all billable metrics

This endpoint retrieves all existing billable metrics that represent pricing components of your application.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BillableMetricsAPIFindAllBillableMetricsRequest
*/
func (a *BillableMetricsAPIService) FindAllBillableMetrics(ctx context.Context) BillableMetricsAPIFindAllBillableMetricsRequest {
	return BillableMetricsAPIFindAllBillableMetricsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BillableMetricsPaginated
func (a *BillableMetricsAPIService) FindAllBillableMetricsExecute(r BillableMetricsAPIFindAllBillableMetricsRequest) (*BillableMetricsPaginated, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BillableMetricsPaginated
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillableMetricsAPIService.FindAllBillableMetrics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/billable_metrics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiErrorUnauthorized
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type BillableMetricsAPIFindBillableMetricRequest struct {
	ctx context.Context
	ApiService *BillableMetricsAPIService
	code string
}

func (r BillableMetricsAPIFindBillableMetricRequest) Execute() (*BillableMetric, *http.Response, error) {
	return r.ApiService.FindBillableMetricExecute(r)
}

/*
FindBillableMetric Retrieve a billable metric

This endpoint retrieves an existing billable metric that represents a pricing component of your application. The billable metric is identified by its unique code.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param code Code of the existing billable metric.
 @return BillableMetricsAPIFindBillableMetricRequest
*/
func (a *BillableMetricsAPIService) FindBillableMetric(ctx context.Context, code string) BillableMetricsAPIFindBillableMetricRequest {
	return BillableMetricsAPIFindBillableMetricRequest{
		ApiService: a,
		ctx: ctx,
		code: code,
	}
}

// Execute executes the request
//  @return BillableMetric
func (a *BillableMetricsAPIService) FindBillableMetricExecute(r BillableMetricsAPIFindBillableMetricRequest) (*BillableMetric, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BillableMetric
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillableMetricsAPIService.FindBillableMetric")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/billable_metrics/{code}"
	localVarPath = strings.Replace(localVarPath, "{"+"code"+"}", url.PathEscape(parameterValueToString(r.code, "code")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiErrorUnauthorized
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiErrorNotFound
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type BillableMetricsAPIUpdateBillableMetricRequest struct {
	ctx context.Context
	ApiService *BillableMetricsAPIService
	code string
	billableMetricUpdateInput *BillableMetricUpdateInput
}

// Billable metric payload
func (r BillableMetricsAPIUpdateBillableMetricRequest) BillableMetricUpdateInput(billableMetricUpdateInput BillableMetricUpdateInput) BillableMetricsAPIUpdateBillableMetricRequest {
	r.billableMetricUpdateInput = &billableMetricUpdateInput
	return r
}

func (r BillableMetricsAPIUpdateBillableMetricRequest) Execute() (*BillableMetric, *http.Response, error) {
	return r.ApiService.UpdateBillableMetricExecute(r)
}

/*
UpdateBillableMetric Update a billable metric

This endpoint updates an existing billable metric representing a pricing component of your application.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param code Code of the existing billable metric.
 @return BillableMetricsAPIUpdateBillableMetricRequest
*/
func (a *BillableMetricsAPIService) UpdateBillableMetric(ctx context.Context, code string) BillableMetricsAPIUpdateBillableMetricRequest {
	return BillableMetricsAPIUpdateBillableMetricRequest{
		ApiService: a,
		ctx: ctx,
		code: code,
	}
}

// Execute executes the request
//  @return BillableMetric
func (a *BillableMetricsAPIService) UpdateBillableMetricExecute(r BillableMetricsAPIUpdateBillableMetricRequest) (*BillableMetric, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BillableMetric
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BillableMetricsAPIService.UpdateBillableMetric")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/billable_metrics/{code}"
	localVarPath = strings.Replace(localVarPath, "{"+"code"+"}", url.PathEscape(parameterValueToString(r.code, "code")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.billableMetricUpdateInput == nil {
		return localVarReturnValue, nil, reportError("billableMetricUpdateInput is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.billableMetricUpdateInput
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiErrorBadRequest
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ApiErrorUnauthorized
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiErrorNotFound
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v ApiErrorUnprocessableEntity
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
