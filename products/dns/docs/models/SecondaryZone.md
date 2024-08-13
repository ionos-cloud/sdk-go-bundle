# SecondaryZone

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ZoneName** | **string** | The zone name | |
|**Description** | Pointer to **string** | The hosted zone is used for... | [optional] |
|**PrimaryIps** | **[]string** | Indicates IP addresses of primary nameservers for a secondary zone. Accepts IPv4 and IPv6 addresses | |

## Methods

### NewSecondaryZone

`func NewSecondaryZone(zoneName string, primaryIps []string, ) *SecondaryZone`

NewSecondaryZone instantiates a new SecondaryZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecondaryZoneWithDefaults

`func NewSecondaryZoneWithDefaults() *SecondaryZone`

NewSecondaryZoneWithDefaults instantiates a new SecondaryZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneName

`func (o *SecondaryZone) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *SecondaryZone) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *SecondaryZone) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.


### GetDescription

`func (o *SecondaryZone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecondaryZone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecondaryZone) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecondaryZone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrimaryIps

`func (o *SecondaryZone) GetPrimaryIps() []string`

GetPrimaryIps returns the PrimaryIps field if non-nil, zero value otherwise.

### GetPrimaryIpsOk

`func (o *SecondaryZone) GetPrimaryIpsOk() (*[]string, bool)`

GetPrimaryIpsOk returns a tuple with the PrimaryIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIps

`func (o *SecondaryZone) SetPrimaryIps(v []string)`

SetPrimaryIps sets PrimaryIps field to given value.



