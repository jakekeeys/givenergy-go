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

// InverterControlApiService InverterControlApi service
type InverterControlApiService service

type ApiInverterInverterPresetsGetRequest struct {
	ctx _context.Context
	ApiService *InverterControlApiService
	inverter string
	authorization *string
	contentType *string
	accept *string
}

func (r ApiInverterInverterPresetsGetRequest) Authorization(authorization string) ApiInverterInverterPresetsGetRequest {
	r.authorization = &authorization
	return r
}
func (r ApiInverterInverterPresetsGetRequest) ContentType(contentType string) ApiInverterInverterPresetsGetRequest {
	r.contentType = &contentType
	return r
}
func (r ApiInverterInverterPresetsGetRequest) Accept(accept string) ApiInverterInverterPresetsGetRequest {
	r.accept = &accept
	return r
}

func (r ApiInverterInverterPresetsGetRequest) Execute() (string, *_nethttp.Response, error) {
	return r.ApiService.InverterInverterPresetsGetExecute(r)
}

/*
InverterInverterPresetsGet Get Setting Presets

Retrieves a list of available setting presets for a given inverter

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inverter The serial number of the inverter
 @return ApiInverterInverterPresetsGetRequest
*/
func (a *InverterControlApiService) InverterInverterPresetsGet(ctx _context.Context, inverter string) ApiInverterInverterPresetsGetRequest {
	return ApiInverterInverterPresetsGetRequest{
		ApiService: a,
		ctx: ctx,
		inverter: inverter,
	}
}

// Execute executes the request
//  @return string
func (a *InverterControlApiService) InverterInverterPresetsGetExecute(r ApiInverterInverterPresetsGetRequest) (string, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InverterControlApiService.InverterInverterPresetsGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inverter/{inverter}/presets"
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
	localVarHTTPHeaderAccepts := []string{"text/plain"}

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

type ApiInverterInverterPresetsIdPostRequest struct {
	ctx _context.Context
	ApiService *InverterControlApiService
	inverter string
	id string
	authorization *string
	contentType *string
	accept *string
}

func (r ApiInverterInverterPresetsIdPostRequest) Authorization(authorization string) ApiInverterInverterPresetsIdPostRequest {
	r.authorization = &authorization
	return r
}
func (r ApiInverterInverterPresetsIdPostRequest) ContentType(contentType string) ApiInverterInverterPresetsIdPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiInverterInverterPresetsIdPostRequest) Accept(accept string) ApiInverterInverterPresetsIdPostRequest {
	r.accept = &accept
	return r
}

func (r ApiInverterInverterPresetsIdPostRequest) Execute() (InlineResponse2007, *_nethttp.Response, error) {
	return r.ApiService.InverterInverterPresetsIdPostExecute(r)
}

/*
InverterInverterPresetsIdPost Modify Preset

Modify one or more inverter settings using a given preset

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inverter The serial number of the inverter
 @param id The preset ID
 @return ApiInverterInverterPresetsIdPostRequest
*/
func (a *InverterControlApiService) InverterInverterPresetsIdPost(ctx _context.Context, inverter string, id string) ApiInverterInverterPresetsIdPostRequest {
	return ApiInverterInverterPresetsIdPostRequest{
		ApiService: a,
		ctx: ctx,
		inverter: inverter,
		id: id,
	}
}

