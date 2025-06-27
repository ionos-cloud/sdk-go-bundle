# \CentralApi

All URIs are relative to *https://logging.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**CentralFindById**](CentralApi.md#CentralFindById) | **Get** /central/{centralId} | Retrieve CentralLogging|
|[**CentralGet**](CentralApi.md#CentralGet) | **Get** /central | Retrieve all CentralLogging|
|[**CentralPut**](CentralApi.md#CentralPut) | **Put** /central/{centralId} | Ensure CentralLogging|



## CentralFindById

```go
var result CentralLoggingRead = CentralFindById(ctx, centralId)
                      .Execute()
```

Retrieve CentralLogging



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
    centralId := "40740a56-ee77-5bff-8abc-2dda26b3144f" // string | The ID (UUID) of the CentralLogging.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := logging.NewAPIClient(configuration)
    resource, resp, err := apiClient.CentralApi.CentralFindById(context.Background(), centralId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CentralApi.CentralFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `CentralFindById`: CentralLoggingRead
    fmt.Fprintf(os.Stdout, "Response from `CentralApi.CentralFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**centralId** | **string** | The ID (UUID) of the CentralLogging. | |

### Other Parameters

Other parameters are passed through a pointer to an apiCentralFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**CentralLoggingRead**](../models/CentralLoggingRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## CentralGet

```go
var result CentralLoggingReadList = CentralGet(ctx)
                      .Execute()
```

Retrieve all CentralLogging



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
    resource, resp, err := apiClient.CentralApi.CentralGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CentralApi.CentralGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `CentralGet`: CentralLoggingReadList
    fmt.Fprintf(os.Stdout, "Response from `CentralApi.CentralGet`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiCentralGetRequest struct via the builder pattern


### Return type

[**CentralLoggingReadList**](../models/CentralLoggingReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## CentralPut

```go
var result CentralLoggingRead = CentralPut(ctx, centralId)
                      .CentralLoggingEnsure(centralLoggingEnsure)
                      .Execute()
```

Ensure CentralLogging



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
    centralId := "40740a56-ee77-5bff-8abc-2dda26b3144f" // string | The ID (UUID) of the CentralLogging.
    centralLoggingEnsure := *openapiclient.NewCentralLoggingEnsure(*openapiclient.NewCentralLogging(false)) // CentralLoggingEnsure | update CentralLogging

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := logging.NewAPIClient(configuration)
    resource, resp, err := apiClient.CentralApi.CentralPut(context.Background(), centralId).CentralLoggingEnsure(centralLoggingEnsure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CentralApi.CentralPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `CentralPut`: CentralLoggingRead
    fmt.Fprintf(os.Stdout, "Response from `CentralApi.CentralPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**centralId** | **string** | The ID (UUID) of the CentralLogging. | |

### Other Parameters

Other parameters are passed through a pointer to an apiCentralPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **centralLoggingEnsure** | [**CentralLoggingEnsure**](../models/CentralLoggingEnsure.md) | update CentralLogging | |

### Return type

[**CentralLoggingRead**](../models/CentralLoggingRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


