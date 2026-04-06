# TrafficTrafficObjNic

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**VdcUUID** | Pointer to **string** |  | [optional] |
|**VdcName** | Pointer to **string** |  | [optional] |
|**Ip** | Pointer to **string** |  | [optional] |
|**Dates** | Pointer to [**[]TrafficEntry**](TrafficEntry.md) |  | [optional] |

## Methods

### NewTrafficTrafficObjNic

`func NewTrafficTrafficObjNic() *TrafficTrafficObjNic`

NewTrafficTrafficObjNic instantiates a new TrafficTrafficObjNic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrafficTrafficObjNicWithDefaults

`func NewTrafficTrafficObjNicWithDefaults() *TrafficTrafficObjNic`

NewTrafficTrafficObjNicWithDefaults instantiates a new TrafficTrafficObjNic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVdcUUID

`func (o *TrafficTrafficObjNic) GetVdcUUID() string`

GetVdcUUID returns the VdcUUID field if non-nil, zero value otherwise.

### GetVdcUUIDOk

`func (o *TrafficTrafficObjNic) GetVdcUUIDOk() (*string, bool)`

GetVdcUUIDOk returns a tuple with the VdcUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdcUUID

`func (o *TrafficTrafficObjNic) SetVdcUUID(v string)`

SetVdcUUID sets VdcUUID field to given value.

### HasVdcUUID

`func (o *TrafficTrafficObjNic) HasVdcUUID() bool`

HasVdcUUID returns a boolean if a field has been set.

### GetVdcName

`func (o *TrafficTrafficObjNic) GetVdcName() string`

GetVdcName returns the VdcName field if non-nil, zero value otherwise.

### GetVdcNameOk

`func (o *TrafficTrafficObjNic) GetVdcNameOk() (*string, bool)`

GetVdcNameOk returns a tuple with the VdcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdcName

`func (o *TrafficTrafficObjNic) SetVdcName(v string)`

SetVdcName sets VdcName field to given value.

### HasVdcName

`func (o *TrafficTrafficObjNic) HasVdcName() bool`

HasVdcName returns a boolean if a field has been set.

### GetIp

`func (o *TrafficTrafficObjNic) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *TrafficTrafficObjNic) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *TrafficTrafficObjNic) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *TrafficTrafficObjNic) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetDates

`func (o *TrafficTrafficObjNic) GetDates() []TrafficEntry`

GetDates returns the Dates field if non-nil, zero value otherwise.

### GetDatesOk

`func (o *TrafficTrafficObjNic) GetDatesOk() (*[]TrafficEntry, bool)`

GetDatesOk returns a tuple with the Dates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDates

`func (o *TrafficTrafficObjNic) SetDates(v []TrafficEntry)`

SetDates sets Dates field to given value.

### HasDates

`func (o *TrafficTrafficObjNic) HasDates() bool`

HasDates returns a boolean if a field has been set.


