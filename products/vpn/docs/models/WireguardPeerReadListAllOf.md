# WireguardPeerReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of WireguardPeer resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of WireguardPeer resources. | |
|**Items** | Pointer to [**[]WireguardPeerRead**](WireguardPeerRead.md) | The list of WireguardPeer resources. | [optional] |

## Methods

### NewWireguardPeerReadListAllOf

`func NewWireguardPeerReadListAllOf(id string, type_ string, href string, ) *WireguardPeerReadListAllOf`

NewWireguardPeerReadListAllOf instantiates a new WireguardPeerReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireguardPeerReadListAllOfWithDefaults

`func NewWireguardPeerReadListAllOfWithDefaults() *WireguardPeerReadListAllOf`

NewWireguardPeerReadListAllOfWithDefaults instantiates a new WireguardPeerReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WireguardPeerReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WireguardPeerReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WireguardPeerReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *WireguardPeerReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WireguardPeerReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WireguardPeerReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *WireguardPeerReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *WireguardPeerReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *WireguardPeerReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *WireguardPeerReadListAllOf) GetItems() []WireguardPeerRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *WireguardPeerReadListAllOf) GetItemsOk() (*[]WireguardPeerRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *WireguardPeerReadListAllOf) SetItems(v []WireguardPeerRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *WireguardPeerReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


