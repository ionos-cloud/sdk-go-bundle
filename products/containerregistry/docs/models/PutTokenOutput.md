# PutTokenOutput

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Href** | Pointer to **string** |  | [optional] |
|**Id** | Pointer to **string** |  | [optional] |
|**Metadata** | [**ApiResourceMetadata**](ApiResourceMetadata.md) |  | |
|**Properties** | [**TokenProperties**](TokenProperties.md) |  | |
|**Type** | Pointer to **string** |  | [optional] |

## Methods

### NewPutTokenOutput

`func NewPutTokenOutput(metadata ApiResourceMetadata, properties TokenProperties, ) *PutTokenOutput`

NewPutTokenOutput instantiates a new PutTokenOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutTokenOutputWithDefaults

`func NewPutTokenOutputWithDefaults() *PutTokenOutput`

NewPutTokenOutputWithDefaults instantiates a new PutTokenOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *PutTokenOutput) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PutTokenOutput) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PutTokenOutput) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *PutTokenOutput) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetId

`func (o *PutTokenOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutTokenOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutTokenOutput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PutTokenOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *PutTokenOutput) GetMetadata() ApiResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PutTokenOutput) GetMetadataOk() (*ApiResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PutTokenOutput) SetMetadata(v ApiResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *PutTokenOutput) GetProperties() TokenProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PutTokenOutput) GetPropertiesOk() (*TokenProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PutTokenOutput) SetProperties(v TokenProperties)`

SetProperties sets Properties field to given value.


### GetType

`func (o *PutTokenOutput) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PutTokenOutput) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PutTokenOutput) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PutTokenOutput) HasType() bool`

HasType returns a boolean if a field has been set.


