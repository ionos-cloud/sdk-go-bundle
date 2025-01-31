# UserReadList

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of User resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of User resources. | |
|**Items** | Pointer to [**[]UserRead**](UserRead.md) | The list of User resources. | [optional] |

## Methods

### NewUserReadList

`func NewUserReadList(id string, type_ string, href string, ) *UserReadList`

NewUserReadList instantiates a new UserReadList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserReadListWithDefaults

`func NewUserReadListWithDefaults() *UserReadList`

NewUserReadListWithDefaults instantiates a new UserReadList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserReadList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserReadList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserReadList) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *UserReadList) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserReadList) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserReadList) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *UserReadList) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *UserReadList) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *UserReadList) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *UserReadList) GetItems() []UserRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UserReadList) GetItemsOk() (*[]UserRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UserReadList) SetItems(v []UserRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *UserReadList) HasItems() bool`

HasItems returns a boolean if a field has been set.


