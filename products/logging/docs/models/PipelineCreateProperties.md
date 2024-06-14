# PipelineCreateProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The friendly name of your pipeline. | |
|**Logs** | [**[]Processor**](Processor.md) | The information of the log pipelines | |

## Methods

### NewPipelineCreateProperties

`func NewPipelineCreateProperties(name string, logs []Processor, ) *PipelineCreateProperties`

NewPipelineCreateProperties instantiates a new PipelineCreateProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineCreatePropertiesWithDefaults

`func NewPipelineCreatePropertiesWithDefaults() *PipelineCreateProperties`

NewPipelineCreatePropertiesWithDefaults instantiates a new PipelineCreateProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PipelineCreateProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PipelineCreateProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PipelineCreateProperties) SetName(v string)`

SetName sets Name field to given value.


### GetLogs

`func (o *PipelineCreateProperties) GetLogs() []Processor`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *PipelineCreateProperties) GetLogsOk() (*[]Processor, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *PipelineCreateProperties) SetLogs(v []Processor)`

SetLogs sets Logs field to given value.



