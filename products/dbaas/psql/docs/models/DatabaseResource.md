# DatabaseResource

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Type** | [**ResourceType**](ResourceType.md) |  | |
|**Id** | **string** | The unique ID of the resource. | |
|**Href** | **string** | Absolute URL of the resource. | |
|**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] |
|**Properties** | [**DatabaseProperties**](DatabaseProperties.md) |  | |

## Methods

### NewDatabaseResource

`func NewDatabaseResource(type_ ResourceType, id string, href string, properties DatabaseProperties, ) *DatabaseResource`

NewDatabaseResource instantiates a new DatabaseResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabaseResourceWithDefaults

`func NewDatabaseResourceWithDefaults() *DatabaseResource`

NewDatabaseResourceWithDefaults instantiates a new DatabaseResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DatabaseResource) GetType() ResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatabaseResource) GetTypeOk() (*ResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatabaseResource) SetType(v ResourceType)`

SetType sets Type field to given value.


### GetId

`func (o *DatabaseResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatabaseResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatabaseResource) SetId(v string)`

SetId sets Id field to given value.


### GetHref

`func (o *DatabaseResource) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DatabaseResource) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DatabaseResource) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *DatabaseResource) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DatabaseResource) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DatabaseResource) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DatabaseResource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *DatabaseResource) GetProperties() DatabaseProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DatabaseResource) GetPropertiesOk() (*DatabaseProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DatabaseResource) SetProperties(v DatabaseProperties)`

SetProperties sets Properties field to given value.



