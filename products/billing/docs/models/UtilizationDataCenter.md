# UtilizationDataCenter

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** |  | [optional] |
|**Name** | Pointer to **string** |  | [optional] |
|**Meters** | Pointer to [**[]UtilizationMeter**](UtilizationMeter.md) |  | [optional] |

## Methods

### NewUtilizationDataCenter

`func NewUtilizationDataCenter() *UtilizationDataCenter`

NewUtilizationDataCenter instantiates a new UtilizationDataCenter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilizationDataCenterWithDefaults

`func NewUtilizationDataCenterWithDefaults() *UtilizationDataCenter`

NewUtilizationDataCenterWithDefaults instantiates a new UtilizationDataCenter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UtilizationDataCenter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UtilizationDataCenter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UtilizationDataCenter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UtilizationDataCenter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UtilizationDataCenter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UtilizationDataCenter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UtilizationDataCenter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UtilizationDataCenter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMeters

`func (o *UtilizationDataCenter) GetMeters() []UtilizationMeter`

GetMeters returns the Meters field if non-nil, zero value otherwise.

### GetMetersOk

`func (o *UtilizationDataCenter) GetMetersOk() (*[]UtilizationMeter, bool)`

GetMetersOk returns a tuple with the Meters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeters

`func (o *UtilizationDataCenter) SetMeters(v []UtilizationMeter)`

SetMeters sets Meters field to given value.

### HasMeters

`func (o *UtilizationDataCenter) HasMeters() bool`

HasMeters returns a boolean if a field has been set.


