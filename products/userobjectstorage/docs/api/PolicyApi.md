# \PolicyApi

All URIs are relative to *https://s3.eu-central-1.ionoscloud.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**DeleteBucketPolicy**](PolicyApi.md#DeleteBucketPolicy) | **Delete** /{Bucket}?policy | DeleteBucketPolicy|
|[**GetBucketPolicy**](PolicyApi.md#GetBucketPolicy) | **Get** /{Bucket}?policy | GetBucketPolicy|
|[**GetBucketPolicyStatus**](PolicyApi.md#GetBucketPolicyStatus) | **Get** /{Bucket}?policyStatus | GetBucketPolicyStatus|
|[**PutBucketPolicy**](PolicyApi.md#PutBucketPolicy) | **Put** /{Bucket}?policy | PutBucketPolicy|



## DeleteBucketPolicy

```go
var result  = DeleteBucketPolicy(ctx, bucket)
                      .Policy(policy)
                      .Execute()
```

DeleteBucketPolicy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    userobjectstorage "github.com/ionos-cloud/sdk-go-bundle/userobjectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    bucket := "bucket_example" // string | 
    policy := true // bool | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.PolicyApi.DeleteBucketPolicy(context.Background(), bucket).Policy(policy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.DeleteBucketPolicy``: %v\n", err)
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

Other parameters are passed through a pointer to an apiDeleteBucketPolicyRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **policy** | **bool** |  | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"PolicyApiService.DeleteBucketPolicy"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "PolicyApiService.DeleteBucketPolicy": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "PolicyApiService.DeleteBucketPolicy": {
    "port": "8443",
},
})
```


## GetBucketPolicy

```go
var result BucketPolicy = GetBucketPolicy(ctx, bucket)
                      .Policy(policy)
                      .Execute()
```

GetBucketPolicy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    userobjectstorage "github.com/ionos-cloud/sdk-go-bundle/userobjectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    bucket := "bucket_example" // string | 
    policy := true // bool | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.PolicyApi.GetBucketPolicy(context.Background(), bucket).Policy(policy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.GetBucketPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetBucketPolicy`: BucketPolicy
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.GetBucketPolicy`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiGetBucketPolicyRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **policy** | **bool** |  | |

### Return type

[**BucketPolicy**](../models/BucketPolicy.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/jsonapplication/xml


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"PolicyApiService.GetBucketPolicy"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "PolicyApiService.GetBucketPolicy": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "PolicyApiService.GetBucketPolicy": {
    "port": "8443",
},
})
```


## GetBucketPolicyStatus

```go
var result GetBucketPolicyStatusOutput = GetBucketPolicyStatus(ctx, bucket)
                      .PolicyStatus(policyStatus)
                      .Execute()
```

GetBucketPolicyStatus



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    userobjectstorage "github.com/ionos-cloud/sdk-go-bundle/userobjectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    bucket := "bucket_example" // string | 
    policyStatus := true // bool | 

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.PolicyApi.GetBucketPolicyStatus(context.Background(), bucket).PolicyStatus(policyStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.GetBucketPolicyStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `GetBucketPolicyStatus`: GetBucketPolicyStatusOutput
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.GetBucketPolicyStatus`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**bucket** | **string** |  | |

### Other Parameters

Other parameters are passed through a pointer to an apiGetBucketPolicyStatusRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **policyStatus** | **bool** |  | |

### Return type

[**GetBucketPolicyStatusOutput**](../models/GetBucketPolicyStatusOutput.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"PolicyApiService.GetBucketPolicyStatus"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "PolicyApiService.GetBucketPolicyStatus": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "PolicyApiService.GetBucketPolicyStatus": {
    "port": "8443",
},
})
```


## PutBucketPolicy

```go
var result  = PutBucketPolicy(ctx, bucket)
                      .Policy(policy)
                      .BucketPolicy(bucketPolicy)
                      .ContentMD5(contentMD5)
                      .Execute()
```

PutBucketPolicy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    userobjectstorage "github.com/ionos-cloud/sdk-go-bundle/userobjectstorage"
    "github.com/ionos-cloud/sdk-go-bundle/shared"
)

func main() {
    bucket := "bucket_example" // string | 
    policy := true // bool | 
    bucketPolicy := *openapiclient.NewBucketPolicy([]openapiclient.BucketPolicyStatement{*openapiclient.NewBucketPolicyStatement([]string{"Action_example"}, "Effect_example", []string{"Resource_example"}, *openapiclient.NewBucketPolicyStatementPrincipal())}) // BucketPolicy | 
    contentMD5 := "contentMD5_example" // string |  (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := userobjectstorage.NewAPIClient(configuration)
    resource, resp, err := apiClient.PolicyApi.PutBucketPolicy(context.Background(), bucket).Policy(policy).BucketPolicy(bucketPolicy).ContentMD5(contentMD5).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.PutBucketPolicy``: %v\n", err)
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

Other parameters are passed through a pointer to an apiPutBucketPolicyRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **policy** | **bool** |  | |
| **bucketPolicy** | [**BucketPolicy**](../models/BucketPolicy.md) |  | |
| **contentMD5** | **string** |  | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/xml


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"PolicyApiService.PutBucketPolicy"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), shared.ContextOperationServerIndices, map[string]int{
    "PolicyApiService.PutBucketPolicy": 2,
})
ctx = context.WithValue(context.Background(), shared.ContextOperationServerVariables, map[string]map[string]string{
    "PolicyApiService.PutBucketPolicy": {
    "port": "8443",
},
})
```

