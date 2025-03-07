# ShareReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The Share identifier (UUID) | |
|**Type** | **string** | The resource type | |
|**Href** | **string** | The URL of the Share. | |
|**Items** | Pointer to [**[]ShareRead**](ShareRead.md) | The list of share resources. | [optional] |

## Methods

### NewShareReadListAllOf

`func NewShareReadListAllOf(id string, type_ string, href string, ) *ShareReadListAllOf`

NewShareReadListAllOf instantiates a new ShareReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareReadListAllOfWithDefaults

`func NewShareReadListAllOfWithDefaults() *ShareReadListAllOf`

NewShareReadListAllOfWithDefaults instantiates a new ShareReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShareReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShareReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShareReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ShareReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShareReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShareReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *ShareReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ShareReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ShareReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *ShareReadListAllOf) GetItems() []ShareRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ShareReadListAllOf) GetItemsOk() (*[]ShareRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ShareReadListAllOf) SetItems(v []ShareRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *ShareReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


