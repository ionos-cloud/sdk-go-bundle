# ProviderReadList

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of Provider resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of Provider resources. | |
|**Items** | Pointer to [**[]ProviderRead**](ProviderRead.md) | The list of Provider resources. | [optional] |
|**Offset** | **int32** | The offset specified in the request (if none was specified, the default offset is 0).  | [readonly] |
|**Limit** | **int32** | The limit specified in the request (if none was specified, use the endpoint&#39;s default pagination limit).  | [readonly] |
|**Links** | [**Links**](Links.md) |  | |

## Methods

### NewProviderReadList

`func NewProviderReadList(id string, type_ string, href string, offset int32, limit int32, links Links, ) *ProviderReadList`

NewProviderReadList instantiates a new ProviderReadList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderReadListWithDefaults

`func NewProviderReadListWithDefaults() *ProviderReadList`

NewProviderReadListWithDefaults instantiates a new ProviderReadList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProviderReadList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderReadList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderReadList) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ProviderReadList) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProviderReadList) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProviderReadList) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *ProviderReadList) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ProviderReadList) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ProviderReadList) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *ProviderReadList) GetItems() []ProviderRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ProviderReadList) GetItemsOk() (*[]ProviderRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ProviderReadList) SetItems(v []ProviderRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *ProviderReadList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOffset

`func (o *ProviderReadList) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ProviderReadList) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ProviderReadList) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *ProviderReadList) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ProviderReadList) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ProviderReadList) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetLinks

`func (o *ProviderReadList) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProviderReadList) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProviderReadList) SetLinks(v Links)`

SetLinks sets Links field to given value.



