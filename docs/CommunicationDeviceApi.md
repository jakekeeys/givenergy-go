# \CommunicationDeviceApi

All URIs are relative to *https://api.givenergy.cloud/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommunicationDeviceGet**](CommunicationDeviceApi.md#CommunicationDeviceGet) | **Get** /communication-device | Get Your Communication Devices
[**CommunicationDeviceIdGet**](CommunicationDeviceApi.md#CommunicationDeviceIdGet) | **Get** /communication-device/{id} | Get Communication Device Information by Serial Number



## CommunicationDeviceGet

> InlineResponse2005 CommunicationDeviceGet(ctx).Page(page).PageSize(pageSize).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()

Get Your Communication Devices



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
    resp, r, err := api_client.CommunicationDeviceApi.CommunicationDeviceGet(context.Background()).Page(page).PageSize(pageSize).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationDeviceApi.CommunicationDeviceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationDeviceGet`: InlineResponse2005
    fmt.Fprintf(os.Stdout, "Response from `CommunicationDeviceApi.CommunicationDeviceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationDeviceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number to return | 
 **pageSize** | **int32** | Number of items to return in a page. Defaults to 15 | 
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**InlineResponse2005**](InlineResponse2005.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommunicationDeviceIdGet

> InlineResponse2006 CommunicationDeviceIdGet(ctx, id).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()

Get Communication Device Information by Serial Number



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
    id := "consequatur" // string | The serial number of the communication device
    authorization := "Bearer {YOUR_API_KEY}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommunicationDeviceApi.CommunicationDeviceIdGet(context.Background(), id).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommunicationDeviceApi.CommunicationDeviceIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommunicationDeviceIdGet`: InlineResponse2006
    fmt.Fprintf(os.Stdout, "Response from `CommunicationDeviceApi.CommunicationDeviceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The serial number of the communication device | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommunicationDeviceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**InlineResponse2006**](InlineResponse2006.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

