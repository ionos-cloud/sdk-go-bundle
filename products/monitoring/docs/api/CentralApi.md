# \CentralApi

All URIs are relative to *https://monitoring.de-fra.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**CentralFindById**](CentralApi.md#CentralFindById) | **Get** /central/{centralId} | Retrieve CentralMonitoring|
|[**CentralGet**](CentralApi.md#CentralGet) | **Get** /central | Retrieve all CentralMonitoring|
|[**CentralPut**](CentralApi.md#CentralPut) | **Put** /central/{centralId} | Ensure CentralMonitoring|



## CentralFindById

```go
var result CentralMonitoringRead = CentralFindById(ctx, centralId)
                      .Execute()
```

Retrieve CentralMonitoring



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    monitoring "github.com/ionos-cloud/sdk-go-bundle/products/monitoring"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    centralId := "23e6183d-6ab3-54de-b165-9b44c3589f4f" // string | The ID (UUID) of the CentralMonitoring.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := monitoring.NewAPIClient(configuration)
    resource, resp, err := apiClient.CentralApi.CentralFindById(context.Background(), centralId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CentralApi.CentralFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `CentralFindById`: CentralMonitoringRead
    fmt.Fprintf(os.Stdout, "Response from `CentralApi.CentralFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**centralId** | **string** | The ID (UUID) of the CentralMonitoring. | |

### Other Parameters

Other parameters are passed through a pointer to an apiCentralFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**CentralMonitoringRead**](../models/CentralMonitoringRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## CentralGet

```go
var result CentralMonitoringReadList = CentralGet(ctx)
                      .Execute()
```

Retrieve all CentralMonitoring



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    monitoring "github.com/ionos-cloud/sdk-go-bundle/products/monitoring"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := monitoring.NewAPIClient(configuration)
    resource, resp, err := apiClient.CentralApi.CentralGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CentralApi.CentralGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `CentralGet`: CentralMonitoringReadList
    fmt.Fprintf(os.Stdout, "Response from `CentralApi.CentralGet`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiCentralGetRequest struct via the builder pattern


### Return type

[**CentralMonitoringReadList**](../models/CentralMonitoringReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## CentralPut

```go
var result CentralMonitoringRead = CentralPut(ctx, centralId)
                      .CentralMonitoringEnsure(centralMonitoringEnsure)
                      .Execute()
```

Ensure CentralMonitoring



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    monitoring "github.com/ionos-cloud/sdk-go-bundle/products/monitoring"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    centralId := "23e6183d-6ab3-54de-b165-9b44c3589f4f" // string | The ID (UUID) of the CentralMonitoring.
    centralMonitoringEnsure := *openapiclient.NewCentralMonitoringEnsure(*openapiclient.NewCentralMonitoring(false)) // CentralMonitoringEnsure | update CentralMonitoring

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := monitoring.NewAPIClient(configuration)
    resource, resp, err := apiClient.CentralApi.CentralPut(context.Background(), centralId).CentralMonitoringEnsure(centralMonitoringEnsure).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CentralApi.CentralPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `CentralPut`: CentralMonitoringRead
    fmt.Fprintf(os.Stdout, "Response from `CentralApi.CentralPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**centralId** | **string** | The ID (UUID) of the CentralMonitoring. | |

### Other Parameters

Other parameters are passed through a pointer to an apiCentralPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **centralMonitoringEnsure** | [**CentralMonitoringEnsure**](../models/CentralMonitoringEnsure.md) | update CentralMonitoring | |

### Return type

[**CentralMonitoringRead**](../models/CentralMonitoringRead.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


