# ResponsePipeline

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

### NewResponsePipeline

`func NewResponsePipeline() *ResponsePipeline`

NewResponsePipeline instantiates a new ResponsePipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponsePipelineWithDefaults

`func NewResponsePipelineWithDefaults() *ResponsePipeline`

NewResponsePipelineWithDefaults instantiates a new ResponsePipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublic

`func (o *ResponsePipeline) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *ResponsePipeline) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *ResponsePipeline) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *ResponsePipeline) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetSource

`func (o *ResponsePipeline) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ResponsePipeline) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ResponsePipeline) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ResponsePipeline) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTag

`func (o *ResponsePipeline) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ResponsePipeline) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ResponsePipeline) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ResponsePipeline) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetProtocol

`func (o *ResponsePipeline) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *ResponsePipeline) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *ResponsePipeline) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *ResponsePipeline) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetLabels

`func (o *ResponsePipeline) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ResponsePipeline) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ResponsePipeline) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ResponsePipeline) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetDestinations

`func (o *ResponsePipeline) GetDestinations() []Destination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *ResponsePipeline) GetDestinationsOk() (*[]Destination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *ResponsePipeline) SetDestinations(v []Destination)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *ResponsePipeline) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.


