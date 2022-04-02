# \SmartDeviceApi

All URIs are relative to *https://api.givenergy.cloud/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SmartDeviceGet**](SmartDeviceApi.md#SmartDeviceGet) | **Get** /smart-device | Get Your Smart Devices
[**SmartDeviceIdGet**](SmartDeviceApi.md#SmartDeviceIdGet) | **Get** /smart-device/{id} | Get Smart Device by ID
[**SmartDevicePost**](SmartDeviceApi.md#SmartDevicePost) | **Post** /smart-device | Create Smart Device
[**SmartDeviceSmartDeviceIdDataGet**](SmartDeviceApi.md#SmartDeviceSmartDeviceIdDataGet) | **Get** /smart-device/{smart-device_id}/data | Get Smart Device Data Points by ID
[**SmartDeviceSmartDeviceUuidDataPost**](SmartDeviceApi.md#SmartDeviceSmartDeviceUuidDataPost) | **Post** /smart-device/{smart_device_uuid}/data | Create Smart Device Data Point



## SmartDeviceGet

> InlineResponse20017 SmartDeviceGet(ctx).Page(page).PageSize(pageSize).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()

Get Your Smart Devices



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
    page := int32(1) // int32 | Page number to return (optional)
    pageSize := int32(56) // int32 | Number of items to return in a page. Defaults to 15 (optional)
    authorization := "Bearer {YOUR_API_KEY}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SmartDeviceApi.SmartDeviceGet(context.Background()).Page(page).PageSize(pageSize).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartDeviceApi.SmartDeviceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SmartDeviceGet`: InlineResponse20017
    fmt.Fprintf(os.Stdout, "Response from `SmartDeviceApi.SmartDeviceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSmartDeviceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number to return | 
 **pageSize** | **int32** | Number of items to return in a page. Defaults to 15 | 
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**InlineResponse20017**](InlineResponse20017.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartDeviceIdGet

> InlineResponse20019 SmartDeviceIdGet(ctx, id).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()

Get Smart Device by ID



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
    id := int32(17) // int32 | The ID of the smart device.
    authorization := "Bearer {YOUR_API_KEY}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SmartDeviceApi.SmartDeviceIdGet(context.Background(), id).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartDeviceApi.SmartDeviceIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SmartDeviceIdGet`: InlineResponse20019
    fmt.Fprintf(os.Stdout, "Response from `SmartDeviceApi.SmartDeviceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the smart device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSmartDeviceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**InlineResponse20019**](InlineResponse20019.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartDevicePost

> InlineResponse20018 SmartDevicePost(ctx).InlineObject2(inlineObject2).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()

Create Smart Device



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
    inlineObject2 := *openapiclient.NewInlineObject2("consequatur") // InlineObject2 | 
    authorization := "Bearer {YOUR_API_KEY}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SmartDeviceApi.SmartDevicePost(context.Background()).InlineObject2(inlineObject2).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartDeviceApi.SmartDevicePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SmartDevicePost`: InlineResponse20018
    fmt.Fprintf(os.Stdout, "Response from `SmartDeviceApi.SmartDevicePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSmartDevicePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inlineObject2** | [**InlineObject2**](InlineObject2.md) |  | 
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**InlineResponse20018**](InlineResponse20018.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartDeviceSmartDeviceIdDataGet

> InlineResponse20020 SmartDeviceSmartDeviceIdDataGet(ctx, smartDeviceId).Page(page).PageSize(pageSize).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()

Get Smart Device Data Points by ID



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
    smartDeviceId := "consequatur" // string | The ID of the smart-device.
    page := int32(1) // int32 | Page number to return (optional)
    pageSize := int32(56) // int32 | Number of items to return in a page. Defaults to 15 (optional)
    authorization := "Bearer {YOUR_API_KEY}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SmartDeviceApi.SmartDeviceSmartDeviceIdDataGet(context.Background(), smartDeviceId).Page(page).PageSize(pageSize).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartDeviceApi.SmartDeviceSmartDeviceIdDataGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SmartDeviceSmartDeviceIdDataGet`: InlineResponse20020
    fmt.Fprintf(os.Stdout, "Response from `SmartDeviceApi.SmartDeviceSmartDeviceIdDataGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smartDeviceId** | **string** | The ID of the smart-device. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSmartDeviceSmartDeviceIdDataGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number to return | 
 **pageSize** | **int32** | Number of items to return in a page. Defaults to 15 | 
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**InlineResponse20020**](InlineResponse20020.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartDeviceSmartDeviceUuidDataPost

> InlineResponse20021 SmartDeviceSmartDeviceUuidDataPost(ctx, smartDeviceUuid).InlineObject3(inlineObject3).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()

Create Smart Device Data Point



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
    smartDeviceUuid := "consequatur" // string | 
    inlineObject3 := *openapiclient.NewInlineObject3("TODO", int32(17)) // InlineObject3 | 
    authorization := "Bearer {YOUR_API_KEY}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SmartDeviceApi.SmartDeviceSmartDeviceUuidDataPost(context.Background(), smartDeviceUuid).InlineObject3(inlineObject3).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SmartDeviceApi.SmartDeviceSmartDeviceUuidDataPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SmartDeviceSmartDeviceUuidDataPost`: InlineResponse20021
    fmt.Fprintf(os.Stdout, "Response from `SmartDeviceApi.SmartDeviceSmartDeviceUuidDataPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**smartDeviceUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSmartDeviceSmartDeviceUuidDataPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject3** | [**InlineObject3**](InlineObject3.md) |  | 
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**InlineResponse20021**](InlineResponse20021.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

