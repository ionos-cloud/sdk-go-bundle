# \PublicAccessBlockApi

All URIs are relative to *https://s3.eu-central-1.ionoscloud.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**DeletePublicAccessBlock**](PublicAccessBlockApi.md#DeletePublicAccessBlock) | **Delete** /{Bucket}?publicAccessBlock | DeletePublicAccessBlock|
|[**GetPublicAccessBlock**](PublicAccessBlockApi.md#GetPublicAccessBlock) | **Get** /{Bucket}?publicAccessBlock | GetPublicAccessBlock|
|[**PutPublicAccessBlock**](PublicAccessBlockApi.md#PutPublicAccessBlock) | **Put** /{Bucket}?publicAccessBlock | PutPublicAccessBlock|



## DeletePublicAccessBlock

```go
var result  = DeletePublicAccessBlock(ctx, bucket)
                      .Policy(policy)
                      .Execute()
```

DeletePublicAccessBlock



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    userobjectstorage "github.com/ionos-cloud/sdk-go-bundle/products/userobjectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    bucket := "bucket_example" // string | 
    policy := true // bool | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.PublicAccessBlockApi.DeletePublicAccessBlock(context.Background(), bucket).Policy(policy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicAccessBlockApi.DeletePublicAccessBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiDeletePublicAccessBlockRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **policy** | **bool** |  | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml



## GetPublicAccessBlock

```go
var result BlockPublicAccessOutput = GetPublicAccessBlock(ctx, bucket)
                      .Policy(policy)
                      .Execute()
```

GetPublicAccessBlock



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    userobjectstorage "github.com/ionos-cloud/sdk-go-bundle/products/userobjectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    bucket := "bucket_example" // string | 
    policy := true // bool | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.PublicAccessBlockApi.GetPublicAccessBlock(context.Background(), bucket).Policy(policy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicAccessBlockApi.GetPublicAccessBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetPublicAccessBlock`: BlockPublicAccessOutput
    fmt.Fprintf(os.Stdout, "Response from `PublicAccessBlockApi.GetPublicAccessBlock`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiGetPublicAccessBlockRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **policy** | **bool** |  | |

### Return type

[**BlockPublicAccessOutput**](../models/BlockPublicAccessOutput.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/jsonapplication/xml



## PutPublicAccessBlock

```go
var result  = PutPublicAccessBlock(ctx, bucket)
                      .Policy(policy)
                      .BlockPublicAccessPayload(blockPublicAccessPayload)
                      .ContentMD5(contentMD5)
                      .Execute()
```

PutPublicAccessBlock



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    userobjectstorage "github.com/ionos-cloud/sdk-go-bundle/products/userobjectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    bucket := "bucket_example" // string | 
    policy := true // bool | 
    blockPublicAccessPayload := *openapiclient.NewBlockPublicAccessPayload() // BlockPublicAccessPayload | 
    contentMD5 := "contentMD5_example" // string |  (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.PublicAccessBlockApi.PutPublicAccessBlock(context.Background(), bucket).Policy(policy).BlockPublicAccessPayload(blockPublicAccessPayload).ContentMD5(contentMD5).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PublicAccessBlockApi.PutPublicAccessBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiPutPublicAccessBlockRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **policy** | **bool** |  | |
| **blockPublicAccessPayload** | [**BlockPublicAccessPayload**](../models/BlockPublicAccessPayload.md) |  | |
| **contentMD5** | **string** |  | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: application/xml
- **Accept**: application/xml


