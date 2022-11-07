# \MetadataApi

All URIs are relative to *https://api.ionos.com/databases/postgresql*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**InfosVersionGet**](MetadataApi.md#InfosVersionGet) | **Get** /infos/version | Get the current API version|
|[**InfosVersionsGet**](MetadataApi.md#InfosVersionsGet) | **Get** /infos/versions | Fetch all API versions|



## InfosVersionGet

```go
var result APIVersion = InfosVersionGet(ctx)
                      .Execute()
```

Get the current API version



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    psql "github.com/ionos-cloud/sdk-go-bundle/services/psql"
    "github.com/ionos-cloud/sdk-go-bundle/common"
)

func main() {

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := psql.NewAPIClient(configuration)
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

[**APIVersion**](APIVersion.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## InfosVersionsGet

```go
var result []APIVersion = InfosVersionsGet(ctx)
                      .Execute()
```

Fetch all API versions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    psql "github.com/ionos-cloud/sdk-go-bundle/services/psql"
    "github.com/ionos-cloud/sdk-go-bundle/common"
)

func main() {

    configuration := common.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := psql.NewAPIClient(configuration)
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

[**[]APIVersion**](APIVersion.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


