# PipelineProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The friendly name of your pipeline. | [optional] |
|**Logs** | Pointer to [**[]PipelineResponse**](PipelineResponse.md) | The information of the log aggregator | [optional] |
|**TcpAddress** | Pointer to **string** | The address to connect fluentBit compatible logging agents to | [optional] |
|**HttpAddress** | Pointer to **string** | The address to post logs to using JSON with basic auth | [optional] |
|**GrafanaAddress** | Pointer to **string** | The address of the client&#39;s grafana instance | [optional] |

## Methods

### NewPipelineProperties

`func NewPipelineProperties() *PipelineProperties`

NewPipelineProperties instantiates a new PipelineProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelinePropertiesWithDefaults

`func NewPipelinePropertiesWithDefaults() *PipelineProperties`

NewPipelinePropertiesWithDefaults instantiates a new PipelineProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PipelineProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PipelineProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PipelineProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PipelineProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLogs

`func (o *PipelineProperties) GetLogs() []PipelineResponse`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *PipelineProperties) GetLogsOk() (*[]PipelineResponse, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *PipelineProperties) SetLogs(v []PipelineResponse)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *PipelineProperties) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetTcpAddress

`func (o *PipelineProperties) GetTcpAddress() string`

GetTcpAddress returns the TcpAddress field if non-nil, zero value otherwise.

### GetTcpAddressOk

`func (o *PipelineProperties) GetTcpAddressOk() (*string, bool)`

GetTcpAddressOk returns a tuple with the TcpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpAddress

`func (o *PipelineProperties) SetTcpAddress(v string)`

SetTcpAddress sets TcpAddress field to given value.

### HasTcpAddress

`func (o *PipelineProperties) HasTcpAddress() bool`

HasTcpAddress returns a boolean if a field has been set.

### GetHttpAddress

`func (o *PipelineProperties) GetHttpAddress() string`

GetHttpAddress returns the HttpAddress field if non-nil, zero value otherwise.

### GetHttpAddressOk

`func (o *PipelineProperties) GetHttpAddressOk() (*string, bool)`

GetHttpAddressOk returns a tuple with the HttpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpAddress

`func (o *PipelineProperties) SetHttpAddress(v string)`

SetHttpAddress sets HttpAddress field to given value.

### HasHttpAddress

`func (o *PipelineProperties) HasHttpAddress() bool`

HasHttpAddress returns a boolean if a field has been set.

### GetGrafanaAddress

`func (o *PipelineProperties) GetGrafanaAddress() string`

GetGrafanaAddress returns the GrafanaAddress field if non-nil, zero value otherwise.

### GetGrafanaAddressOk

`func (o *PipelineProperties) GetGrafanaAddressOk() (*string, bool)`

GetGrafanaAddressOk returns a tuple with the GrafanaAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrafanaAddress

`func (o *PipelineProperties) SetGrafanaAddress(v string)`

SetGrafanaAddress sets GrafanaAddress field to given value.

### HasGrafanaAddress

`func (o *PipelineProperties) HasGrafanaAddress() bool`

HasGrafanaAddress returns a boolean if a field has been set.


