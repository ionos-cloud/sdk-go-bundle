# VersionRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the Version. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the Version. | |
|**Metadata** | **map[string]interface{}** |  | [readonly] |
|**Properties** | [**SupportedVersion**](SupportedVersion.md) |  | |

## Methods

### NewVersionRead

`func NewVersionRead(id string, type_ string, href string, metadata map[string]interface{}, properties SupportedVersion, ) *VersionRead`

NewVersionRead instantiates a new VersionRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionReadWithDefaults

`func NewVersionReadWithDefaults() *VersionRead`

NewVersionReadWithDefaults instantiates a new VersionRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VersionRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VersionRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VersionRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *VersionRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VersionRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VersionRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *VersionRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *VersionRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *VersionRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *VersionRead) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *VersionRead) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *VersionRead) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *VersionRead) GetProperties() SupportedVersion`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *VersionRead) GetPropertiesOk() (*SupportedVersion, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *VersionRead) SetProperties(v SupportedVersion)`

SetProperties sets Properties field to given value.



