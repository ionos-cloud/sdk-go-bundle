# SecondaryZoneReadList

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID (UUID) created to identify this action. | |
|**Type** | **string** |  | |
|**Href** | **string** |  | [readonly] |
|**Offset** | **float32** | Pagination offset. | [readonly] |
|**Limit** | **float32** | Pagination limit. | [readonly] |
|**Links** | [**Links**](Links.md) |  | |
|**Items** | [**[]SecondaryZoneRead**](SecondaryZoneRead.md) |  | |

## Methods

### NewSecondaryZoneReadList

`func NewSecondaryZoneReadList(id string, type_ string, href string, offset float32, limit float32, links Links, items []SecondaryZoneRead, ) *SecondaryZoneReadList`

NewSecondaryZoneReadList instantiates a new SecondaryZoneReadList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecondaryZoneReadListWithDefaults

`func NewSecondaryZoneReadListWithDefaults() *SecondaryZoneReadList`

NewSecondaryZoneReadListWithDefaults instantiates a new SecondaryZoneReadList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecondaryZoneReadList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecondaryZoneReadList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecondaryZoneReadList) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SecondaryZoneReadList) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecondaryZoneReadList) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecondaryZoneReadList) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *SecondaryZoneReadList) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SecondaryZoneReadList) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SecondaryZoneReadList) SetHref(v string)`

SetHref sets Href field to given value.


### GetOffset

`func (o *SecondaryZoneReadList) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SecondaryZoneReadList) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SecondaryZoneReadList) SetOffset(v float32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *SecondaryZoneReadList) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SecondaryZoneReadList) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SecondaryZoneReadList) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetLinks

`func (o *SecondaryZoneReadList) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SecondaryZoneReadList) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SecondaryZoneReadList) SetLinks(v Links)`

SetLinks sets Links field to given value.


### GetItems

`func (o *SecondaryZoneReadList) GetItems() []SecondaryZoneRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SecondaryZoneReadList) GetItemsOk() (*[]SecondaryZoneRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SecondaryZoneReadList) SetItems(v []SecondaryZoneRead)`

SetItems sets Items field to given value.



