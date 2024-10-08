# CertificateRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the Certificate. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the Certificate. | |
|**Metadata** | [**MetadataWithCertificateInformation**](MetadataWithCertificateInformation.md) |  | |
|**Properties** | [**Certificate**](Certificate.md) |  | |

## Methods

### NewCertificateRead

`func NewCertificateRead(id string, type_ string, href string, metadata MetadataWithCertificateInformation, properties Certificate, ) *CertificateRead`

NewCertificateRead instantiates a new CertificateRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateReadWithDefaults

`func NewCertificateReadWithDefaults() *CertificateRead`

NewCertificateReadWithDefaults instantiates a new CertificateRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *CertificateRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificateRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificateRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *CertificateRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CertificateRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CertificateRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *CertificateRead) GetMetadata() MetadataWithCertificateInformation`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CertificateRead) GetMetadataOk() (*MetadataWithCertificateInformation, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CertificateRead) SetMetadata(v MetadataWithCertificateInformation)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *CertificateRead) GetProperties() Certificate`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CertificateRead) GetPropertiesOk() (*Certificate, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CertificateRead) SetProperties(v Certificate)`

SetProperties sets Properties field to given value.



