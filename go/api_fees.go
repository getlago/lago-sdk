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
	"time"
)


// FeesAPIService FeesAPI service
type FeesAPIService service

type FeesAPIFindAllFeesRequest struct {
	ctx context.Context
	ApiService *FeesAPIService
	page *int32
	perPage *int32
	externalCustomerId *string
	externalSubscriptionId *string
	currency *Currency
	feeType *string
	billableMetricCode *string
	paymentStatus *string
	createdAtFrom *time.Time
	createdAtTo *time.Time
	succeededAtFrom *time.Time
	succeededAtTo *time.Time
	failedAtFrom *time.Time
	failedAtTo *time.Time
	refundedAtFrom *time.Time
	refundedAtTo *time.Time
}

// Page number.
func (r FeesAPIFindAllFeesRequest) Page(page int32) FeesAPIFindAllFeesRequest {
	r.page = &page
	return r
}

// Number of records per page.
func (r FeesAPIFindAllFeesRequest) PerPage(perPage int32) FeesAPIFindAllFeesRequest {
	r.perPage = &perPage
	return r
}

// Unique identifier assigned to the customer in your application.
func (r FeesAPIFindAllFeesRequest) ExternalCustomerId(externalCustomerId string) FeesAPIFindAllFeesRequest {
	r.externalCustomerId = &externalCustomerId
	return r
}

// External subscription ID
func (r FeesAPIFindAllFeesRequest) ExternalSubscriptionId(externalSubscriptionId string) FeesAPIFindAllFeesRequest {
	r.externalSubscriptionId = &externalSubscriptionId
	return r
}

// Filter results by fee’s currency.
func (r FeesAPIFindAllFeesRequest) Currency(currency Currency) FeesAPIFindAllFeesRequest {
	r.currency = &currency
	return r
}

// The fee type. Possible values are &#x60;add-on&#x60;, &#x60;charge&#x60;, &#x60;credit&#x60; or &#x60;subscription&#x60;.
func (r FeesAPIFindAllFeesRequest) FeeType(feeType string) FeesAPIFindAllFeesRequest {
	r.feeType = &feeType
	return r
}

// Filter results by the &#x60;code&#x60; of the billable metric attached to the fee. Only applies to &#x60;charge&#x60; types.
func (r FeesAPIFindAllFeesRequest) BillableMetricCode(billableMetricCode string) FeesAPIFindAllFeesRequest {
	r.billableMetricCode = &billableMetricCode
	return r
}

// Indicates the payment status of the fee. It represents the current status of the payment associated with the fee. The possible values for this field are &#x60;pending&#x60;, &#x60;succeeded&#x60;, &#x60;failed&#x60; and refunded&#x60;.
func (r FeesAPIFindAllFeesRequest) PaymentStatus(paymentStatus string) FeesAPIFindAllFeesRequest {
	r.paymentStatus = &paymentStatus
	return r
}

// Filter results created after creation date and time in UTC.
func (r FeesAPIFindAllFeesRequest) CreatedAtFrom(createdAtFrom time.Time) FeesAPIFindAllFeesRequest {
	r.createdAtFrom = &createdAtFrom
	return r
}

// Filter results created before creation date and time in UTC.
func (r FeesAPIFindAllFeesRequest) CreatedAtTo(createdAtTo time.Time) FeesAPIFindAllFeesRequest {
	r.createdAtTo = &createdAtTo
	return r
}

// Filter results with payment success after creation date and time in UTC.
func (r FeesAPIFindAllFeesRequest) SucceededAtFrom(succeededAtFrom time.Time) FeesAPIFindAllFeesRequest {
	r.succeededAtFrom = &succeededAtFrom
	return r
}

// Filter results with payment success after creation date and time in UTC.
func (r FeesAPIFindAllFeesRequest) SucceededAtTo(succeededAtTo time.Time) FeesAPIFindAllFeesRequest {
	r.succeededAtTo = &succeededAtTo
	return r
}

// Filter results with payment failure after creation date and time in UTC.
func (r FeesAPIFindAllFeesRequest) FailedAtFrom(failedAtFrom time.Time) FeesAPIFindAllFeesRequest {
	r.failedAtFrom = &failedAtFrom
	return r
}

// Filter results with payment failure after creation date and time in UTC.
func (r FeesAPIFindAllFeesRequest) FailedAtTo(failedAtTo time.Time) FeesAPIFindAllFeesRequest {
	r.failedAtTo = &failedAtTo
	return r
}

// Filter results with payment refund after creation date and time in UTC.
func (r FeesAPIFindAllFeesRequest) RefundedAtFrom(refundedAtFrom time.Time) FeesAPIFindAllFeesRequest {
	r.refundedAtFrom = &refundedAtFrom
	return r
}

