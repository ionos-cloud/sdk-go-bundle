# PipelineRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the Pipeline. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the Pipeline. | |
|**Metadata** | [**MetadataWithEndpoint**](MetadataWithEndpoint.md) |  | |
|**Properties** | [**Pipeline**](Pipeline.md) |  | |

## Methods

### NewPipelineRead

`func NewPipelineRead(id string, type_ string, href string, metadata MetadataWithEndpoint, properties Pipeline, ) *PipelineRead`

NewPipelineRead instantiates a new PipelineRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineReadWithDefaults

`func NewPipelineReadWithDefaults() *PipelineRead`

NewPipelineReadWithDefaults instantiates a new PipelineRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PipelineRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PipelineRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PipelineRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PipelineRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PipelineRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PipelineRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *PipelineRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PipelineRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PipelineRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *PipelineRead) GetMetadata() MetadataWithEndpoint`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PipelineRead) GetMetadataOk() (*MetadataWithEndpoint, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PipelineRead) SetMetadata(v MetadataWithEndpoint)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *PipelineRead) GetProperties() Pipeline`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PipelineRead) GetPropertiesOk() (*Pipeline, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PipelineRead) SetProperties(v Pipeline)`

SetProperties sets Properties field to given value.



