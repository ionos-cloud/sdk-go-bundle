# DatacenterProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The name of the  resource. | [optional] |
|**Description** | Pointer to **string** | A description for the datacenter, such as staging, production. | [optional] |
|**Location** | **string** | The physical location where the datacenter will be created. This will be where all of your servers live. Property cannot be modified after datacenter creation (disallowed in update requests). | |
|**Version** | Pointer to **int32** | The version of the data center; incremented with every change. | [optional] [readonly] |
|**Features** | Pointer to **[]string** | List of features supported by the location where this data center is provisioned. | [optional] [readonly] |
|**SecAuthProtection** | Pointer to **bool** | Boolean value representing if the data center requires extra protection, such as two-step verification. | [optional] |
|**CpuArchitecture** | Pointer to [**[]CpuArchitectureProperties**](CpuArchitectureProperties.md) | Array of features and CPU families available in a location | [optional] [readonly] |
|**Ipv6CidrBlock** | Pointer to **NullableString** | This value is either &#39;null&#39; or contains an automatically-assigned /56 IPv6 CIDR block if IPv6 is enabled on this virtual data center. It can neither be changed nor removed. | [optional] [readonly] |
|**DefaultSecurityGroupId** | Pointer to **string** | Optional property to define the default security group of the datacenter. | [optional] |

## Methods

### NewDatacenterProperties

`func NewDatacenterProperties(location string, ) *DatacenterProperties`

NewDatacenterProperties instantiates a new DatacenterProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatacenterPropertiesWithDefaults

`func NewDatacenterPropertiesWithDefaults() *DatacenterProperties`

NewDatacenterPropertiesWithDefaults instantiates a new DatacenterProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatacenterProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatacenterProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatacenterProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DatacenterProperties) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *DatacenterProperties) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatacenterProperties) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatacenterProperties) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatacenterProperties) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocation

`func (o *DatacenterProperties) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DatacenterProperties) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DatacenterProperties) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetVersion

`func (o *DatacenterProperties) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DatacenterProperties) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DatacenterProperties) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DatacenterProperties) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetFeatures

`func (o *DatacenterProperties) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *DatacenterProperties) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *DatacenterProperties) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *DatacenterProperties) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetSecAuthProtection

`func (o *DatacenterProperties) GetSecAuthProtection() bool`

GetSecAuthProtection returns the SecAuthProtection field if non-nil, zero value otherwise.

### GetSecAuthProtectionOk

`func (o *DatacenterProperties) GetSecAuthProtectionOk() (*bool, bool)`

GetSecAuthProtectionOk returns a tuple with the SecAuthProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecAuthProtection

`func (o *DatacenterProperties) SetSecAuthProtection(v bool)`

SetSecAuthProtection sets SecAuthProtection field to given value.

### HasSecAuthProtection

`func (o *DatacenterProperties) HasSecAuthProtection() bool`

HasSecAuthProtection returns a boolean if a field has been set.

### GetCpuArchitecture

`func (o *DatacenterProperties) GetCpuArchitecture() []CpuArchitectureProperties`

GetCpuArchitecture returns the CpuArchitecture field if non-nil, zero value otherwise.

### GetCpuArchitectureOk

`func (o *DatacenterProperties) GetCpuArchitectureOk() (*[]CpuArchitectureProperties, bool)`

GetCpuArchitectureOk returns a tuple with the CpuArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuArchitecture

`func (o *DatacenterProperties) SetCpuArchitecture(v []CpuArchitectureProperties)`

SetCpuArchitecture sets CpuArchitecture field to given value.

### HasCpuArchitecture

`func (o *DatacenterProperties) HasCpuArchitecture() bool`

HasCpuArchitecture returns a boolean if a field has been set.

### GetIpv6CidrBlock

`func (o *DatacenterProperties) GetIpv6CidrBlock() string`

GetIpv6CidrBlock returns the Ipv6CidrBlock field if non-nil, zero value otherwise.

### GetIpv6CidrBlockOk

`func (o *DatacenterProperties) GetIpv6CidrBlockOk() (*string, bool)`

GetIpv6CidrBlockOk returns a tuple with the Ipv6CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6CidrBlock

`func (o *DatacenterProperties) SetIpv6CidrBlock(v string)`

SetIpv6CidrBlock sets Ipv6CidrBlock field to given value.

### HasIpv6CidrBlock

`func (o *DatacenterProperties) HasIpv6CidrBlock() bool`

HasIpv6CidrBlock returns a boolean if a field has been set.

### SetIpv6CidrBlockNil

`func (o *DatacenterProperties) SetIpv6CidrBlockNil(b bool)`

 SetIpv6CidrBlockNil sets the value for Ipv6CidrBlock to be an explicit nil

### UnsetIpv6CidrBlock
`func (o *DatacenterProperties) UnsetIpv6CidrBlock()`

UnsetIpv6CidrBlock ensures that no value is present for Ipv6CidrBlock, not even an explicit nil
### GetDefaultSecurityGroupId

`func (o *DatacenterProperties) GetDefaultSecurityGroupId() string`

GetDefaultSecurityGroupId returns the DefaultSecurityGroupId field if non-nil, zero value otherwise.

### GetDefaultSecurityGroupIdOk

`func (o *DatacenterProperties) GetDefaultSecurityGroupIdOk() (*string, bool)`

GetDefaultSecurityGroupIdOk returns a tuple with the DefaultSecurityGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSecurityGroupId

`func (o *DatacenterProperties) SetDefaultSecurityGroupId(v string)`

SetDefaultSecurityGroupId sets DefaultSecurityGroupId field to given value.

### HasDefaultSecurityGroupId

`func (o *DatacenterProperties) HasDefaultSecurityGroupId() bool`

HasDefaultSecurityGroupId returns a boolean if a field has been set.


