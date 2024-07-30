# WireguardGatewayReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of WireguardGateway resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of WireguardGateway resources. | |
|**Items** | Pointer to [**[]WireguardGatewayRead**](WireguardGatewayRead.md) | The list of WireguardGateway resources. | [optional] |

## Methods

### NewWireguardGatewayReadListAllOf

`func NewWireguardGatewayReadListAllOf(id string, type_ string, href string, ) *WireguardGatewayReadListAllOf`

NewWireguardGatewayReadListAllOf instantiates a new WireguardGatewayReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireguardGatewayReadListAllOfWithDefaults

`func NewWireguardGatewayReadListAllOfWithDefaults() *WireguardGatewayReadListAllOf`

NewWireguardGatewayReadListAllOfWithDefaults instantiates a new WireguardGatewayReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WireguardGatewayReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WireguardGatewayReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WireguardGatewayReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *WireguardGatewayReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WireguardGatewayReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WireguardGatewayReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *WireguardGatewayReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *WireguardGatewayReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *WireguardGatewayReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *WireguardGatewayReadListAllOf) GetItems() []WireguardGatewayRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *WireguardGatewayReadListAllOf) GetItemsOk() (*[]WireguardGatewayRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *WireguardGatewayReadListAllOf) SetItems(v []WireguardGatewayRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *WireguardGatewayReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


