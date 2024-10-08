# LocationsResponse

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Href** | Pointer to **string** |  | [optional] |
|**Id** | Pointer to **string** |  | [optional] |
|**Items** | Pointer to [**[]Location**](Location.md) |  | [optional] |
|**Type** | Pointer to **string** |  | [optional] |

## Methods

### NewLocationsResponse

`func NewLocationsResponse() *LocationsResponse`

NewLocationsResponse instantiates a new LocationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationsResponseWithDefaults

`func NewLocationsResponseWithDefaults() *LocationsResponse`

NewLocationsResponseWithDefaults instantiates a new LocationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *LocationsResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LocationsResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LocationsResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *LocationsResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *LocationsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocationsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocationsResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LocationsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItems

`func (o *LocationsResponse) GetItems() []Location`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *LocationsResponse) GetItemsOk() (*[]Location, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *LocationsResponse) SetItems(v []Location)`

SetItems sets Items field to given value.

### HasItems

`func (o *LocationsResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetType

`func (o *LocationsResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LocationsResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LocationsResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LocationsResponse) HasType() bool`

HasType returns a boolean if a field has been set.


