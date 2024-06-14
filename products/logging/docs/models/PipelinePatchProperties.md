# PipelinePatchProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The friendly name of your pipeline. | [optional] |
|**Logs** | Pointer to [**[]Processor**](Processor.md) | The information of the log pipelines | [optional] |

## Methods

### NewPipelinePatchProperties

`func NewPipelinePatchProperties() *PipelinePatchProperties`

NewPipelinePatchProperties instantiates a new PipelinePatchProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelinePatchPropertiesWithDefaults

`func NewPipelinePatchPropertiesWithDefaults() *PipelinePatchProperties`

NewPipelinePatchPropertiesWithDefaults instantiates a new PipelinePatchProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PipelinePatchProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PipelinePatchProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PipelinePatchProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PipelinePatchProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLogs

`func (o *PipelinePatchProperties) GetLogs() []Processor`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *PipelinePatchProperties) GetLogsOk() (*[]Processor, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *PipelinePatchProperties) SetLogs(v []Processor)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *PipelinePatchProperties) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