// Execute executes the request
//  @return InlineResponse2007
func (a *InverterControlApiService) InverterInverterPresetsIdPostExecute(r ApiInverterInverterPresetsIdPostRequest) (InlineResponse2007, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2007
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InverterControlApiService.InverterInverterPresetsIdPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inverter/{inverter}/presets/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"inverter"+"}", _neturl.PathEscape(parameterToString(r.inverter, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

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

type ApiInverterInverterSettingsGetRequest struct {
	ctx _context.Context
	ApiService *InverterControlApiService
	inverter string
	authorization *string
	contentType *string
	accept *string
}

func (r ApiInverterInverterSettingsGetRequest) Authorization(authorization string) ApiInverterInverterSettingsGetRequest {
	r.authorization = &authorization
	return r
}
func (r ApiInverterInverterSettingsGetRequest) ContentType(contentType string) ApiInverterInverterSettingsGetRequest {
	r.contentType = &contentType
	return r
}
func (r ApiInverterInverterSettingsGetRequest) Accept(accept string) ApiInverterInverterSettingsGetRequest {
	r.accept = &accept
	return r
}

func (r ApiInverterInverterSettingsGetRequest) Execute() (InlineResponse2008, *_nethttp.Response, error) {
	return r.ApiService.InverterInverterSettingsGetExecute(r)
}

/*
InverterInverterSettingsGet Get Settings List

Returns a set of inverter settings available to your account

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inverter The serial number of the inverter
 @return ApiInverterInverterSettingsGetRequest
*/
func (a *InverterControlApiService) InverterInverterSettingsGet(ctx _context.Context, inverter string) ApiInverterInverterSettingsGetRequest {
	return ApiInverterInverterSettingsGetRequest{
		ApiService: a,
		ctx: ctx,
		inverter: inverter,
	}
}

// Execute executes the request
//  @return InlineResponse2008
func (a *InverterControlApiService) InverterInverterSettingsGetExecute(r ApiInverterInverterSettingsGetRequest) (InlineResponse2008, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2008
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InverterControlApiService.InverterInverterSettingsGet")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inverter/{inverter}/settings"
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

type ApiInverterInverterSettingsSettingReadPostRequest struct {
	ctx _context.Context
	ApiService *InverterControlApiService
	inverter string
	setting int32
	authorization *string
	contentType *string
	accept *string
}

func (r ApiInverterInverterSettingsSettingReadPostRequest) Authorization(authorization string) ApiInverterInverterSettingsSettingReadPostRequest {
	r.authorization = &authorization
	return r
}
func (r ApiInverterInverterSettingsSettingReadPostRequest) ContentType(contentType string) ApiInverterInverterSettingsSettingReadPostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiInverterInverterSettingsSettingReadPostRequest) Accept(accept string) ApiInverterInverterSettingsSettingReadPostRequest {
	r.accept = &accept
	return r
}

func (r ApiInverterInverterSettingsSettingReadPostRequest) Execute() (InlineResponse2009, *_nethttp.Response, error) {
	return r.ApiService.InverterInverterSettingsSettingReadPostExecute(r)
}

/*
InverterInverterSettingsSettingReadPost Read Setting

Read a specific setting on the inverter

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inverter The serial number of the inverter
 @param setting
 @return ApiInverterInverterSettingsSettingReadPostRequest
*/
func (a *InverterControlApiService) InverterInverterSettingsSettingReadPost(ctx _context.Context, inverter string, setting int32) ApiInverterInverterSettingsSettingReadPostRequest {
	return ApiInverterInverterSettingsSettingReadPostRequest{
		ApiService: a,
		ctx: ctx,
		inverter: inverter,
		setting: setting,
	}
}

// Execute executes the request
//  @return InlineResponse2009
func (a *InverterControlApiService) InverterInverterSettingsSettingReadPostExecute(r ApiInverterInverterSettingsSettingReadPostRequest) (InlineResponse2009, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2009
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InverterControlApiService.InverterInverterSettingsSettingReadPost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inverter/{inverter}/settings/{setting}/read"
	localVarPath = strings.Replace(localVarPath, "{"+"inverter"+"}", _neturl.PathEscape(parameterToString(r.inverter, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"setting"+"}", _neturl.PathEscape(parameterToString(r.setting, "")), -1)

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

type ApiInverterInverterSettingsSettingWritePostRequest struct {
	ctx _context.Context
	ApiService *InverterControlApiService
	inverter string
	setting int32
	inlineObject *InlineObject
	authorization *string
	contentType *string
	accept *string
}

func (r ApiInverterInverterSettingsSettingWritePostRequest) InlineObject(inlineObject InlineObject) ApiInverterInverterSettingsSettingWritePostRequest {
	r.inlineObject = &inlineObject
	return r
}
func (r ApiInverterInverterSettingsSettingWritePostRequest) Authorization(authorization string) ApiInverterInverterSettingsSettingWritePostRequest {
	r.authorization = &authorization
	return r
}
func (r ApiInverterInverterSettingsSettingWritePostRequest) ContentType(contentType string) ApiInverterInverterSettingsSettingWritePostRequest {
	r.contentType = &contentType
	return r
}
func (r ApiInverterInverterSettingsSettingWritePostRequest) Accept(accept string) ApiInverterInverterSettingsSettingWritePostRequest {
	r.accept = &accept
	return r
}

func (r ApiInverterInverterSettingsSettingWritePostRequest) Execute() (InlineResponse20010, *_nethttp.Response, error) {
	return r.ApiService.InverterInverterSettingsSettingWritePostExecute(r)
}

/*
InverterInverterSettingsSettingWritePost Modify Setting

Write a value to the setting on the inverter

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param inverter The serial number of the inverter
 @param setting
 @return ApiInverterInverterSettingsSettingWritePostRequest
*/
func (a *InverterControlApiService) InverterInverterSettingsSettingWritePost(ctx _context.Context, inverter string, setting int32) ApiInverterInverterSettingsSettingWritePostRequest {
	return ApiInverterInverterSettingsSettingWritePostRequest{
		ApiService: a,
		ctx: ctx,
		inverter: inverter,
		setting: setting,
	}
}

// Execute executes the request
//  @return InlineResponse20010
func (a *InverterControlApiService) InverterInverterSettingsSettingWritePostExecute(r ApiInverterInverterSettingsSettingWritePostRequest) (InlineResponse20010, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse20010
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InverterControlApiService.InverterInverterSettingsSettingWritePost")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/inverter/{inverter}/settings/{setting}/write"
	localVarPath = strings.Replace(localVarPath, "{"+"inverter"+"}", _neturl.PathEscape(parameterToString(r.inverter, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"setting"+"}", _neturl.PathEscape(parameterToString(r.setting, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject == nil {
		return localVarReturnValue, nil, reportError("inlineObject is required and must be specified")
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
	localVarPostBody = r.inlineObject
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
