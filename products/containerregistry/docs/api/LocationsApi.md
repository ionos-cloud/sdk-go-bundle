# \LocationsApi

All URIs are relative to *https://api.ionos.com/containerregistries*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**LocationsGet**](LocationsApi.md#LocationsGet) | **Get** /locations | Get container registry locations|



## LocationsGet

```go
var result LocationsResponse = LocationsGet(ctx)
                      .Execute()
```

Get container registry locations

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

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := containerregistry.NewAPIClient(configuration)
    resource, resp, err := apiClient.LocationsApi.LocationsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.LocationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `LocationsGet`: LocationsResponse
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.LocationsGet`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiLocationsGetRequest struct via the builder pattern


### Return type

[**LocationsResponse**](../models/LocationsResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


