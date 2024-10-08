# MetadataWithCertificateInformationAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**AutoCertificate** | Pointer to **string** | The ID of the auto-certificate that caused issuing of the certificate. | [optional] |
|**LastIssuedCertificate** | Pointer to **string** | The ID of the last issued certificate that belongs to the same auto-certificate. | [optional] |
|**Expired** | **bool** | Indicates if the certificate is expired. | |
|**NotBefore** | [**time.Time**](time.Time.md) | The start date of the certificate. | |
|**NotAfter** | [**time.Time**](time.Time.md) | The end date of the certificate. | |
|**SerialNumber** | **string** | The serial number of the certificate in hex. | |
|**CommonName** | **string** | The common name (DNS) of the certificate. | |
|**SubjectAlternativeNames** | **[]string** | Optional additional names added to the issued certificate. The additional names needs to be part of a zone in IONOS Cloud DNS.  | |

## Methods

### NewMetadataWithCertificateInformationAllOf

`func NewMetadataWithCertificateInformationAllOf(expired bool, notBefore time.Time, notAfter time.Time, serialNumber string, commonName string, subjectAlternativeNames []string, ) *MetadataWithCertificateInformationAllOf`

NewMetadataWithCertificateInformationAllOf instantiates a new MetadataWithCertificateInformationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithCertificateInformationAllOfWithDefaults

`func NewMetadataWithCertificateInformationAllOfWithDefaults() *MetadataWithCertificateInformationAllOf`

NewMetadataWithCertificateInformationAllOfWithDefaults instantiates a new MetadataWithCertificateInformationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoCertificate

`func (o *MetadataWithCertificateInformationAllOf) GetAutoCertificate() string`

GetAutoCertificate returns the AutoCertificate field if non-nil, zero value otherwise.

### GetAutoCertificateOk

`func (o *MetadataWithCertificateInformationAllOf) GetAutoCertificateOk() (*string, bool)`

GetAutoCertificateOk returns a tuple with the AutoCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCertificate

`func (o *MetadataWithCertificateInformationAllOf) SetAutoCertificate(v string)`

SetAutoCertificate sets AutoCertificate field to given value.

### HasAutoCertificate

`func (o *MetadataWithCertificateInformationAllOf) HasAutoCertificate() bool`

HasAutoCertificate returns a boolean if a field has been set.

### GetLastIssuedCertificate

`func (o *MetadataWithCertificateInformationAllOf) GetLastIssuedCertificate() string`

GetLastIssuedCertificate returns the LastIssuedCertificate field if non-nil, zero value otherwise.

### GetLastIssuedCertificateOk

`func (o *MetadataWithCertificateInformationAllOf) GetLastIssuedCertificateOk() (*string, bool)`

GetLastIssuedCertificateOk returns a tuple with the LastIssuedCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIssuedCertificate

`func (o *MetadataWithCertificateInformationAllOf) SetLastIssuedCertificate(v string)`

SetLastIssuedCertificate sets LastIssuedCertificate field to given value.

### HasLastIssuedCertificate

`func (o *MetadataWithCertificateInformationAllOf) HasLastIssuedCertificate() bool`

HasLastIssuedCertificate returns a boolean if a field has been set.

### GetExpired

`func (o *MetadataWithCertificateInformationAllOf) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *MetadataWithCertificateInformationAllOf) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *MetadataWithCertificateInformationAllOf) SetExpired(v bool)`

SetExpired sets Expired field to given value.


### GetNotBefore

`func (o *MetadataWithCertificateInformationAllOf) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *MetadataWithCertificateInformationAllOf) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *MetadataWithCertificateInformationAllOf) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.


### GetNotAfter

`func (o *MetadataWithCertificateInformationAllOf) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *MetadataWithCertificateInformationAllOf) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *MetadataWithCertificateInformationAllOf) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.


### GetSerialNumber

`func (o *MetadataWithCertificateInformationAllOf) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *MetadataWithCertificateInformationAllOf) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *MetadataWithCertificateInformationAllOf) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetCommonName

`func (o *MetadataWithCertificateInformationAllOf) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *MetadataWithCertificateInformationAllOf) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *MetadataWithCertificateInformationAllOf) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### GetSubjectAlternativeNames

`func (o *MetadataWithCertificateInformationAllOf) GetSubjectAlternativeNames() []string`

GetSubjectAlternativeNames returns the SubjectAlternativeNames field if non-nil, zero value otherwise.

### GetSubjectAlternativeNamesOk

`func (o *MetadataWithCertificateInformationAllOf) GetSubjectAlternativeNamesOk() (*[]string, bool)`

GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeNames

`func (o *MetadataWithCertificateInformationAllOf) SetSubjectAlternativeNames(v []string)`

SetSubjectAlternativeNames sets SubjectAlternativeNames field to given value.



