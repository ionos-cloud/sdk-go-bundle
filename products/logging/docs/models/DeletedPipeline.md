# DeletedPipeline

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The unique ID of the resource. | [optional] |
|**Metadata** | Pointer to [**DeletedMetadata**](DeletedMetadata.md) |  | [optional] |
|**Properties** | Pointer to [**PipelineProperties**](PipelineProperties.md) |  | [optional] |

## Methods

### NewDeletedPipeline

`func NewDeletedPipeline() *DeletedPipeline`

NewDeletedPipeline instantiates a new DeletedPipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletedPipelineWithDefaults

`func NewDeletedPipelineWithDefaults() *DeletedPipeline`

NewDeletedPipelineWithDefaults instantiates a new DeletedPipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeletedPipeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeletedPipeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeletedPipeline) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeletedPipeline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *DeletedPipeline) GetMetadata() DeletedMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeletedPipeline) GetMetadataOk() (*DeletedMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeletedPipeline) SetMetadata(v DeletedMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DeletedPipeline) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *DeletedPipeline) GetProperties() PipelineProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DeletedPipeline) GetPropertiesOk() (*PipelineProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DeletedPipeline) SetProperties(v PipelineProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DeletedPipeline) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


