# RegionReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of Region resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of Region resources. | |
|**Items** | Pointer to [**[]RegionRead**](RegionRead.md) | The list of Region resources. | [optional] |

## Methods

### NewRegionReadListAllOf

`func NewRegionReadListAllOf(id string, type_ string, href string, ) *RegionReadListAllOf`

NewRegionReadListAllOf instantiates a new RegionReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionReadListAllOfWithDefaults

`func NewRegionReadListAllOfWithDefaults() *RegionReadListAllOf`

NewRegionReadListAllOfWithDefaults instantiates a new RegionReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RegionReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegionReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegionReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *RegionReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegionReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegionReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *RegionReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RegionReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RegionReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *RegionReadListAllOf) GetItems() []RegionRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *RegionReadListAllOf) GetItemsOk() (*[]RegionRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *RegionReadListAllOf) SetItems(v []RegionRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *RegionReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


