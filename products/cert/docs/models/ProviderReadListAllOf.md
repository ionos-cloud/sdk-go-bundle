# ProviderReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of Provider resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of Provider resources. | |
|**Items** | Pointer to [**[]ProviderRead**](ProviderRead.md) | The list of Provider resources. | [optional] |

## Methods

### NewProviderReadListAllOf

`func NewProviderReadListAllOf(id string, type_ string, href string, ) *ProviderReadListAllOf`

NewProviderReadListAllOf instantiates a new ProviderReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderReadListAllOfWithDefaults

`func NewProviderReadListAllOfWithDefaults() *ProviderReadListAllOf`

NewProviderReadListAllOfWithDefaults instantiates a new ProviderReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProviderReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ProviderReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProviderReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProviderReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *ProviderReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ProviderReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ProviderReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *ProviderReadListAllOf) GetItems() []ProviderRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ProviderReadListAllOf) GetItemsOk() (*[]ProviderRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ProviderReadListAllOf) SetItems(v []ProviderRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *ProviderReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


