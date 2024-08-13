# ZoneTransferPrimaryIpStatus

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**PrimaryIp** | **string** | one single IP from the primaryIps field for secondary zones | |
|**Status** | **string** | Human readable status of the zone transfer status for the IP | |
|**ErrorMessage** | Pointer to **string** | Human readable explanation of the error when status is not ok | [optional] |

## Methods

### NewZoneTransferPrimaryIpStatus

`func NewZoneTransferPrimaryIpStatus(primaryIp string, status string, ) *ZoneTransferPrimaryIpStatus`

NewZoneTransferPrimaryIpStatus instantiates a new ZoneTransferPrimaryIpStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneTransferPrimaryIpStatusWithDefaults

`func NewZoneTransferPrimaryIpStatusWithDefaults() *ZoneTransferPrimaryIpStatus`

NewZoneTransferPrimaryIpStatusWithDefaults instantiates a new ZoneTransferPrimaryIpStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryIp

`func (o *ZoneTransferPrimaryIpStatus) GetPrimaryIp() string`

GetPrimaryIp returns the PrimaryIp field if non-nil, zero value otherwise.

### GetPrimaryIpOk

`func (o *ZoneTransferPrimaryIpStatus) GetPrimaryIpOk() (*string, bool)`

GetPrimaryIpOk returns a tuple with the PrimaryIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryIp

`func (o *ZoneTransferPrimaryIpStatus) SetPrimaryIp(v string)`

SetPrimaryIp sets PrimaryIp field to given value.


### GetStatus

`func (o *ZoneTransferPrimaryIpStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ZoneTransferPrimaryIpStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ZoneTransferPrimaryIpStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetErrorMessage

`func (o *ZoneTransferPrimaryIpStatus) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ZoneTransferPrimaryIpStatus) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ZoneTransferPrimaryIpStatus) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ZoneTransferPrimaryIpStatus) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


