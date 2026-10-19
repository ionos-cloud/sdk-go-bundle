# SnapshotReadList

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of Snapshot resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of Snapshot resources. | |
|**Items** | Pointer to [**[]SnapshotRead**](SnapshotRead.md) | The list of Snapshot resources. | [optional] |
|**Offset** | **int32** | The offset specified in the request (if none was specified, the default offset is 0).  | [readonly] |
|**Limit** | **int32** | The limit specified in the request (if none was specified, use the endpoint&#39;s default pagination limit).  | [readonly] |
|**Links** | [**Links**](Links.md) |  | |

## Methods

### NewSnapshotReadList

`func NewSnapshotReadList(id string, type_ string, href string, offset int32, limit int32, links Links, ) *SnapshotReadList`

NewSnapshotReadList instantiates a new SnapshotReadList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotReadListWithDefaults

`func NewSnapshotReadListWithDefaults() *SnapshotReadList`

NewSnapshotReadListWithDefaults instantiates a new SnapshotReadList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SnapshotReadList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnapshotReadList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnapshotReadList) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *SnapshotReadList) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SnapshotReadList) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SnapshotReadList) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *SnapshotReadList) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *SnapshotReadList) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *SnapshotReadList) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *SnapshotReadList) GetItems() []SnapshotRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SnapshotReadList) GetItemsOk() (*[]SnapshotRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SnapshotReadList) SetItems(v []SnapshotRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *SnapshotReadList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOffset

`func (o *SnapshotReadList) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SnapshotReadList) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SnapshotReadList) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *SnapshotReadList) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SnapshotReadList) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SnapshotReadList) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetLinks

`func (o *SnapshotReadList) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SnapshotReadList) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SnapshotReadList) SetLinks(v Links)`

SetLinks sets Links field to given value.



