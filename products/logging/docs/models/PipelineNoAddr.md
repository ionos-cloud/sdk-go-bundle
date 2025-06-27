# PipelineNoAddr

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The name of the pipeline. Must be not more that 20 characters long.  | |
|**Logs** | Pointer to [**[]PipelineNoAddrLogs**](PipelineNoAddrLogs.md) |  | [optional] |

## Methods

### NewPipelineNoAddr

`func NewPipelineNoAddr(name string, ) *PipelineNoAddr`

NewPipelineNoAddr instantiates a new PipelineNoAddr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineNoAddrWithDefaults

`func NewPipelineNoAddrWithDefaults() *PipelineNoAddr`

NewPipelineNoAddrWithDefaults instantiates a new PipelineNoAddr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PipelineNoAddr) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PipelineNoAddr) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PipelineNoAddr) SetName(v string)`

SetName sets Name field to given value.


### GetLogs

`func (o *PipelineNoAddr) GetLogs() []PipelineNoAddrLogs`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *PipelineNoAddr) GetLogsOk() (*[]PipelineNoAddrLogs, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *PipelineNoAddr) SetLogs(v []PipelineNoAddrLogs)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *PipelineNoAddr) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


