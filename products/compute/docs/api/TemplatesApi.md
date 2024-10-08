# \TemplatesApi

All URIs are relative to *https://api.ionos.com/cloudapi/v6*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**TemplatesFindById**](TemplatesApi.md#TemplatesFindById) | **Get** /templates/{templateId} | Get Cubes Template by ID|
|[**TemplatesGet**](TemplatesApi.md#TemplatesGet) | **Get** /templates | Get Cubes Templates|



## TemplatesFindById

```go
var result Template = TemplatesFindById(ctx, templateId)
                      .Depth(depth)
                      .Execute()
```

Get Cubes Template by ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    templateId := "templateId_example" // string | The unique template ID.
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.TemplatesApi.TemplatesFindById(context.Background(), templateId).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `TemplatesFindById`: Template
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**templateId** | **string** | The unique template ID. | |

### Other Parameters

Other parameters are passed through a pointer to an apiTemplatesFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|

### Return type

[**Template**](../models/Template.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## TemplatesGet

```go
var result Templates = TemplatesGet(ctx)
                      .Depth(depth)
                      .Execute()
```

Get Cubes Templates



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    compute "github.com/ionos-cloud/sdk-go-bundle/products/compute"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.TemplatesApi.TemplatesGet(context.Background()).Depth(depth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TemplatesApi.TemplatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `TemplatesGet`: Templates
    fmt.Fprintf(os.Stdout, "Response from `TemplatesApi.TemplatesGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiTemplatesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|

### Return type

[**Templates**](../models/Templates.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


