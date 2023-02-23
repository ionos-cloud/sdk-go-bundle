# PatchClusterRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Properties** | [**PatchClusterProperties**](PatchClusterProperties.md) |  | |

## Methods

### NewPatchClusterRequest

`func NewPatchClusterRequest(properties PatchClusterProperties, ) *PatchClusterRequest`

NewPatchClusterRequest instantiates a new PatchClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchClusterRequestWithDefaults

`func NewPatchClusterRequestWithDefaults() *PatchClusterRequest`

NewPatchClusterRequestWithDefaults instantiates a new PatchClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *PatchClusterRequest) GetProperties() PatchClusterProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PatchClusterRequest) GetPropertiesOk() (*PatchClusterProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PatchClusterRequest) SetProperties(v PatchClusterProperties)`

SetProperties sets Properties field to given value.



