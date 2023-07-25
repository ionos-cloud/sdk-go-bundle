# CreateRequestProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The friendly name of your pipeline. | |
|**Logs** | [**[]CreateRequestPipeline**](CreateRequestPipeline.md) | The information of the log pipelines | |

## Methods

### NewCreateRequestProperties

`func NewCreateRequestProperties(name string, logs []CreateRequestPipeline, ) *CreateRequestProperties`

NewCreateRequestProperties instantiates a new CreateRequestProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRequestPropertiesWithDefaults

`func NewCreateRequestPropertiesWithDefaults() *CreateRequestProperties`

NewCreateRequestPropertiesWithDefaults instantiates a new CreateRequestProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRequestProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRequestProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRequestProperties) SetName(v string)`

SetName sets Name field to given value.


### GetLogs

`func (o *CreateRequestProperties) GetLogs() []CreateRequestPipeline`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *CreateRequestProperties) GetLogsOk() (*[]CreateRequestPipeline, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *CreateRequestProperties) SetLogs(v []CreateRequestPipeline)`

SetLogs sets Logs field to given value.



