# Connection

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**DatacenterId** | **string** | The datacenter to connect your VPN Gateway to.  | |
|**LanId** | **string** | The numeric LAN ID to connect your VPN Gateway to. | |
|**Ipv4CIDR** | **string** | Describes a range of IP V4 addresses in CIDR notation.  | |
|**Ipv6CIDR** | Pointer to **string** | Describes a range of IP V6 addresses in CIDR notation.  | [optional] |

## Methods

### NewConnection

`func NewConnection(datacenterId string, lanId string, ipv4CIDR string, ) *Connection`

NewConnection instantiates a new Connection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionWithDefaults

`func NewConnectionWithDefaults() *Connection`

NewConnectionWithDefaults instantiates a new Connection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenterId

`func (o *Connection) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *Connection) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *Connection) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.


### GetLanId

`func (o *Connection) GetLanId() string`

GetLanId returns the LanId field if non-nil, zero value otherwise.

### GetLanIdOk

`func (o *Connection) GetLanIdOk() (*string, bool)`

GetLanIdOk returns a tuple with the LanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanId

`func (o *Connection) SetLanId(v string)`

SetLanId sets LanId field to given value.


### GetIpv4CIDR

`func (o *Connection) GetIpv4CIDR() string`

GetIpv4CIDR returns the Ipv4CIDR field if non-nil, zero value otherwise.

### GetIpv4CIDROk

`func (o *Connection) GetIpv4CIDROk() (*string, bool)`

GetIpv4CIDROk returns a tuple with the Ipv4CIDR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4CIDR

`func (o *Connection) SetIpv4CIDR(v string)`

SetIpv4CIDR sets Ipv4CIDR field to given value.


### GetIpv6CIDR

`func (o *Connection) GetIpv6CIDR() string`

GetIpv6CIDR returns the Ipv6CIDR field if non-nil, zero value otherwise.

### GetIpv6CIDROk

`func (o *Connection) GetIpv6CIDROk() (*string, bool)`

GetIpv6CIDROk returns a tuple with the Ipv6CIDR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6CIDR

`func (o *Connection) SetIpv6CIDR(v string)`

SetIpv6CIDR sets Ipv6CIDR field to given value.

### HasIpv6CIDR

`func (o *Connection) HasIpv6CIDR() bool`

HasIpv6CIDR returns a boolean if a field has been set.


