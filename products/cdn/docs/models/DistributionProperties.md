# DistributionProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Domain** | **string** | The domain of the distribution. | |
|**CertificateId** | Pointer to **string** | The ID of the certificate to use for the distribution. | [optional] |
|**RoutingRules** | [**[]RoutingRule**](RoutingRule.md) | The routing rules for the distribution. | |

## Methods

### NewDistributionProperties

`func NewDistributionProperties(domain string, routingRules []RoutingRule, ) *DistributionProperties`

NewDistributionProperties instantiates a new DistributionProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributionPropertiesWithDefaults

`func NewDistributionPropertiesWithDefaults() *DistributionProperties`

NewDistributionPropertiesWithDefaults instantiates a new DistributionProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DistributionProperties) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DistributionProperties) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DistributionProperties) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetCertificateId

`func (o *DistributionProperties) GetCertificateId() string`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *DistributionProperties) GetCertificateIdOk() (*string, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *DistributionProperties) SetCertificateId(v string)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *DistributionProperties) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetRoutingRules

`func (o *DistributionProperties) GetRoutingRules() []RoutingRule`

GetRoutingRules returns the RoutingRules field if non-nil, zero value otherwise.

### GetRoutingRulesOk

`func (o *DistributionProperties) GetRoutingRulesOk() (*[]RoutingRule, bool)`

GetRoutingRulesOk returns a tuple with the RoutingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingRules

`func (o *DistributionProperties) SetRoutingRules(v []RoutingRule)`

SetRoutingRules sets RoutingRules field to given value.



