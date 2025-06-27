# PipelineAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**TcpAddress** | Pointer to **string** | The TCP address of the pipeline. This is the address to which logs are sent using the TCP protocol.  | [optional] |
|**HttpAddress** | Pointer to **string** | The HTTP address of the pipeline. This is the address to which logs are sent using the HTTP protocol.  | [optional] |
|**GrafanaAddress** | Pointer to **string** | The Grafana address is where user can access their logs, create dashboards, and set up alerts.  | [optional] |
|**ResourceTier** | Pointer to **string** |  | [optional] |
|**Key** | Pointer to **string** | The key is shared once and is used to authenticate the logs sent to the pipeline.  | [optional] |

## Methods

### NewPipelineAllOf

`func NewPipelineAllOf() *PipelineAllOf`

NewPipelineAllOf instantiates a new PipelineAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineAllOfWithDefaults

`func NewPipelineAllOfWithDefaults() *PipelineAllOf`

NewPipelineAllOfWithDefaults instantiates a new PipelineAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTcpAddress

`func (o *PipelineAllOf) GetTcpAddress() string`

GetTcpAddress returns the TcpAddress field if non-nil, zero value otherwise.

### GetTcpAddressOk

`func (o *PipelineAllOf) GetTcpAddressOk() (*string, bool)`

GetTcpAddressOk returns a tuple with the TcpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpAddress

`func (o *PipelineAllOf) SetTcpAddress(v string)`

SetTcpAddress sets TcpAddress field to given value.

### HasTcpAddress

`func (o *PipelineAllOf) HasTcpAddress() bool`

HasTcpAddress returns a boolean if a field has been set.

### GetHttpAddress

`func (o *PipelineAllOf) GetHttpAddress() string`

GetHttpAddress returns the HttpAddress field if non-nil, zero value otherwise.

### GetHttpAddressOk

`func (o *PipelineAllOf) GetHttpAddressOk() (*string, bool)`

GetHttpAddressOk returns a tuple with the HttpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpAddress

`func (o *PipelineAllOf) SetHttpAddress(v string)`

SetHttpAddress sets HttpAddress field to given value.

### HasHttpAddress

`func (o *PipelineAllOf) HasHttpAddress() bool`

HasHttpAddress returns a boolean if a field has been set.

### GetGrafanaAddress

`func (o *PipelineAllOf) GetGrafanaAddress() string`

GetGrafanaAddress returns the GrafanaAddress field if non-nil, zero value otherwise.

### GetGrafanaAddressOk

`func (o *PipelineAllOf) GetGrafanaAddressOk() (*string, bool)`

GetGrafanaAddressOk returns a tuple with the GrafanaAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaAddress

`func (o *PipelineAllOf) SetGrafanaAddress(v string)`

SetGrafanaAddress sets GrafanaAddress field to given value.

### HasGrafanaAddress

`func (o *PipelineAllOf) HasGrafanaAddress() bool`

HasGrafanaAddress returns a boolean if a field has been set.

### GetResourceTier

`func (o *PipelineAllOf) GetResourceTier() string`

GetResourceTier returns the ResourceTier field if non-nil, zero value otherwise.

### GetResourceTierOk

`func (o *PipelineAllOf) GetResourceTierOk() (*string, bool)`

GetResourceTierOk returns a tuple with the ResourceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceTier

`func (o *PipelineAllOf) SetResourceTier(v string)`

SetResourceTier sets ResourceTier field to given value.

### HasResourceTier

`func (o *PipelineAllOf) HasResourceTier() bool`

HasResourceTier returns a boolean if a field has been set.

### GetKey

`func (o *PipelineAllOf) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PipelineAllOf) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PipelineAllOf) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PipelineAllOf) HasKey() bool`

HasKey returns a boolean if a field has been set.


