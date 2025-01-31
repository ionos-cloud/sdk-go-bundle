# UserAccessMetadataAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**CertificateAuthority** | Pointer to **string** | PEM for the certificate authority. | [optional] |
|**PrivateKey** | Pointer to **string** | PEM for the private key. | [optional] |
|**Certificate** | Pointer to **string** | PEM for the certificate. | [optional] |

## Methods

### NewUserAccessMetadataAllOf

`func NewUserAccessMetadataAllOf() *UserAccessMetadataAllOf`

NewUserAccessMetadataAllOf instantiates a new UserAccessMetadataAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAccessMetadataAllOfWithDefaults

`func NewUserAccessMetadataAllOfWithDefaults() *UserAccessMetadataAllOf`

NewUserAccessMetadataAllOfWithDefaults instantiates a new UserAccessMetadataAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateAuthority

`func (o *UserAccessMetadataAllOf) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *UserAccessMetadataAllOf) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *UserAccessMetadataAllOf) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *UserAccessMetadataAllOf) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### GetPrivateKey

`func (o *UserAccessMetadataAllOf) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *UserAccessMetadataAllOf) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *UserAccessMetadataAllOf) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *UserAccessMetadataAllOf) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetCertificate

`func (o *UserAccessMetadataAllOf) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *UserAccessMetadataAllOf) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *UserAccessMetadataAllOf) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *UserAccessMetadataAllOf) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.


