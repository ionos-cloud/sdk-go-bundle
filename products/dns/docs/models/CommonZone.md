# CommonZone

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ZoneName** | **string** | The zone name | |
|**Description** | Pointer to **string** | The hosted zone is used for... | [optional] |

## Methods

### NewCommonZone

`func NewCommonZone(zoneName string, ) *CommonZone`

NewCommonZone instantiates a new CommonZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonZoneWithDefaults

`func NewCommonZoneWithDefaults() *CommonZone`

NewCommonZoneWithDefaults instantiates a new CommonZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneName

`func (o *CommonZone) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *CommonZone) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *CommonZone) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.


### GetDescription

`func (o *CommonZone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CommonZone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CommonZone) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CommonZone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


