# \TaggingApi

All URIs are relative to *https://s3.eu-central-1.ionoscloud.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**DeleteBucketTagging**](TaggingApi.md#DeleteBucketTagging) | **Delete** /{Bucket}?tagging | DeleteBucketTagging|
|[**DeleteObjectTagging**](TaggingApi.md#DeleteObjectTagging) | **Delete** /{Bucket}/{Key}?tagging | DeleteObjectTagging|
|[**GetBucketTagging**](TaggingApi.md#GetBucketTagging) | **Get** /{Bucket}?tagging | GetBucketTagging|
|[**GetObjectTagging**](TaggingApi.md#GetObjectTagging) | **Get** /{Bucket}/{Key}?tagging | GetObjectTagging|
|[**PutBucketTagging**](TaggingApi.md#PutBucketTagging) | **Put** /{Bucket}?tagging | PutBucketTagging|
|[**PutObjectTagging**](TaggingApi.md#PutObjectTagging) | **Put** /{Bucket}/{Key}?tagging | PutObjectTagging|



## DeleteBucketTagging

```go
var result  = DeleteBucketTagging(ctx, bucket)
                      .Tagging(tagging)
                      .Execute()
```

DeleteBucketTagging



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
    tagging := true // bool | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.TaggingApi.DeleteBucketTagging(context.Background(), bucket).Tagging(tagging).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggingApi.DeleteBucketTagging``: %v\n", err)
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

Other parameters are passed through a pointer to an apiDeleteBucketTaggingRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **tagging** | **bool** |  | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined



## DeleteObjectTagging

```go
var result map[string]interface{} = DeleteObjectTagging(ctx, bucket, key)
                      .Tagging(tagging)
                      .VersionId(versionId)
                      .Execute()
```

DeleteObjectTagging



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
    key := "key_example" // string | The key that identifies the object in the bucket from which to remove all tags.
    tagging := true // bool | 
    versionId := "versionId_example" // string | The versionId of the object that the tag-set will be removed from. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.TaggingApi.DeleteObjectTagging(context.Background(), bucket, key).Tagging(tagging).VersionId(versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggingApi.DeleteObjectTagging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DeleteObjectTagging`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TaggingApi.DeleteObjectTagging`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** |  | |
|**key** | **string** | The key that identifies the object in the bucket from which to remove all tags. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDeleteObjectTaggingRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **tagging** | **bool** |  | |
| **versionId** | **string** | The versionId of the object that the tag-set will be removed from. | |

### Return type

**map[string]interface{}**

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml



## GetBucketTagging

```go
var result GetBucketTaggingOutput = GetBucketTagging(ctx, bucket)
                      .Tagging(tagging)
                      .Execute()
```

GetBucketTagging



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
    tagging := true // bool | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.TaggingApi.GetBucketTagging(context.Background(), bucket).Tagging(tagging).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggingApi.GetBucketTagging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetBucketTagging`: GetBucketTaggingOutput
    fmt.Fprintf(os.Stdout, "Response from `TaggingApi.GetBucketTagging`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiGetBucketTaggingRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **tagging** | **bool** |  | |

### Return type

[**GetBucketTaggingOutput**](../models/GetBucketTaggingOutput.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml



## GetObjectTagging

```go
var result GetObjectTaggingOutput = GetObjectTagging(ctx, bucket, key)
                      .Tagging(tagging)
                      .VersionId(versionId)
                      .Execute()
```

GetObjectTagging



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
    key := "key_example" // string | Object key for which to get the tagging information.
    tagging := true // bool | 
    versionId := "versionId_example" // string | The versionId of the object for which to get the tagging information. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.TaggingApi.GetObjectTagging(context.Background(), bucket, key).Tagging(tagging).VersionId(versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggingApi.GetObjectTagging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetObjectTagging`: GetObjectTaggingOutput
    fmt.Fprintf(os.Stdout, "Response from `TaggingApi.GetObjectTagging`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** |  | |
|**key** | **string** | Object key for which to get the tagging information. | |

### Other Parameters

Other parameters are passed through a pointer to an apiGetObjectTaggingRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **tagging** | **bool** |  | |
| **versionId** | **string** | The versionId of the object for which to get the tagging information. | |

### Return type

[**GetObjectTaggingOutput**](../models/GetObjectTaggingOutput.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml



## PutBucketTagging

```go
var result  = PutBucketTagging(ctx, bucket)
                      .Tagging(tagging)
                      .PutBucketTaggingRequest(putBucketTaggingRequest)
                      .ContentMD5(contentMD5)
                      .Execute()
```

PutBucketTagging



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
    tagging := true // bool | 
    putBucketTaggingRequest := *openapiclient.NewPutBucketTaggingRequest(*openapiclient.NewPutBucketTaggingRequestTagging()) // PutBucketTaggingRequest | 
    contentMD5 := "contentMD5_example" // string |  (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.TaggingApi.PutBucketTagging(context.Background(), bucket).Tagging(tagging).PutBucketTaggingRequest(putBucketTaggingRequest).ContentMD5(contentMD5).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggingApi.PutBucketTagging``: %v\n", err)
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

Other parameters are passed through a pointer to an apiPutBucketTaggingRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **tagging** | **bool** |  | |
| **putBucketTaggingRequest** | [**PutBucketTaggingRequest**](../models/PutBucketTaggingRequest.md) |  | |
| **contentMD5** | **string** |  | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: application/xml
- **Accept**: Not defined



## PutObjectTagging

```go
var result map[string]interface{} = PutObjectTagging(ctx, bucket, key)
                      .Tagging(tagging)
                      .PutBucketTaggingRequest(putBucketTaggingRequest)
                      .VersionId(versionId)
                      .ContentMD5(contentMD5)
                      .Execute()
```

PutObjectTagging



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
    key := "key_example" // string | Name of the object key.
    tagging := true // bool | 
    putBucketTaggingRequest := *openapiclient.NewPutBucketTaggingRequest(*openapiclient.NewPutBucketTaggingRequestTagging()) // PutBucketTaggingRequest | 
    versionId := "versionId_example" // string | The versionId of the object that the tag-set will be added to. (optional)
    contentMD5 := "contentMD5_example" // string |  (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.TaggingApi.PutObjectTagging(context.Background(), bucket, key).Tagging(tagging).PutBucketTaggingRequest(putBucketTaggingRequest).VersionId(versionId).ContentMD5(contentMD5).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaggingApi.PutObjectTagging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `PutObjectTagging`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `TaggingApi.PutObjectTagging`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** |  | |
|**key** | **string** | Name of the object key. | |

### Other Parameters

Other parameters are passed through a pointer to an apiPutObjectTaggingRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **tagging** | **bool** |  | |
| **putBucketTaggingRequest** | [**PutBucketTaggingRequest**](../models/PutBucketTaggingRequest.md) |  | |
| **versionId** | **string** | The versionId of the object that the tag-set will be added to. | |
| **contentMD5** | **string** |  | |

### Return type

**map[string]interface{}**

### HTTP request headers

- **Content-Type**: application/xml
- **Accept**: application/xml


