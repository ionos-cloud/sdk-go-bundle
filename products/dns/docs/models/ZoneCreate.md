# ZoneCreate

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Properties** | [**Zone**](Zone.md) |  | |

## Methods

### NewZoneCreate

`func NewZoneCreate(properties Zone, ) *ZoneCreate`

NewZoneCreate instantiates a new ZoneCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneCreateWithDefaults

`func NewZoneCreateWithDefaults() *ZoneCreate`

NewZoneCreateWithDefaults instantiates a new ZoneCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *ZoneCreate) GetProperties() Zone`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ZoneCreate) GetPropertiesOk() (*Zone, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ZoneCreate) SetProperties(v Zone)`

SetProperties sets Properties field to given value.



