# \InformationApi

All URIs are relative to *https://api.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**GetInfo**](InformationApi.md#GetInfo) | **Get** /certificatemanager | Get the Service API Information|
|[**GetJsonOpenApiSpec**](InformationApi.md#GetJsonOpenApiSpec) | **Get** /certificatemanager/openapi.json | Get the Open API Documentation JSON|
|[**GetYamlOpenApiSpec**](InformationApi.md#GetYamlOpenApiSpec) | **Get** /certificatemanager/openapi.yaml | Get the Open API Documentation YAML|



## GetInfo

```go
var result ApiInfoDto = GetInfo(ctx)
                      .Execute()
```

Get the Service API Information



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    cert "github.com/ionos-cloud/sdk-go-bundle/services/cert"
)

func main() {

    configuration := cert.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cert.NewAPIClient(configuration)
    resource, resp, err := apiClient.InformationApi.GetInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationApi.GetInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetInfo`: ApiInfoDto
    fmt.Fprintf(os.Stdout, "Response from `InformationApi.GetInfo`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiGetInfoRequest struct via the builder pattern


### Return type

[**ApiInfoDto**](ApiInfoDto.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## GetJsonOpenApiSpec

```go
var result *os.File = GetJsonOpenApiSpec(ctx)
                      .Execute()
```

Get the Open API Documentation JSON



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    cert "github.com/ionos-cloud/sdk-go-bundle/services/cert"
)

func main() {

    configuration := cert.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cert.NewAPIClient(configuration)
    resource, resp, err := apiClient.InformationApi.GetJsonOpenApiSpec(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationApi.GetJsonOpenApiSpec``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetJsonOpenApiSpec`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `InformationApi.GetJsonOpenApiSpec`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiGetJsonOpenApiSpecRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## GetYamlOpenApiSpec

```go
var result *os.File = GetYamlOpenApiSpec(ctx)
                      .Execute()
```

Get the Open API Documentation YAML



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    cert "github.com/ionos-cloud/sdk-go-bundle/services/cert"
)

func main() {

    configuration := cert.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := cert.NewAPIClient(configuration)
    resource, resp, err := apiClient.InformationApi.GetYamlOpenApiSpec(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InformationApi.GetYamlOpenApiSpec``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetYamlOpenApiSpec`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `InformationApi.GetYamlOpenApiSpec`: %v\n", resource)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to an apiGetYamlOpenApiSpecRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/yaml


