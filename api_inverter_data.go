/*
GivEnergy API Documentation (v1.5.1)

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package givenergy

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// InverterDataApiService InverterDataApi service
type InverterDataApiService service

type ApiInverterInverterDataPointsDateGetRequest struct {
	ctx _context.Context
	ApiService *InverterDataApiService
	inverter string
	date ISO8601Date
	page *int32
	pageSize *int32
	authorization *string
	contentType *string
	accept *string
}

// Page number to return
func (r ApiInverterInverterDataPointsDateGetRequest) Page(page int32) ApiInverterInverterDataPointsDateGetRequest {
	r.page = &page
	return r
}
// Number of items to return in a page. Defaults to 15
func (r ApiInverterInverterDataPointsDateGetRequest) PageSize(pageSize int32) ApiInverterInverterDataPointsDateGetRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiInverterInverterDataPointsDateGetRequest) Authorization(authorization string) ApiInverterInverterDataPointsDateGetRequest {
	r.authorization = &authorization
	return r
}
func (r ApiInverterInverterDataPointsDateGetRequest) ContentType(contentType string) ApiInverterInverterDataPointsDateGetRequest {
	r.contentType = &contentType
	return r
}
func (r ApiInverterInverterDataPointsDateGetRequest) Accept(accept string) ApiInverterInverterDataPointsDateGetRequest {
	r.accept = &accept
	return r
}

func (r ApiInverterInverterDataPointsDateGetRequest) Execute() (InlineResponse20014, *_nethttp.Response, error) {
	return r.ApiService.InverterInverterDataPointsDateGetExecute(r)
}

/*
InverterInverterDataPointsDateGet Get Data Points

Displays the entire data packet set from the chosen date

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inverter The serial number of the inverter
 @param date The date to fetch the data. This assumes the date at the system's local time
 @return ApiInverterInverterDataPointsDateGetRequest
*/
func (a *InverterDataApiService) InverterInverterDataPointsDateGet(ctx _context.Context, inverter string, date ISO8601Date) ApiInverterInverterDataPointsDateGetRequest {
	return ApiInverterInverterDataPointsDateGetRequest{
		ApiService: a,
		ctx: ctx,
		inverter: inverter,
		date: date,
	}
}

