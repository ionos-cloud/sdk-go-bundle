# SnapshotLocationReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of SnapshotLocation resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of SnapshotLocation resources. | |
|**Items** | Pointer to [**[]SnapshotLocationRead**](SnapshotLocationRead.md) | The list of SnapshotLocation resources. | [optional] |

## Methods

### NewSnapshotLocationReadListAllOf

`func NewSnapshotLocationReadListAllOf(id string, type_ string, href string, ) *SnapshotLocationReadListAllOf`

NewSnapshotLocationReadListAllOf instantiates a new SnapshotLocationReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotLocationReadListAllOfWithDefaults

`func NewSnapshotLocationReadListAllOfWithDefaults() *SnapshotLocationReadListAllOf`

NewSnapshotLocationReadListAllOfWithDefaults instantiates a new SnapshotLocationReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SnapshotLocationReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnapshotLocationReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnapshotLocationReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SnapshotLocationReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnapshotLocationReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnapshotLocationReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *SnapshotLocationReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SnapshotLocationReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SnapshotLocationReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *SnapshotLocationReadListAllOf) GetItems() []SnapshotLocationRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SnapshotLocationReadListAllOf) GetItemsOk() (*[]SnapshotLocationRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SnapshotLocationReadListAllOf) SetItems(v []SnapshotLocationRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *SnapshotLocationReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


