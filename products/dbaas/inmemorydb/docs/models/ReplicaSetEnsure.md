# ReplicaSetEnsure

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the ReplicaSet. | |
|**Metadata** | Pointer to **map[string]interface{}** | Metadata | [optional] |
|**Properties** | [**ReplicaSet**](ReplicaSet.md) |  | |

## Methods

### NewReplicaSetEnsure

`func NewReplicaSetEnsure(id string, properties ReplicaSet, ) *ReplicaSetEnsure`

NewReplicaSetEnsure instantiates a new ReplicaSetEnsure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicaSetEnsureWithDefaults

`func NewReplicaSetEnsureWithDefaults() *ReplicaSetEnsure`

NewReplicaSetEnsureWithDefaults instantiates a new ReplicaSetEnsure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReplicaSetEnsure) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReplicaSetEnsure) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReplicaSetEnsure) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *ReplicaSetEnsure) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ReplicaSetEnsure) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ReplicaSetEnsure) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ReplicaSetEnsure) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *ReplicaSetEnsure) GetProperties() ReplicaSet`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ReplicaSetEnsure) GetPropertiesOk() (*ReplicaSet, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ReplicaSetEnsure) SetProperties(v ReplicaSet)`

SetProperties sets Properties field to given value.



