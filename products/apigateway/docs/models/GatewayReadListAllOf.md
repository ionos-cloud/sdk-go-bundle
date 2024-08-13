# GatewayReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the Gateway. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the Gateway. | |
|**Items** | Pointer to [**[]GatewayRead**](GatewayRead.md) | The list of Gateway resources. | [optional] |

## Methods

### NewGatewayReadListAllOf

`func NewGatewayReadListAllOf(id string, type_ string, href string, ) *GatewayReadListAllOf`

NewGatewayReadListAllOf instantiates a new GatewayReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayReadListAllOfWithDefaults

`func NewGatewayReadListAllOfWithDefaults() *GatewayReadListAllOf`

NewGatewayReadListAllOfWithDefaults instantiates a new GatewayReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GatewayReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *GatewayReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *GatewayReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *GatewayReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *GatewayReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *GatewayReadListAllOf) GetItems() []GatewayRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GatewayReadListAllOf) GetItemsOk() (*[]GatewayRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GatewayReadListAllOf) SetItems(v []GatewayRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *GatewayReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


