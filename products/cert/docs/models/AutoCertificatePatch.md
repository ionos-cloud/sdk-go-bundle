# AutoCertificatePatch

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Metadata** | Pointer to **map[string]interface{}** | Metadata | [optional] |
|**Properties** | [**PatchName**](PatchName.md) |  | |

## Methods

### NewAutoCertificatePatch

`func NewAutoCertificatePatch(properties PatchName, ) *AutoCertificatePatch`

NewAutoCertificatePatch instantiates a new AutoCertificatePatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoCertificatePatchWithDefaults

`func NewAutoCertificatePatchWithDefaults() *AutoCertificatePatch`

NewAutoCertificatePatchWithDefaults instantiates a new AutoCertificatePatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *AutoCertificatePatch) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AutoCertificatePatch) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AutoCertificatePatch) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AutoCertificatePatch) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *AutoCertificatePatch) GetProperties() PatchName`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AutoCertificatePatch) GetPropertiesOk() (*PatchName, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AutoCertificatePatch) SetProperties(v PatchName)`

SetProperties sets Properties field to given value.



