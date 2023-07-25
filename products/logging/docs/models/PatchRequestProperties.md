# PatchRequestProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The friendly name of your pipeline. | [optional] |
|**Logs** | Pointer to [**[]PatchRequestPipeline**](PatchRequestPipeline.md) | The information of the log pipelines | [optional] |

## Methods

### NewPatchRequestProperties

`func NewPatchRequestProperties() *PatchRequestProperties`

NewPatchRequestProperties instantiates a new PatchRequestProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchRequestPropertiesWithDefaults

`func NewPatchRequestPropertiesWithDefaults() *PatchRequestProperties`

NewPatchRequestPropertiesWithDefaults instantiates a new PatchRequestProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PatchRequestProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PatchRequestProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PatchRequestProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PatchRequestProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLogs

`func (o *PatchRequestProperties) GetLogs() []PatchRequestPipeline`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *PatchRequestProperties) GetLogsOk() (*[]PatchRequestPipeline, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *PatchRequestProperties) SetLogs(v []PatchRequestPipeline)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *PatchRequestProperties) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


