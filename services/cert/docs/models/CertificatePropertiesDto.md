# CertificatePropertiesDto

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The certificate name. | [optional] |
|**Certificate** | Pointer to **string** | The certificate body. | [optional] |
|**CertificateChain** | Pointer to **string** | Optional. The certificate chain. | [optional] |

## Methods

### NewCertificatePropertiesDto

`func NewCertificatePropertiesDto() *CertificatePropertiesDto`

NewCertificatePropertiesDto instantiates a new CertificatePropertiesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatePropertiesDtoWithDefaults

`func NewCertificatePropertiesDtoWithDefaults() *CertificatePropertiesDto`

NewCertificatePropertiesDtoWithDefaults instantiates a new CertificatePropertiesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CertificatePropertiesDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificatePropertiesDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificatePropertiesDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CertificatePropertiesDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCertificate

`func (o *CertificatePropertiesDto) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificatePropertiesDto) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificatePropertiesDto) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CertificatePropertiesDto) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetCertificateChain

`func (o *CertificatePropertiesDto) GetCertificateChain() string`

GetCertificateChain returns the CertificateChain field if non-nil, zero value otherwise.

### GetCertificateChainOk

`func (o *CertificatePropertiesDto) GetCertificateChainOk() (*string, bool)`

GetCertificateChainOk returns a tuple with the CertificateChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateChain

`func (o *CertificatePropertiesDto) SetCertificateChain(v string)`

SetCertificateChain sets CertificateChain field to given value.

### HasCertificateChain

`func (o *CertificatePropertiesDto) HasCertificateChain() bool`

HasCertificateChain returns a boolean if a field has been set.


