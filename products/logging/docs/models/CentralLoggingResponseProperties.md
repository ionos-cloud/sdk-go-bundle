# CentralLoggingResponseProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Enabled** | Pointer to **bool** |  | [optional] [default to false]|

## Methods

### NewCentralLoggingResponseProperties

`func NewCentralLoggingResponseProperties() *CentralLoggingResponseProperties`

NewCentralLoggingResponseProperties instantiates a new CentralLoggingResponseProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCentralLoggingResponsePropertiesWithDefaults

`func NewCentralLoggingResponsePropertiesWithDefaults() *CentralLoggingResponseProperties`

NewCentralLoggingResponsePropertiesWithDefaults instantiates a new CentralLoggingResponseProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *CentralLoggingResponseProperties) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CentralLoggingResponseProperties) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CentralLoggingResponseProperties) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CentralLoggingResponseProperties) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


