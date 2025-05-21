# RestoreMetadataAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**RestoreTime** | Pointer to [**time.Time**](time.Time.md) | The time the snapshot was dumped from the replica set. | [optional] |
|**RestoredSnapshotId** | Pointer to **string** | The ID of the snapshot that was restored. | [optional] |

## Methods

### NewRestoreMetadataAllOf

`func NewRestoreMetadataAllOf() *RestoreMetadataAllOf`

NewRestoreMetadataAllOf instantiates a new RestoreMetadataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreMetadataAllOfWithDefaults

`func NewRestoreMetadataAllOfWithDefaults() *RestoreMetadataAllOf`

NewRestoreMetadataAllOfWithDefaults instantiates a new RestoreMetadataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestoreTime

`func (o *RestoreMetadataAllOf) GetRestoreTime() time.Time`

GetRestoreTime returns the RestoreTime field if non-nil, zero value otherwise.

### GetRestoreTimeOk

`func (o *RestoreMetadataAllOf) GetRestoreTimeOk() (*time.Time, bool)`

GetRestoreTimeOk returns a tuple with the RestoreTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreTime

`func (o *RestoreMetadataAllOf) SetRestoreTime(v time.Time)`

SetRestoreTime sets RestoreTime field to given value.

### HasRestoreTime

`func (o *RestoreMetadataAllOf) HasRestoreTime() bool`

HasRestoreTime returns a boolean if a field has been set.

### GetRestoredSnapshotId

`func (o *RestoreMetadataAllOf) GetRestoredSnapshotId() string`

GetRestoredSnapshotId returns the RestoredSnapshotId field if non-nil, zero value otherwise.

### GetRestoredSnapshotIdOk

`func (o *RestoreMetadataAllOf) GetRestoredSnapshotIdOk() (*string, bool)`

GetRestoredSnapshotIdOk returns a tuple with the RestoredSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredSnapshotId

`func (o *RestoreMetadataAllOf) SetRestoredSnapshotId(v string)`

SetRestoredSnapshotId sets RestoredSnapshotId field to given value.

### HasRestoredSnapshotId

`func (o *RestoreMetadataAllOf) HasRestoredSnapshotId() bool`

HasRestoredSnapshotId returns a boolean if a field has been set.


