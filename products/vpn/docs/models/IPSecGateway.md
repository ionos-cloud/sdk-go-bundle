# IPSecGateway

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The human readable name of your IPSecGateway. | |
|**Description** | Pointer to **string** | Human readable description of the IPSecGateway. | [optional] |
|**GatewayIP** | **string** | Public IP address to be assigned to the gateway. __Note__: This must be an IP address in the same datacenter as the connections.  | |
|**Connections** | [**[]Connection**](Connection.md) | The network connection for your gateway. __Note__: all connections must belong to the same datacenterId.  | |
|**Version** | Pointer to **string** | The IKE version that is permitted for the VPN tunnels.\\ Options:  - IKEv2  | [optional] [default to "IKEv2"]|

## Methods

### NewIPSecGateway

`func NewIPSecGateway(name string, gatewayIP string, connections []Connection, ) *IPSecGateway`

NewIPSecGateway instantiates a new IPSecGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecGatewayWithDefaults

`func NewIPSecGatewayWithDefaults() *IPSecGateway`

NewIPSecGatewayWithDefaults instantiates a new IPSecGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IPSecGateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IPSecGateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IPSecGateway) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *IPSecGateway) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IPSecGateway) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IPSecGateway) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IPSecGateway) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGatewayIP

`func (o *IPSecGateway) GetGatewayIP() string`

GetGatewayIP returns the GatewayIP field if non-nil, zero value otherwise.

### GetGatewayIPOk

`func (o *IPSecGateway) GetGatewayIPOk() (*string, bool)`

GetGatewayIPOk returns a tuple with the GatewayIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayIP

`func (o *IPSecGateway) SetGatewayIP(v string)`

SetGatewayIP sets GatewayIP field to given value.


### GetConnections

`func (o *IPSecGateway) GetConnections() []Connection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *IPSecGateway) GetConnectionsOk() (*[]Connection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *IPSecGateway) SetConnections(v []Connection)`

SetConnections sets Connections field to given value.


### GetVersion

`func (o *IPSecGateway) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *IPSecGateway) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *IPSecGateway) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *IPSecGateway) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


