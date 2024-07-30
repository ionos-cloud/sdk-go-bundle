# IPSecTunnelReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of IPSecTunnel resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of IPSecTunnel resources. | |
|**Items** | Pointer to [**[]IPSecTunnelRead**](IPSecTunnelRead.md) | The list of IPSecTunnel resources. | [optional] |

## Methods

### NewIPSecTunnelReadListAllOf

`func NewIPSecTunnelReadListAllOf(id string, type_ string, href string, ) *IPSecTunnelReadListAllOf`

NewIPSecTunnelReadListAllOf instantiates a new IPSecTunnelReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecTunnelReadListAllOfWithDefaults

`func NewIPSecTunnelReadListAllOfWithDefaults() *IPSecTunnelReadListAllOf`

NewIPSecTunnelReadListAllOfWithDefaults instantiates a new IPSecTunnelReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IPSecTunnelReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IPSecTunnelReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IPSecTunnelReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *IPSecTunnelReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IPSecTunnelReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IPSecTunnelReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *IPSecTunnelReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *IPSecTunnelReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *IPSecTunnelReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *IPSecTunnelReadListAllOf) GetItems() []IPSecTunnelRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *IPSecTunnelReadListAllOf) GetItemsOk() (*[]IPSecTunnelRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *IPSecTunnelReadListAllOf) SetItems(v []IPSecTunnelRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *IPSecTunnelReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


