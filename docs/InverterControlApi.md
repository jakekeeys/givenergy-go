# \InverterControlApi

All URIs are relative to *https://api.givenergy.cloud/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InverterInverterPresetsGet**](InverterControlApi.md#InverterInverterPresetsGet) | **Get** /inverter/{inverter}/presets | Get Setting Presets
[**InverterInverterPresetsIdPost**](InverterControlApi.md#InverterInverterPresetsIdPost) | **Post** /inverter/{inverter}/presets/{id} | Modify Preset
[**InverterInverterSettingsGet**](InverterControlApi.md#InverterInverterSettingsGet) | **Get** /inverter/{inverter}/settings | Get Settings List
[**InverterInverterSettingsSettingReadPost**](InverterControlApi.md#InverterInverterSettingsSettingReadPost) | **Post** /inverter/{inverter}/settings/{setting}/read | Read Setting
[**InverterInverterSettingsSettingWritePost**](InverterControlApi.md#InverterInverterSettingsSettingWritePost) | **Post** /inverter/{inverter}/settings/{setting}/write | Modify Setting



## InverterInverterPresetsGet

> string InverterInverterPresetsGet(ctx, inverter).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()

Get Setting Presets



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inverter := "consequatur" // string | The serial number of the inverter
    authorization := "Bearer {YOUR_API_KEY}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InverterControlApi.InverterInverterPresetsGet(context.Background(), inverter).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InverterControlApi.InverterInverterPresetsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InverterInverterPresetsGet`: string
    fmt.Fprintf(os.Stdout, "Response from `InverterControlApi.InverterInverterPresetsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inverter** | **string** | The serial number of the inverter | 

### Other Parameters

Other parameters are passed through a pointer to a apiInverterInverterPresetsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

### Return type

**string**

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InverterInverterPresetsIdPost

> InlineResponse2007 InverterInverterPresetsIdPost(ctx, inverter, id).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()

Modify Preset



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inverter := "consequatur" // string | The serial number of the inverter
    id := "consequatur" // string | The preset ID
    authorization := "Bearer {YOUR_API_KEY}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InverterControlApi.InverterInverterPresetsIdPost(context.Background(), inverter, id).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InverterControlApi.InverterInverterPresetsIdPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InverterInverterPresetsIdPost`: InlineResponse2007
    fmt.Fprintf(os.Stdout, "Response from `InverterControlApi.InverterInverterPresetsIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inverter** | **string** | The serial number of the inverter | 
**id** | **string** | The preset ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiInverterInverterPresetsIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**InlineResponse2007**](InlineResponse2007.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InverterInverterSettingsGet

> InlineResponse2008 InverterInverterSettingsGet(ctx, inverter).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()

Get Settings List



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inverter := "consequatur" // string | The serial number of the inverter
    authorization := "Bearer {YOUR_API_KEY}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InverterControlApi.InverterInverterSettingsGet(context.Background(), inverter).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InverterControlApi.InverterInverterSettingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InverterInverterSettingsGet`: InlineResponse2008
    fmt.Fprintf(os.Stdout, "Response from `InverterControlApi.InverterInverterSettingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inverter** | **string** | The serial number of the inverter | 

### Other Parameters

Other parameters are passed through a pointer to a apiInverterInverterSettingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**InlineResponse2008**](InlineResponse2008.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InverterInverterSettingsSettingReadPost

> InlineResponse2009 InverterInverterSettingsSettingReadPost(ctx, inverter, setting).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()

Read Setting



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inverter := "consequatur" // string | The serial number of the inverter
    setting := int32(1) // int32 | 
    authorization := "Bearer {YOUR_API_KEY}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InverterControlApi.InverterInverterSettingsSettingReadPost(context.Background(), inverter, setting).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InverterControlApi.InverterInverterSettingsSettingReadPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InverterInverterSettingsSettingReadPost`: InlineResponse2009
    fmt.Fprintf(os.Stdout, "Response from `InverterControlApi.InverterInverterSettingsSettingReadPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inverter** | **string** | The serial number of the inverter | 
**setting** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInverterInverterSettingsSettingReadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**InlineResponse2009**](InlineResponse2009.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InverterInverterSettingsSettingWritePost

> InlineResponse20010 InverterInverterSettingsSettingWritePost(ctx, inverter, setting).InlineObject(inlineObject).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()

Modify Setting



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inverter := "consequatur" // string | The serial number of the inverter
    setting := int32(1) // int32 | 
    inlineObject := *openapiclient.NewInlineObject("TODO") // InlineObject | 
    authorization := "Bearer {YOUR_API_KEY}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InverterControlApi.InverterInverterSettingsSettingWritePost(context.Background(), inverter, setting).InlineObject(inlineObject).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InverterControlApi.InverterInverterSettingsSettingWritePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InverterInverterSettingsSettingWritePost`: InlineResponse20010
    fmt.Fprintf(os.Stdout, "Response from `InverterControlApi.InverterInverterSettingsSettingWritePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inverter** | **string** | The serial number of the inverter | 
**setting** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiInverterInverterSettingsSettingWritePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inlineObject** | [**InlineObject**](InlineObject.md) |  | 
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**InlineResponse20010**](InlineResponse20010.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

