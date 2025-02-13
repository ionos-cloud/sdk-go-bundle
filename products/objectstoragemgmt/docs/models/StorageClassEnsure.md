# StorageClassEnsure

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The StorageClass of the StorageClass. | |
|**Metadata** | Pointer to **map[string]interface{}** | Metadata | [optional] |
|**Properties** | [**StorageClass**](StorageClass.md) |  | |

## Methods

### NewStorageClassEnsure

`func NewStorageClassEnsure(id string, properties StorageClass, ) *StorageClassEnsure`

NewStorageClassEnsure instantiates a new StorageClassEnsure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageClassEnsureWithDefaults

`func NewStorageClassEnsureWithDefaults() *StorageClassEnsure`

NewStorageClassEnsureWithDefaults instantiates a new StorageClassEnsure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageClassEnsure) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageClassEnsure) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageClassEnsure) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *StorageClassEnsure) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StorageClassEnsure) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StorageClassEnsure) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *StorageClassEnsure) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *StorageClassEnsure) GetProperties() StorageClass`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *StorageClassEnsure) GetPropertiesOk() (*StorageClass, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *StorageClassEnsure) SetProperties(v StorageClass)`

SetProperties sets Properties field to given value.



