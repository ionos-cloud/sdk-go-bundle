# PatchClusterRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] |
|**Properties** | Pointer to [**PatchClusterProperties**](PatchClusterProperties.md) |  | [optional] |

## Methods

### NewPatchClusterRequest

`func NewPatchClusterRequest() *PatchClusterRequest`

NewPatchClusterRequest instantiates a new PatchClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchClusterRequestWithDefaults

`func NewPatchClusterRequestWithDefaults() *PatchClusterRequest`

NewPatchClusterRequestWithDefaults instantiates a new PatchClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *PatchClusterRequest) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchClusterRequest) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchClusterRequest) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchClusterRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

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

### HasProperties

`func (o *PatchClusterRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


