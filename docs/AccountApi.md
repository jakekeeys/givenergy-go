# \AccountApi

All URIs are relative to *https://api.givenergy.cloud/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountChildrenAccountGet**](AccountApi.md#AccountChildrenAccountGet) | **Get** /account-children/{account} | Get Account Children Information by ID
[**AccountChildrenGet**](AccountApi.md#AccountChildrenGet) | **Get** /account-children | Get Your Account Children Information
[**AccountGet**](AccountApi.md#AccountGet) | **Get** /account | Get Your Account Information
[**AccountIdGet**](AccountApi.md#AccountIdGet) | **Get** /account/{id} | Get Account Information by ID
[**AccountSearchAccountUserAccountGet**](AccountApi.md#AccountSearchAccountUserAccountGet) | **Get** /account/search/{account_user_account} | Get Account Information by Username



## AccountChildrenAccountGet

> InlineResponse2004 AccountChildrenAccountGet(ctx, account).Page(page).PageSize(pageSize).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()

Get Account Children Information by ID



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
    account := int32(17) // int32 | The ID of the account
    page := int32(1) // int32 | Page number to return (optional)
    pageSize := int32(56) // int32 | Number of items to return in a page. Defaults to 15 (optional)
    authorization := "Bearer {YOUR_API_KEY}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.AccountChildrenAccountGet(context.Background(), account).Page(page).PageSize(pageSize).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AccountChildrenAccountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountChildrenAccountGet`: InlineResponse2004
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.AccountChildrenAccountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**account** | **int32** | The ID of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountChildrenAccountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number to return | 
 **pageSize** | **int32** | Number of items to return in a page. Defaults to 15 | 
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**InlineResponse2004**](InlineResponse2004.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountChildrenGet

> InlineResponse2003 AccountChildrenGet(ctx).Page(page).PageSize(pageSize).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()

Get Your Account Children Information



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
    resp, r, err := api_client.AccountApi.AccountChildrenGet(context.Background()).Page(page).PageSize(pageSize).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AccountChildrenGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountChildrenGet`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.AccountChildrenGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountChildrenGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number to return | 
 **pageSize** | **int32** | Number of items to return in a page. Defaults to 15 | 
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountGet

> InlineResponse200 AccountGet(ctx).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()

Get Your Account Information



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
    authorization := "Bearer {YOUR_API_KEY}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.AccountGet(context.Background()).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AccountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountGet`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.AccountGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountIdGet

> InlineResponse2001 AccountIdGet(ctx, id).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()

Get Account Information by ID



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
    id := int32(17) // int32 | The ID of the account.
    authorization := "Bearer {YOUR_API_KEY}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.AccountIdGet(context.Background(), id).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AccountIdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountIdGet`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.AccountIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the account. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountSearchAccountUserAccountGet

> InlineResponse2002 AccountSearchAccountUserAccountGet(ctx, accountUserAccount).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()

Get Account Information by Username



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
    accountUserAccount := "consequatur" // string | The username of the account
    authorization := "Bearer {YOUR_API_KEY}" // string |  (optional)
    contentType := "application/json" // string |  (optional)
    accept := "application/json" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountApi.AccountSearchAccountUserAccountGet(context.Background(), accountUserAccount).Authorization(authorization).ContentType(contentType).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.AccountSearchAccountUserAccountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountSearchAccountUserAccountGet`: InlineResponse2002
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.AccountSearchAccountUserAccountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountUserAccount** | **string** | The username of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccountSearchAccountUserAccountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **accept** | **string** |  | 

### Return type

[**InlineResponse2002**](InlineResponse2002.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

