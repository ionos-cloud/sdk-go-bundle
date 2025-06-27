# PipelineNoAddrLogs

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Source** | **string** | The source of the logs. Could be one of the following: &#x60;generic&#x60;, &#x60;kubernetes&#x60;, &#x60;docker&#x60;, &#x60;systemd&#x60;.  | |
|**Tag** | **string** | The tag of the logs. Tag represents a short alphanumeric badge (3–20 characters). It must contain only letters and digits, with no special characters or spaces.  | |
|**Protocol** | **string** | The protocol used to send logs. Could be one of the following: &#x60;http&#x60;, &#x60;tcp&#x60;.  | |
|**Destinations** | [**[]PipelineNoAddrLogsDestinations**](PipelineNoAddrLogsDestinations.md) |  | |

## Methods

### NewPipelineNoAddrLogs

`func NewPipelineNoAddrLogs(source string, tag string, protocol string, destinations []PipelineNoAddrLogsDestinations, ) *PipelineNoAddrLogs`

NewPipelineNoAddrLogs instantiates a new PipelineNoAddrLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineNoAddrLogsWithDefaults

`func NewPipelineNoAddrLogsWithDefaults() *PipelineNoAddrLogs`

NewPipelineNoAddrLogsWithDefaults instantiates a new PipelineNoAddrLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *PipelineNoAddrLogs) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PipelineNoAddrLogs) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PipelineNoAddrLogs) SetSource(v string)`

SetSource sets Source field to given value.


### GetTag

`func (o *PipelineNoAddrLogs) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *PipelineNoAddrLogs) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *PipelineNoAddrLogs) SetTag(v string)`

SetTag sets Tag field to given value.


### GetProtocol

`func (o *PipelineNoAddrLogs) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PipelineNoAddrLogs) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PipelineNoAddrLogs) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetDestinations

`func (o *PipelineNoAddrLogs) GetDestinations() []PipelineNoAddrLogsDestinations`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *PipelineNoAddrLogs) GetDestinationsOk() (*[]PipelineNoAddrLogsDestinations, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *PipelineNoAddrLogs) SetDestinations(v []PipelineNoAddrLogsDestinations)`

SetDestinations sets Destinations field to given value.



