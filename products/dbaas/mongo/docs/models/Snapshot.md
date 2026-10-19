# Snapshot

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ClusterId** | Pointer to **string** | The unique identifier of the cluster this snapshot belongs to. | [optional] |
|**ClusterName** | Pointer to **string** | The name of the MongoDB cluster this snapshot belongs to. | [optional] [readonly] |
|**MongoVersion** | Pointer to **string** |  | [optional] |
|**EarliestRecoveryTargetTime** | Pointer to [**time.Time**](time.Time.md) | The earliest point in time to which the cluster can be restored from this snapshot. | [optional] |
|**LatestRecoveryTargetTime** | Pointer to [**NullableTime**](time.Time.md) | The latest possible point in time to which the cluster can be restored. If the cluster can be restored up to the current time, this field will be null.  | [optional] |
|**Location** | Pointer to **string** | The Object Storage location where the snapshots will be created. For added data safety, use a different location than the cluster.  | [optional] |
|**TotalSize** | Pointer to **int32** | The size of the snapshot in Mebibytes | [optional] |
|**RequiredStorageSize** | Pointer to **int32** | The minimum storage size required to restore this snapshot in Mebibytes | [optional] |

## Methods

### NewSnapshot

`func NewSnapshot() *Snapshot`

NewSnapshot instantiates a new Snapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotWithDefaults

`func NewSnapshotWithDefaults() *Snapshot`

NewSnapshotWithDefaults instantiates a new Snapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterId

`func (o *Snapshot) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *Snapshot) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *Snapshot) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.

### HasClusterId

`func (o *Snapshot) HasClusterId() bool`

HasClusterId returns a boolean if a field has been set.

### GetClusterName

`func (o *Snapshot) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *Snapshot) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *Snapshot) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *Snapshot) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetMongoVersion

`func (o *Snapshot) GetMongoVersion() string`

GetMongoVersion returns the MongoVersion field if non-nil, zero value otherwise.

### GetMongoVersionOk

`func (o *Snapshot) GetMongoVersionOk() (*string, bool)`

GetMongoVersionOk returns a tuple with the MongoVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoVersion

`func (o *Snapshot) SetMongoVersion(v string)`

SetMongoVersion sets MongoVersion field to given value.

### HasMongoVersion

`func (o *Snapshot) HasMongoVersion() bool`

HasMongoVersion returns a boolean if a field has been set.

### GetEarliestRecoveryTargetTime

`func (o *Snapshot) GetEarliestRecoveryTargetTime() time.Time`

GetEarliestRecoveryTargetTime returns the EarliestRecoveryTargetTime field if non-nil, zero value otherwise.

### GetEarliestRecoveryTargetTimeOk

`func (o *Snapshot) GetEarliestRecoveryTargetTimeOk() (*time.Time, bool)`

GetEarliestRecoveryTargetTimeOk returns a tuple with the EarliestRecoveryTargetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestRecoveryTargetTime

`func (o *Snapshot) SetEarliestRecoveryTargetTime(v time.Time)`

SetEarliestRecoveryTargetTime sets EarliestRecoveryTargetTime field to given value.

### HasEarliestRecoveryTargetTime

`func (o *Snapshot) HasEarliestRecoveryTargetTime() bool`

HasEarliestRecoveryTargetTime returns a boolean if a field has been set.

### GetLatestRecoveryTargetTime

`func (o *Snapshot) GetLatestRecoveryTargetTime() time.Time`

GetLatestRecoveryTargetTime returns the LatestRecoveryTargetTime field if non-nil, zero value otherwise.

### GetLatestRecoveryTargetTimeOk

`func (o *Snapshot) GetLatestRecoveryTargetTimeOk() (*time.Time, bool)`

GetLatestRecoveryTargetTimeOk returns a tuple with the LatestRecoveryTargetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRecoveryTargetTime

`func (o *Snapshot) SetLatestRecoveryTargetTime(v time.Time)`

SetLatestRecoveryTargetTime sets LatestRecoveryTargetTime field to given value.

### HasLatestRecoveryTargetTime

`func (o *Snapshot) HasLatestRecoveryTargetTime() bool`

HasLatestRecoveryTargetTime returns a boolean if a field has been set.

### SetLatestRecoveryTargetTimeNil

`func (o *Snapshot) SetLatestRecoveryTargetTimeNil(b bool)`

 SetLatestRecoveryTargetTimeNil sets the value for LatestRecoveryTargetTime to be an explicit nil

### UnsetLatestRecoveryTargetTime
`func (o *Snapshot) UnsetLatestRecoveryTargetTime()`

UnsetLatestRecoveryTargetTime ensures that no value is present for LatestRecoveryTargetTime, not even an explicit nil
### GetLocation

`func (o *Snapshot) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Snapshot) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Snapshot) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Snapshot) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetTotalSize

`func (o *Snapshot) GetTotalSize() int32`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *Snapshot) GetTotalSizeOk() (*int32, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *Snapshot) SetTotalSize(v int32)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *Snapshot) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### GetRequiredStorageSize

`func (o *Snapshot) GetRequiredStorageSize() int32`

GetRequiredStorageSize returns the RequiredStorageSize field if non-nil, zero value otherwise.

### GetRequiredStorageSizeOk

`func (o *Snapshot) GetRequiredStorageSizeOk() (*int32, bool)`

GetRequiredStorageSizeOk returns a tuple with the RequiredStorageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredStorageSize

`func (o *Snapshot) SetRequiredStorageSize(v int32)`

SetRequiredStorageSize sets RequiredStorageSize field to given value.

### HasRequiredStorageSize

`func (o *Snapshot) HasRequiredStorageSize() bool`

HasRequiredStorageSize returns a boolean if a field has been set.


