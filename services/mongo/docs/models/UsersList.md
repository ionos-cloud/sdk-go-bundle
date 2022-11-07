# UsersList

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Type** | Pointer to [**ResourceType**](ResourceType.md) |  | [optional] |
|**Id** | Pointer to **string** | The unique ID of the resource. | [optional] |
|**Items** | Pointer to [**[]User**](User.md) |  | [optional] |

## Methods

### NewUsersList

`func NewUsersList() *UsersList`

NewUsersList instantiates a new UsersList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersListWithDefaults

`func NewUsersListWithDefaults() *UsersList`

NewUsersListWithDefaults instantiates a new UsersList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UsersList) GetType() ResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UsersList) GetTypeOk() (*ResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UsersList) SetType(v ResourceType)`

SetType sets Type field to given value.

### HasType

`func (o *UsersList) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *UsersList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsersList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsersList) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UsersList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *UsersList) GetItems() []User`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UsersList) GetItemsOk() (*[]User, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UsersList) SetItems(v []User)`

SetItems sets Items field to given value.

### HasItems

`func (o *UsersList) HasItems() bool`

HasItems returns a boolean if a field has been set.


