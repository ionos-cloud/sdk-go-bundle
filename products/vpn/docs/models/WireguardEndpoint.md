# WireguardEndpoint

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Host** | **string** | Hostname or IPV4 address that the WireGuard Server will connect to. | |
|**Port** | Pointer to **int32** | IP port number | [optional] [default to 51820]|

## Methods

### NewWireguardEndpoint

`func NewWireguardEndpoint(host string, ) *WireguardEndpoint`

NewWireguardEndpoint instantiates a new WireguardEndpoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireguardEndpointWithDefaults

`func NewWireguardEndpointWithDefaults() *WireguardEndpoint`

NewWireguardEndpointWithDefaults instantiates a new WireguardEndpoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *WireguardEndpoint) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *WireguardEndpoint) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *WireguardEndpoint) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *WireguardEndpoint) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *WireguardEndpoint) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *WireguardEndpoint) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *WireguardEndpoint) HasPort() bool`

HasPort returns a boolean if a field has been set.


