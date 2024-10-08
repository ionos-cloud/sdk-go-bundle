# PrivateCrossConnectProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The name of the Cross Connect. | [optional] |
|**Description** | Pointer to **string** | Human-readable description of the Cross Connect. | [optional] |
|**Peers** | Pointer to [**[]Peer**](Peer.md) | Read-Only attribute. Lists LAN&#39;s connected to this Cross Connect. | [optional] [readonly] |
|**ConnectableDatacenters** | Pointer to [**[]ConnectableDatacenter**](ConnectableDatacenter.md) | Read-Only attribute. Lists data centers that can be connected to this Cross Connect. If the Cross Connect is not connected to any LANs it lists all VDCs the user has access to. If the Cross Connect is connected to at least 1 LAN it lists all VDCs the user has access to in the location of the connected LAN. | [optional] [readonly] |

## Methods

### NewPrivateCrossConnectProperties

`func NewPrivateCrossConnectProperties() *PrivateCrossConnectProperties`

NewPrivateCrossConnectProperties instantiates a new PrivateCrossConnectProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateCrossConnectPropertiesWithDefaults

`func NewPrivateCrossConnectPropertiesWithDefaults() *PrivateCrossConnectProperties`

NewPrivateCrossConnectPropertiesWithDefaults instantiates a new PrivateCrossConnectProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PrivateCrossConnectProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateCrossConnectProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateCrossConnectProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrivateCrossConnectProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PrivateCrossConnectProperties) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrivateCrossConnectProperties) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrivateCrossConnectProperties) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrivateCrossConnectProperties) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPeers

`func (o *PrivateCrossConnectProperties) GetPeers() []Peer`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *PrivateCrossConnectProperties) GetPeersOk() (*[]Peer, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *PrivateCrossConnectProperties) SetPeers(v []Peer)`

SetPeers sets Peers field to given value.

### HasPeers

`func (o *PrivateCrossConnectProperties) HasPeers() bool`

HasPeers returns a boolean if a field has been set.

### GetConnectableDatacenters

`func (o *PrivateCrossConnectProperties) GetConnectableDatacenters() []ConnectableDatacenter`

GetConnectableDatacenters returns the ConnectableDatacenters field if non-nil, zero value otherwise.

### GetConnectableDatacentersOk

`func (o *PrivateCrossConnectProperties) GetConnectableDatacentersOk() (*[]ConnectableDatacenter, bool)`

GetConnectableDatacentersOk returns a tuple with the ConnectableDatacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectableDatacenters

`func (o *PrivateCrossConnectProperties) SetConnectableDatacenters(v []ConnectableDatacenter)`

SetConnectableDatacenters sets ConnectableDatacenters field to given value.

### HasConnectableDatacenters

`func (o *PrivateCrossConnectProperties) HasConnectableDatacenters() bool`

HasConnectableDatacenters returns a boolean if a field has been set.


