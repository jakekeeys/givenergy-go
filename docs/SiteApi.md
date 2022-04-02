# \SiteApi

All URIs are relative to *https://api.givenergy.cloud/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SiteGet**](SiteApi.md#SiteGet) | **Get** /site | Get Your Sites
[**SiteIdGet**](SiteApi.md#SiteIdGet) | **Get** /site/{id} | Get Single Site by ID



## SiteGet

> InlineResponse20015 SiteGet(ctx).Page(page).PageSize(pageSize).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()

Get Your Sites



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
    resp, r, err := api_client.SiteApi.SiteGet(context.Background()).Page(page).PageSize(pageSize).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteApi.SiteGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SiteGet`: InlineResponse20015
    fmt.Fprintf(os.Stdout, "Response from `SiteApi.SiteGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSiteGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number to return | 
 **pageSize** | **int32** | Number of items to return in a page. Defaults to 15 | 
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**InlineResponse20015**](InlineResponse20015.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SiteIdGet

> InlineResponse20016 SiteIdGet(ctx, id).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()

Get Single Site by ID



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
    id := int32(17) // int32 | The ID of the site.
    authorization := "Bearer {YOUR_API_KEY}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SiteApi.SiteIdGet(context.Background(), id).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SiteApi.SiteIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SiteIdGet`: InlineResponse20016
    fmt.Fprintf(os.Stdout, "Response from `SiteApi.SiteIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the site. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSiteIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**InlineResponse20016**](InlineResponse20016.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

