# RegistriesResponse

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Links** | [**PaginationLinks**](PaginationLinks.md) |  | |
|**Href** | Pointer to **string** |  | [optional] |
|**Id** | Pointer to **string** |  | [optional] |
|**Items** | Pointer to [**[]RegistryResponse**](RegistryResponse.md) |  | [optional] |
|**Limit** | **int32** |  | |
|**NextPageToken** | **string** |  | |
|**Type** | Pointer to **string** |  | [optional] |

## Methods

### NewRegistriesResponse

`func NewRegistriesResponse(links PaginationLinks, limit int32, nextPageToken string, ) *RegistriesResponse`

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
### GetLimit

`func (o *RegistriesResponse) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RegistriesResponse) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RegistriesResponse) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetNextPageToken

`func (o *RegistriesResponse) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *RegistriesResponse) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *RegistriesResponse) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.


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


