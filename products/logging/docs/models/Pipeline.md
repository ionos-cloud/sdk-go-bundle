# Pipeline

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The unique ID of the resource. | [optional] |
|**Type** | Pointer to **string** |  | [optional] |
|**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] |
|**Properties** | Pointer to [**PipelineProperties**](PipelineProperties.md) |  | [optional] |

## Methods

### NewPipeline

`func NewPipeline() *Pipeline`

NewPipeline instantiates a new Pipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineWithDefaults

`func NewPipelineWithDefaults() *Pipeline`

NewPipelineWithDefaults instantiates a new Pipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Pipeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Pipeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Pipeline) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Pipeline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Pipeline) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Pipeline) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Pipeline) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Pipeline) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMetadata

`func (o *Pipeline) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Pipeline) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Pipeline) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Pipeline) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *Pipeline) GetProperties() PipelineProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Pipeline) GetPropertiesOk() (*PipelineProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Pipeline) SetProperties(v PipelineProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Pipeline) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