// Execute executes the request
//  @return InlineResponse20014
func (a *InverterDataApiService) InverterInverterDataPointsDateGetExecute(r ApiInverterInverterDataPointsDateGetRequest) (InlineResponse20014, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20014
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InverterDataApiService.InverterInverterDataPointsDateGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inverter/{inverter}/data-points/{date}"
	localVarPath = strings.Replace(localVarPath, "{"+"inverter"+"}", _neturl.PathEscape(parameterToString(r.inverter, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"date"+"}", _neturl.PathEscape(parameterToString(r.date, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
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
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	if r.contentType != nil {
		localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
	}
	if r.accept != nil {
		localVarHeaderParams["Accept"] = parameterToString(*r.accept, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiInverterInverterEventsGetRequest struct {
	ctx _context.Context
	ApiService *InverterDataApiService
	inverter string
	start *ISO8601Datetime
	end *ISO8601Datetime
	cleared *bool
	page *int32
	pageSize *int32
	authorization *string
	contentType *string
	accept *string
	inlineObject1 *InlineObject1
}

// Start time
func (r ApiInverterInverterEventsGetRequest) Start(start ISO8601Datetime) ApiInverterInverterEventsGetRequest {
	r.start = &start
	return r
}
// End time
func (r ApiInverterInverterEventsGetRequest) End(end ISO8601Datetime) ApiInverterInverterEventsGetRequest {
	r.end = &end
	return r
}
// Whether &#39;cleared&#39; events should be included with the data. Default is false
func (r ApiInverterInverterEventsGetRequest) Cleared(cleared bool) ApiInverterInverterEventsGetRequest {
	r.cleared = &cleared
	return r
}
// Page number to return
func (r ApiInverterInverterEventsGetRequest) Page(page int32) ApiInverterInverterEventsGetRequest {
	r.page = &page
	return r
}
// Number of items to return in a page. Defaults to 15
func (r ApiInverterInverterEventsGetRequest) PageSize(pageSize int32) ApiInverterInverterEventsGetRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiInverterInverterEventsGetRequest) Authorization(authorization string) ApiInverterInverterEventsGetRequest {
	r.authorization = &authorization
	return r
}
func (r ApiInverterInverterEventsGetRequest) ContentType(contentType string) ApiInverterInverterEventsGetRequest {
	r.contentType = &contentType
	return r
}
func (r ApiInverterInverterEventsGetRequest) Accept(accept string) ApiInverterInverterEventsGetRequest {
	r.accept = &accept
	return r
}
func (r ApiInverterInverterEventsGetRequest) InlineObject1(inlineObject1 InlineObject1) ApiInverterInverterEventsGetRequest {
	r.inlineObject1 = &inlineObject1
	return r
}

func (r ApiInverterInverterEventsGetRequest) Execute() (InlineResponse20013, *_nethttp.Response, error) {
	return r.ApiService.InverterInverterEventsGetExecute(r)
}

/*
InverterInverterEventsGet Get Events

Retrieves a list of faults that were triggered from the inverter and when they were cleared

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inverter The serial number of the inverter
 @return ApiInverterInverterEventsGetRequest
*/
func (a *InverterDataApiService) InverterInverterEventsGet(ctx _context.Context, inverter string) ApiInverterInverterEventsGetRequest {
	return ApiInverterInverterEventsGetRequest{
		ApiService: a,
		ctx: ctx,
		inverter: inverter,
	}
}

// Execute executes the request
//  @return InlineResponse20013
func (a *InverterDataApiService) InverterInverterEventsGetExecute(r ApiInverterInverterEventsGetRequest) (InlineResponse20013, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20013
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InverterDataApiService.InverterInverterEventsGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inverter/{inverter}/events"
	localVarPath = strings.Replace(localVarPath, "{"+"inverter"+"}", _neturl.PathEscape(parameterToString(r.inverter, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.start != nil {
		localVarQueryParams.Add("start", parameterToString(*r.start, ""))
	}
	if r.end != nil {
		localVarQueryParams.Add("end", parameterToString(*r.end, ""))
	}
	if r.cleared != nil {
		localVarQueryParams.Add("cleared", parameterToString(*r.cleared, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("pageSize", parameterToString(*r.pageSize, ""))
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
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	if r.contentType != nil {
		localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
	}
	if r.accept != nil {
		localVarHeaderParams["Accept"] = parameterToString(*r.accept, "")
	}
	// body params
	localVarPostBody = r.inlineObject1
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiInverterInverterMeterDataLatestGetRequest struct {
	ctx _context.Context
	ApiService *InverterDataApiService
	inverter string
	authorization *string
	contentType *string
	accept *string
}

func (r ApiInverterInverterMeterDataLatestGetRequest) Authorization(authorization string) ApiInverterInverterMeterDataLatestGetRequest {
	r.authorization = &authorization
	return r
}
func (r ApiInverterInverterMeterDataLatestGetRequest) ContentType(contentType string) ApiInverterInverterMeterDataLatestGetRequest {
	r.contentType = &contentType
	return r
}
func (r ApiInverterInverterMeterDataLatestGetRequest) Accept(accept string) ApiInverterInverterMeterDataLatestGetRequest {
	r.accept = &accept
	return r
}

func (r ApiInverterInverterMeterDataLatestGetRequest) Execute() (InlineResponse20012, *_nethttp.Response, error) {
	return r.ApiService.InverterInverterMeterDataLatestGetExecute(r)
}

/*
InverterInverterMeterDataLatestGet Get Latest Meter Data

Retrieves the latest meter data from the inverter

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inverter The serial number of the inverter
 @return ApiInverterInverterMeterDataLatestGetRequest
*/
func (a *InverterDataApiService) InverterInverterMeterDataLatestGet(ctx _context.Context, inverter string) ApiInverterInverterMeterDataLatestGetRequest {
	return ApiInverterInverterMeterDataLatestGetRequest{
		ApiService: a,
		ctx: ctx,
		inverter: inverter,
	}
}

// Execute executes the request
//  @return InlineResponse20012
func (a *InverterDataApiService) InverterInverterMeterDataLatestGetExecute(r ApiInverterInverterMeterDataLatestGetRequest) (InlineResponse20012, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20012
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InverterDataApiService.InverterInverterMeterDataLatestGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inverter/{inverter}/meter-data/latest"
	localVarPath = strings.Replace(localVarPath, "{"+"inverter"+"}", _neturl.PathEscape(parameterToString(r.inverter, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	if r.contentType != nil {
		localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
	}
	if r.accept != nil {
		localVarHeaderParams["Accept"] = parameterToString(*r.accept, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiInverterInverterSystemDataLatestGetRequest struct {
	ctx _context.Context
	ApiService *InverterDataApiService
	inverter string
	authorization *string
	contentType *string
	accept *string
}

func (r ApiInverterInverterSystemDataLatestGetRequest) Authorization(authorization string) ApiInverterInverterSystemDataLatestGetRequest {
	r.authorization = &authorization
	return r
}
func (r ApiInverterInverterSystemDataLatestGetRequest) ContentType(contentType string) ApiInverterInverterSystemDataLatestGetRequest {
	r.contentType = &contentType
	return r
}
func (r ApiInverterInverterSystemDataLatestGetRequest) Accept(accept string) ApiInverterInverterSystemDataLatestGetRequest {
	r.accept = &accept
	return r
}

func (r ApiInverterInverterSystemDataLatestGetRequest) Execute() (InlineResponse20011, *_nethttp.Response, error) {
	return r.ApiService.InverterInverterSystemDataLatestGetExecute(r)
}

/*
InverterInverterSystemDataLatestGet Get Latest System Data

Retrieves the latest system data from the inverter

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inverter The serial number of the inverter
 @return ApiInverterInverterSystemDataLatestGetRequest
*/
func (a *InverterDataApiService) InverterInverterSystemDataLatestGet(ctx _context.Context, inverter string) ApiInverterInverterSystemDataLatestGetRequest {
	return ApiInverterInverterSystemDataLatestGetRequest{
		ApiService: a,
		ctx: ctx,
		inverter: inverter,
	}
}

// Execute executes the request
//  @return InlineResponse20011
func (a *InverterDataApiService) InverterInverterSystemDataLatestGetExecute(r ApiInverterInverterSystemDataLatestGetRequest) (InlineResponse20011, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20011
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InverterDataApiService.InverterInverterSystemDataLatestGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inverter/{inverter}/system-data/latest"
	localVarPath = strings.Replace(localVarPath, "{"+"inverter"+"}", _neturl.PathEscape(parameterToString(r.inverter, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	if r.authorization != nil {
		localVarHeaderParams["Authorization"] = parameterToString(*r.authorization, "")
	}
	if r.contentType != nil {
		localVarHeaderParams["Content-Type"] = parameterToString(*r.contentType, "")
	}
	if r.accept != nil {
		localVarHeaderParams["Accept"] = parameterToString(*r.accept, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
