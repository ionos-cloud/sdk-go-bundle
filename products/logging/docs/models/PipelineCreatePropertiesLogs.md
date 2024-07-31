# PipelineCreatePropertiesLogs

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Source** | Pointer to **string** | The source parser to be used | [optional] |
|**Tag** | Pointer to **string** | Tag is to distinguish different pipelines. must be unique amongst the pipeline&#39;s array items. | [optional] |
|**Protocol** | Pointer to **string** | Protocol to use as intake | [optional] |
|**Labels** | Pointer to **[]string** | Optional custom labels to filter and report logs | [optional] |
|**Destinations** | Pointer to [**[]Destination**](Destination.md) | The configuration of the logs datastore | [optional] |

## Methods

### NewPipelineCreatePropertiesLogs

`func NewPipelineCreatePropertiesLogs() *PipelineCreatePropertiesLogs`

NewPipelineCreatePropertiesLogs instantiates a new PipelineCreatePropertiesLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineCreatePropertiesLogsWithDefaults

`func NewPipelineCreatePropertiesLogsWithDefaults() *PipelineCreatePropertiesLogs`

NewPipelineCreatePropertiesLogsWithDefaults instantiates a new PipelineCreatePropertiesLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *PipelineCreatePropertiesLogs) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PipelineCreatePropertiesLogs) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PipelineCreatePropertiesLogs) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PipelineCreatePropertiesLogs) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTag

`func (o *PipelineCreatePropertiesLogs) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PipelineCreatePropertiesLogs) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PipelineCreatePropertiesLogs) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *PipelineCreatePropertiesLogs) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetProtocol

`func (o *PipelineCreatePropertiesLogs) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PipelineCreatePropertiesLogs) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PipelineCreatePropertiesLogs) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PipelineCreatePropertiesLogs) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetLabels

`func (o *PipelineCreatePropertiesLogs) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PipelineCreatePropertiesLogs) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PipelineCreatePropertiesLogs) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *PipelineCreatePropertiesLogs) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetDestinations

`func (o *PipelineCreatePropertiesLogs) GetDestinations() []Destination`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *PipelineCreatePropertiesLogs) GetDestinationsOk() (*[]Destination, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *PipelineCreatePropertiesLogs) SetDestinations(v []Destination)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *PipelineCreatePropertiesLogs) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.


