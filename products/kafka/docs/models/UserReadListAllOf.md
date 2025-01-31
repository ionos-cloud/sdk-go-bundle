# UserReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of User resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of User resources. | |
|**Items** | Pointer to [**[]UserRead**](UserRead.md) | The list of User resources. | [optional] |

## Methods

### NewUserReadListAllOf

`func NewUserReadListAllOf(id string, type_ string, href string, ) *UserReadListAllOf`

NewUserReadListAllOf instantiates a new UserReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserReadListAllOfWithDefaults

`func NewUserReadListAllOfWithDefaults() *UserReadListAllOf`

NewUserReadListAllOfWithDefaults instantiates a new UserReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *UserReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *UserReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *UserReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *UserReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *UserReadListAllOf) GetItems() []UserRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UserReadListAllOf) GetItemsOk() (*[]UserRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UserReadListAllOf) SetItems(v []UserRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *UserReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


