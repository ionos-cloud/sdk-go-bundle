# \ContractsApi

All URIs are relative to *https://api.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**GetAvailableContracts**](ContractsApi.md#GetAvailableContracts) | **Get** /activitylog/v1/contracts | List all accessible contracts|
|[**GetByContract**](ContractsApi.md#GetByContract) | **Get** /activitylog/v1/contracts/{contractNumber} | Download Activity Log entries.|



## GetAvailableContracts

```go
var result []ReferenceById = GetAvailableContracts(ctx)
                      .Execute()
```

List all accessible contracts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    activitylog "github.com/ionos-cloud/sdk-go-bundle/products/activitylog"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := activitylog.NewAPIClient(configuration)
    resource, resp, err := apiClient.ContractsApi.GetAvailableContracts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.GetAvailableContracts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetAvailableContracts`: []ReferenceById
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.GetAvailableContracts`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiGetAvailableContractsRequest struct via the builder pattern


### Return type

[**[]ReferenceById**](../models/ReferenceById.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## GetByContract

```go
var result GetByContractResponse = GetByContract(ctx, contractNumber)
                      .DateStart(dateStart)
                      .DateEnd(dateEnd)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

Download Activity Log entries.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    activitylog "github.com/ionos-cloud/sdk-go-bundle/products/activitylog"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    contractNumber := int32(56) // int32 | Contract number thats activity log entires should be downloaded.
    dateStart := "dateStart_example" // string | Start date for the Activity Log entries (inclusive) (e.g. '2021-09-23T11:43:51Z' or '2016-07-01') (optional)
    dateEnd := "dateEnd_example" // string | End date for the Activity Log entries (exclusive) (e.g. '2021-09-27T10:31:15Z' or '2016-07-30') (optional)
    offset := int32(56) // int32 | Page index ( the number of hits to skip ) (optional) (default to 0)
    limit := int32(56) // int32 | Page size ( the maximum number of hits to return ) (optional) (default to 20)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := activitylog.NewAPIClient(configuration)
    resource, resp, err := apiClient.ContractsApi.GetByContract(context.Background(), contractNumber).DateStart(dateStart).DateEnd(dateEnd).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContractsApi.GetByContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetByContract`: GetByContractResponse
    fmt.Fprintf(os.Stdout, "Response from `ContractsApi.GetByContract`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**contractNumber** | **int32** | Contract number thats activity log entires should be downloaded. | |

### Other Parameters

Other parameters are passed through a pointer to an apiGetByContractRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **dateStart** | **string** | Start date for the Activity Log entries (inclusive) (e.g. &#39;2021-09-23T11:43:51Z&#39; or &#39;2016-07-01&#39;) | |
| **dateEnd** | **string** | End date for the Activity Log entries (exclusive) (e.g. &#39;2021-09-27T10:31:15Z&#39; or &#39;2016-07-30&#39;) | |
| **offset** | **int32** | Page index ( the number of hits to skip ) | [default to 0]|
| **limit** | **int32** | Page size ( the maximum number of hits to return ) | [default to 20]|

### Return type

[**GetByContractResponse**](../models/GetByContractResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


