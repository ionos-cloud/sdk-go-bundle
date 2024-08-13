# Gateway

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** |  | |
|**Logs** | Pointer to **bool** | This field enables or disables the collection and reporting of logs for observability of this instance. | [optional] [default to false]|
|**Metrics** | Pointer to **bool** | This field enables or disables the collection and reporting of metrics for observability of this instance. | [optional] [default to false]|
|**CustomDomains** | Pointer to [**[]GatewayCustomDomains**](GatewayCustomDomains.md) |  | [optional] |

## Methods

### NewGateway

`func NewGateway(name string, ) *Gateway`

NewGateway instantiates a new Gateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayWithDefaults

`func NewGatewayWithDefaults() *Gateway`

NewGatewayWithDefaults instantiates a new Gateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Gateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Gateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Gateway) SetName(v string)`

SetName sets Name field to given value.


### GetLogs

`func (o *Gateway) GetLogs() bool`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *Gateway) GetLogsOk() (*bool, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *Gateway) SetLogs(v bool)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *Gateway) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetMetrics

`func (o *Gateway) GetMetrics() bool`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *Gateway) GetMetricsOk() (*bool, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *Gateway) SetMetrics(v bool)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *Gateway) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetCustomDomains

`func (o *Gateway) GetCustomDomains() []GatewayCustomDomains`

GetCustomDomains returns the CustomDomains field if non-nil, zero value otherwise.

### GetCustomDomainsOk

`func (o *Gateway) GetCustomDomainsOk() (*[]GatewayCustomDomains, bool)`

GetCustomDomainsOk returns a tuple with the CustomDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomDomains

`func (o *Gateway) SetCustomDomains(v []GatewayCustomDomains)`

SetCustomDomains sets CustomDomains field to given value.

### HasCustomDomains

`func (o *Gateway) HasCustomDomains() bool`

HasCustomDomains returns a boolean if a field has been set.


