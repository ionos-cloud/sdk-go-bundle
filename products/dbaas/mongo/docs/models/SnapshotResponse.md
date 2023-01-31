# SnapshotResponse

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Type** | Pointer to [**ResourceType**](ResourceType.md) |  | [optional] |
|**Id** | Pointer to **string** | The unique ID of the resource. | [optional] |
|**Properties** | Pointer to [**SnapshotProperties**](SnapshotProperties.md) |  | [optional] |

## Methods

### NewSnapshotResponse

`func NewSnapshotResponse() *SnapshotResponse`

NewSnapshotResponse instantiates a new SnapshotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotResponseWithDefaults

`func NewSnapshotResponseWithDefaults() *SnapshotResponse`

NewSnapshotResponseWithDefaults instantiates a new SnapshotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SnapshotResponse) GetType() ResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnapshotResponse) GetTypeOk() (*ResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnapshotResponse) SetType(v ResourceType)`

SetType sets Type field to given value.

### HasType

`func (o *SnapshotResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *SnapshotResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnapshotResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnapshotResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SnapshotResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProperties

`func (o *SnapshotResponse) GetProperties() SnapshotProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SnapshotResponse) GetPropertiesOk() (*SnapshotProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SnapshotResponse) SetProperties(v SnapshotProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SnapshotResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


