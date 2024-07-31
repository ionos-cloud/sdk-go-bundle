# ProvisioningPipeline

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** | The unique ID of the resource. | [optional] |
|**Metadata** | Pointer to [**ProvisioningMetadata**](ProvisioningMetadata.md) |  | [optional] |
|**Properties** | Pointer to [**PipelineProperties**](PipelineProperties.md) |  | [optional] |

## Methods

### NewProvisioningPipeline

`func NewProvisioningPipeline() *ProvisioningPipeline`

NewProvisioningPipeline instantiates a new ProvisioningPipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningPipelineWithDefaults

`func NewProvisioningPipelineWithDefaults() *ProvisioningPipeline`

NewProvisioningPipelineWithDefaults instantiates a new ProvisioningPipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProvisioningPipeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvisioningPipeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvisioningPipeline) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProvisioningPipeline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *ProvisioningPipeline) GetMetadata() ProvisioningMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProvisioningPipeline) GetMetadataOk() (*ProvisioningMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProvisioningPipeline) SetMetadata(v ProvisioningMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ProvisioningPipeline) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *ProvisioningPipeline) GetProperties() PipelineProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ProvisioningPipeline) GetPropertiesOk() (*PipelineProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ProvisioningPipeline) SetProperties(v PipelineProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ProvisioningPipeline) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


