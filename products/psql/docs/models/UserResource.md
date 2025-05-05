# UserResource

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Type** | [**ResourceType**](ResourceType.md) |  | |
|**Id** | **string** | The unique ID of the resource. | |
|**Href** | **string** | Absolute URL of the resource. | |
|**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] |
|**Properties** | [**UserProperties**](UserProperties.md) |  | |

## Methods

### NewUserResource

`func NewUserResource(type_ ResourceType, id string, href string, properties UserProperties, ) *UserResource`

NewUserResource instantiates a new UserResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResourceWithDefaults

`func NewUserResourceWithDefaults() *UserResource`

NewUserResourceWithDefaults instantiates a new UserResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UserResource) GetType() ResourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserResource) GetTypeOk() (*ResourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserResource) SetType(v ResourceType)`

SetType sets Type field to given value.


### GetId

`func (o *UserResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserResource) SetId(v string)`

SetId sets Id field to given value.


### GetHref

`func (o *UserResource) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *UserResource) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *UserResource) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *UserResource) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UserResource) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UserResource) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UserResource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *UserResource) GetProperties() UserProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *UserResource) GetPropertiesOk() (*UserProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *UserResource) SetProperties(v UserProperties)`

SetProperties sets Properties field to given value.



