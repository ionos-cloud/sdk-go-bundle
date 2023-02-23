# PatchNodePoolRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Properties** | [**PatchNodePoolProperties**](PatchNodePoolProperties.md) |  | |

## Methods

### NewPatchNodePoolRequest

`func NewPatchNodePoolRequest(properties PatchNodePoolProperties, ) *PatchNodePoolRequest`

NewPatchNodePoolRequest instantiates a new PatchNodePoolRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchNodePoolRequestWithDefaults

`func NewPatchNodePoolRequestWithDefaults() *PatchNodePoolRequest`

NewPatchNodePoolRequestWithDefaults instantiates a new PatchNodePoolRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *PatchNodePoolRequest) GetProperties() PatchNodePoolProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PatchNodePoolRequest) GetPropertiesOk() (*PatchNodePoolProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PatchNodePoolRequest) SetProperties(v PatchNodePoolProperties)`

SetProperties sets Properties field to given value.



