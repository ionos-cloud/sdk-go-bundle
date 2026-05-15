# DeprecatedPagination

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Offset** | Pointer to **int32** | The offset specified in the request (if none was specified, the default offset is 0).  | [optional] [default to 0]|
|**Limit** | Pointer to **int32** | The limit specified in the request (if none was specified, the default limit is 100).  | [optional] [default to 100]|
|**Links** | Pointer to [**PaginationLinks**](PaginationLinks.md) |  | [optional] |

## Methods

### NewDeprecatedPagination

`func NewDeprecatedPagination() *DeprecatedPagination`

NewDeprecatedPagination instantiates a new DeprecatedPagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeprecatedPaginationWithDefaults

`func NewDeprecatedPaginationWithDefaults() *DeprecatedPagination`

NewDeprecatedPaginationWithDefaults instantiates a new DeprecatedPagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffset

`func (o *DeprecatedPagination) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *DeprecatedPagination) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *DeprecatedPagination) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *DeprecatedPagination) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *DeprecatedPagination) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *DeprecatedPagination) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *DeprecatedPagination) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *DeprecatedPagination) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLinks

`func (o *DeprecatedPagination) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DeprecatedPagination) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DeprecatedPagination) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DeprecatedPagination) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


