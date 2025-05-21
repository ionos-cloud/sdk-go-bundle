# SnapshotMetadataAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ReplicasetId** | **string** | The ID of the In-Memory DB replica set the snapshot is taken from.  | |
|**SnapshotTime** | Pointer to [**time.Time**](time.Time.md) | The time the snapshot was dumped from the replica set. | [optional] |
|**DatacenterId** | **string** | The ID of the datacenter the snapshot was created in. Please mind, that the snapshot is not available in other datacenters.  | |

## Methods

### NewSnapshotMetadataAllOf

`func NewSnapshotMetadataAllOf(replicasetId string, datacenterId string, ) *SnapshotMetadataAllOf`

NewSnapshotMetadataAllOf instantiates a new SnapshotMetadataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotMetadataAllOfWithDefaults

`func NewSnapshotMetadataAllOfWithDefaults() *SnapshotMetadataAllOf`

NewSnapshotMetadataAllOfWithDefaults instantiates a new SnapshotMetadataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplicasetId

`func (o *SnapshotMetadataAllOf) GetReplicasetId() string`

GetReplicasetId returns the ReplicasetId field if non-nil, zero value otherwise.

### GetReplicasetIdOk

`func (o *SnapshotMetadataAllOf) GetReplicasetIdOk() (*string, bool)`

GetReplicasetIdOk returns a tuple with the ReplicasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicasetId

`func (o *SnapshotMetadataAllOf) SetReplicasetId(v string)`

SetReplicasetId sets ReplicasetId field to given value.


### GetSnapshotTime

`func (o *SnapshotMetadataAllOf) GetSnapshotTime() time.Time`

GetSnapshotTime returns the SnapshotTime field if non-nil, zero value otherwise.

### GetSnapshotTimeOk

`func (o *SnapshotMetadataAllOf) GetSnapshotTimeOk() (*time.Time, bool)`

GetSnapshotTimeOk returns a tuple with the SnapshotTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotTime

`func (o *SnapshotMetadataAllOf) SetSnapshotTime(v time.Time)`

SetSnapshotTime sets SnapshotTime field to given value.

### HasSnapshotTime

`func (o *SnapshotMetadataAllOf) HasSnapshotTime() bool`

HasSnapshotTime returns a boolean if a field has been set.

### GetDatacenterId

`func (o *SnapshotMetadataAllOf) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *SnapshotMetadataAllOf) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *SnapshotMetadataAllOf) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.



