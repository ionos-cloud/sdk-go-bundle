# MetadataWithCertificateInformation

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**CreatedDate** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 creation timestamp. | [optional] [readonly] |
|**CreatedBy** | Pointer to **string** | Unique name of the identity that created the resource. | [optional] [readonly] |
|**CreatedByUserId** | Pointer to **string** | Unique id of the identity that created the resource. | [optional] [readonly] |
|**LastModifiedDate** | Pointer to [**time.Time**](time.Time.md) | The ISO 8601 modified timestamp. | [optional] [readonly] |
|**LastModifiedBy** | Pointer to **string** | Unique name of the identity that last modified the resource. | [optional] [readonly] |
|**LastModifiedByUserId** | Pointer to **string** | Unique id of the identity that last modified the resource. | [optional] [readonly] |
|**ResourceURN** | Pointer to **string** | Unique name of the resource. | [optional] [readonly] |
|**State** | **string** | The resource state. | |
|**Message** | **string** | A human readable message describing the current state. In case of an error, the message will contain a detailed error message.  | |
|**AutoCertificate** | Pointer to **string** | The ID of the auto-certificate that caused issuing of the certificate. | [optional] |
|**LastIssuedCertificate** | Pointer to **string** | The ID of the last issued certificate that belongs to the same auto-certificate. | [optional] |
|**Expired** | **bool** | Indicates if the certificate is expired. | |
|**NotBefore** | [**time.Time**](time.Time.md) | The start date of the certificate. | |
|**NotAfter** | [**time.Time**](time.Time.md) | The end date of the certificate. | |
|**SerialNumber** | **string** | The serial number of the certificate in hex. | |
|**CommonName** | **string** | The common name (DNS) of the certificate. | |
|**SubjectAlternativeNames** | **[]string** | Optional additional names added to the issued certificate. The additional names needs to be part of a zone in IONOS Cloud DNS.  | |

## Methods

### NewMetadataWithCertificateInformation

`func NewMetadataWithCertificateInformation(state string, message string, expired bool, notBefore time.Time, notAfter time.Time, serialNumber string, commonName string, subjectAlternativeNames []string, ) *MetadataWithCertificateInformation`

NewMetadataWithCertificateInformation instantiates a new MetadataWithCertificateInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithCertificateInformationWithDefaults

`func NewMetadataWithCertificateInformationWithDefaults() *MetadataWithCertificateInformation`

NewMetadataWithCertificateInformationWithDefaults instantiates a new MetadataWithCertificateInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedDate

`func (o *MetadataWithCertificateInformation) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *MetadataWithCertificateInformation) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *MetadataWithCertificateInformation) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *MetadataWithCertificateInformation) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetCreatedBy

`func (o *MetadataWithCertificateInformation) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MetadataWithCertificateInformation) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MetadataWithCertificateInformation) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *MetadataWithCertificateInformation) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *MetadataWithCertificateInformation) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *MetadataWithCertificateInformation) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *MetadataWithCertificateInformation) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *MetadataWithCertificateInformation) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetLastModifiedDate

`func (o *MetadataWithCertificateInformation) GetLastModifiedDate() time.Time`

GetLastModifiedDate returns the LastModifiedDate field if non-nil, zero value otherwise.

### GetLastModifiedDateOk

`func (o *MetadataWithCertificateInformation) GetLastModifiedDateOk() (*time.Time, bool)`

GetLastModifiedDateOk returns a tuple with the LastModifiedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDate

`func (o *MetadataWithCertificateInformation) SetLastModifiedDate(v time.Time)`

SetLastModifiedDate sets LastModifiedDate field to given value.

### HasLastModifiedDate

`func (o *MetadataWithCertificateInformation) HasLastModifiedDate() bool`

HasLastModifiedDate returns a boolean if a field has been set.

### GetLastModifiedBy

