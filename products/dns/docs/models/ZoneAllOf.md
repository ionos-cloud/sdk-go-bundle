# ZoneAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Enabled** | Pointer to **bool** | Users can activate and deactivate zones. | [optional] [default to true]|

## Methods

### NewZoneAllOf

`func NewZoneAllOf() *ZoneAllOf`

NewZoneAllOf instantiates a new ZoneAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneAllOfWithDefaults

`func NewZoneAllOfWithDefaults() *ZoneAllOf`

NewZoneAllOfWithDefaults instantiates a new ZoneAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ZoneAllOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ZoneAllOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ZoneAllOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ZoneAllOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


