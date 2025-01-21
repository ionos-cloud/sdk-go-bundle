# AutoCertificate

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Provider** | **string** | The certificate provider used to issue the certificates. | |
|**CommonName** | **string** | The common name (DNS) of the certificate to issue. The common name needs to be part of a zone in IONOS Cloud DNS.  | |
|**KeyAlgorithm** | **string** | The key algorithm used to generate the certificate. | |
|**Name** | **string** | A certificate name used for management purposes. | |
|**SubjectAlternativeNames** | Pointer to **[]string** | Optional additional names to be added to the issued certificate. The additional names needs to be part of a zone in IONOS Cloud DNS.  | [optional] |

## Methods

### NewAutoCertificate

`func NewAutoCertificate(provider string, commonName string, keyAlgorithm string, name string, ) *AutoCertificate`

NewAutoCertificate instantiates a new AutoCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoCertificateWithDefaults

`func NewAutoCertificateWithDefaults() *AutoCertificate`

NewAutoCertificateWithDefaults instantiates a new AutoCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *AutoCertificate) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AutoCertificate) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AutoCertificate) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetCommonName

`func (o *AutoCertificate) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *AutoCertificate) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *AutoCertificate) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### GetKeyAlgorithm

`func (o *AutoCertificate) GetKeyAlgorithm() string`

GetKeyAlgorithm returns the KeyAlgorithm field if non-nil, zero value otherwise.

### GetKeyAlgorithmOk

`func (o *AutoCertificate) GetKeyAlgorithmOk() (*string, bool)`

GetKeyAlgorithmOk returns a tuple with the KeyAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyAlgorithm

`func (o *AutoCertificate) SetKeyAlgorithm(v string)`

SetKeyAlgorithm sets KeyAlgorithm field to given value.


### GetName

`func (o *AutoCertificate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutoCertificate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutoCertificate) SetName(v string)`

SetName sets Name field to given value.


### GetSubjectAlternativeNames

`func (o *AutoCertificate) GetSubjectAlternativeNames() []string`

GetSubjectAlternativeNames returns the SubjectAlternativeNames field if non-nil, zero value otherwise.

### GetSubjectAlternativeNamesOk

`func (o *AutoCertificate) GetSubjectAlternativeNamesOk() (*[]string, bool)`

GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeNames

`func (o *AutoCertificate) SetSubjectAlternativeNames(v []string)`

SetSubjectAlternativeNames sets SubjectAlternativeNames field to given value.

### HasSubjectAlternativeNames

`func (o *AutoCertificate) HasSubjectAlternativeNames() bool`

HasSubjectAlternativeNames returns a boolean if a field has been set.


