# \InverterDataApi

All URIs are relative to *https://api.givenergy.cloud/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InverterInverterDataPointsDateGet**](InverterDataApi.md#InverterInverterDataPointsDateGet) | **Get** /inverter/{inverter}/data-points/{date} | Get Data Points
[**InverterInverterEventsGet**](InverterDataApi.md#InverterInverterEventsGet) | **Get** /inverter/{inverter}/events | Get Events
[**InverterInverterMeterDataLatestGet**](InverterDataApi.md#InverterInverterMeterDataLatestGet) | **Get** /inverter/{inverter}/meter-data/latest | Get Latest Meter Data
[**InverterInverterSystemDataLatestGet**](InverterDataApi.md#InverterInverterSystemDataLatestGet) | **Get** /inverter/{inverter}/system-data/latest | Get Latest System Data



## InverterInverterDataPointsDateGet

> InlineResponse20014 InverterInverterDataPointsDateGet(ctx, inverter, date).Page(page).PageSize(pageSize).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()

Get Data Points



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
    date := TODO // ISO8601Date | The date to fetch the data. This assumes the date at the system's local time
    page := int32(1) // int32 | Page number to return (optional)
    pageSize := int32(56) // int32 | Number of items to return in a page. Defaults to 15 (optional)
    authorization := "Bearer {YOUR_API_KEY}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InverterDataApi.InverterInverterDataPointsDateGet(context.Background(), inverter, date).Page(page).PageSize(pageSize).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InverterDataApi.InverterInverterDataPointsDateGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InverterInverterDataPointsDateGet`: InlineResponse20014
    fmt.Fprintf(os.Stdout, "Response from `InverterDataApi.InverterInverterDataPointsDateGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inverter** | **string** | The serial number of the inverter | 
**date** | [**ISO8601Date**](.md) | The date to fetch the data. This assumes the date at the system&#39;s local time | 

### Other Parameters

Other parameters are passed through a pointer to a apiInverterInverterDataPointsDateGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Page number to return | 
 **pageSize** | **int32** | Number of items to return in a page. Defaults to 15 | 
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**InlineResponse20014**](InlineResponse20014.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InverterInverterEventsGet

> InlineResponse20013 InverterInverterEventsGet(ctx, inverter).Start(start).End(end).Cleared(cleared).Page(page).PageSize(pageSize).Authorization(authorization).ContentType(contentType).Accept(accept).InlineObject1(inlineObject1).Execute()

Get Events



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
    start := TODO // ISO8601Datetime | Start time (optional)
    end := TODO // ISO8601Datetime | End time (optional)
    cleared := false // bool | Whether 'cleared' events should be included with the data. Default is false (optional)
    page := int32(1) // int32 | Page number to return (optional)
    pageSize := int32(56) // int32 | Number of items to return in a page. Defaults to 15 (optional)
    authorization := "Bearer {YOUR_API_KEY}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    accept := "application/json" // string |  (optional)
    inlineObject1 := *openapiclient.NewInlineObject1() // InlineObject1 |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InverterDataApi.InverterInverterEventsGet(context.Background(), inverter).Start(start).End(end).Cleared(cleared).Page(page).PageSize(pageSize).Authorization(authorization).ContentType(contentType).Accept(accept).InlineObject1(inlineObject1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InverterDataApi.InverterInverterEventsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InverterInverterEventsGet`: InlineResponse20013
    fmt.Fprintf(os.Stdout, "Response from `InverterDataApi.InverterInverterEventsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inverter** | **string** | The serial number of the inverter | 

### Other Parameters

Other parameters are passed through a pointer to a apiInverterInverterEventsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | [**ISO8601Datetime**](ISO8601Datetime.md) | Start time | 
 **end** | [**ISO8601Datetime**](ISO8601Datetime.md) | End time | 
 **cleared** | **bool** | Whether &#39;cleared&#39; events should be included with the data. Default is false | 
 **page** | **int32** | Page number to return | 
 **pageSize** | **int32** | Number of items to return in a page. Defaults to 15 | 
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 
 **inlineObject1** | [**InlineObject1**](InlineObject1.md) |  | 

### Return type

[**InlineResponse20013**](InlineResponse20013.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InverterInverterMeterDataLatestGet

> InlineResponse20012 InverterInverterMeterDataLatestGet(ctx, inverter).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()

Get Latest Meter Data



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
    resp, r, err := api_client.InverterDataApi.InverterInverterMeterDataLatestGet(context.Background(), inverter).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InverterDataApi.InverterInverterMeterDataLatestGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InverterInverterMeterDataLatestGet`: InlineResponse20012
    fmt.Fprintf(os.Stdout, "Response from `InverterDataApi.InverterInverterMeterDataLatestGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inverter** | **string** | The serial number of the inverter | 

### Other Parameters

Other parameters are passed through a pointer to a apiInverterInverterMeterDataLatestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**InlineResponse20012**](InlineResponse20012.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InverterInverterSystemDataLatestGet

> InlineResponse20011 InverterInverterSystemDataLatestGet(ctx, inverter).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()

Get Latest System Data



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
    resp, r, err := api_client.InverterDataApi.InverterInverterSystemDataLatestGet(context.Background(), inverter).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InverterDataApi.InverterInverterSystemDataLatestGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InverterInverterSystemDataLatestGet`: InlineResponse20011
    fmt.Fprintf(os.Stdout, "Response from `InverterDataApi.InverterInverterSystemDataLatestGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inverter** | **string** | The serial number of the inverter | 

### Other Parameters

Other parameters are passed through a pointer to a apiInverterInverterSystemDataLatestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**InlineResponse20011**](InlineResponse20011.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

