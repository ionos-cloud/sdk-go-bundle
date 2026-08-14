# VersionReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of Version resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of Version resources. | |
|**Items** | Pointer to [**[]VersionRead**](VersionRead.md) | The list of Version resources. | [optional] |

## Methods

### NewVersionReadListAllOf

`func NewVersionReadListAllOf(id string, type_ string, href string, ) *VersionReadListAllOf`

NewVersionReadListAllOf instantiates a new VersionReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionReadListAllOfWithDefaults

`func NewVersionReadListAllOfWithDefaults() *VersionReadListAllOf`

NewVersionReadListAllOfWithDefaults instantiates a new VersionReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VersionReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VersionReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VersionReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *VersionReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VersionReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VersionReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *VersionReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *VersionReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *VersionReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *VersionReadListAllOf) GetItems() []VersionRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *VersionReadListAllOf) GetItemsOk() (*[]VersionRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *VersionReadListAllOf) SetItems(v []VersionRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *VersionReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


