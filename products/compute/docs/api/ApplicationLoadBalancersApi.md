# \ApplicationLoadBalancersApi

All URIs are relative to *https://api.ionos.com/cloudapi/v6*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**DatacentersApplicationloadbalancersDelete**](ApplicationLoadBalancersApi.md#DatacentersApplicationloadbalancersDelete) | **Delete** /datacenters/{datacenterId}/applicationloadbalancers/{applicationLoadBalancerId} | Delete an Application Load Balancer by ID|
|[**DatacentersApplicationloadbalancersFindByApplicationLoadBalancerId**](ApplicationLoadBalancersApi.md#DatacentersApplicationloadbalancersFindByApplicationLoadBalancerId) | **Get** /datacenters/{datacenterId}/applicationloadbalancers/{applicationLoadBalancerId} | Get an Application Load Balancer by ID|
|[**DatacentersApplicationloadbalancersFlowlogsDelete**](ApplicationLoadBalancersApi.md#DatacentersApplicationloadbalancersFlowlogsDelete) | **Delete** /datacenters/{datacenterId}/applicationloadbalancers/{applicationLoadBalancerId}/flowlogs/{flowLogId} | Delete an ALB Flow Log by ID|
|[**DatacentersApplicationloadbalancersFlowlogsFindByFlowLogId**](ApplicationLoadBalancersApi.md#DatacentersApplicationloadbalancersFlowlogsFindByFlowLogId) | **Get** /datacenters/{datacenterId}/applicationloadbalancers/{applicationLoadBalancerId}/flowlogs/{flowLogId} | Get an ALB Flow Log by ID|
|[**DatacentersApplicationloadbalancersFlowlogsGet**](ApplicationLoadBalancersApi.md#DatacentersApplicationloadbalancersFlowlogsGet) | **Get** /datacenters/{datacenterId}/applicationloadbalancers/{applicationLoadBalancerId}/flowlogs | Get ALB Flow Logs|
|[**DatacentersApplicationloadbalancersFlowlogsPatch**](ApplicationLoadBalancersApi.md#DatacentersApplicationloadbalancersFlowlogsPatch) | **Patch** /datacenters/{datacenterId}/applicationloadbalancers/{applicationLoadBalancerId}/flowlogs/{flowLogId} | Partially Modify an ALB Flow Log by ID|
|[**DatacentersApplicationloadbalancersFlowlogsPost**](ApplicationLoadBalancersApi.md#DatacentersApplicationloadbalancersFlowlogsPost) | **Post** /datacenters/{datacenterId}/applicationloadbalancers/{applicationLoadBalancerId}/flowlogs | Create an ALB Flow Log|
|[**DatacentersApplicationloadbalancersFlowlogsPut**](ApplicationLoadBalancersApi.md#DatacentersApplicationloadbalancersFlowlogsPut) | **Put** /datacenters/{datacenterId}/applicationloadbalancers/{applicationLoadBalancerId}/flowlogs/{flowLogId} | Modify an ALB Flow Log by ID|
|[**DatacentersApplicationloadbalancersForwardingrulesDelete**](ApplicationLoadBalancersApi.md#DatacentersApplicationloadbalancersForwardingrulesDelete) | **Delete** /datacenters/{datacenterId}/applicationloadbalancers/{applicationLoadBalancerId}/forwardingrules/{forwardingRuleId} | Delete an ALB Forwarding Rule by ID|
|[**DatacentersApplicationloadbalancersForwardingrulesFindByForwardingRuleId**](ApplicationLoadBalancersApi.md#DatacentersApplicationloadbalancersForwardingrulesFindByForwardingRuleId) | **Get** /datacenters/{datacenterId}/applicationloadbalancers/{applicationLoadBalancerId}/forwardingrules/{forwardingRuleId} | Get an ALB Forwarding Rule by ID|
|[**DatacentersApplicationloadbalancersForwardingrulesGet**](ApplicationLoadBalancersApi.md#DatacentersApplicationloadbalancersForwardingrulesGet) | **Get** /datacenters/{datacenterId}/applicationloadbalancers/{applicationLoadBalancerId}/forwardingrules | Get ALB Forwarding Rules|
|[**DatacentersApplicationloadbalancersForwardingrulesPatch**](ApplicationLoadBalancersApi.md#DatacentersApplicationloadbalancersForwardingrulesPatch) | **Patch** /datacenters/{datacenterId}/applicationloadbalancers/{applicationLoadBalancerId}/forwardingrules/{forwardingRuleId} | Partially modify an ALB Forwarding Rule by ID|
|[**DatacentersApplicationloadbalancersForwardingrulesPost**](ApplicationLoadBalancersApi.md#DatacentersApplicationloadbalancersForwardingrulesPost) | **Post** /datacenters/{datacenterId}/applicationloadbalancers/{applicationLoadBalancerId}/forwardingrules | Create an ALB Forwarding Rule|
|[**DatacentersApplicationloadbalancersForwardingrulesPut**](ApplicationLoadBalancersApi.md#DatacentersApplicationloadbalancersForwardingrulesPut) | **Put** /datacenters/{datacenterId}/applicationloadbalancers/{applicationLoadBalancerId}/forwardingrules/{forwardingRuleId} | Modify an ALB Forwarding Rule by ID|
|[**DatacentersApplicationloadbalancersGet**](ApplicationLoadBalancersApi.md#DatacentersApplicationloadbalancersGet) | **Get** /datacenters/{datacenterId}/applicationloadbalancers | Get Application Load Balancers|
|[**DatacentersApplicationloadbalancersPatch**](ApplicationLoadBalancersApi.md#DatacentersApplicationloadbalancersPatch) | **Patch** /datacenters/{datacenterId}/applicationloadbalancers/{applicationLoadBalancerId} | Partially Modify an Application Load Balancer by ID|
|[**DatacentersApplicationloadbalancersPost**](ApplicationLoadBalancersApi.md#DatacentersApplicationloadbalancersPost) | **Post** /datacenters/{datacenterId}/applicationloadbalancers | Create an Application Load Balancer|
|[**DatacentersApplicationloadbalancersPut**](ApplicationLoadBalancersApi.md#DatacentersApplicationloadbalancersPut) | **Put** /datacenters/{datacenterId}/applicationloadbalancers/{applicationLoadBalancerId} | Modify an Application Load Balancer by ID|



## DatacentersApplicationloadbalancersDelete

```go
var result  = DatacentersApplicationloadbalancersDelete(ctx, datacenterId, applicationLoadBalancerId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete an Application Load Balancer by ID



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    applicationLoadBalancerId := "applicationLoadBalancerId_example" // string | The unique ID of the Application Load Balancer.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resp, err := apiClient.ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersDelete(context.Background(), datacenterId, applicationLoadBalancerId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**applicationLoadBalancerId** | **string** | The unique ID of the Application Load Balancer. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersApplicationloadbalancersDeleteRequest struct via the builder pattern


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



## DatacentersApplicationloadbalancersFindByApplicationLoadBalancerId

```go
var result ApplicationLoadBalancer = DatacentersApplicationloadbalancersFindByApplicationLoadBalancerId(ctx, datacenterId, applicationLoadBalancerId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Get an Application Load Balancer by ID



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    applicationLoadBalancerId := "applicationLoadBalancerId_example" // string | The unique ID of the Application Load Balancer.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersFindByApplicationLoadBalancerId(context.Background(), datacenterId, applicationLoadBalancerId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersFindByApplicationLoadBalancerId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersApplicationloadbalancersFindByApplicationLoadBalancerId`: ApplicationLoadBalancer
    fmt.Fprintf(os.Stdout, "Response from `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersFindByApplicationLoadBalancerId`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**applicationLoadBalancerId** | **string** | The unique ID of the Application Load Balancer. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersApplicationloadbalancersFindByApplicationLoadBalancerIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**ApplicationLoadBalancer**](../models/ApplicationLoadBalancer.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersApplicationloadbalancersFlowlogsDelete

```go
var result  = DatacentersApplicationloadbalancersFlowlogsDelete(ctx, datacenterId, applicationLoadBalancerId, flowLogId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete an ALB Flow Log by ID



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    applicationLoadBalancerId := "applicationLoadBalancerId_example" // string | The unique ID of the Application Load Balancer.
    flowLogId := "flowLogId_example" // string | The unique ID of the flow log.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resp, err := apiClient.ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersFlowlogsDelete(context.Background(), datacenterId, applicationLoadBalancerId, flowLogId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersFlowlogsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**applicationLoadBalancerId** | **string** | The unique ID of the Application Load Balancer. | |
|**flowLogId** | **string** | The unique ID of the flow log. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersApplicationloadbalancersFlowlogsDeleteRequest struct via the builder pattern


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



## DatacentersApplicationloadbalancersFlowlogsFindByFlowLogId

```go
var result FlowLog = DatacentersApplicationloadbalancersFlowlogsFindByFlowLogId(ctx, datacenterId, applicationLoadBalancerId, flowLogId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Get an ALB Flow Log by ID



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    applicationLoadBalancerId := "applicationLoadBalancerId_example" // string | The unique ID of the Application Load Balancer.
    flowLogId := "flowLogId_example" // string | The unique ID of the flow log.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersFlowlogsFindByFlowLogId(context.Background(), datacenterId, applicationLoadBalancerId, flowLogId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersFlowlogsFindByFlowLogId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersApplicationloadbalancersFlowlogsFindByFlowLogId`: FlowLog
    fmt.Fprintf(os.Stdout, "Response from `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersFlowlogsFindByFlowLogId`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**applicationLoadBalancerId** | **string** | The unique ID of the Application Load Balancer. | |
|**flowLogId** | **string** | The unique ID of the flow log. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersApplicationloadbalancersFlowlogsFindByFlowLogIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**FlowLog**](../models/FlowLog.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersApplicationloadbalancersFlowlogsGet

```go
var result FlowLogs = DatacentersApplicationloadbalancersFlowlogsGet(ctx, datacenterId, applicationLoadBalancerId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Get ALB Flow Logs



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    applicationLoadBalancerId := "applicationLoadBalancerId_example" // string | The unique ID of the Application Load Balancer.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersFlowlogsGet(context.Background(), datacenterId, applicationLoadBalancerId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersFlowlogsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersApplicationloadbalancersFlowlogsGet`: FlowLogs
    fmt.Fprintf(os.Stdout, "Response from `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersFlowlogsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**applicationLoadBalancerId** | **string** | The unique ID of the Application Load Balancer. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersApplicationloadbalancersFlowlogsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**FlowLogs**](../models/FlowLogs.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersApplicationloadbalancersFlowlogsPatch

```go
var result FlowLog = DatacentersApplicationloadbalancersFlowlogsPatch(ctx, datacenterId, applicationLoadBalancerId, flowLogId)
                      .ApplicationLoadBalancerFlowLogProperties(applicationLoadBalancerFlowLogProperties)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Partially Modify an ALB Flow Log by ID



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    applicationLoadBalancerId := "applicationLoadBalancerId_example" // string | The unique ID of the Application Load Balancer.
    flowLogId := "flowLogId_example" // string | The unique ID of the flow log.
    applicationLoadBalancerFlowLogProperties := *openapiclient.NewFlowLogProperties("My resource", "ACCEPTED", "INGRESS", "bucketName/key") // FlowLogProperties | The properties of the ALB flow log to be updated.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersFlowlogsPatch(context.Background(), datacenterId, applicationLoadBalancerId, flowLogId).ApplicationLoadBalancerFlowLogProperties(applicationLoadBalancerFlowLogProperties).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersFlowlogsPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersApplicationloadbalancersFlowlogsPatch`: FlowLog
    fmt.Fprintf(os.Stdout, "Response from `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersFlowlogsPatch`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**applicationLoadBalancerId** | **string** | The unique ID of the Application Load Balancer. | |
|**flowLogId** | **string** | The unique ID of the flow log. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersApplicationloadbalancersFlowlogsPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **applicationLoadBalancerFlowLogProperties** | [**FlowLogProperties**](../models/FlowLogProperties.md) | The properties of the ALB flow log to be updated. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**FlowLog**](../models/FlowLog.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersApplicationloadbalancersFlowlogsPost

```go
var result FlowLog = DatacentersApplicationloadbalancersFlowlogsPost(ctx, datacenterId, applicationLoadBalancerId)
                      .ApplicationLoadBalancerFlowLog(applicationLoadBalancerFlowLog)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create an ALB Flow Log



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    applicationLoadBalancerId := "applicationLoadBalancerId_example" // string | The unique ID of the Application Load Balancer.
    applicationLoadBalancerFlowLog := *openapiclient.NewFlowLog(*openapiclient.NewFlowLogProperties("My resource", "ACCEPTED", "INGRESS", "bucketName/key")) // FlowLog | The flow log to create.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersFlowlogsPost(context.Background(), datacenterId, applicationLoadBalancerId).ApplicationLoadBalancerFlowLog(applicationLoadBalancerFlowLog).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersFlowlogsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersApplicationloadbalancersFlowlogsPost`: FlowLog
    fmt.Fprintf(os.Stdout, "Response from `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersFlowlogsPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**applicationLoadBalancerId** | **string** | The unique ID of the Application Load Balancer. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersApplicationloadbalancersFlowlogsPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **applicationLoadBalancerFlowLog** | [**FlowLog**](../models/FlowLog.md) | The flow log to create. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**FlowLog**](../models/FlowLog.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersApplicationloadbalancersFlowlogsPut

```go
var result FlowLog = DatacentersApplicationloadbalancersFlowlogsPut(ctx, datacenterId, applicationLoadBalancerId, flowLogId)
                      .ApplicationLoadBalancerFlowLog(applicationLoadBalancerFlowLog)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify an ALB Flow Log by ID



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    applicationLoadBalancerId := "applicationLoadBalancerId_example" // string | The unique ID of the Application Load Balancer.
    flowLogId := "flowLogId_example" // string | The unique ID of the flow log.
    applicationLoadBalancerFlowLog := *openapiclient.NewFlowLogPut(*openapiclient.NewFlowLogProperties("My resource", "ACCEPTED", "INGRESS", "bucketName/key")) // FlowLogPut | The modified ALB flow log.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersFlowlogsPut(context.Background(), datacenterId, applicationLoadBalancerId, flowLogId).ApplicationLoadBalancerFlowLog(applicationLoadBalancerFlowLog).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersFlowlogsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersApplicationloadbalancersFlowlogsPut`: FlowLog
    fmt.Fprintf(os.Stdout, "Response from `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersFlowlogsPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**applicationLoadBalancerId** | **string** | The unique ID of the Application Load Balancer. | |
|**flowLogId** | **string** | The unique ID of the flow log. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersApplicationloadbalancersFlowlogsPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **applicationLoadBalancerFlowLog** | [**FlowLogPut**](../models/FlowLogPut.md) | The modified ALB flow log. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**FlowLog**](../models/FlowLog.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersApplicationloadbalancersForwardingrulesDelete

```go
var result  = DatacentersApplicationloadbalancersForwardingrulesDelete(ctx, datacenterId, applicationLoadBalancerId, forwardingRuleId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Delete an ALB Forwarding Rule by ID



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    applicationLoadBalancerId := "applicationLoadBalancerId_example" // string | The unique ID of the Application Load Balancer.
    forwardingRuleId := "forwardingRuleId_example" // string | The unique ID of the forwarding rule.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resp, err := apiClient.ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersForwardingrulesDelete(context.Background(), datacenterId, applicationLoadBalancerId, forwardingRuleId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersForwardingrulesDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**applicationLoadBalancerId** | **string** | The unique ID of the Application Load Balancer. | |
|**forwardingRuleId** | **string** | The unique ID of the forwarding rule. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersApplicationloadbalancersForwardingrulesDeleteRequest struct via the builder pattern


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



## DatacentersApplicationloadbalancersForwardingrulesFindByForwardingRuleId

```go
var result ApplicationLoadBalancerForwardingRule = DatacentersApplicationloadbalancersForwardingrulesFindByForwardingRuleId(ctx, datacenterId, applicationLoadBalancerId, forwardingRuleId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Get an ALB Forwarding Rule by ID



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    applicationLoadBalancerId := "applicationLoadBalancerId_example" // string | The unique ID of the Application Load Balancer.
    forwardingRuleId := "forwardingRuleId_example" // string | The unique ID of the forwarding rule.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersForwardingrulesFindByForwardingRuleId(context.Background(), datacenterId, applicationLoadBalancerId, forwardingRuleId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersForwardingrulesFindByForwardingRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersApplicationloadbalancersForwardingrulesFindByForwardingRuleId`: ApplicationLoadBalancerForwardingRule
    fmt.Fprintf(os.Stdout, "Response from `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersForwardingrulesFindByForwardingRuleId`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**applicationLoadBalancerId** | **string** | The unique ID of the Application Load Balancer. | |
|**forwardingRuleId** | **string** | The unique ID of the forwarding rule. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersApplicationloadbalancersForwardingrulesFindByForwardingRuleIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**ApplicationLoadBalancerForwardingRule**](../models/ApplicationLoadBalancerForwardingRule.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersApplicationloadbalancersForwardingrulesGet

```go
var result ApplicationLoadBalancerForwardingRules = DatacentersApplicationloadbalancersForwardingrulesGet(ctx, datacenterId, applicationLoadBalancerId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Get ALB Forwarding Rules



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    applicationLoadBalancerId := "applicationLoadBalancerId_example" // string | The unique ID of the Application Load Balancer.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersForwardingrulesGet(context.Background(), datacenterId, applicationLoadBalancerId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersForwardingrulesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersApplicationloadbalancersForwardingrulesGet`: ApplicationLoadBalancerForwardingRules
    fmt.Fprintf(os.Stdout, "Response from `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersForwardingrulesGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**applicationLoadBalancerId** | **string** | The unique ID of the Application Load Balancer. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersApplicationloadbalancersForwardingrulesGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**ApplicationLoadBalancerForwardingRules**](../models/ApplicationLoadBalancerForwardingRules.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersApplicationloadbalancersForwardingrulesPatch

```go
var result ApplicationLoadBalancerForwardingRule = DatacentersApplicationloadbalancersForwardingrulesPatch(ctx, datacenterId, applicationLoadBalancerId, forwardingRuleId)
                      .ApplicationLoadBalancerForwardingRuleProperties(applicationLoadBalancerForwardingRuleProperties)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Partially modify an ALB Forwarding Rule by ID



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    applicationLoadBalancerId := "applicationLoadBalancerId_example" // string | The unique ID of the Application Load Balancer.
    forwardingRuleId := "forwardingRuleId_example" // string | The unique ID of the forwarding rule.
    applicationLoadBalancerForwardingRuleProperties := *openapiclient.NewApplicationLoadBalancerForwardingRuleProperties("My Application Load Balancer forwarding rule", "HTTP", "81.173.1.2", int32(8080)) // ApplicationLoadBalancerForwardingRuleProperties | The properties of the forwarding rule to be updated.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersForwardingrulesPatch(context.Background(), datacenterId, applicationLoadBalancerId, forwardingRuleId).ApplicationLoadBalancerForwardingRuleProperties(applicationLoadBalancerForwardingRuleProperties).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersForwardingrulesPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersApplicationloadbalancersForwardingrulesPatch`: ApplicationLoadBalancerForwardingRule
    fmt.Fprintf(os.Stdout, "Response from `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersForwardingrulesPatch`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**applicationLoadBalancerId** | **string** | The unique ID of the Application Load Balancer. | |
|**forwardingRuleId** | **string** | The unique ID of the forwarding rule. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersApplicationloadbalancersForwardingrulesPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **applicationLoadBalancerForwardingRuleProperties** | [**ApplicationLoadBalancerForwardingRuleProperties**](../models/ApplicationLoadBalancerForwardingRuleProperties.md) | The properties of the forwarding rule to be updated. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**ApplicationLoadBalancerForwardingRule**](../models/ApplicationLoadBalancerForwardingRule.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersApplicationloadbalancersForwardingrulesPost

```go
var result ApplicationLoadBalancerForwardingRule = DatacentersApplicationloadbalancersForwardingrulesPost(ctx, datacenterId, applicationLoadBalancerId)
                      .ApplicationLoadBalancerForwardingRule(applicationLoadBalancerForwardingRule)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create an ALB Forwarding Rule



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    applicationLoadBalancerId := "applicationLoadBalancerId_example" // string | The unique ID of the Application Load Balancer.
    applicationLoadBalancerForwardingRule := *openapiclient.NewApplicationLoadBalancerForwardingRule(*openapiclient.NewApplicationLoadBalancerForwardingRuleProperties("My Application Load Balancer forwarding rule", "HTTP", "81.173.1.2", int32(8080))) // ApplicationLoadBalancerForwardingRule | The forwarding rule to create.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersForwardingrulesPost(context.Background(), datacenterId, applicationLoadBalancerId).ApplicationLoadBalancerForwardingRule(applicationLoadBalancerForwardingRule).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersForwardingrulesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersApplicationloadbalancersForwardingrulesPost`: ApplicationLoadBalancerForwardingRule
    fmt.Fprintf(os.Stdout, "Response from `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersForwardingrulesPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**applicationLoadBalancerId** | **string** | The unique ID of the Application Load Balancer. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersApplicationloadbalancersForwardingrulesPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **applicationLoadBalancerForwardingRule** | [**ApplicationLoadBalancerForwardingRule**](../models/ApplicationLoadBalancerForwardingRule.md) | The forwarding rule to create. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**ApplicationLoadBalancerForwardingRule**](../models/ApplicationLoadBalancerForwardingRule.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersApplicationloadbalancersForwardingrulesPut

```go
var result ApplicationLoadBalancerForwardingRule = DatacentersApplicationloadbalancersForwardingrulesPut(ctx, datacenterId, applicationLoadBalancerId, forwardingRuleId)
                      .ApplicationLoadBalancerForwardingRule(applicationLoadBalancerForwardingRule)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify an ALB Forwarding Rule by ID



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    applicationLoadBalancerId := "applicationLoadBalancerId_example" // string | The unique ID of the Application Load Balancer.
    forwardingRuleId := "forwardingRuleId_example" // string | The unique ID of the forwarding rule.
    applicationLoadBalancerForwardingRule := *openapiclient.NewApplicationLoadBalancerForwardingRulePut(*openapiclient.NewApplicationLoadBalancerForwardingRuleProperties("My Application Load Balancer forwarding rule", "HTTP", "81.173.1.2", int32(8080))) // ApplicationLoadBalancerForwardingRulePut | The modified ALB forwarding rule.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersForwardingrulesPut(context.Background(), datacenterId, applicationLoadBalancerId, forwardingRuleId).ApplicationLoadBalancerForwardingRule(applicationLoadBalancerForwardingRule).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersForwardingrulesPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersApplicationloadbalancersForwardingrulesPut`: ApplicationLoadBalancerForwardingRule
    fmt.Fprintf(os.Stdout, "Response from `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersForwardingrulesPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**applicationLoadBalancerId** | **string** | The unique ID of the Application Load Balancer. | |
|**forwardingRuleId** | **string** | The unique ID of the forwarding rule. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersApplicationloadbalancersForwardingrulesPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **applicationLoadBalancerForwardingRule** | [**ApplicationLoadBalancerForwardingRulePut**](../models/ApplicationLoadBalancerForwardingRulePut.md) | The modified ALB forwarding rule. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**ApplicationLoadBalancerForwardingRule**](../models/ApplicationLoadBalancerForwardingRule.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersApplicationloadbalancersGet

```go
var result ApplicationLoadBalancers = DatacentersApplicationloadbalancersGet(ctx, datacenterId)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Offset(offset)
                      .Limit(limit)
                      .Execute()
```

Get Application Load Balancers



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)
    offset := int32(56) // int32 | The first element (from the complete list of the elements) to include in the response (used together with <b><i>limit</i></b> for pagination). (optional) (default to 0)
    limit := int32(56) // int32 | The maximum number of elements to return (use together with offset for pagination). (optional) (default to 1000)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersGet(context.Background(), datacenterId).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersApplicationloadbalancersGet`: ApplicationLoadBalancers
    fmt.Fprintf(os.Stdout, "Response from `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersApplicationloadbalancersGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |
| **offset** | **int32** | The first element (from the complete list of the elements) to include in the response (used together with &lt;b&gt;&lt;i&gt;limit&lt;/i&gt;&lt;/b&gt; for pagination). | [default to 0]|
| **limit** | **int32** | The maximum number of elements to return (use together with offset for pagination). | [default to 1000]|

### Return type

[**ApplicationLoadBalancers**](../models/ApplicationLoadBalancers.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersApplicationloadbalancersPatch

```go
var result ApplicationLoadBalancer = DatacentersApplicationloadbalancersPatch(ctx, datacenterId, applicationLoadBalancerId)
                      .ApplicationLoadBalancerProperties(applicationLoadBalancerProperties)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Partially Modify an Application Load Balancer by ID



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    applicationLoadBalancerId := "applicationLoadBalancerId_example" // string | The unique ID of the Application Load Balancer.
    applicationLoadBalancerProperties := *openapiclient.NewApplicationLoadBalancerProperties("My Application Load Balancer", int32(1), int32(2)) // ApplicationLoadBalancerProperties | The Application Load Balancer properties to be updated.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersPatch(context.Background(), datacenterId, applicationLoadBalancerId).ApplicationLoadBalancerProperties(applicationLoadBalancerProperties).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersApplicationloadbalancersPatch`: ApplicationLoadBalancer
    fmt.Fprintf(os.Stdout, "Response from `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersPatch`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**applicationLoadBalancerId** | **string** | The unique ID of the Application Load Balancer. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersApplicationloadbalancersPatchRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **applicationLoadBalancerProperties** | [**ApplicationLoadBalancerProperties**](../models/ApplicationLoadBalancerProperties.md) | The Application Load Balancer properties to be updated. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**ApplicationLoadBalancer**](../models/ApplicationLoadBalancer.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json



## DatacentersApplicationloadbalancersPost

```go
var result ApplicationLoadBalancer = DatacentersApplicationloadbalancersPost(ctx, datacenterId)
                      .ApplicationLoadBalancer(applicationLoadBalancer)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Create an Application Load Balancer



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    applicationLoadBalancer := *openapiclient.NewApplicationLoadBalancer(*openapiclient.NewApplicationLoadBalancerProperties("My Application Load Balancer", int32(1), int32(2))) // ApplicationLoadBalancer | The Application Load Balancer to create.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersPost(context.Background(), datacenterId).ApplicationLoadBalancer(applicationLoadBalancer).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersApplicationloadbalancersPost`: ApplicationLoadBalancer
    fmt.Fprintf(os.Stdout, "Response from `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersPost`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersApplicationloadbalancersPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **applicationLoadBalancer** | [**ApplicationLoadBalancer**](../models/ApplicationLoadBalancer.md) | The Application Load Balancer to create. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**ApplicationLoadBalancer**](../models/ApplicationLoadBalancer.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json



## DatacentersApplicationloadbalancersPut

```go
var result ApplicationLoadBalancer = DatacentersApplicationloadbalancersPut(ctx, datacenterId, applicationLoadBalancerId)
                      .ApplicationLoadBalancer(applicationLoadBalancer)
                      .Pretty(pretty)
                      .Depth(depth)
                      .XContractNumber(xContractNumber)
                      .Execute()
```

Modify an Application Load Balancer by ID



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
    datacenterId := "datacenterId_example" // string | The unique ID of the data center.
    applicationLoadBalancerId := "applicationLoadBalancerId_example" // string | The unique ID of the Application Load Balancer.
    applicationLoadBalancer := *openapiclient.NewApplicationLoadBalancerPut(*openapiclient.NewApplicationLoadBalancerProperties("My Application Load Balancer", int32(1), int32(2))) // ApplicationLoadBalancerPut | The modified Application Load Balancer.
    pretty := true // bool | Controls whether the response is pretty-printed (with indentations and new lines). (optional) (default to true)
    depth := int32(56) // int32 | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth=0: Only direct properties are included; children (servers and other elements) are not included.  - depth=1: Direct properties and children references are included.  - depth=2: Direct properties and children properties are included.  - depth=3: Direct properties and children properties and children's children are included.  - depth=... and so on (optional) (default to 0)
    xContractNumber := int32(56) // int32 | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. (optional)

    configuration := shared.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := compute.NewAPIClient(configuration)
    resource, resp, err := apiClient.ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersPut(context.Background(), datacenterId, applicationLoadBalancerId).ApplicationLoadBalancer(applicationLoadBalancer).Pretty(pretty).Depth(depth).XContractNumber(xContractNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `DatacentersApplicationloadbalancersPut`: ApplicationLoadBalancer
    fmt.Fprintf(os.Stdout, "Response from `ApplicationLoadBalancersApi.DatacentersApplicationloadbalancersPut`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**datacenterId** | **string** | The unique ID of the data center. | |
|**applicationLoadBalancerId** | **string** | The unique ID of the Application Load Balancer. | |

### Other Parameters

Other parameters are passed through a pointer to an apiDatacentersApplicationloadbalancersPutRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **applicationLoadBalancer** | [**ApplicationLoadBalancerPut**](../models/ApplicationLoadBalancerPut.md) | The modified Application Load Balancer. | |
| **pretty** | **bool** | Controls whether the response is pretty-printed (with indentations and new lines). | [default to true]|
| **depth** | **int32** | Controls the detail depth of the response objects.  GET /datacenters/[ID]  - depth&#x3D;0: Only direct properties are included; children (servers and other elements) are not included.  - depth&#x3D;1: Direct properties and children references are included.  - depth&#x3D;2: Direct properties and children properties are included.  - depth&#x3D;3: Direct properties and children properties and children&#39;s children are included.  - depth&#x3D;... and so on | [default to 0]|
| **xContractNumber** | **int32** | Users with multiple contracts must provide the contract number, for which all API requests are to be executed. | |

### Return type

[**ApplicationLoadBalancer**](../models/ApplicationLoadBalancer.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


