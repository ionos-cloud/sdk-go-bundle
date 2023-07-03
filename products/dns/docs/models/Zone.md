# Zone

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ZoneName** | **string** | The zone name | |
|**Description** | Pointer to **string** | The hosted zone is used for... | [optional] |
|**Enabled** | Pointer to **bool** | Users can activate and deactivate zones. | [optional] [default to true]|

## Methods

### NewZone

`func NewZone(zoneName string, ) *Zone`

NewZone instantiates a new Zone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneWithDefaults

`func NewZoneWithDefaults() *Zone`

NewZoneWithDefaults instantiates a new Zone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneName

`func (o *Zone) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *Zone) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *Zone) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.


### GetDescription

`func (o *Zone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Zone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Zone) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Zone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *Zone) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Zone) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Zone) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Zone) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


