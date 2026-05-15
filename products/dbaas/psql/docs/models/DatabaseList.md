# DatabaseList

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Offset** | Pointer to **int32** | The offset specified in the request (if none was specified, the default offset is 0).  | [optional] [default to 0]|
|**Limit** | Pointer to **int32** | The limit specified in the request (if none was specified, the default limit is 100).  | [optional] [default to 100]|
|**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] |
|**Type** | [**ResourceType**](ResourceType.md) |  | |
|**Id** | **string** | The unique ID of the resource. | |
|**Href** | **string** | Absolute URL of the resource. | |
|**Items** | [**[]DatabaseResource**](DatabaseResource.md) |  | |

## Methods

### NewDatabaseList

`func NewDatabaseList(type_ ResourceType, id string, href string, items []DatabaseResource, ) *DatabaseList`

NewDatabaseList instantiates a new DatabaseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseListWithDefaults

`func NewDatabaseListWithDefaults() *DatabaseList`

NewDatabaseListWithDefaults instantiates a new DatabaseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *DatabaseList) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *DatabaseList) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *DatabaseList) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *DatabaseList) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *DatabaseList) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *DatabaseList) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *DatabaseList) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *DatabaseList) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLinks

`func (o *DatabaseList) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DatabaseList) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DatabaseList) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DatabaseList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetType

`func (o *DatabaseList) GetType() ResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatabaseList) GetTypeOk() (*ResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatabaseList) SetType(v ResourceType)`

SetType sets Type field to given value.


### GetId

`func (o *DatabaseList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseList) SetId(v string)`

SetId sets Id field to given value.


### GetHref

`func (o *DatabaseList) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DatabaseList) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DatabaseList) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *DatabaseList) GetItems() []DatabaseResource`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DatabaseList) GetItemsOk() (*[]DatabaseResource, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DatabaseList) SetItems(v []DatabaseResource)`

SetItems sets Items field to given value.



