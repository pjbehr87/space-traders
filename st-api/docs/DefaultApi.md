# \DefaultApi

All URIs are relative to *https://api.spacetraders.io/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStatus**](DefaultApi.md#GetStatus) | **Get** / | Get Status
[**Register**](DefaultApi.md#Register) | **Post** /register | Register New Agent



## GetStatus

> GetStatus200Response GetStatus(ctx).Execute()

Get Status



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pjbehr87/space-traders/stapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatus`: GetStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusRequest struct via the builder pattern


### Return type

[**GetStatus200Response**](GetStatus200Response.md)

### Authorization

[AgentToken](../README.md#AgentToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Register

> Register201Response Register(ctx).RegisterRequest(registerRequest).Execute()

Register New Agent



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/pjbehr87/space-traders/stapi"
)

func main() {
    registerRequest := *openapiclient.NewRegisterRequest(openapiclient.FactionSymbols("COSMIC"), "BADGER") // RegisterRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.Register(context.Background()).RegisterRequest(registerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.Register``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Register`: Register201Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.Register`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerRequest** | [**RegisterRequest**](RegisterRequest.md) |  | 

### Return type

[**Register201Response**](Register201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

