# WireguardPeer

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The human readable name of your WireguardGateway Peer. | |
|**Description** | Pointer to **string** | Human readable description of the WireguardGateway Peer. | [optional] |
|**Endpoint** | Pointer to [**WireguardEndpoint**](WireguardEndpoint.md) |  | [optional] |
|**AllowedIPs** | **[]string** | The subnet CIDRs that are allowed to connect to the WireGuard Gateway.  Specify \&quot;a.b.c.d/32\&quot; for an individual IP address.  Specify \&quot;0.0.0.0/0\&quot; or \&quot;::/0\&quot; for all addresses.  | |
|**PublicKey** | **string** | WireGuard public key of the connecting peer | |

## Methods

### NewWireguardPeer

`func NewWireguardPeer(name string, allowedIPs []string, publicKey string, ) *WireguardPeer`

NewWireguardPeer instantiates a new WireguardPeer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireguardPeerWithDefaults

`func NewWireguardPeerWithDefaults() *WireguardPeer`

NewWireguardPeerWithDefaults instantiates a new WireguardPeer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WireguardPeer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WireguardPeer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WireguardPeer) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WireguardPeer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WireguardPeer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WireguardPeer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WireguardPeer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEndpoint

`func (o *WireguardPeer) GetEndpoint() WireguardEndpoint`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *WireguardPeer) GetEndpointOk() (*WireguardEndpoint, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *WireguardPeer) SetEndpoint(v WireguardEndpoint)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *WireguardPeer) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetAllowedIPs

`func (o *WireguardPeer) GetAllowedIPs() []string`

GetAllowedIPs returns the AllowedIPs field if non-nil, zero value otherwise.

### GetAllowedIPsOk

`func (o *WireguardPeer) GetAllowedIPsOk() (*[]string, bool)`

GetAllowedIPsOk returns a tuple with the AllowedIPs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedIPs

`func (o *WireguardPeer) SetAllowedIPs(v []string)`

SetAllowedIPs sets AllowedIPs field to given value.


### GetPublicKey

`func (o *WireguardPeer) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *WireguardPeer) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *WireguardPeer) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.



