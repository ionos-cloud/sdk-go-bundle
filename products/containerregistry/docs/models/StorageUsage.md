# StorageUsage

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Bytes** | **int64** |  | |
|**UpdatedAt** | Pointer to [**time.Time**](time.Time.md) |  | [optional] |

## Methods

### NewStorageUsage

`func NewStorageUsage(bytes int64, ) *StorageUsage`

NewStorageUsage instantiates a new StorageUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageUsageWithDefaults

`func NewStorageUsageWithDefaults() *StorageUsage`

NewStorageUsageWithDefaults instantiates a new StorageUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *StorageUsage) GetBytes() int64`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *StorageUsage) GetBytesOk() (*int64, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *StorageUsage) SetBytes(v int64)`

SetBytes sets Bytes field to given value.


### GetUpdatedAt

`func (o *StorageUsage) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *StorageUsage) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *StorageUsage) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *StorageUsage) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


