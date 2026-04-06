# EvnDatacenters

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**VdcUUID** | Pointer to **string** |  | [optional] |
|**Name** | Pointer to **string** |  | [optional] |
|**Data** | Pointer to [**[]EvnItem**](EvnItem.md) |  | [optional] |

## Methods

### NewEvnDatacenters

`func NewEvnDatacenters() *EvnDatacenters`

NewEvnDatacenters instantiates a new EvnDatacenters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvnDatacentersWithDefaults

`func NewEvnDatacentersWithDefaults() *EvnDatacenters`

NewEvnDatacentersWithDefaults instantiates a new EvnDatacenters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVdcUUID

`func (o *EvnDatacenters) GetVdcUUID() string`

GetVdcUUID returns the VdcUUID field if non-nil, zero value otherwise.

### GetVdcUUIDOk

`func (o *EvnDatacenters) GetVdcUUIDOk() (*string, bool)`

GetVdcUUIDOk returns a tuple with the VdcUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdcUUID

`func (o *EvnDatacenters) SetVdcUUID(v string)`

SetVdcUUID sets VdcUUID field to given value.

### HasVdcUUID

`func (o *EvnDatacenters) HasVdcUUID() bool`

HasVdcUUID returns a boolean if a field has been set.

### GetName

`func (o *EvnDatacenters) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EvnDatacenters) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EvnDatacenters) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EvnDatacenters) HasName() bool`

HasName returns a boolean if a field has been set.

### GetData

`func (o *EvnDatacenters) GetData() []EvnItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EvnDatacenters) GetDataOk() (*[]EvnItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EvnDatacenters) SetData(v []EvnItem)`

SetData sets Data field to given value.

### HasData

`func (o *EvnDatacenters) HasData() bool`

HasData returns a boolean if a field has been set.


