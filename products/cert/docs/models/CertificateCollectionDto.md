# CertificateCollectionDto

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The resource ID. | [optional] [readonly] |
|**Type** | Pointer to **string** | The resource type. | [optional] [readonly] |
|**Href** | Pointer to **string** | The URL to the object representation (absolute path). | [optional] [readonly] |
|**Items** | Pointer to [**[]CertificateDto**](CertificateDto.md) | The list of certificates. | [optional] [readonly] |
|**Offset** | Pointer to **int32** | The pagination offset. | [optional] [readonly] |
|**Limit** | Pointer to **int32** | The pagination limit. | [optional] [readonly] |
|**Links** | Pointer to [**CertificateCollectionDtoLinks**](CertificateCollectionDtoLinks.md) |  | [optional] |

## Methods

### NewCertificateCollectionDto

`func NewCertificateCollectionDto() *CertificateCollectionDto`

NewCertificateCollectionDto instantiates a new CertificateCollectionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateCollectionDtoWithDefaults

`func NewCertificateCollectionDtoWithDefaults() *CertificateCollectionDto`

NewCertificateCollectionDtoWithDefaults instantiates a new CertificateCollectionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateCollectionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateCollectionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateCollectionDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CertificateCollectionDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CertificateCollectionDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificateCollectionDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificateCollectionDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CertificateCollectionDto) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *CertificateCollectionDto) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *CertificateCollectionDto) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *CertificateCollectionDto) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *CertificateCollectionDto) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetItems

`func (o *CertificateCollectionDto) GetItems() []CertificateDto`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CertificateCollectionDto) GetItemsOk() (*[]CertificateDto, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CertificateCollectionDto) SetItems(v []CertificateDto)`

SetItems sets Items field to given value.

### HasItems

`func (o *CertificateCollectionDto) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetOffset

`func (o *CertificateCollectionDto) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *CertificateCollectionDto) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *CertificateCollectionDto) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *CertificateCollectionDto) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *CertificateCollectionDto) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *CertificateCollectionDto) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *CertificateCollectionDto) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *CertificateCollectionDto) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLinks

`func (o *CertificateCollectionDto) GetLinks() CertificateCollectionDtoLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CertificateCollectionDto) GetLinksOk() (*CertificateCollectionDtoLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CertificateCollectionDto) SetLinks(v CertificateCollectionDtoLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CertificateCollectionDto) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


