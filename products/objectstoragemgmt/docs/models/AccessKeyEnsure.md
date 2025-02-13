# AccessKeyEnsure

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the AccessKey. | |
|**Metadata** | Pointer to **map[string]interface{}** | Metadata | [optional] |
|**Properties** | [**AccessKey**](AccessKey.md) |  | |

## Methods

### NewAccessKeyEnsure

`func NewAccessKeyEnsure(id string, properties AccessKey, ) *AccessKeyEnsure`

NewAccessKeyEnsure instantiates a new AccessKeyEnsure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessKeyEnsureWithDefaults

`func NewAccessKeyEnsureWithDefaults() *AccessKeyEnsure`

NewAccessKeyEnsureWithDefaults instantiates a new AccessKeyEnsure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessKeyEnsure) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessKeyEnsure) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessKeyEnsure) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *AccessKeyEnsure) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccessKeyEnsure) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccessKeyEnsure) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AccessKeyEnsure) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *AccessKeyEnsure) GetProperties() AccessKey`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AccessKeyEnsure) GetPropertiesOk() (*AccessKey, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AccessKeyEnsure) SetProperties(v AccessKey)`

SetProperties sets Properties field to given value.



