# StorageClassRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The StorageClass of the StorageClass. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the StorageClass. | |
|**Metadata** | **map[string]interface{}** |  | |
|**Properties** | [**StorageClass**](StorageClass.md) |  | |

## Methods

### NewStorageClassRead

`func NewStorageClassRead(id string, type_ string, href string, metadata map[string]interface{}, properties StorageClass, ) *StorageClassRead`

NewStorageClassRead instantiates a new StorageClassRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageClassReadWithDefaults

`func NewStorageClassReadWithDefaults() *StorageClassRead`

NewStorageClassReadWithDefaults instantiates a new StorageClassRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageClassRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageClassRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageClassRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *StorageClassRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageClassRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageClassRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *StorageClassRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *StorageClassRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *StorageClassRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *StorageClassRead) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StorageClassRead) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StorageClassRead) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *StorageClassRead) GetProperties() StorageClass`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *StorageClassRead) GetPropertiesOk() (*StorageClass, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *StorageClassRead) SetProperties(v StorageClass)`

SetProperties sets Properties field to given value.



