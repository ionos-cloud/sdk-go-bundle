# RestoreClusterFromSnapshot

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**SourceSnapshotId** | **string** | The UUID of the snapshot to restore from. | |
|**RecoveryTargetDatetime** | Pointer to [**time.Time**](time.Time.md) | Provide an ISO 8601 timestamp to restore from the most recent snapshot taken at or before that time, within the snapshot&#39;s recovery window. In-Memory DB does not provide continuous point-in-time recovery; the nearest preceding snapshot is used. If omitted, the cluster is restored from the latest available snapshot.  | [optional] |

## Methods

### NewRestoreClusterFromSnapshot

`func NewRestoreClusterFromSnapshot(sourceSnapshotId string, ) *RestoreClusterFromSnapshot`

NewRestoreClusterFromSnapshot instantiates a new RestoreClusterFromSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreClusterFromSnapshotWithDefaults

`func NewRestoreClusterFromSnapshotWithDefaults() *RestoreClusterFromSnapshot`

NewRestoreClusterFromSnapshotWithDefaults instantiates a new RestoreClusterFromSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceSnapshotId

`func (o *RestoreClusterFromSnapshot) GetSourceSnapshotId() string`

GetSourceSnapshotId returns the SourceSnapshotId field if non-nil, zero value otherwise.

### GetSourceSnapshotIdOk

`func (o *RestoreClusterFromSnapshot) GetSourceSnapshotIdOk() (*string, bool)`

GetSourceSnapshotIdOk returns a tuple with the SourceSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSnapshotId

`func (o *RestoreClusterFromSnapshot) SetSourceSnapshotId(v string)`

SetSourceSnapshotId sets SourceSnapshotId field to given value.


### GetRecoveryTargetDatetime

`func (o *RestoreClusterFromSnapshot) GetRecoveryTargetDatetime() time.Time`

GetRecoveryTargetDatetime returns the RecoveryTargetDatetime field if non-nil, zero value otherwise.

### GetRecoveryTargetDatetimeOk

`func (o *RestoreClusterFromSnapshot) GetRecoveryTargetDatetimeOk() (*time.Time, bool)`

GetRecoveryTargetDatetimeOk returns a tuple with the RecoveryTargetDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetDatetime

`func (o *RestoreClusterFromSnapshot) SetRecoveryTargetDatetime(v time.Time)`

SetRecoveryTargetDatetime sets RecoveryTargetDatetime field to given value.

### HasRecoveryTargetDatetime

`func (o *RestoreClusterFromSnapshot) HasRecoveryTargetDatetime() bool`

HasRecoveryTargetDatetime returns a boolean if a field has been set.


