# TokensResponse

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Links** | [**PaginationLinks**](PaginationLinks.md) |  | |
|**Count** | **int32** |  | |
|**Href** | Pointer to **string** |  | [optional] |
|**Id** | Pointer to **string** |  | [optional] |
|**Items** | Pointer to [**[]TokenResponse**](TokenResponse.md) |  | [optional] |
|**Limit** | **int32** |  | |
|**Offset** | **int32** |  | |
|**Total** | **int32** |  | |
|**Type** | Pointer to **string** |  | [optional] |

## Methods

### NewTokensResponse

`func NewTokensResponse(links PaginationLinks, count int32, limit int32, offset int32, total int32, ) *TokensResponse`

NewTokensResponse instantiates a new TokensResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokensResponseWithDefaults

`func NewTokensResponseWithDefaults() *TokensResponse`

NewTokensResponseWithDefaults instantiates a new TokensResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *TokensResponse) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TokensResponse) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TokensResponse) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.


### GetCount

`func (o *TokensResponse) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TokensResponse) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TokensResponse) SetCount(v int32)`

SetCount sets Count field to given value.


### GetHref

`func (o *TokensResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *TokensResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *TokensResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *TokensResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *TokensResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokensResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokensResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TokensResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *TokensResponse) GetItems() []TokenResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TokensResponse) GetItemsOk() (*[]TokenResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TokensResponse) SetItems(v []TokenResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *TokensResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLimit

`func (o *TokensResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TokensResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TokensResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetOffset

`func (o *TokensResponse) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TokensResponse) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TokensResponse) SetOffset(v int32)`

SetOffset sets Offset field to given value.


### GetTotal

`func (o *TokensResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TokensResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TokensResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetType

`func (o *TokensResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokensResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokensResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TokensResponse) HasType() bool`

HasType returns a boolean if a field has been set.


