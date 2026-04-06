# TrafficTrafficObjVdc

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**VdcUUID** | Pointer to **string** |  | [optional] |
|**VdcName** | Pointer to **string** |  | [optional] |
|**Dates** | Pointer to [**[]TrafficEntry**](TrafficEntry.md) |  | [optional] |

## Methods

### NewTrafficTrafficObjVdc

`func NewTrafficTrafficObjVdc() *TrafficTrafficObjVdc`

NewTrafficTrafficObjVdc instantiates a new TrafficTrafficObjVdc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficTrafficObjVdcWithDefaults

`func NewTrafficTrafficObjVdcWithDefaults() *TrafficTrafficObjVdc`

NewTrafficTrafficObjVdcWithDefaults instantiates a new TrafficTrafficObjVdc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVdcUUID

`func (o *TrafficTrafficObjVdc) GetVdcUUID() string`

GetVdcUUID returns the VdcUUID field if non-nil, zero value otherwise.

### GetVdcUUIDOk

`func (o *TrafficTrafficObjVdc) GetVdcUUIDOk() (*string, bool)`

GetVdcUUIDOk returns a tuple with the VdcUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdcUUID

`func (o *TrafficTrafficObjVdc) SetVdcUUID(v string)`

SetVdcUUID sets VdcUUID field to given value.

### HasVdcUUID

`func (o *TrafficTrafficObjVdc) HasVdcUUID() bool`

HasVdcUUID returns a boolean if a field has been set.

### GetVdcName

`func (o *TrafficTrafficObjVdc) GetVdcName() string`

GetVdcName returns the VdcName field if non-nil, zero value otherwise.

### GetVdcNameOk

`func (o *TrafficTrafficObjVdc) GetVdcNameOk() (*string, bool)`

GetVdcNameOk returns a tuple with the VdcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdcName

`func (o *TrafficTrafficObjVdc) SetVdcName(v string)`

SetVdcName sets VdcName field to given value.

### HasVdcName

`func (o *TrafficTrafficObjVdc) HasVdcName() bool`

HasVdcName returns a boolean if a field has been set.

### GetDates

`func (o *TrafficTrafficObjVdc) GetDates() []TrafficEntry`

GetDates returns the Dates field if non-nil, zero value otherwise.

### GetDatesOk

`func (o *TrafficTrafficObjVdc) GetDatesOk() (*[]TrafficEntry, bool)`

GetDatesOk returns a tuple with the Dates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDates

`func (o *TrafficTrafficObjVdc) SetDates(v []TrafficEntry)`

SetDates sets Dates field to given value.

### HasDates

`func (o *TrafficTrafficObjVdc) HasDates() bool`

HasDates returns a boolean if a field has been set.


