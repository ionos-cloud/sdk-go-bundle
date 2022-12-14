# RegistriesResponse

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Links** | [**PaginationLinks**](PaginationLinks.md) |  | |
|**Href** | Pointer to **string** |  | [optional] |
|**Id** | Pointer to **string** |  | [optional] |
|**Items** | Pointer to [**[]RegistryResponse**](RegistryResponse.md) |  | [optional] |
|**Pagination** | [**Pagination**](Pagination.md) |  | |
|**Type** | Pointer to **string** |  | [optional] |

## Methods

### NewRegistriesResponse

`func NewRegistriesResponse(links PaginationLinks, pagination Pagination, ) *RegistriesResponse`

NewRegistriesResponse instantiates a new RegistriesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistriesResponseWithDefaults

`func NewRegistriesResponseWithDefaults() *RegistriesResponse`

NewRegistriesResponseWithDefaults instantiates a new RegistriesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *RegistriesResponse) GetLinks() PaginationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RegistriesResponse) GetLinksOk() (*PaginationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RegistriesResponse) SetLinks(v PaginationLinks)`

SetLinks sets Links field to given value.


### GetHref

`func (o *RegistriesResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RegistriesResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RegistriesResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RegistriesResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *RegistriesResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegistriesResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegistriesResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RegistriesResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *RegistriesResponse) GetItems() []RegistryResponse`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RegistriesResponse) GetItemsOk() (*[]RegistryResponse, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RegistriesResponse) SetItems(v []RegistryResponse)`

SetItems sets Items field to given value.

### HasItems

`func (o *RegistriesResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### SetItemsNil

`func (o *RegistriesResponse) SetItemsNil(b bool)`

 SetItemsNil sets the value for Items to be an explicit nil

### UnsetItems
`func (o *RegistriesResponse) UnsetItems()`

UnsetItems ensures that no value is present for Items, not even an explicit nil
### GetPagination

`func (o *RegistriesResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *RegistriesResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *RegistriesResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.


### GetType

`func (o *RegistriesResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegistriesResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegistriesResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RegistriesResponse) HasType() bool`

HasType returns a boolean if a field has been set.


