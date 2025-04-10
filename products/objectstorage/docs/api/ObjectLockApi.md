# \ObjectLockApi

All URIs are relative to *https://s3.eu-central-3.ionoscloud.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**GetObjectLegalHold**](ObjectLockApi.md#GetObjectLegalHold) | **Get** /{Bucket}/{Key}?legal-hold | GetObjectLegalHold|
|[**GetObjectLockConfiguration**](ObjectLockApi.md#GetObjectLockConfiguration) | **Get** /{Bucket}?object-lock | GetObjectLockConfiguration|
|[**GetObjectRetention**](ObjectLockApi.md#GetObjectRetention) | **Get** /{Bucket}/{Key}?retention | GetObjectRetention|
|[**PutObjectLegalHold**](ObjectLockApi.md#PutObjectLegalHold) | **Put** /{Bucket}/{Key}?legal-hold | PutObjectLegalHold|
|[**PutObjectLockConfiguration**](ObjectLockApi.md#PutObjectLockConfiguration) | **Put** /{Bucket}?object-lock | PutObjectLockConfiguration|
|[**PutObjectRetention**](ObjectLockApi.md#PutObjectRetention) | **Put** /{Bucket}/{Key}?retention | PutObjectRetention|



## GetObjectLegalHold

```go
var result ObjectLegalHoldConfiguration = GetObjectLegalHold(ctx, bucket, key)
                      .VersionId(versionId)
                      .Execute()
```

GetObjectLegalHold



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    objectstorage "github.com/ionos-cloud/sdk-go-bundle/objectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    bucket := "bucket_example" // string | 
    key := "key_example" // string | The key name of the object whose Legal Hold status you want to retrieve.
    versionId := "versionId_example" // string | The version ID of the object whose Legal Hold status you want to retrieve. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := objectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.ObjectLockApi.GetObjectLegalHold(context.Background(), bucket, key).VersionId(versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectLockApi.GetObjectLegalHold``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetObjectLegalHold`: ObjectLegalHoldConfiguration
    fmt.Fprintf(os.Stdout, "Response from `ObjectLockApi.GetObjectLegalHold`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** |  | |
|**key** | **string** | The key name of the object whose Legal Hold status you want to retrieve. | |

### Other Parameters

Other parameters are passed through a pointer to an apiGetObjectLegalHoldRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **versionId** | **string** | The version ID of the object whose Legal Hold status you want to retrieve. | |

### Return type

[**ObjectLegalHoldConfiguration**](../models/ObjectLegalHoldConfiguration.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ObjectLockApiService.GetObjectLegalHold"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "ObjectLockApiService.GetObjectLegalHold": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "ObjectLockApiService.GetObjectLegalHold": {
    "port": "8443",
},
})
```


## GetObjectLockConfiguration

```go
var result GetObjectLockConfigurationOutput = GetObjectLockConfiguration(ctx, bucket)
                      .Execute()
```

GetObjectLockConfiguration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    objectstorage "github.com/ionos-cloud/sdk-go-bundle/objectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    bucket := "bucket_example" // string | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := objectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.ObjectLockApi.GetObjectLockConfiguration(context.Background(), bucket).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectLockApi.GetObjectLockConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetObjectLockConfiguration`: GetObjectLockConfigurationOutput
    fmt.Fprintf(os.Stdout, "Response from `ObjectLockApi.GetObjectLockConfiguration`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiGetObjectLockConfigurationRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**GetObjectLockConfigurationOutput**](../models/GetObjectLockConfigurationOutput.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ObjectLockApiService.GetObjectLockConfiguration"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "ObjectLockApiService.GetObjectLockConfiguration": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "ObjectLockApiService.GetObjectLockConfiguration": {
    "port": "8443",
},
})
```


## GetObjectRetention

```go
var result ObjectLockRetention = GetObjectRetention(ctx, bucket, key)
                      .VersionId(versionId)
                      .Execute()
```

GetObjectRetention



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    objectstorage "github.com/ionos-cloud/sdk-go-bundle/objectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    bucket := "bucket_example" // string | 
    key := "key_example" // string | The key name of the object whose retention settings you want to retrieve.
    versionId := "versionId_example" // string | The version ID of the object whose retention settings you want to retrieve. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := objectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.ObjectLockApi.GetObjectRetention(context.Background(), bucket, key).VersionId(versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectLockApi.GetObjectRetention``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetObjectRetention`: ObjectLockRetention
    fmt.Fprintf(os.Stdout, "Response from `ObjectLockApi.GetObjectRetention`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** |  | |
|**key** | **string** | The key name of the object whose retention settings you want to retrieve. | |

### Other Parameters

Other parameters are passed through a pointer to an apiGetObjectRetentionRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **versionId** | **string** | The version ID of the object whose retention settings you want to retrieve. | |

### Return type

[**ObjectLockRetention**](../models/ObjectLockRetention.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ObjectLockApiService.GetObjectRetention"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "ObjectLockApiService.GetObjectRetention": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "ObjectLockApiService.GetObjectRetention": {
    "port": "8443",
},
})
```


## PutObjectLegalHold

```go
var result  = PutObjectLegalHold(ctx, bucket, key)
                      .ObjectLegalHoldConfiguration(objectLegalHoldConfiguration)
                      .VersionId(versionId)
                      .ContentMD5(contentMD5)
                      .Execute()
```

PutObjectLegalHold



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    objectstorage "github.com/ionos-cloud/sdk-go-bundle/objectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    bucket := "bucket_example" // string | 
    key := "key_example" // string | The key name of the object on which you want to place a Legal Hold.
    objectLegalHoldConfiguration := *openapiclient.NewObjectLegalHoldConfiguration() // ObjectLegalHoldConfiguration | 
    versionId := "versionId_example" // string | The version ID of the object on which you want to place a Legal Hold. (optional)
    contentMD5 := "contentMD5_example" // string |  (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := objectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.ObjectLockApi.PutObjectLegalHold(context.Background(), bucket, key).ObjectLegalHoldConfiguration(objectLegalHoldConfiguration).VersionId(versionId).ContentMD5(contentMD5).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectLockApi.PutObjectLegalHold``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** |  | |
|**key** | **string** | The key name of the object on which you want to place a Legal Hold. | |

### Other Parameters

Other parameters are passed through a pointer to an apiPutObjectLegalHoldRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **objectLegalHoldConfiguration** | [**ObjectLegalHoldConfiguration**](../models/ObjectLegalHoldConfiguration.md) |  | |
| **versionId** | **string** | The version ID of the object on which you want to place a Legal Hold. | |
| **contentMD5** | **string** |  | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: application/xml
- **Accept**: application/xml


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ObjectLockApiService.PutObjectLegalHold"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "ObjectLockApiService.PutObjectLegalHold": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "ObjectLockApiService.PutObjectLegalHold": {
    "port": "8443",
},
})
```


## PutObjectLockConfiguration

```go
var result  = PutObjectLockConfiguration(ctx, bucket)
                      .ContentMD5(contentMD5)
                      .PutObjectLockConfigurationRequest(putObjectLockConfigurationRequest)
                      .Execute()
```

PutObjectLockConfiguration



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    objectstorage "github.com/ionos-cloud/sdk-go-bundle/objectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    bucket := "bucket_example" // string | 
    contentMD5 := "contentMD5_example" // string | 
    putObjectLockConfigurationRequest := *openapiclient.NewPutObjectLockConfigurationRequest() // PutObjectLockConfigurationRequest | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := objectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.ObjectLockApi.PutObjectLockConfiguration(context.Background(), bucket).ContentMD5(contentMD5).PutObjectLockConfigurationRequest(putObjectLockConfigurationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectLockApi.PutObjectLockConfiguration``: %v\n", err)
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

Other parameters are passed through a pointer to an apiPutObjectLockConfigurationRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **contentMD5** | **string** |  | |
| **putObjectLockConfigurationRequest** | [**PutObjectLockConfigurationRequest**](../models/PutObjectLockConfigurationRequest.md) |  | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: application/xml
- **Accept**: application/xml


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ObjectLockApiService.PutObjectLockConfiguration"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "ObjectLockApiService.PutObjectLockConfiguration": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "ObjectLockApiService.PutObjectLockConfiguration": {
    "port": "8443",
},
})
```


## PutObjectRetention

```go
var result  = PutObjectRetention(ctx, bucket, key)
                      .PutObjectRetentionRequest(putObjectRetentionRequest)
                      .VersionId(versionId)
                      .XAmzBypassGovernanceRetention(xAmzBypassGovernanceRetention)
                      .ContentMD5(contentMD5)
                      .Execute()
