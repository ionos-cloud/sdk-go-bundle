# \InstanceTemplatesApi

All URIs are relative to *https://mongodb.gb-bhx.ionos.com/v2*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**InstancetemplatesFindById**](InstanceTemplatesApi.md#InstancetemplatesFindById) | **Get** /instance-templates/{templateId} | Retrieve InstanceTemplate|
|[**InstancetemplatesGet**](InstanceTemplatesApi.md#InstancetemplatesGet) | **Get** /instance-templates | Retrieve all InstanceTemplates|



## InstancetemplatesFindById

```go
var result InstanceTemplateRead = InstancetemplatesFindById(ctx, templateId)
                      .Execute()
```

Retrieve InstanceTemplate



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    mongo "github.com/ionos-cloud/sdk-go-bundle/products/mongo"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    templateId := "591b4c04-524d-5f76-a466-3f369755d0ce" // string | The ID (UUID) of the InstanceTemplate.

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := mongo.NewAPIClient(configuration)
    resource, resp, err := apiClient.InstanceTemplatesApi.InstancetemplatesFindById(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceTemplatesApi.InstancetemplatesFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `InstancetemplatesFindById`: InstanceTemplateRead
    fmt.Fprintf(os.Stdout, "Response from `InstanceTemplatesApi.InstancetemplatesFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**templateId** | **string** | The ID (UUID) of the InstanceTemplate. | |

### Other Parameters

Other parameters are passed through a pointer to an apiInstancetemplatesFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**InstanceTemplateRead**](../models/InstanceTemplateRead.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## InstancetemplatesGet

```go
var result InstanceTemplateReadList = InstancetemplatesGet(ctx)
                      .Offset(offset)
                      .Limit(limit)
                      .FilterEdition(filterEdition)
                      .Execute()
```

Retrieve all InstanceTemplates



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    mongo "github.com/ionos-cloud/sdk-go-bundle/products/mongo"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    offset := int32(0) // int32 | The first element (of the total list of elements) to include in the response. Use this parameter together with the limit for pagination. (optional) (default to 0)
    limit := int32(100) // int32 | The maximum number of elements to return. Use this parameter together with the offset for pagination. (optional) (default to 100)
    filterEdition := openapiclient.MongoEdition("ENTERPRISE") // MongoEdition | Filters resources by edition. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := mongo.NewAPIClient(configuration)
    resource, resp, err := apiClient.InstanceTemplatesApi.InstancetemplatesGet(context.Background()).Offset(offset).Limit(limit).FilterEdition(filterEdition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InstanceTemplatesApi.InstancetemplatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `InstancetemplatesGet`: InstanceTemplateReadList
    fmt.Fprintf(os.Stdout, "Response from `InstanceTemplatesApi.InstancetemplatesGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiInstancetemplatesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **offset** | **int32** | The first element (of the total list of elements) to include in the response. Use this parameter together with the limit for pagination. | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return. Use this parameter together with the offset for pagination. | [default to 100]|
| **filterEdition** | [**MongoEdition**](../models/.md) | Filters resources by edition. | |

### Return type

[**InstanceTemplateReadList**](../models/InstanceTemplateReadList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


