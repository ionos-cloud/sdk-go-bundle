# SnapshotEnsure

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the Snapshot. | |
|**Metadata** | Pointer to **map[string]interface{}** | Metadata | [optional] |
|**Properties** | **map[string]interface{}** | A point in time snapshot of a In-Memory DB replica set.  | |

## Methods

### NewSnapshotEnsure

`func NewSnapshotEnsure(id string, properties map[string]interface{}, ) *SnapshotEnsure`

NewSnapshotEnsure instantiates a new SnapshotEnsure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotEnsureWithDefaults

`func NewSnapshotEnsureWithDefaults() *SnapshotEnsure`

NewSnapshotEnsureWithDefaults instantiates a new SnapshotEnsure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SnapshotEnsure) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnapshotEnsure) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnapshotEnsure) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *SnapshotEnsure) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SnapshotEnsure) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SnapshotEnsure) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SnapshotEnsure) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *SnapshotEnsure) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SnapshotEnsure) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SnapshotEnsure) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.



