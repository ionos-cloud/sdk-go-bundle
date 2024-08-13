# GatewayCustomDomains

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The domain name of the distribution. | [optional] |
|**CertificateId** | Pointer to **string** | The ID of the certificate to use for the distribution. | [optional] |

## Methods

### NewGatewayCustomDomains

`func NewGatewayCustomDomains() *GatewayCustomDomains`

NewGatewayCustomDomains instantiates a new GatewayCustomDomains object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCustomDomainsWithDefaults

`func NewGatewayCustomDomainsWithDefaults() *GatewayCustomDomains`

NewGatewayCustomDomainsWithDefaults instantiates a new GatewayCustomDomains object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GatewayCustomDomains) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCustomDomains) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCustomDomains) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayCustomDomains) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCertificateId

`func (o *GatewayCustomDomains) GetCertificateId() string`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *GatewayCustomDomains) GetCertificateIdOk() (*string, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *GatewayCustomDomains) SetCertificateId(v string)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *GatewayCustomDomains) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.


