# IPSecTunnelAuth

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Method** | **string** | The Authentication Method to use for IPSec Authentication.\\ Options:   - PSK  | [default to "PSK"]|
|**Psk** | Pointer to [**IPSecPSK**](IPSecPSK.md) |  | [optional] |

## Methods

### NewIPSecTunnelAuth

`func NewIPSecTunnelAuth(method string, ) *IPSecTunnelAuth`

NewIPSecTunnelAuth instantiates a new IPSecTunnelAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPSecTunnelAuthWithDefaults

`func NewIPSecTunnelAuthWithDefaults() *IPSecTunnelAuth`

NewIPSecTunnelAuthWithDefaults instantiates a new IPSecTunnelAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *IPSecTunnelAuth) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *IPSecTunnelAuth) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *IPSecTunnelAuth) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetPsk

`func (o *IPSecTunnelAuth) GetPsk() IPSecPSK`

GetPsk returns the Psk field if non-nil, zero value otherwise.

### GetPskOk

`func (o *IPSecTunnelAuth) GetPskOk() (*IPSecPSK, bool)`

GetPskOk returns a tuple with the Psk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsk

`func (o *IPSecTunnelAuth) SetPsk(v IPSecPSK)`

SetPsk sets Psk field to given value.

### HasPsk

`func (o *IPSecTunnelAuth) HasPsk() bool`

HasPsk returns a boolean if a field has been set.


