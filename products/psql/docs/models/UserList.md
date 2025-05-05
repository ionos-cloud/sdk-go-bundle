# UserList

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Offset** | Pointer to **int32** | The offset specified in the request (if none was specified, the default offset is 0).  | [optional] [default to 0]|
|**Limit** | Pointer to **int32** | The limit specified in the request (if none was specified, the default limit is 100).  | [optional] [default to 100]|
|**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] |
|**Type** | [**ResourceType**](ResourceType.md) |  | |
|**Id** | **string** | The unique ID of the resource. | |
|**Href** | **string** | Absolute URL of the resource. | |
|**Items** | [**[]UserResource**](UserResource.md) |  | |

## Methods

### NewUserList

`func NewUserList(type_ ResourceType, id string, href string, items []UserResource, ) *UserList`

NewUserList instantiates a new UserList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserListWithDefaults

`func NewUserListWithDefaults() *UserList`

NewUserListWithDefaults instantiates a new UserList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *UserList) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *UserList) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *UserList) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *UserList) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *UserList) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *UserList) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *UserList) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *UserList) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLinks

`func (o *UserList) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserList) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserList) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetType

`func (o *UserList) GetType() ResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserList) GetTypeOk() (*ResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserList) SetType(v ResourceType)`

SetType sets Type field to given value.


### GetId

`func (o *UserList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserList) SetId(v string)`

SetId sets Id field to given value.


### GetHref

`func (o *UserList) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *UserList) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *UserList) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *UserList) GetItems() []UserResource`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UserList) GetItemsOk() (*[]UserResource, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UserList) SetItems(v []UserResource)`

SetItems sets Items field to given value.



