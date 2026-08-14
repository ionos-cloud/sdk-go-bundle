# ClusterRestoreFromSnapshot

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**SourceSnapshotId** | **string** | The UUID of the snapshot to restore from. | |
|**RecoveryTargetDatetime** | [**time.Time**](time.Time.md) | Provide an ISO 8601 timestamp to restore from the most recent snapshot taken at or before that time. In-Memory DB does not provide continuous point-in-time recovery; the nearest preceding snapshot is used. This field is required for in-place restore.  | |

## Methods

### NewClusterRestoreFromSnapshot

`func NewClusterRestoreFromSnapshot(sourceSnapshotId string, recoveryTargetDatetime time.Time, ) *ClusterRestoreFromSnapshot`

NewClusterRestoreFromSnapshot instantiates a new ClusterRestoreFromSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRestoreFromSnapshotWithDefaults

`func NewClusterRestoreFromSnapshotWithDefaults() *ClusterRestoreFromSnapshot`

NewClusterRestoreFromSnapshotWithDefaults instantiates a new ClusterRestoreFromSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceSnapshotId

`func (o *ClusterRestoreFromSnapshot) GetSourceSnapshotId() string`

GetSourceSnapshotId returns the SourceSnapshotId field if non-nil, zero value otherwise.

### GetSourceSnapshotIdOk

`func (o *ClusterRestoreFromSnapshot) GetSourceSnapshotIdOk() (*string, bool)`

GetSourceSnapshotIdOk returns a tuple with the SourceSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSnapshotId

`func (o *ClusterRestoreFromSnapshot) SetSourceSnapshotId(v string)`

SetSourceSnapshotId sets SourceSnapshotId field to given value.


### GetRecoveryTargetDatetime

`func (o *ClusterRestoreFromSnapshot) GetRecoveryTargetDatetime() time.Time`

GetRecoveryTargetDatetime returns the RecoveryTargetDatetime field if non-nil, zero value otherwise.

### GetRecoveryTargetDatetimeOk

`func (o *ClusterRestoreFromSnapshot) GetRecoveryTargetDatetimeOk() (*time.Time, bool)`

GetRecoveryTargetDatetimeOk returns a tuple with the RecoveryTargetDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetDatetime

`func (o *ClusterRestoreFromSnapshot) SetRecoveryTargetDatetime(v time.Time)`

SetRecoveryTargetDatetime sets RecoveryTargetDatetime field to given value.



