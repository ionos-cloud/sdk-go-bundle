# ZoneReadList

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID (UUID) created to identify this action. | |
|**Type** | **string** |  | |
|**Href** | **string** |  | [readonly] |
|**Offset** | **float32** | Pagination offset. | [readonly] |
|**Limit** | **float32** | Pagination limit. | [readonly] |
|**Links** | [**Links**](Links.md) |  | |
|**Items** | [**[]ZoneRead**](ZoneRead.md) |  | |

## Methods

### NewZoneReadList

`func NewZoneReadList(id string, type_ string, href string, offset float32, limit float32, links Links, items []ZoneRead, ) *ZoneReadList`

NewZoneReadList instantiates a new ZoneReadList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneReadListWithDefaults

`func NewZoneReadListWithDefaults() *ZoneReadList`

NewZoneReadListWithDefaults instantiates a new ZoneReadList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ZoneReadList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ZoneReadList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ZoneReadList) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ZoneReadList) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ZoneReadList) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ZoneReadList) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *ZoneReadList) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ZoneReadList) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ZoneReadList) SetHref(v string)`

SetHref sets Href field to given value.


### GetOffset

`func (o *ZoneReadList) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ZoneReadList) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ZoneReadList) SetOffset(v float32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *ZoneReadList) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ZoneReadList) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ZoneReadList) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetLinks

`func (o *ZoneReadList) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ZoneReadList) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ZoneReadList) SetLinks(v Links)`

SetLinks sets Links field to given value.


### GetItems

`func (o *ZoneReadList) GetItems() []ZoneRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ZoneReadList) GetItemsOk() (*[]ZoneRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ZoneReadList) SetItems(v []ZoneRead)`

SetItems sets Items field to given value.



