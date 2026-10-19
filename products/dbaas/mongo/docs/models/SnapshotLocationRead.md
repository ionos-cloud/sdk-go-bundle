# SnapshotLocationRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the SnapshotLocation. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the SnapshotLocation. | |
|**Metadata** | **map[string]interface{}** |  | [readonly] |
|**Properties** | [**SnapshotLocation**](SnapshotLocation.md) |  | |

## Methods

### NewSnapshotLocationRead

`func NewSnapshotLocationRead(id string, type_ string, href string, metadata map[string]interface{}, properties SnapshotLocation, ) *SnapshotLocationRead`

NewSnapshotLocationRead instantiates a new SnapshotLocationRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotLocationReadWithDefaults

`func NewSnapshotLocationReadWithDefaults() *SnapshotLocationRead`

NewSnapshotLocationReadWithDefaults instantiates a new SnapshotLocationRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SnapshotLocationRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnapshotLocationRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnapshotLocationRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SnapshotLocationRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnapshotLocationRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnapshotLocationRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *SnapshotLocationRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SnapshotLocationRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SnapshotLocationRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *SnapshotLocationRead) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SnapshotLocationRead) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SnapshotLocationRead) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *SnapshotLocationRead) GetProperties() SnapshotLocation`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SnapshotLocationRead) GetPropertiesOk() (*SnapshotLocation, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SnapshotLocationRead) SetProperties(v SnapshotLocation)`

SetProperties sets Properties field to given value.



