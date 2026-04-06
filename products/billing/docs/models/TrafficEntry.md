# TrafficEntry

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Date** | Pointer to **string** |  | [optional] |
|**In** | Pointer to **float32** |  | [optional] |
|**Out** | Pointer to **float32** |  | [optional] |

## Methods

### NewTrafficEntry

`func NewTrafficEntry() *TrafficEntry`

NewTrafficEntry instantiates a new TrafficEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficEntryWithDefaults

`func NewTrafficEntryWithDefaults() *TrafficEntry`

NewTrafficEntryWithDefaults instantiates a new TrafficEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *TrafficEntry) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TrafficEntry) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TrafficEntry) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *TrafficEntry) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetIn

`func (o *TrafficEntry) GetIn() float32`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *TrafficEntry) GetInOk() (*float32, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *TrafficEntry) SetIn(v float32)`

SetIn sets In field to given value.

### HasIn

`func (o *TrafficEntry) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetOut

`func (o *TrafficEntry) GetOut() float32`

GetOut returns the Out field if non-nil, zero value otherwise.

### GetOutOk

`func (o *TrafficEntry) GetOutOk() (*float32, bool)`

GetOutOk returns a tuple with the Out field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOut

`func (o *TrafficEntry) SetOut(v float32)`

SetOut sets Out field to given value.

### HasOut

`func (o *TrafficEntry) HasOut() bool`

HasOut returns a boolean if a field has been set.


