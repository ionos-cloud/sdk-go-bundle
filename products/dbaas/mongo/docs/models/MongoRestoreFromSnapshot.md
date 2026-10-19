# MongoRestoreFromSnapshot

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**SourceSnapshotId** | Pointer to **string** | UUID of the snapshot to restore from. Required for restore on cluster creation; not valid for in-place restore, where the source is inferred from the cluster&#39;s own snapshots.  | [optional] |
|**RecoveryTargetDatetime** | Pointer to [**time.Time**](time.Time.md) | An ISO 8601 timestamp for point-in-time recovery; the cluster is restored to this exact moment (within the available recovery window), not merely to the nearest snapshot. Optional on cluster creation — if omitted, the source snapshot (&#x60;sourceSnapshotId&#x60;) is applied in full; required for in-place restore.  | [optional] |

## Methods

### NewMongoRestoreFromSnapshot

`func NewMongoRestoreFromSnapshot() *MongoRestoreFromSnapshot`

NewMongoRestoreFromSnapshot instantiates a new MongoRestoreFromSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoRestoreFromSnapshotWithDefaults

`func NewMongoRestoreFromSnapshotWithDefaults() *MongoRestoreFromSnapshot`

NewMongoRestoreFromSnapshotWithDefaults instantiates a new MongoRestoreFromSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceSnapshotId

`func (o *MongoRestoreFromSnapshot) GetSourceSnapshotId() string`

GetSourceSnapshotId returns the SourceSnapshotId field if non-nil, zero value otherwise.

### GetSourceSnapshotIdOk

`func (o *MongoRestoreFromSnapshot) GetSourceSnapshotIdOk() (*string, bool)`

GetSourceSnapshotIdOk returns a tuple with the SourceSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSnapshotId

`func (o *MongoRestoreFromSnapshot) SetSourceSnapshotId(v string)`

SetSourceSnapshotId sets SourceSnapshotId field to given value.

### HasSourceSnapshotId

`func (o *MongoRestoreFromSnapshot) HasSourceSnapshotId() bool`

HasSourceSnapshotId returns a boolean if a field has been set.

### GetRecoveryTargetDatetime

`func (o *MongoRestoreFromSnapshot) GetRecoveryTargetDatetime() time.Time`

GetRecoveryTargetDatetime returns the RecoveryTargetDatetime field if non-nil, zero value otherwise.

### GetRecoveryTargetDatetimeOk

`func (o *MongoRestoreFromSnapshot) GetRecoveryTargetDatetimeOk() (*time.Time, bool)`

GetRecoveryTargetDatetimeOk returns a tuple with the RecoveryTargetDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoveryTargetDatetime

`func (o *MongoRestoreFromSnapshot) SetRecoveryTargetDatetime(v time.Time)`

SetRecoveryTargetDatetime sets RecoveryTargetDatetime field to given value.

### HasRecoveryTargetDatetime

`func (o *MongoRestoreFromSnapshot) HasRecoveryTargetDatetime() bool`

HasRecoveryTargetDatetime returns a boolean if a field has been set.


