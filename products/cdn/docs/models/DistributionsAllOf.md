# DistributionsAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of Distribution resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of Distribution resources. | |
|**Items** | Pointer to [**[]Distribution**](Distribution.md) | The list of Distribution resources. | [optional] |

## Methods

### NewDistributionsAllOf

`func NewDistributionsAllOf(id string, type_ string, href string, ) *DistributionsAllOf`

NewDistributionsAllOf instantiates a new DistributionsAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributionsAllOfWithDefaults

`func NewDistributionsAllOfWithDefaults() *DistributionsAllOf`

NewDistributionsAllOfWithDefaults instantiates a new DistributionsAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DistributionsAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DistributionsAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DistributionsAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *DistributionsAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DistributionsAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DistributionsAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *DistributionsAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DistributionsAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DistributionsAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *DistributionsAllOf) GetItems() []Distribution`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DistributionsAllOf) GetItemsOk() (*[]Distribution, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DistributionsAllOf) SetItems(v []Distribution)`

SetItems sets Items field to given value.

### HasItems

`func (o *DistributionsAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


