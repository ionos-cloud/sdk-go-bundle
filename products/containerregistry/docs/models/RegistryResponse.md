# RegistryResponse

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Href** | Pointer to **string** |  | [optional] |
|**Id** | Pointer to **string** |  | [optional] |
|**Metadata** | [**ApiResourceMetadata**](ApiResourceMetadata.md) |  | |
|**Properties** | [**RegistryProperties**](RegistryProperties.md) |  | |
|**Type** | Pointer to **string** |  | [optional] |

## Methods

### NewRegistryResponse

`func NewRegistryResponse(metadata ApiResourceMetadata, properties RegistryProperties, ) *RegistryResponse`

NewRegistryResponse instantiates a new RegistryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistryResponseWithDefaults

`func NewRegistryResponseWithDefaults() *RegistryResponse`

NewRegistryResponseWithDefaults instantiates a new RegistryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *RegistryResponse) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *RegistryResponse) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *RegistryResponse) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *RegistryResponse) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *RegistryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegistryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegistryResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RegistryResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *RegistryResponse) GetMetadata() ApiResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RegistryResponse) GetMetadataOk() (*ApiResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RegistryResponse) SetMetadata(v ApiResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *RegistryResponse) GetProperties() RegistryProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *RegistryResponse) GetPropertiesOk() (*RegistryProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *RegistryResponse) SetProperties(v RegistryProperties)`

SetProperties sets Properties field to given value.


### GetType

`func (o *RegistryResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RegistryResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RegistryResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RegistryResponse) HasType() bool`

HasType returns a boolean if a field has been set.


