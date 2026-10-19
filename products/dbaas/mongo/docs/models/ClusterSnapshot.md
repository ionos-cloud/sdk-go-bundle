# ClusterSnapshot

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Location** | **string** | The Object Storage location where the snapshots will be created. For added data safety, use a different location than the cluster.  | |
|**RetentionDays** | **int32** | Configures how many days cluster snapshots are retained before being automatically deleted. Reducing this value purges existing snapshots that fall outside the new retention window.  | |

## Methods

### NewClusterSnapshot

`func NewClusterSnapshot(location string, retentionDays int32, ) *ClusterSnapshot`

NewClusterSnapshot instantiates a new ClusterSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterSnapshotWithDefaults

`func NewClusterSnapshotWithDefaults() *ClusterSnapshot`

NewClusterSnapshotWithDefaults instantiates a new ClusterSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *ClusterSnapshot) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ClusterSnapshot) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ClusterSnapshot) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetRetentionDays

`func (o *ClusterSnapshot) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *ClusterSnapshot) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *ClusterSnapshot) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.



