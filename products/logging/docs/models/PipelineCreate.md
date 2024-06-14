# PipelineCreate

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Properties** | [**PipelineCreateProperties**](PipelineCreateProperties.md) |  | |

## Methods

### NewPipelineCreate

`func NewPipelineCreate(properties PipelineCreateProperties, ) *PipelineCreate`

NewPipelineCreate instantiates a new PipelineCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineCreateWithDefaults

`func NewPipelineCreateWithDefaults() *PipelineCreate`

NewPipelineCreateWithDefaults instantiates a new PipelineCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *PipelineCreate) GetProperties() PipelineCreateProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PipelineCreate) GetPropertiesOk() (*PipelineCreateProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PipelineCreate) SetProperties(v PipelineCreateProperties)`

SetProperties sets Properties field to given value.



