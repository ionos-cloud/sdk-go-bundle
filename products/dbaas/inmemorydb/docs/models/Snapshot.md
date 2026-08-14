# Snapshot

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ClusterId** | Pointer to **string** | The ID of the cluster this snapshot belongs to. | [optional] |
|**ClusterName** | Pointer to **string** | The name of the In-Memory DB cluster this snapshot belongs to. | [optional] [readonly] |
|**DatacenterId** | Pointer to **string** | The ID of the data center where the snapshot was created. Snapshots are not available across data centers.  | [optional] |
|**EarliestRecoveryTargetTime** | Pointer to [**time.Time**](time.Time.md) | The earliest time for which a snapshot is available to restore from. | [optional] |
|**LatestRecoveryTargetTime** | Pointer to [**NullableTime**](time.Time.md) | The most recent time for which a snapshot is available to restore from. If a snapshot is available up to the current time, this field will be null.  | [optional] |
|**Location** | Pointer to **string** | The Object Storage location where snapshots will be stored. For added data safety, use a different location than the cluster. A list of supported locations is provided by the SnapshotLocations endpoint.  | [optional] |
|**ClusterVersion** | Pointer to **string** | The version for the cluster. Use GET /versions to retrieve the list of supported versions. To upgrade, provide a version listed in canUpgradeTo for the current version. Downgrades are not supported.  | [optional] |
|**SnapshotSize** | Pointer to **float32** | The size of the snapshot in gigabytes (GB). | [optional] |
|**RequiredSizeForRestore** | Pointer to **float32** | The minimum storage size in gigabytes (GB) required on the target cluster to restore from this snapshot.  | [optional] |

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

### GetDatacenterId

`func (o *Snapshot) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *Snapshot) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *Snapshot) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *Snapshot) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

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

### GetClusterVersion

`func (o *Snapshot) GetClusterVersion() string`

GetClusterVersion returns the ClusterVersion field if non-nil, zero value otherwise.

### GetClusterVersionOk

`func (o *Snapshot) GetClusterVersionOk() (*string, bool)`

GetClusterVersionOk returns a tuple with the ClusterVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterVersion

`func (o *Snapshot) SetClusterVersion(v string)`

SetClusterVersion sets ClusterVersion field to given value.

### HasClusterVersion

`func (o *Snapshot) HasClusterVersion() bool`

HasClusterVersion returns a boolean if a field has been set.

### GetSnapshotSize

`func (o *Snapshot) GetSnapshotSize() float32`

GetSnapshotSize returns the SnapshotSize field if non-nil, zero value otherwise.

### GetSnapshotSizeOk

`func (o *Snapshot) GetSnapshotSizeOk() (*float32, bool)`

GetSnapshotSizeOk returns a tuple with the SnapshotSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotSize

`func (o *Snapshot) SetSnapshotSize(v float32)`

SetSnapshotSize sets SnapshotSize field to given value.

### HasSnapshotSize

`func (o *Snapshot) HasSnapshotSize() bool`

HasSnapshotSize returns a boolean if a field has been set.

### GetRequiredSizeForRestore

`func (o *Snapshot) GetRequiredSizeForRestore() float32`

GetRequiredSizeForRestore returns the RequiredSizeForRestore field if non-nil, zero value otherwise.

### GetRequiredSizeForRestoreOk

`func (o *Snapshot) GetRequiredSizeForRestoreOk() (*float32, bool)`

GetRequiredSizeForRestoreOk returns a tuple with the RequiredSizeForRestore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSizeForRestore

`func (o *Snapshot) SetRequiredSizeForRestore(v float32)`

SetRequiredSizeForRestore sets RequiredSizeForRestore field to given value.

### HasRequiredSizeForRestore

`func (o *Snapshot) HasRequiredSizeForRestore() bool`

HasRequiredSizeForRestore returns a boolean if a field has been set.


