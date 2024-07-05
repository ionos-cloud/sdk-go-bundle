# UserItems

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Items** | [**[]UserResource**](UserResource.md) |  | |

## Methods

### NewUserItems

`func NewUserItems(items []UserResource, ) *UserItems`

NewUserItems instantiates a new UserItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserItemsWithDefaults

`func NewUserItemsWithDefaults() *UserItems`

NewUserItemsWithDefaults instantiates a new UserItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *UserItems) GetItems() []UserResource`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UserItems) GetItemsOk() (*[]UserResource, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UserItems) SetItems(v []UserResource)`

SetItems sets Items field to given value.



