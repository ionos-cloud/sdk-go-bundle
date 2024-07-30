# TemplateList

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Type** | Pointer to [**ResourceType**](ResourceType.md) |  | [optional] |
|**Id** | Pointer to **string** | The unique ID of the resource. | [optional] |
|**Items** | Pointer to [**[]TemplateResponse**](TemplateResponse.md) |  | [optional] |
|**Offset** | Pointer to **int32** | The offset specified in the request (if none was specified, the default offset is 0).  | [optional] [default to 0]|
|**Limit** | Pointer to **int32** | The limit specified in the request (if none was specified, the default limit is 100).  | [optional] [default to 100]|
|**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] |

## Methods

### NewTemplateList

`func NewTemplateList() *TemplateList`

NewTemplateList instantiates a new TemplateList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateListWithDefaults

`func NewTemplateListWithDefaults() *TemplateList`

NewTemplateListWithDefaults instantiates a new TemplateList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TemplateList) GetType() ResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TemplateList) GetTypeOk() (*ResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TemplateList) SetType(v ResourceType)`

SetType sets Type field to given value.

### HasType

`func (o *TemplateList) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *TemplateList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TemplateList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TemplateList) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TemplateList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *TemplateList) GetItems() []TemplateResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TemplateList) GetItemsOk() (*[]TemplateResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TemplateList) SetItems(v []TemplateResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *TemplateList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOffset

`func (o *TemplateList) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TemplateList) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TemplateList) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TemplateList) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *TemplateList) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TemplateList) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TemplateList) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TemplateList) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLinks

`func (o *TemplateList) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TemplateList) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TemplateList) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TemplateList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


