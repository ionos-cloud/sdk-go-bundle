# \TargetGroupsApi

All URIs are relative to *https://api.ionos.com/cloudapi/v6*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**TargetGroupsDelete**](TargetGroupsApi.md#TargetGroupsDelete) | **Delete** /targetgroups/{targetGroupId} | Delete a Target Group by ID|
|[**TargetgroupsFindByTargetGroupId**](TargetGroupsApi.md#TargetgroupsFindByTargetGroupId) | **Get** /targetgroups/{targetGroupId} | Get a Target Group by ID|
|[**TargetgroupsGet**](TargetGroupsApi.md#TargetgroupsGet) | **Get** /targetgroups | Get Target Groups|
|[**TargetgroupsPatch**](TargetGroupsApi.md#TargetgroupsPatch) | **Patch** /targetgroups/{targetGroupId} | Partially Modify a Target Group by ID|
|[**TargetgroupsPost**](TargetGroupsApi.md#TargetgroupsPost) | **Post** /targetgroups | Create a Target Group|
|[**TargetgroupsPut**](TargetGroupsApi.md#TargetgroupsPut) | **Put** /targetgroups/{targetGroupId} | Modify a Target Group by ID|



## TargetGroupsDelete

```go
var result  = TargetGroupsDelete(ctx, targetGroupId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete a Target Group by ID



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
    targetGroupId := "targetGroupId_example" // string | The unique ID of the target group.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resp, err := apiClient.TargetGroupsApi.TargetGroupsDelete(context.Background(), targetGroupId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetGroupsApi.TargetGroupsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**targetGroupId** | **string** | The unique ID of the target group. | |

### Other Parameters

Other parameters are passed through a pointer to an apiTargetGroupsDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## TargetgroupsFindByTargetGroupId

```go
var result TargetGroup = TargetgroupsFindByTargetGroupId(ctx, targetGroupId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Get a Target Group by ID



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
    targetGroupId := "targetGroupId_example" // string | The unique ID of the target group.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.TargetGroupsApi.TargetgroupsFindByTargetGroupId(context.Background(), targetGroupId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetGroupsApi.TargetgroupsFindByTargetGroupId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `TargetgroupsFindByTargetGroupId`: TargetGroup
    fmt.Fprintf(os.Stdout, "Response from `TargetGroupsApi.TargetgroupsFindByTargetGroupId`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**targetGroupId** | **string** | The unique ID of the target group. | |

### Other Parameters

Other parameters are passed through a pointer to an apiTargetgroupsFindByTargetGroupIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**TargetGroup**](../models/TargetGroup.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## TargetgroupsGet

```go
var result TargetGroups = TargetgroupsGet(ctx)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

Get Target Groups



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
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)
    offset := int32(56) // int32 | The first element (from the complete list of the elements) to include in the response (used together with <b><i>limit</i></b> for pagination). (optional) (default to 0)
    limit := int32(56) // int32 | The maximum number of elements to return (used together with <b><i>offset</i></b> for pagination). It must not exceed <b><i>200</i></b>. (optional) (default to 100)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.TargetGroupsApi.TargetgroupsGet(context.Background()).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetGroupsApi.TargetgroupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `TargetgroupsGet`: TargetGroups
    fmt.Fprintf(os.Stdout, "Response from `TargetGroupsApi.TargetgroupsGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiTargetgroupsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |
| **offset** | **int32** | The first element (from the complete list of the elements) to include in the response (used together with &lt;b&gt;&lt;i&gt;limit&lt;/i&gt;&lt;/b&gt; for pagination). | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return (used together with &lt;b&gt;&lt;i&gt;offset&lt;/i&gt;&lt;/b&gt; for pagination). It must not exceed &lt;b&gt;&lt;i&gt;200&lt;/i&gt;&lt;/b&gt;. | [default to 100]|

### Return type

[**TargetGroups**](../models/TargetGroups.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## TargetgroupsPatch

```go
var result TargetGroup = TargetgroupsPatch(ctx, targetGroupId)
                      .TargetGroupProperties(targetGroupProperties)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Partially Modify a Target Group by ID



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
    targetGroupId := "targetGroupId_example" // string | The unique ID of the target group.
    targetGroupProperties := *openapiclient.NewTargetGroupProperties("My target group", "ROUND_ROBIN", "HTTP") // TargetGroupProperties | The target group properties to be updated.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.TargetGroupsApi.TargetgroupsPatch(context.Background(), targetGroupId).TargetGroupProperties(targetGroupProperties).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetGroupsApi.TargetgroupsPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `TargetgroupsPatch`: TargetGroup
    fmt.Fprintf(os.Stdout, "Response from `TargetGroupsApi.TargetgroupsPatch`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**targetGroupId** | **string** | The unique ID of the target group. | |

### Other Parameters

Other parameters are passed through a pointer to an apiTargetgroupsPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **targetGroupProperties** | [**TargetGroupProperties**](../models/TargetGroupProperties.md) | The target group properties to be updated. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**TargetGroup**](../models/TargetGroup.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## TargetgroupsPost

```go
var result TargetGroup = TargetgroupsPost(ctx)
                      .TargetGroup(targetGroup)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create a Target Group



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
    targetGroup := *openapiclient.NewTargetGroup(*openapiclient.NewTargetGroupProperties("My target group", "ROUND_ROBIN", "HTTP")) // TargetGroup | The target group to create.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.TargetGroupsApi.TargetgroupsPost(context.Background()).TargetGroup(targetGroup).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetGroupsApi.TargetgroupsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `TargetgroupsPost`: TargetGroup
    fmt.Fprintf(os.Stdout, "Response from `TargetGroupsApi.TargetgroupsPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiTargetgroupsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **targetGroup** | [**TargetGroup**](../models/TargetGroup.md) | The target group to create. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**TargetGroup**](../models/TargetGroup.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## TargetgroupsPut

```go
var result TargetGroup = TargetgroupsPut(ctx, targetGroupId)
                      .TargetGroup(targetGroup)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify a Target Group by ID



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
    targetGroupId := "targetGroupId_example" // string | The unique ID of the target group.
    targetGroup := *openapiclient.NewTargetGroupPut(*openapiclient.NewTargetGroupProperties("My target group", "ROUND_ROBIN", "HTTP")) // TargetGroupPut | The modified target group.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.TargetGroupsApi.TargetgroupsPut(context.Background(), targetGroupId).TargetGroup(targetGroup).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TargetGroupsApi.TargetgroupsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `TargetgroupsPut`: TargetGroup
    fmt.Fprintf(os.Stdout, "Response from `TargetGroupsApi.TargetgroupsPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**targetGroupId** | **string** | The unique ID of the target group. | |

### Other Parameters

Other parameters are passed through a pointer to an apiTargetgroupsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **targetGroup** | [**TargetGroupPut**](../models/TargetGroupPut.md) | The modified target group. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**TargetGroup**](../models/TargetGroup.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


