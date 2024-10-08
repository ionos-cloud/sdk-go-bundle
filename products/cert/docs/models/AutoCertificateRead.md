# AutoCertificateRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the AutoCertificate. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the AutoCertificate. | |
|**Metadata** | [**MetadataWithAutoCertificateInformation**](MetadataWithAutoCertificateInformation.md) |  | |
|**Properties** | [**AutoCertificate**](AutoCertificate.md) |  | |

## Methods

### NewAutoCertificateRead

`func NewAutoCertificateRead(id string, type_ string, href string, metadata MetadataWithAutoCertificateInformation, properties AutoCertificate, ) *AutoCertificateRead`

NewAutoCertificateRead instantiates a new AutoCertificateRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoCertificateReadWithDefaults

`func NewAutoCertificateReadWithDefaults() *AutoCertificateRead`

NewAutoCertificateReadWithDefaults instantiates a new AutoCertificateRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoCertificateRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoCertificateRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoCertificateRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *AutoCertificateRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutoCertificateRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutoCertificateRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *AutoCertificateRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AutoCertificateRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AutoCertificateRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *AutoCertificateRead) GetMetadata() MetadataWithAutoCertificateInformation`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AutoCertificateRead) GetMetadataOk() (*MetadataWithAutoCertificateInformation, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AutoCertificateRead) SetMetadata(v MetadataWithAutoCertificateInformation)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *AutoCertificateRead) GetProperties() AutoCertificate`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AutoCertificateRead) GetPropertiesOk() (*AutoCertificate, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AutoCertificateRead) SetProperties(v AutoCertificate)`

SetProperties sets Properties field to given value.