// Filter results with payment refund after creation date and time in UTC.
func (r FeesAPIFindAllFeesRequest) RefundedAtTo(refundedAtTo time.Time) FeesAPIFindAllFeesRequest {
	r.refundedAtTo = &refundedAtTo
	return r
}

func (r FeesAPIFindAllFeesRequest) Execute() (*FeesPaginated, *http.Response, error) {
	return r.ApiService.FindAllFeesExecute(r)
}

/*
FindAllFees List all fees

This endpoint is used for retrieving all fees that has been issued.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return FeesAPIFindAllFeesRequest
*/
func (a *FeesAPIService) FindAllFees(ctx context.Context) FeesAPIFindAllFeesRequest {
	return FeesAPIFindAllFeesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FeesPaginated
func (a *FeesAPIService) FindAllFeesExecute(r FeesAPIFindAllFeesRequest) (*FeesPaginated, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FeesPaginated
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeesAPIService.FindAllFees")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fees"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.perPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "per_page", r.perPage, "")
	}
	if r.externalCustomerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "external_customer_id", r.externalCustomerId, "")
	}
	if r.externalSubscriptionId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "external_subscription_id", r.externalSubscriptionId, "")
	}
	if r.currency != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "currency", r.currency, "")
	}
	if r.feeType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fee_type", r.feeType, "")
	}
	if r.billableMetricCode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "billable_metric_code", r.billableMetricCode, "")
	}
	if r.paymentStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "payment_status", r.paymentStatus, "")
	}
	if r.createdAtFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "created_at_from", r.createdAtFrom, "")
	}
	if r.createdAtTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "created_at_to", r.createdAtTo, "")
	}
	if r.succeededAtFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "succeeded_at_from", r.succeededAtFrom, "")
	}
	if r.succeededAtTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "succeeded_at_to", r.succeededAtTo, "")
	}
	if r.failedAtFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "failed_at_from", r.failedAtFrom, "")
	}
	if r.failedAtTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "failed_at_to", r.failedAtTo, "")
	}
	if r.refundedAtFrom != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "refunded_at_from", r.refundedAtFrom, "")
	}
	if r.refundedAtTo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "refunded_at_to", r.refundedAtTo, "")
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

type FeesAPIFindFeeRequest struct {
	ctx context.Context
	ApiService *FeesAPIService
	lagoId string
}

func (r FeesAPIFindFeeRequest) Execute() (*Fee, *http.Response, error) {
	return r.ApiService.FindFeeExecute(r)
}

/*
FindFee Retrieve a specific fee

This endpoint is used for retrieving a specific fee that has been issued.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param lagoId Unique identifier assigned to the fee within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the fee’s record within the Lago system.
 @return FeesAPIFindFeeRequest
*/
func (a *FeesAPIService) FindFee(ctx context.Context, lagoId string) FeesAPIFindFeeRequest {
	return FeesAPIFindFeeRequest{
		ApiService: a,
		ctx: ctx,
		lagoId: lagoId,
	}
}

// Execute executes the request
//  @return Fee
func (a *FeesAPIService) FindFeeExecute(r FeesAPIFindFeeRequest) (*Fee, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Fee
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeesAPIService.FindFee")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fees/{lago_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"lago_id"+"}", url.PathEscape(parameterValueToString(r.lagoId, "lagoId")), -1)

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

type FeesAPIUpdateFeeRequest struct {
	ctx context.Context
	ApiService *FeesAPIService
	lagoId string
	feeUpdateInput *FeeUpdateInput
}

// Fee payload
func (r FeesAPIUpdateFeeRequest) FeeUpdateInput(feeUpdateInput FeeUpdateInput) FeesAPIUpdateFeeRequest {
	r.feeUpdateInput = &feeUpdateInput
	return r
}

func (r FeesAPIUpdateFeeRequest) Execute() (*Fee, *http.Response, error) {
	return r.ApiService.UpdateFeeExecute(r)
}

/*
UpdateFee Update a fee

This endpoint is used for updating a specific fee that has been issued.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param lagoId Unique identifier assigned to the fee within the Lago application. This ID is exclusively created by Lago and serves as a unique identifier for the fee’s record within the Lago system.
 @return FeesAPIUpdateFeeRequest
*/
func (a *FeesAPIService) UpdateFee(ctx context.Context, lagoId string) FeesAPIUpdateFeeRequest {
	return FeesAPIUpdateFeeRequest{
		ApiService: a,
		ctx: ctx,
		lagoId: lagoId,
	}
}

// Execute executes the request
//  @return Fee
func (a *FeesAPIService) UpdateFeeExecute(r FeesAPIUpdateFeeRequest) (*Fee, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Fee
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeesAPIService.UpdateFee")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fees/{lago_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"lago_id"+"}", url.PathEscape(parameterValueToString(r.lagoId, "lagoId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.feeUpdateInput
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
