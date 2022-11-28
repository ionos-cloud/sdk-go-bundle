# CertificateDto

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The certificate ID. | [optional] [readonly] |
|**Type** | Pointer to **string** | The resource type. | [optional] [readonly] |
|**Href** | Pointer to **string** | The URL to the object representation (absolute path). | [optional] [readonly] |
|**Metadata** | Pointer to [**ResourceMetadataDto**](ResourceMetadataDto.md) |  | [optional] |
|**Properties** | Pointer to [**CertificatePropertiesDto**](CertificatePropertiesDto.md) |  | [optional] |

## Methods

### NewCertificateDto

`func NewCertificateDto() *CertificateDto`

NewCertificateDto instantiates a new CertificateDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateDtoWithDefaults

`func NewCertificateDtoWithDefaults() *CertificateDto`

NewCertificateDtoWithDefaults instantiates a new CertificateDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CertificateDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CertificateDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificateDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificateDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CertificateDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *CertificateDto) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CertificateDto) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CertificateDto) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CertificateDto) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetMetadata

`func (o *CertificateDto) GetMetadata() ResourceMetadataDto`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CertificateDto) GetMetadataOk() (*ResourceMetadataDto, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CertificateDto) SetMetadata(v ResourceMetadataDto)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CertificateDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *CertificateDto) GetProperties() CertificatePropertiesDto`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CertificateDto) GetPropertiesOk() (*CertificatePropertiesDto, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CertificateDto) SetProperties(v CertificatePropertiesDto)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CertificateDto) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


