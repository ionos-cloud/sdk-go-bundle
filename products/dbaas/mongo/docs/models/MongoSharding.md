# MongoSharding

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Enabled** | Pointer to **bool** | A flag to activate or deactivate sharding for this cluster. Can only be set to true if the MongoDB edition is ENTERPRISE. | [optional] [default to false]|
|**Shards** | Pointer to **int32** | Number of shards. If sharding is enabled, this field specifies the number of shards to create in the cluster. The number of shards must be between 2 and 5. When sharding is enabled, the instances count signifies the number of replicas per shard.  | [optional] [default to 2]|

## Methods

### NewMongoSharding

`func NewMongoSharding() *MongoSharding`

NewMongoSharding instantiates a new MongoSharding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoShardingWithDefaults

`func NewMongoShardingWithDefaults() *MongoSharding`

NewMongoShardingWithDefaults instantiates a new MongoSharding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *MongoSharding) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MongoSharding) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MongoSharding) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MongoSharding) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetShards

`func (o *MongoSharding) GetShards() int32`

GetShards returns the Shards field if non-nil, zero value otherwise.

### GetShardsOk

`func (o *MongoSharding) GetShardsOk() (*int32, bool)`

GetShardsOk returns a tuple with the Shards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShards

`func (o *MongoSharding) SetShards(v int32)`

SetShards sets Shards field to given value.

### HasShards

`func (o *MongoSharding) HasShards() bool`

HasShards returns a boolean if a field has been set.


