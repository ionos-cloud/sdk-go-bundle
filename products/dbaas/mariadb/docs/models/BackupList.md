# BackupList

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The unique ID of the resource. | [optional] |
|**Items** | Pointer to [**[]BackupResponse**](BackupResponse.md) |  | [optional] |
|**Offset** | Pointer to **int32** | The offset specified in the request (if none was specified, the default offset is 0).  | [optional] [default to 0]|
|**Limit** | Pointer to **int32** | The limit specified in the request (if none was specified, the default limit is 100).  | [optional] [default to 100]|
|**Total** | Pointer to **int32** | The total number of elements matching the request (without pagination).  | [optional] [default to 0]|
|**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] |

## Methods

### NewBackupList

`func NewBackupList() *BackupList`

NewBackupList instantiates a new BackupList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupListWithDefaults

`func NewBackupListWithDefaults() *BackupList`

NewBackupListWithDefaults instantiates a new BackupList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BackupList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BackupList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BackupList) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BackupList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *BackupList) GetItems() []BackupResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BackupList) GetItemsOk() (*[]BackupResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BackupList) SetItems(v []BackupResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *BackupList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOffset

`func (o *BackupList) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *BackupList) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *BackupList) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *BackupList) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *BackupList) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *BackupList) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *BackupList) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *BackupList) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotal

`func (o *BackupList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BackupList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BackupList) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BackupList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetLinks

`func (o *BackupList) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BackupList) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BackupList) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BackupList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


