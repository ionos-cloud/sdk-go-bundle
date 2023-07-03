# RecordReadList

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The resource&#39;s unique identifier. | [readonly] |
|**Type** | **string** |  | |
|**Href** | **string** |  | [readonly] |
|**Items** | [**[]RecordRead**](RecordRead.md) |  | |
|**Offset** | **float32** | Pagination offset. | [readonly] |
|**Limit** | **float32** | Pagination limit. | [readonly] |
|**Links** | [**Links**](Links.md) |  | |

## Methods

### NewRecordReadList

`func NewRecordReadList(id string, type_ string, href string, items []RecordRead, offset float32, limit float32, links Links, ) *RecordReadList`

NewRecordReadList instantiates a new RecordReadList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordReadListWithDefaults

`func NewRecordReadListWithDefaults() *RecordReadList`

NewRecordReadListWithDefaults instantiates a new RecordReadList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RecordReadList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RecordReadList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RecordReadList) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *RecordReadList) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RecordReadList) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RecordReadList) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *RecordReadList) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RecordReadList) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RecordReadList) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *RecordReadList) GetItems() []RecordRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RecordReadList) GetItemsOk() (*[]RecordRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RecordReadList) SetItems(v []RecordRead)`

SetItems sets Items field to given value.


### GetOffset

`func (o *RecordReadList) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *RecordReadList) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *RecordReadList) SetOffset(v float32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *RecordReadList) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RecordReadList) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RecordReadList) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetLinks

`func (o *RecordReadList) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RecordReadList) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RecordReadList) SetLinks(v Links)`

SetLinks sets Links field to given value.



