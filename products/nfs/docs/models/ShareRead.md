# ShareRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The share identifier (UUID) | |
|**Type** | **string** | The resource type. | |
|**Href** | **string** | The URL of the share. | |
|**Metadata** | [**MetadataWithPath**](MetadataWithPath.md) |  | |
|**Properties** | [**Share**](Share.md) |  | |

## Methods

### NewShareRead

`func NewShareRead(id string, type_ string, href string, metadata MetadataWithPath, properties Share, ) *ShareRead`

NewShareRead instantiates a new ShareRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShareReadWithDefaults

`func NewShareReadWithDefaults() *ShareRead`

NewShareReadWithDefaults instantiates a new ShareRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShareRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShareRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShareRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ShareRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ShareRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ShareRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *ShareRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ShareRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ShareRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *ShareRead) GetMetadata() MetadataWithPath`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ShareRead) GetMetadataOk() (*MetadataWithPath, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ShareRead) SetMetadata(v MetadataWithPath)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *ShareRead) GetProperties() Share`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ShareRead) GetPropertiesOk() (*Share, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ShareRead) SetProperties(v Share)`

SetProperties sets Properties field to given value.



