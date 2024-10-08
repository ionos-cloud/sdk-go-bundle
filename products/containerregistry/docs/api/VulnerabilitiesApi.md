# \VulnerabilitiesApi

All URIs are relative to *https://api.ionos.com/containerregistries*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**VulnerabilitiesFindByID**](VulnerabilitiesApi.md#VulnerabilitiesFindByID) | **Get** /vulnerabilities/{vulnerabilityId} | Retrieve Vulnerability|



## VulnerabilitiesFindByID

```go
var result VulnerabilityRead = VulnerabilitiesFindByID(ctx, vulnerabilityId)
                      .Execute()
```

Retrieve Vulnerability



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    containerregistry "github.com/ionos-cloud/sdk-go-bundle/products/containerregistry"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    vulnerabilityId := "CVE-2019-1234" // string | The ID of the Vulnerability.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resource, resp, err := apiClient.VulnerabilitiesApi.VulnerabilitiesFindByID(context.Background(), vulnerabilityId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VulnerabilitiesApi.VulnerabilitiesFindByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `VulnerabilitiesFindByID`: VulnerabilityRead
    fmt.Fprintf(os.Stdout, "Response from `VulnerabilitiesApi.VulnerabilitiesFindByID`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**vulnerabilityId** | **string** | The ID of the Vulnerability. | |

### Other Parameters

Other parameters are passed through a pointer to an apiVulnerabilitiesFindByIDRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**VulnerabilityRead**](../models/VulnerabilityRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


