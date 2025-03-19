# AccessKeyRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the AccessKey. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the AccessKey. | |
|**Metadata** | [**MetadataWithSupportedRegions**](MetadataWithSupportedRegions.md) |  | |
|**Properties** | [**AccessKey**](AccessKey.md) |  | |

## Methods

### NewAccessKeyRead

`func NewAccessKeyRead(id string, type_ string, href string, metadata MetadataWithSupportedRegions, properties AccessKey, ) *AccessKeyRead`

NewAccessKeyRead instantiates a new AccessKeyRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessKeyReadWithDefaults

`func NewAccessKeyReadWithDefaults() *AccessKeyRead`

NewAccessKeyReadWithDefaults instantiates a new AccessKeyRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessKeyRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessKeyRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessKeyRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *AccessKeyRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccessKeyRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccessKeyRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *AccessKeyRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *AccessKeyRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *AccessKeyRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *AccessKeyRead) GetMetadata() MetadataWithSupportedRegions`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccessKeyRead) GetMetadataOk() (*MetadataWithSupportedRegions, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccessKeyRead) SetMetadata(v MetadataWithSupportedRegions)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *AccessKeyRead) GetProperties() AccessKey`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AccessKeyRead) GetPropertiesOk() (*AccessKey, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AccessKeyRead) SetProperties(v AccessKey)`

SetProperties sets Properties field to given value.