`func (o *MetadataWithCertificateInformation) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *MetadataWithCertificateInformation) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *MetadataWithCertificateInformation) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.

### HasLastModifiedBy

`func (o *MetadataWithCertificateInformation) HasLastModifiedBy() bool`

HasLastModifiedBy returns a boolean if a field has been set.

### GetLastModifiedByUserId

`func (o *MetadataWithCertificateInformation) GetLastModifiedByUserId() string`

GetLastModifiedByUserId returns the LastModifiedByUserId field if non-nil, zero value otherwise.

### GetLastModifiedByUserIdOk

`func (o *MetadataWithCertificateInformation) GetLastModifiedByUserIdOk() (*string, bool)`

GetLastModifiedByUserIdOk returns a tuple with the LastModifiedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedByUserId

`func (o *MetadataWithCertificateInformation) SetLastModifiedByUserId(v string)`

SetLastModifiedByUserId sets LastModifiedByUserId field to given value.

### HasLastModifiedByUserId

`func (o *MetadataWithCertificateInformation) HasLastModifiedByUserId() bool`

HasLastModifiedByUserId returns a boolean if a field has been set.

### GetResourceURN

`func (o *MetadataWithCertificateInformation) GetResourceURN() string`

GetResourceURN returns the ResourceURN field if non-nil, zero value otherwise.

### GetResourceURNOk

`func (o *MetadataWithCertificateInformation) GetResourceURNOk() (*string, bool)`

GetResourceURNOk returns a tuple with the ResourceURN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceURN

`func (o *MetadataWithCertificateInformation) SetResourceURN(v string)`

SetResourceURN sets ResourceURN field to given value.

### HasResourceURN

`func (o *MetadataWithCertificateInformation) HasResourceURN() bool`

HasResourceURN returns a boolean if a field has been set.

### GetState

`func (o *MetadataWithCertificateInformation) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MetadataWithCertificateInformation) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MetadataWithCertificateInformation) SetState(v string)`

SetState sets State field to given value.


### GetMessage

`func (o *MetadataWithCertificateInformation) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MetadataWithCertificateInformation) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MetadataWithCertificateInformation) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAutoCertificate

`func (o *MetadataWithCertificateInformation) GetAutoCertificate() string`

GetAutoCertificate returns the AutoCertificate field if non-nil, zero value otherwise.

### GetAutoCertificateOk

`func (o *MetadataWithCertificateInformation) GetAutoCertificateOk() (*string, bool)`

GetAutoCertificateOk returns a tuple with the AutoCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCertificate

`func (o *MetadataWithCertificateInformation) SetAutoCertificate(v string)`

SetAutoCertificate sets AutoCertificate field to given value.

### HasAutoCertificate

`func (o *MetadataWithCertificateInformation) HasAutoCertificate() bool`

HasAutoCertificate returns a boolean if a field has been set.

### GetLastIssuedCertificate

`func (o *MetadataWithCertificateInformation) GetLastIssuedCertificate() string`

GetLastIssuedCertificate returns the LastIssuedCertificate field if non-nil, zero value otherwise.

### GetLastIssuedCertificateOk

`func (o *MetadataWithCertificateInformation) GetLastIssuedCertificateOk() (*string, bool)`

GetLastIssuedCertificateOk returns a tuple with the LastIssuedCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIssuedCertificate

`func (o *MetadataWithCertificateInformation) SetLastIssuedCertificate(v string)`

SetLastIssuedCertificate sets LastIssuedCertificate field to given value.

### HasLastIssuedCertificate

`func (o *MetadataWithCertificateInformation) HasLastIssuedCertificate() bool`

HasLastIssuedCertificate returns a boolean if a field has been set.

### GetExpired

`func (o *MetadataWithCertificateInformation) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *MetadataWithCertificateInformation) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *MetadataWithCertificateInformation) SetExpired(v bool)`

SetExpired sets Expired field to given value.


### GetNotBefore

`func (o *MetadataWithCertificateInformation) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *MetadataWithCertificateInformation) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *MetadataWithCertificateInformation) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.


### GetNotAfter

`func (o *MetadataWithCertificateInformation) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *MetadataWithCertificateInformation) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *MetadataWithCertificateInformation) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.


### GetSerialNumber

`func (o *MetadataWithCertificateInformation) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *MetadataWithCertificateInformation) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *MetadataWithCertificateInformation) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetCommonName

`func (o *MetadataWithCertificateInformation) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *MetadataWithCertificateInformation) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *MetadataWithCertificateInformation) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### GetSubjectAlternativeNames

`func (o *MetadataWithCertificateInformation) GetSubjectAlternativeNames() []string`

GetSubjectAlternativeNames returns the SubjectAlternativeNames field if non-nil, zero value otherwise.

### GetSubjectAlternativeNamesOk

`func (o *MetadataWithCertificateInformation) GetSubjectAlternativeNamesOk() (*[]string, bool)`

GetSubjectAlternativeNamesOk returns a tuple with the SubjectAlternativeNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAlternativeNames

`func (o *MetadataWithCertificateInformation) SetSubjectAlternativeNames(v []string)`

SetSubjectAlternativeNames sets SubjectAlternativeNames field to given value.



