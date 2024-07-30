# IPSecTunnel

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The human readable name of your IPSec Gateway Tunnel. | |
|**Description** | Pointer to **string** | Human readable description of the IPSec Gateway Tunnel. | [optional] |
|**RemoteHost** | **string** | The remote peer host fully qualified domain name or IPV4 IP to connect to. * __Note__: This should be the public IP of the remote peer. * Tunnels only support IPV4 or hostname (fully qualified DNS name).  | |
|**Auth** | [**IPSecTunnelAuth**](IPSecTunnelAuth.md) |  | |
|**Ike** | [**IKEEncryption**](IKEEncryption.md) |  | |
|**Esp** | [**ESPEncryption**](ESPEncryption.md) |  | |
|**CloudNetworkCIDRs** | **[]string** | The network CIDRs on the \&quot;Left\&quot; side that are allowed to connect to the IPSec tunnel, i.e the CIDRs within your IONOS Cloud LAN.  Specify \&quot;0.0.0.0/0\&quot; or \&quot;::/0\&quot; for all addresses.  | |
|**PeerNetworkCIDRs** | **[]string** | The network CIDRs on the \&quot;Right\&quot; side that are allowed to connect to the IPSec tunnel.  Specify \&quot;0.0.0.0/0\&quot; or \&quot;::/0\&quot; for all addresses. | |

## Methods

### NewIPSecTunnel

`func NewIPSecTunnel(name string, remoteHost string, auth IPSecTunnelAuth, ike IKEEncryption, esp ESPEncryption, cloudNetworkCIDRs []string, peerNetworkCIDRs []string, ) *IPSecTunnel`

NewIPSecTunnel instantiates a new IPSecTunnel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecTunnelWithDefaults

`func NewIPSecTunnelWithDefaults() *IPSecTunnel`

NewIPSecTunnelWithDefaults instantiates a new IPSecTunnel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IPSecTunnel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IPSecTunnel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IPSecTunnel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *IPSecTunnel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IPSecTunnel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IPSecTunnel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IPSecTunnel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRemoteHost

`func (o *IPSecTunnel) GetRemoteHost() string`

GetRemoteHost returns the RemoteHost field if non-nil, zero value otherwise.

### GetRemoteHostOk

`func (o *IPSecTunnel) GetRemoteHostOk() (*string, bool)`

GetRemoteHostOk returns a tuple with the RemoteHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteHost

`func (o *IPSecTunnel) SetRemoteHost(v string)`

SetRemoteHost sets RemoteHost field to given value.


### GetAuth

`func (o *IPSecTunnel) GetAuth() IPSecTunnelAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *IPSecTunnel) GetAuthOk() (*IPSecTunnelAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *IPSecTunnel) SetAuth(v IPSecTunnelAuth)`

SetAuth sets Auth field to given value.


### GetIke

`func (o *IPSecTunnel) GetIke() IKEEncryption`

GetIke returns the Ike field if non-nil, zero value otherwise.

### GetIkeOk

`func (o *IPSecTunnel) GetIkeOk() (*IKEEncryption, bool)`

GetIkeOk returns a tuple with the Ike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIke

`func (o *IPSecTunnel) SetIke(v IKEEncryption)`

SetIke sets Ike field to given value.


### GetEsp

`func (o *IPSecTunnel) GetEsp() ESPEncryption`

GetEsp returns the Esp field if non-nil, zero value otherwise.

### GetEspOk

`func (o *IPSecTunnel) GetEspOk() (*ESPEncryption, bool)`

GetEspOk returns a tuple with the Esp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsp

`func (o *IPSecTunnel) SetEsp(v ESPEncryption)`

SetEsp sets Esp field to given value.


### GetCloudNetworkCIDRs

`func (o *IPSecTunnel) GetCloudNetworkCIDRs() []string`

GetCloudNetworkCIDRs returns the CloudNetworkCIDRs field if non-nil, zero value otherwise.

### GetCloudNetworkCIDRsOk

`func (o *IPSecTunnel) GetCloudNetworkCIDRsOk() (*[]string, bool)`

GetCloudNetworkCIDRsOk returns a tuple with the CloudNetworkCIDRs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudNetworkCIDRs

`func (o *IPSecTunnel) SetCloudNetworkCIDRs(v []string)`

SetCloudNetworkCIDRs sets CloudNetworkCIDRs field to given value.


### GetPeerNetworkCIDRs

`func (o *IPSecTunnel) GetPeerNetworkCIDRs() []string`

GetPeerNetworkCIDRs returns the PeerNetworkCIDRs field if non-nil, zero value otherwise.

### GetPeerNetworkCIDRsOk

`func (o *IPSecTunnel) GetPeerNetworkCIDRsOk() (*[]string, bool)`

GetPeerNetworkCIDRsOk returns a tuple with the PeerNetworkCIDRs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerNetworkCIDRs

`func (o *IPSecTunnel) SetPeerNetworkCIDRs(v []string)`

SetPeerNetworkCIDRs sets PeerNetworkCIDRs field to given value.



