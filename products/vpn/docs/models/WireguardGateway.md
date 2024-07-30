# WireguardGateway

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The human readable name of your WireguardGateway. | |
|**Description** | Pointer to **string** | Human readable description of the WireguardGateway. | [optional] |
|**GatewayIP** | **string** | Public IP address to be assigned to the gateway. __Note__: This must be an IP address in the same datacenter as the connections.  | |
|**InterfaceIPv4CIDR** | Pointer to **string** | Describes a range of IP V4 addresses in CIDR notation.  | [optional] |
|**InterfaceIPv6CIDR** | Pointer to **string** | Describes a range of IP V6 addresses in CIDR notation.  | [optional] |
|**Connections** | [**[]Connection**](Connection.md) | The network connection for your gateway. __Note__: all connections must belong to the same datacenterId.  | |
|**PrivateKey** | **string** | PrivateKey used for WireGuard Server  | |
|**ListenPort** | Pointer to **int32** | IP port number | [optional] [default to 51820]|

## Methods

### NewWireguardGateway

`func NewWireguardGateway(name string, gatewayIP string, connections []Connection, privateKey string, ) *WireguardGateway`

NewWireguardGateway instantiates a new WireguardGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireguardGatewayWithDefaults

`func NewWireguardGatewayWithDefaults() *WireguardGateway`

NewWireguardGatewayWithDefaults instantiates a new WireguardGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WireguardGateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WireguardGateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WireguardGateway) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *WireguardGateway) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WireguardGateway) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WireguardGateway) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WireguardGateway) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGatewayIP

`func (o *WireguardGateway) GetGatewayIP() string`

GetGatewayIP returns the GatewayIP field if non-nil, zero value otherwise.

### GetGatewayIPOk

`func (o *WireguardGateway) GetGatewayIPOk() (*string, bool)`

GetGatewayIPOk returns a tuple with the GatewayIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIP

`func (o *WireguardGateway) SetGatewayIP(v string)`

SetGatewayIP sets GatewayIP field to given value.


### GetInterfaceIPv4CIDR

`func (o *WireguardGateway) GetInterfaceIPv4CIDR() string`

GetInterfaceIPv4CIDR returns the InterfaceIPv4CIDR field if non-nil, zero value otherwise.

### GetInterfaceIPv4CIDROk

`func (o *WireguardGateway) GetInterfaceIPv4CIDROk() (*string, bool)`

GetInterfaceIPv4CIDROk returns a tuple with the InterfaceIPv4CIDR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIPv4CIDR

`func (o *WireguardGateway) SetInterfaceIPv4CIDR(v string)`

SetInterfaceIPv4CIDR sets InterfaceIPv4CIDR field to given value.

### HasInterfaceIPv4CIDR

`func (o *WireguardGateway) HasInterfaceIPv4CIDR() bool`

HasInterfaceIPv4CIDR returns a boolean if a field has been set.

### GetInterfaceIPv6CIDR

`func (o *WireguardGateway) GetInterfaceIPv6CIDR() string`

GetInterfaceIPv6CIDR returns the InterfaceIPv6CIDR field if non-nil, zero value otherwise.

### GetInterfaceIPv6CIDROk

`func (o *WireguardGateway) GetInterfaceIPv6CIDROk() (*string, bool)`

GetInterfaceIPv6CIDROk returns a tuple with the InterfaceIPv6CIDR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIPv6CIDR

`func (o *WireguardGateway) SetInterfaceIPv6CIDR(v string)`

SetInterfaceIPv6CIDR sets InterfaceIPv6CIDR field to given value.

### HasInterfaceIPv6CIDR

`func (o *WireguardGateway) HasInterfaceIPv6CIDR() bool`

HasInterfaceIPv6CIDR returns a boolean if a field has been set.

### GetConnections

`func (o *WireguardGateway) GetConnections() []Connection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *WireguardGateway) GetConnectionsOk() (*[]Connection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *WireguardGateway) SetConnections(v []Connection)`

SetConnections sets Connections field to given value.


### GetPrivateKey

`func (o *WireguardGateway) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *WireguardGateway) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *WireguardGateway) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetListenPort

`func (o *WireguardGateway) GetListenPort() int32`

GetListenPort returns the ListenPort field if non-nil, zero value otherwise.

### GetListenPortOk

`func (o *WireguardGateway) GetListenPortOk() (*int32, bool)`

GetListenPortOk returns a tuple with the ListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPort

`func (o *WireguardGateway) SetListenPort(v int32)`

SetListenPort sets ListenPort field to given value.

### HasListenPort

`func (o *WireguardGateway) HasListenPort() bool`

HasListenPort returns a boolean if a field has been set.


