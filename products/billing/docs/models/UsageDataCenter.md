# UsageDataCenter

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** |  | [optional] |
|**Name** | Pointer to **string** |  | [optional] |
|**Location** | Pointer to **string** |  | [optional] |
|**Meters** | Pointer to [**[]UsageMeter**](UsageMeter.md) |  | [optional] |

## Methods

### NewUsageDataCenter

`func NewUsageDataCenter() *UsageDataCenter`

NewUsageDataCenter instantiates a new UsageDataCenter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageDataCenterWithDefaults

`func NewUsageDataCenterWithDefaults() *UsageDataCenter`

NewUsageDataCenterWithDefaults instantiates a new UsageDataCenter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UsageDataCenter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsageDataCenter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsageDataCenter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UsageDataCenter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UsageDataCenter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UsageDataCenter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UsageDataCenter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UsageDataCenter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLocation

`func (o *UsageDataCenter) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UsageDataCenter) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UsageDataCenter) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UsageDataCenter) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMeters

`func (o *UsageDataCenter) GetMeters() []UsageMeter`

GetMeters returns the Meters field if non-nil, zero value otherwise.

### GetMetersOk

`func (o *UsageDataCenter) GetMetersOk() (*[]UsageMeter, bool)`

GetMetersOk returns a tuple with the Meters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeters

`func (o *UsageDataCenter) SetMeters(v []UsageMeter)`

SetMeters sets Meters field to given value.

### HasMeters

`func (o *UsageDataCenter) HasMeters() bool`

HasMeters returns a boolean if a field has been set.


