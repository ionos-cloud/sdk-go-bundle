# PipelineNoAddrLogsDestinations

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Type** | **string** | The type of the destination. Could be one of the following: &#x60;loki&#x60;.  | |
|**RetentionInDays** | **int32** | The retention period of the logs in days. Could be one of the following: 0, 7, 14, 30.  | |

## Methods

### NewPipelineNoAddrLogsDestinations

`func NewPipelineNoAddrLogsDestinations(type_ string, retentionInDays int32, ) *PipelineNoAddrLogsDestinations`

NewPipelineNoAddrLogsDestinations instantiates a new PipelineNoAddrLogsDestinations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineNoAddrLogsDestinationsWithDefaults

`func NewPipelineNoAddrLogsDestinationsWithDefaults() *PipelineNoAddrLogsDestinations`

NewPipelineNoAddrLogsDestinationsWithDefaults instantiates a new PipelineNoAddrLogsDestinations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PipelineNoAddrLogsDestinations) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PipelineNoAddrLogsDestinations) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PipelineNoAddrLogsDestinations) SetType(v string)`

SetType sets Type field to given value.


### GetRetentionInDays

`func (o *PipelineNoAddrLogsDestinations) GetRetentionInDays() int32`

GetRetentionInDays returns the RetentionInDays field if non-nil, zero value otherwise.

### GetRetentionInDaysOk

`func (o *PipelineNoAddrLogsDestinations) GetRetentionInDaysOk() (*int32, bool)`

GetRetentionInDaysOk returns a tuple with the RetentionInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionInDays

`func (o *PipelineNoAddrLogsDestinations) SetRetentionInDays(v int32)`

SetRetentionInDays sets RetentionInDays field to given value.



