# PatchRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Properties** | [**PatchRequestProperties**](PatchRequestProperties.md) |  | |

## Methods

### NewPatchRequest

`func NewPatchRequest(properties PatchRequestProperties, ) *PatchRequest`

NewPatchRequest instantiates a new PatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchRequestWithDefaults

`func NewPatchRequestWithDefaults() *PatchRequest`

NewPatchRequestWithDefaults instantiates a new PatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *PatchRequest) GetProperties() PatchRequestProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PatchRequest) GetPropertiesOk() (*PatchRequestProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PatchRequest) SetProperties(v PatchRequestProperties)`

SetProperties sets Properties field to given value.