```

PutObjectRetention



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    objectstorage "github.com/ionos-cloud/sdk-go-bundle/objectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    bucket := "bucket_example" // string | 
    key := "key_example" // string | The key name of the object to which you want to apply the Object Retention configuration.
    putObjectRetentionRequest := *openapiclient.NewPutObjectRetentionRequest() // PutObjectRetentionRequest | 
    versionId := "versionId_example" // string | The version ID of the object to which you want to apply the Object Retention configuration. (optional)
    xAmzBypassGovernanceRetention := true // bool | Indicates whether this operation should bypass Governance mode's restrictions. (optional) (default to false)
    contentMD5 := "contentMD5_example" // string |  (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := objectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.ObjectLockApi.PutObjectRetention(context.Background(), bucket, key).PutObjectRetentionRequest(putObjectRetentionRequest).VersionId(versionId).XAmzBypassGovernanceRetention(xAmzBypassGovernanceRetention).ContentMD5(contentMD5).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectLockApi.PutObjectRetention``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** |  | |
|**key** | **string** | The key name of the object to which you want to apply the Object Retention configuration. | |

### Other Parameters

Other parameters are passed through a pointer to an apiPutObjectRetentionRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **putObjectRetentionRequest** | [**PutObjectRetentionRequest**](../models/PutObjectRetentionRequest.md) |  | |
| **versionId** | **string** | The version ID of the object to which you want to apply the Object Retention configuration. | |
| **xAmzBypassGovernanceRetention** | **bool** | Indicates whether this operation should bypass Governance mode&#39;s restrictions. | [default to false]|
| **contentMD5** | **string** |  | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: application/xml
- **Accept**: application/xml


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ObjectLockApiService.PutObjectRetention"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "ObjectLockApiService.PutObjectRetention": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "ObjectLockApiService.PutObjectRetention": {
    "port": "8443",
},
})
```

