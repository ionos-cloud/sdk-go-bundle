# PatchUserRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Metadata** | Pointer to [**UserMetadata**](UserMetadata.md) |  | [optional] |
|**Properties** | Pointer to [**PatchUserProperties**](PatchUserProperties.md) |  | [optional] |

## Methods

### NewPatchUserRequest

`func NewPatchUserRequest() *PatchUserRequest`

NewPatchUserRequest instantiates a new PatchUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchUserRequestWithDefaults

`func NewPatchUserRequestWithDefaults() *PatchUserRequest`

NewPatchUserRequestWithDefaults instantiates a new PatchUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *PatchUserRequest) GetMetadata() UserMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PatchUserRequest) GetMetadataOk() (*UserMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PatchUserRequest) SetMetadata(v UserMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PatchUserRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *PatchUserRequest) GetProperties() PatchUserProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PatchUserRequest) GetPropertiesOk() (*PatchUserProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PatchUserRequest) SetProperties(v PatchUserProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *PatchUserRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


