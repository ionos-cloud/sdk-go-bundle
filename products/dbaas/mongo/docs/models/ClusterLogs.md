# ClusterLogs

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Instances** | Pointer to [**[]ClusterLogsInstances**](ClusterLogsInstances.md) |  | [optional] |

## Methods

### NewClusterLogs

`func NewClusterLogs() *ClusterLogs`

NewClusterLogs instantiates a new ClusterLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLogsWithDefaults

`func NewClusterLogsWithDefaults() *ClusterLogs`

NewClusterLogsWithDefaults instantiates a new ClusterLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *ClusterLogs) GetInstances() []ClusterLogsInstances`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *ClusterLogs) GetInstancesOk() (*[]ClusterLogsInstances, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *ClusterLogs) SetInstances(v []ClusterLogsInstances)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *ClusterLogs) HasInstances() bool`

HasInstances returns a boolean if a field has been set.


