# \EncryptionApi

All URIs are relative to *https://s3.eu-central-1.ionoscloud.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**DeleteBucketEncryption**](EncryptionApi.md#DeleteBucketEncryption) | **Delete** /{Bucket}?encryption | DeleteBucketEncryption|
|[**GetBucketEncryption**](EncryptionApi.md#GetBucketEncryption) | **Get** /{Bucket}?encryption | GetBucketEncryption|
|[**PutBucketEncryption**](EncryptionApi.md#PutBucketEncryption) | **Put** /{Bucket}?encryption | PutBucketEncryption|



## DeleteBucketEncryption

```go
var result  = DeleteBucketEncryption(ctx, bucket)
                      .Encryption(encryption)
                      .Execute()
```

DeleteBucketEncryption



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
    encryption := true // bool | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.EncryptionApi.DeleteBucketEncryption(context.Background(), bucket).Encryption(encryption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionApi.DeleteBucketEncryption``: %v\n", err)
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

Other parameters are passed through a pointer to an apiDeleteBucketEncryptionRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **encryption** | **bool** |  | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined



## GetBucketEncryption

```go
var result GetBucketEncryptionOutput = GetBucketEncryption(ctx, bucket)
                      .Encryption(encryption)
                      .Execute()
```

GetBucketEncryption



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
    encryption := true // bool | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.EncryptionApi.GetBucketEncryption(context.Background(), bucket).Encryption(encryption).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionApi.GetBucketEncryption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetBucketEncryption`: GetBucketEncryptionOutput
    fmt.Fprintf(os.Stdout, "Response from `EncryptionApi.GetBucketEncryption`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiGetBucketEncryptionRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **encryption** | **bool** |  | |

### Return type

[**GetBucketEncryptionOutput**](../models/GetBucketEncryptionOutput.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml



## PutBucketEncryption

```go
var result  = PutBucketEncryption(ctx, bucket)
                      .Encryption(encryption)
                      .PutBucketEncryptionRequest(putBucketEncryptionRequest)
                      .Execute()
```

PutBucketEncryption



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
    encryption := true // bool | 
    putBucketEncryptionRequest := *openapiclient.NewPutBucketEncryptionRequest(*openapiclient.NewPutBucketEncryptionRequestServerSideEncryptionConfiguration()) // PutBucketEncryptionRequest | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.EncryptionApi.PutBucketEncryption(context.Background(), bucket).Encryption(encryption).PutBucketEncryptionRequest(putBucketEncryptionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncryptionApi.PutBucketEncryption``: %v\n", err)
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

Other parameters are passed through a pointer to an apiPutBucketEncryptionRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **encryption** | **bool** |  | |
| **putBucketEncryptionRequest** | [**PutBucketEncryptionRequest**](../models/PutBucketEncryptionRequest.md) |  | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: application/xml
- **Accept**: Not defined


