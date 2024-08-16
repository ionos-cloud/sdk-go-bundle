# IpAddresses

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**PublicEndpointIpv4** | Pointer to **string** | IP of the distribution. It has to be included on the domain DNS Zone as A record. | [optional] [readonly] |
|**PublicEndpointIpv6** | Pointer to **string** | IP of the distribution, it has to be included on the domain DNS Zone as AAAA record. | [optional] [readonly] |

## Methods

### NewIpAddresses

`func NewIpAddresses() *IpAddresses`

NewIpAddresses instantiates a new IpAddresses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpAddressesWithDefaults

`func NewIpAddressesWithDefaults() *IpAddresses`

NewIpAddressesWithDefaults instantiates a new IpAddresses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicEndpointIpv4

`func (o *IpAddresses) GetPublicEndpointIpv4() string`

GetPublicEndpointIpv4 returns the PublicEndpointIpv4 field if non-nil, zero value otherwise.

### GetPublicEndpointIpv4Ok

`func (o *IpAddresses) GetPublicEndpointIpv4Ok() (*string, bool)`

GetPublicEndpointIpv4Ok returns a tuple with the PublicEndpointIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicEndpointIpv4

`func (o *IpAddresses) SetPublicEndpointIpv4(v string)`

SetPublicEndpointIpv4 sets PublicEndpointIpv4 field to given value.

### HasPublicEndpointIpv4

`func (o *IpAddresses) HasPublicEndpointIpv4() bool`

HasPublicEndpointIpv4 returns a boolean if a field has been set.

### GetPublicEndpointIpv6

`func (o *IpAddresses) GetPublicEndpointIpv6() string`

GetPublicEndpointIpv6 returns the PublicEndpointIpv6 field if non-nil, zero value otherwise.

### GetPublicEndpointIpv6Ok

`func (o *IpAddresses) GetPublicEndpointIpv6Ok() (*string, bool)`

GetPublicEndpointIpv6Ok returns a tuple with the PublicEndpointIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicEndpointIpv6

`func (o *IpAddresses) SetPublicEndpointIpv6(v string)`

SetPublicEndpointIpv6 sets PublicEndpointIpv6 field to given value.

### HasPublicEndpointIpv6

`func (o *IpAddresses) HasPublicEndpointIpv6() bool`

HasPublicEndpointIpv6 returns a boolean if a field has been set.


