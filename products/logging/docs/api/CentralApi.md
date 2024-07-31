# \CentralApi

All URIs are relative to *https://logging.de-txl.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**CentralLoggingGet**](CentralApi.md#CentralLoggingGet) | **Get** /central | Gets the central logging properties.|
|[**CentralLoggingToggle**](CentralApi.md#CentralLoggingToggle) | **Put** /central | Toggles the central logging.|



## CentralLoggingGet

```go
var result CentralLoggingResponse = CentralLoggingGet(ctx)
                      .Execute()
```

Gets the central logging properties.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    logging "github.com/ionos-cloud/sdk-go-bundle/products/logging"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := logging.NewAPIClient(configuration)
    resource, resp, err := apiClient.CentralApi.CentralLoggingGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CentralApi.CentralLoggingGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `CentralLoggingGet`: CentralLoggingResponse
    fmt.Fprintf(os.Stdout, "Response from `CentralApi.CentralLoggingGet`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiCentralLoggingGetRequest struct via the builder pattern


### Return type

[**CentralLoggingResponse**](../models/CentralLoggingResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## CentralLoggingToggle

```go
var result CentralLoggingResponse = CentralLoggingToggle(ctx)
                      .CentralLoggingToggle(centralLoggingToggle)
                      .Execute()
```

Toggles the central logging.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    logging "github.com/ionos-cloud/sdk-go-bundle/products/logging"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    centralLoggingToggle := *openapiclient.NewCentralLoggingToggle() // CentralLoggingToggle | Toggle central logging.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := logging.NewAPIClient(configuration)
    resource, resp, err := apiClient.CentralApi.CentralLoggingToggle(context.Background()).CentralLoggingToggle(centralLoggingToggle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CentralApi.CentralLoggingToggle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `CentralLoggingToggle`: CentralLoggingResponse
    fmt.Fprintf(os.Stdout, "Response from `CentralApi.CentralLoggingToggle`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiCentralLoggingToggleRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **centralLoggingToggle** | [**CentralLoggingToggle**](../models/CentralLoggingToggle.md) | Toggle central logging. | |

### Return type

[**CentralLoggingResponse**](../models/CentralLoggingResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


