# \MetadataApi

All URIs are relative to *https://api.ionos.com/databases/mongodb*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**InfosVersionGet**](MetadataApi.md#InfosVersionGet) | **Get** /infos/version | Get API Version|
|[**InfosVersionsGet**](MetadataApi.md#InfosVersionsGet) | **Get** /infos/versions | Get All API Versions|



## InfosVersionGet

```go
var result APIVersion = InfosVersionGet(ctx)
                      .Execute()
```

Get API Version



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    mongo "github.com/ionos-cloud/sdk-go-bundle/products/dbaas/mongo"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := mongo.NewAPIClient(configuration)
    resource, resp, err := apiClient.MetadataApi.InfosVersionGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.InfosVersionGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `InfosVersionGet`: APIVersion
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.InfosVersionGet`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiInfosVersionGetRequest struct via the builder pattern


### Return type

[**APIVersion**](../models/APIVersion.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## InfosVersionsGet

```go
var result []APIVersion = InfosVersionsGet(ctx)
                      .Execute()
```

Get All API Versions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    mongo "github.com/ionos-cloud/sdk-go-bundle/products/dbaas/mongo"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := mongo.NewAPIClient(configuration)
    resource, resp, err := apiClient.MetadataApi.InfosVersionsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.InfosVersionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `InfosVersionsGet`: []APIVersion
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.InfosVersionsGet`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiInfosVersionsGetRequest struct via the builder pattern


### Return type

[**[]APIVersion**](../models/APIVersion.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


