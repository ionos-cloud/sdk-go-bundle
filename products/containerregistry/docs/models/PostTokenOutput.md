# PostTokenOutput

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Href** | Pointer to **string** |  | [optional] |
|**Id** | Pointer to **string** |  | [optional] |
|**Metadata** | [**NullableApiResourceMetadata**](ApiResourceMetadata.md) |  | |
|**Properties** | [**NullableTokenProperties**](TokenProperties.md) |  | |
|**Type** | Pointer to **string** |  | [optional] |

## Methods

### NewPostTokenOutput

`func NewPostTokenOutput(metadata NullableApiResourceMetadata, properties NullableTokenProperties, ) *PostTokenOutput`

NewPostTokenOutput instantiates a new PostTokenOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTokenOutputWithDefaults

`func NewPostTokenOutputWithDefaults() *PostTokenOutput`

NewPostTokenOutputWithDefaults instantiates a new PostTokenOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *PostTokenOutput) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PostTokenOutput) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PostTokenOutput) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PostTokenOutput) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *PostTokenOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostTokenOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostTokenOutput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PostTokenOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *PostTokenOutput) GetMetadata() ApiResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PostTokenOutput) GetMetadataOk() (*ApiResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PostTokenOutput) SetMetadata(v ApiResourceMetadata)`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *PostTokenOutput) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *PostTokenOutput) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetProperties

`func (o *PostTokenOutput) GetProperties() TokenProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PostTokenOutput) GetPropertiesOk() (*TokenProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PostTokenOutput) SetProperties(v TokenProperties)`

SetProperties sets Properties field to given value.


### SetPropertiesNil

`func (o *PostTokenOutput) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *PostTokenOutput) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetType

`func (o *PostTokenOutput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostTokenOutput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostTokenOutput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostTokenOutput) HasType() bool`

HasType returns a boolean if a field has been set.


