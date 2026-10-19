# SnapshotRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the Snapshot. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the Snapshot. | |
|**Metadata** | **map[string]interface{}** |  | [readonly] |
|**Properties** | [**Snapshot**](Snapshot.md) |  | |

## Methods

### NewSnapshotRead

`func NewSnapshotRead(id string, type_ string, href string, metadata map[string]interface{}, properties Snapshot, ) *SnapshotRead`

NewSnapshotRead instantiates a new SnapshotRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotReadWithDefaults

`func NewSnapshotReadWithDefaults() *SnapshotRead`

NewSnapshotReadWithDefaults instantiates a new SnapshotRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SnapshotRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnapshotRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnapshotRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SnapshotRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnapshotRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnapshotRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *SnapshotRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SnapshotRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SnapshotRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *SnapshotRead) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SnapshotRead) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SnapshotRead) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *SnapshotRead) GetProperties() Snapshot`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SnapshotRead) GetPropertiesOk() (*Snapshot, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SnapshotRead) SetProperties(v Snapshot)`

SetProperties sets Properties field to given value.



