# SnapshotReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of Snapshot resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of Snapshot resources. | |
|**Items** | Pointer to [**[]SnapshotRead**](SnapshotRead.md) | The list of Snapshot resources. | [optional] |

## Methods

### NewSnapshotReadListAllOf

`func NewSnapshotReadListAllOf(id string, type_ string, href string, ) *SnapshotReadListAllOf`

NewSnapshotReadListAllOf instantiates a new SnapshotReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotReadListAllOfWithDefaults

`func NewSnapshotReadListAllOfWithDefaults() *SnapshotReadListAllOf`

NewSnapshotReadListAllOfWithDefaults instantiates a new SnapshotReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SnapshotReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnapshotReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnapshotReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SnapshotReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnapshotReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnapshotReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *SnapshotReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SnapshotReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SnapshotReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *SnapshotReadListAllOf) GetItems() []SnapshotRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SnapshotReadListAllOf) GetItemsOk() (*[]SnapshotRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SnapshotReadListAllOf) SetItems(v []SnapshotRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *SnapshotReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


