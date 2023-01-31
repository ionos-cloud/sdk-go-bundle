# SnapshotProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Version** | Pointer to **string** | The MongoDB version this backup was created from. | [optional] |
|**Size** | Pointer to **int32** | The size of the snapshot in Mebibytes. | [optional] |
|**CreationTime** | Pointer to [**time.Time**](time.Time.md) | The date the resource was created. | [optional] |

## Methods

### NewSnapshotProperties

`func NewSnapshotProperties() *SnapshotProperties`

NewSnapshotProperties instantiates a new SnapshotProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotPropertiesWithDefaults

`func NewSnapshotPropertiesWithDefaults() *SnapshotProperties`

NewSnapshotPropertiesWithDefaults instantiates a new SnapshotProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *SnapshotProperties) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SnapshotProperties) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SnapshotProperties) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SnapshotProperties) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSize

`func (o *SnapshotProperties) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SnapshotProperties) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SnapshotProperties) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *SnapshotProperties) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetCreationTime

`func (o *SnapshotProperties) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *SnapshotProperties) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *SnapshotProperties) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *SnapshotProperties) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.


