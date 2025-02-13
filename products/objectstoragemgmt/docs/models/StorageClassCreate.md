# StorageClassCreate

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Metadata** | Pointer to **map[string]interface{}** | Metadata | [optional] |
|**Properties** | [**StorageClass**](StorageClass.md) |  | |

## Methods

### NewStorageClassCreate

`func NewStorageClassCreate(properties StorageClass, ) *StorageClassCreate`

NewStorageClassCreate instantiates a new StorageClassCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageClassCreateWithDefaults

`func NewStorageClassCreateWithDefaults() *StorageClassCreate`

NewStorageClassCreateWithDefaults instantiates a new StorageClassCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *StorageClassCreate) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StorageClassCreate) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StorageClassCreate) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *StorageClassCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *StorageClassCreate) GetProperties() StorageClass`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *StorageClassCreate) GetPropertiesOk() (*StorageClass, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *StorageClassCreate) SetProperties(v StorageClass)`

SetProperties sets Properties field to given value.



