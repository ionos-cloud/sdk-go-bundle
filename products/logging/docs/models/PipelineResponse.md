# PipelineResponse

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Public** | Pointer to **bool** |  | [optional] |
|**Source** | Pointer to **string** | The source parser to be used | [optional] |
|**Tag** | Pointer to **string** | Tag is to distinguish different pipelines. must be unique amongst the pipeline&#39;s array items. | [optional] |
|**Protocol** | Pointer to **string** | Protocol to use as intake | [optional] |
|**Labels** | Pointer to **[]string** | Optional custom labels to filter and report logs | [optional] |
|**Destinations** | Pointer to [**[]Destination**](Destination.md) |  | [optional] |

## Methods

### NewPipelineResponse

`func NewPipelineResponse() *PipelineResponse`

NewPipelineResponse instantiates a new PipelineResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineResponseWithDefaults

`func NewPipelineResponseWithDefaults() *PipelineResponse`

NewPipelineResponseWithDefaults instantiates a new PipelineResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublic

`func (o *PipelineResponse) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *PipelineResponse) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *PipelineResponse) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *PipelineResponse) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetSource

`func (o *PipelineResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PipelineResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PipelineResponse) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PipelineResponse) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTag

`func (o *PipelineResponse) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PipelineResponse) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PipelineResponse) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *PipelineResponse) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetProtocol

`func (o *PipelineResponse) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PipelineResponse) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PipelineResponse) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PipelineResponse) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetLabels

`func (o *PipelineResponse) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PipelineResponse) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PipelineResponse) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PipelineResponse) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetDestinations

`func (o *PipelineResponse) GetDestinations() []Destination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *PipelineResponse) GetDestinationsOk() (*[]Destination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *PipelineResponse) SetDestinations(v []Destination)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *PipelineResponse) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.


