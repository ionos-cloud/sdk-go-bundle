# MariadbVersionReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of MariadbVersion resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of MariadbVersion resources. | |
|**Items** | Pointer to [**[]MariadbVersionRead**](MariadbVersionRead.md) | The list of MariadbVersion resources. | [optional] |

## Methods

### NewMariadbVersionReadListAllOf

`func NewMariadbVersionReadListAllOf(id string, type_ string, href string, ) *MariadbVersionReadListAllOf`

NewMariadbVersionReadListAllOf instantiates a new MariadbVersionReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMariadbVersionReadListAllOfWithDefaults

`func NewMariadbVersionReadListAllOfWithDefaults() *MariadbVersionReadListAllOf`

NewMariadbVersionReadListAllOfWithDefaults instantiates a new MariadbVersionReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MariadbVersionReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MariadbVersionReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MariadbVersionReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *MariadbVersionReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MariadbVersionReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MariadbVersionReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *MariadbVersionReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *MariadbVersionReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *MariadbVersionReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *MariadbVersionReadListAllOf) GetItems() []MariadbVersionRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MariadbVersionReadListAllOf) GetItemsOk() (*[]MariadbVersionRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MariadbVersionReadListAllOf) SetItems(v []MariadbVersionRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *MariadbVersionReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


