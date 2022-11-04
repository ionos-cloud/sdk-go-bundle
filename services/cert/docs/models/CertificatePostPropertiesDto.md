# CertificatePostPropertiesDto

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The certificate name. | |
|**Certificate** | **string** | The certificate body. | |
|**CertificateChain** | **string** | The certificate chain. | |
|**PrivateKey** | **string** | The RSA private key is used for authentication and symmetric key exchange when establishing an SSL session. It is a part of the public key infrastructure generally used with SSL certificates. | |

## Methods

### NewCertificatePostPropertiesDto

`func NewCertificatePostPropertiesDto(name string, certificate string, certificateChain string, privateKey string, ) *CertificatePostPropertiesDto`

NewCertificatePostPropertiesDto instantiates a new CertificatePostPropertiesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatePostPropertiesDtoWithDefaults

`func NewCertificatePostPropertiesDtoWithDefaults() *CertificatePostPropertiesDto`

NewCertificatePostPropertiesDtoWithDefaults instantiates a new CertificatePostPropertiesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CertificatePostPropertiesDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificatePostPropertiesDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificatePostPropertiesDto) SetName(v string)`

SetName sets Name field to given value.


### GetCertificate

`func (o *CertificatePostPropertiesDto) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificatePostPropertiesDto) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificatePostPropertiesDto) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetCertificateChain

`func (o *CertificatePostPropertiesDto) GetCertificateChain() string`

GetCertificateChain returns the CertificateChain field if non-nil, zero value otherwise.

### GetCertificateChainOk

`func (o *CertificatePostPropertiesDto) GetCertificateChainOk() (*string, bool)`

GetCertificateChainOk returns a tuple with the CertificateChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateChain

`func (o *CertificatePostPropertiesDto) SetCertificateChain(v string)`

SetCertificateChain sets CertificateChain field to given value.


### GetPrivateKey

`func (o *CertificatePostPropertiesDto) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CertificatePostPropertiesDto) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CertificatePostPropertiesDto) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.



