# ClusterList

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The unique ID of the resource. | [optional] |
|**Items** | Pointer to [**[]ClusterResponse**](ClusterResponse.md) |  | [optional] |
|**Offset** | Pointer to **int32** | The offset specified in the request (if none was specified, the default offset is 0).  | [optional] [default to 0]|
|**Limit** | Pointer to **int32** | The limit specified in the request (if none was specified, the default limit is 100).  | [optional] [default to 100]|
|**Total** | Pointer to **int32** | The total number of elements matching the request (without pagination).  | [optional] [default to 0]|
|**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] |

## Methods

### NewClusterList

`func NewClusterList() *ClusterList`

NewClusterList instantiates a new ClusterList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterListWithDefaults

`func NewClusterListWithDefaults() *ClusterList`

NewClusterListWithDefaults instantiates a new ClusterList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterList) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *ClusterList) GetItems() []ClusterResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ClusterList) GetItemsOk() (*[]ClusterResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ClusterList) SetItems(v []ClusterResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *ClusterList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOffset

`func (o *ClusterList) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *ClusterList) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *ClusterList) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *ClusterList) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *ClusterList) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ClusterList) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ClusterList) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *ClusterList) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotal

`func (o *ClusterList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ClusterList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ClusterList) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ClusterList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetLinks

`func (o *ClusterList) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ClusterList) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ClusterList) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ClusterList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


