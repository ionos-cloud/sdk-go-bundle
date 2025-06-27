# CentralLogging

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Enabled** | **bool** | The status of the central logging. If &#x60;true&#x60;, the central logging is enabled. If &#x60;false&#x60;, the central logging is disabled.  | [default to false]|

## Methods

### NewCentralLogging

`func NewCentralLogging(enabled bool, ) *CentralLogging`

NewCentralLogging instantiates a new CentralLogging object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCentralLoggingWithDefaults

`func NewCentralLoggingWithDefaults() *CentralLogging`

NewCentralLoggingWithDefaults instantiates a new CentralLogging object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CentralLogging) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CentralLogging) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CentralLogging) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



