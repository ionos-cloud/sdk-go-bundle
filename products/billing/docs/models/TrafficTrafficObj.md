# TrafficTrafficObj

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Vdc** | Pointer to [**[]TrafficTrafficObjVdc**](TrafficTrafficObjVdc.md) |  | [optional] |
|**Nic** | Pointer to [**[]TrafficTrafficObjNic**](TrafficTrafficObjNic.md) |  | [optional] |

## Methods

### NewTrafficTrafficObj

`func NewTrafficTrafficObj() *TrafficTrafficObj`

NewTrafficTrafficObj instantiates a new TrafficTrafficObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficTrafficObjWithDefaults

`func NewTrafficTrafficObjWithDefaults() *TrafficTrafficObj`

NewTrafficTrafficObjWithDefaults instantiates a new TrafficTrafficObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVdc

`func (o *TrafficTrafficObj) GetVdc() []TrafficTrafficObjVdc`

GetVdc returns the Vdc field if non-nil, zero value otherwise.

### GetVdcOk

`func (o *TrafficTrafficObj) GetVdcOk() (*[]TrafficTrafficObjVdc, bool)`

GetVdcOk returns a tuple with the Vdc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdc

`func (o *TrafficTrafficObj) SetVdc(v []TrafficTrafficObjVdc)`

SetVdc sets Vdc field to given value.

### HasVdc

`func (o *TrafficTrafficObj) HasVdc() bool`

HasVdc returns a boolean if a field has been set.

### GetNic

`func (o *TrafficTrafficObj) GetNic() []TrafficTrafficObjNic`

GetNic returns the Nic field if non-nil, zero value otherwise.

### GetNicOk

`func (o *TrafficTrafficObj) GetNicOk() (*[]TrafficTrafficObjNic, bool)`

GetNicOk returns a tuple with the Nic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNic

`func (o *TrafficTrafficObj) SetNic(v []TrafficTrafficObjNic)`

SetNic sets Nic field to given value.

### HasNic

`func (o *TrafficTrafficObj) HasNic() bool`

HasNic returns a boolean if a field has been set.


