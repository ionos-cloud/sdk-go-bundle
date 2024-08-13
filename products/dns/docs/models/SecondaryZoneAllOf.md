# SecondaryZoneAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**PrimaryIps** | **[]string** | Indicates IP addresses of primary nameservers for a secondary zone. Accepts IPv4 and IPv6 addresses | |

## Methods

### NewSecondaryZoneAllOf

`func NewSecondaryZoneAllOf(primaryIps []string, ) *SecondaryZoneAllOf`

NewSecondaryZoneAllOf instantiates a new SecondaryZoneAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecondaryZoneAllOfWithDefaults

`func NewSecondaryZoneAllOfWithDefaults() *SecondaryZoneAllOf`

NewSecondaryZoneAllOfWithDefaults instantiates a new SecondaryZoneAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryIps

`func (o *SecondaryZoneAllOf) GetPrimaryIps() []string`

GetPrimaryIps returns the PrimaryIps field if non-nil, zero value otherwise.

### GetPrimaryIpsOk

`func (o *SecondaryZoneAllOf) GetPrimaryIpsOk() (*[]string, bool)`

GetPrimaryIpsOk returns a tuple with the PrimaryIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIps

`func (o *SecondaryZoneAllOf) SetPrimaryIps(v []string)`

SetPrimaryIps sets PrimaryIps field to given value.



