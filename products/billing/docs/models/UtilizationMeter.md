# UtilizationMeter

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **string** |  | [optional] |
|**ResourceId** | Pointer to **NullableString** |  | [optional] |
|**ServerId** | Pointer to **NullableString** |  | [optional] |
|**Name** | Pointer to **string** |  | [optional] |
|**Type** | Pointer to **string** |  | [optional] |
|**From** | Pointer to [**time.Time**](time.Time.md) |  | [optional] |
|**To** | Pointer to [**time.Time**](time.Time.md) |  | [optional] |
|**Exists** | Pointer to **NullableBool** |  | [optional] |
|**MeterId** | Pointer to **string** |  | [optional] |
|**MeterDesc** | Pointer to **string** |  | [optional] |
|**Region** | Pointer to **string** |  | [optional] |
|**Quantity** | Pointer to [**UtilizationMeterQuantity**](UtilizationMeterQuantity.md) |  | [optional] |

## Methods

### NewUtilizationMeter

`func NewUtilizationMeter() *UtilizationMeter`

NewUtilizationMeter instantiates a new UtilizationMeter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilizationMeterWithDefaults

`func NewUtilizationMeterWithDefaults() *UtilizationMeter`

NewUtilizationMeterWithDefaults instantiates a new UtilizationMeter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UtilizationMeter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UtilizationMeter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UtilizationMeter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UtilizationMeter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourceId

`func (o *UtilizationMeter) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *UtilizationMeter) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *UtilizationMeter) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *UtilizationMeter) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### SetResourceIdNil

`func (o *UtilizationMeter) SetResourceIdNil(b bool)`

 SetResourceIdNil sets the value for ResourceId to be an explicit nil

### UnsetResourceId
`func (o *UtilizationMeter) UnsetResourceId()`

UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
### GetServerId

`func (o *UtilizationMeter) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *UtilizationMeter) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *UtilizationMeter) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *UtilizationMeter) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *UtilizationMeter) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *UtilizationMeter) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetName

`func (o *UtilizationMeter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UtilizationMeter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UtilizationMeter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UtilizationMeter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *UtilizationMeter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UtilizationMeter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UtilizationMeter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UtilizationMeter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFrom

`func (o *UtilizationMeter) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *UtilizationMeter) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *UtilizationMeter) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *UtilizationMeter) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *UtilizationMeter) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *UtilizationMeter) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *UtilizationMeter) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *UtilizationMeter) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetExists

`func (o *UtilizationMeter) GetExists() bool`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *UtilizationMeter) GetExistsOk() (*bool, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *UtilizationMeter) SetExists(v bool)`

SetExists sets Exists field to given value.

### HasExists

`func (o *UtilizationMeter) HasExists() bool`

HasExists returns a boolean if a field has been set.

### SetExistsNil

`func (o *UtilizationMeter) SetExistsNil(b bool)`

 SetExistsNil sets the value for Exists to be an explicit nil

### UnsetExists
`func (o *UtilizationMeter) UnsetExists()`

UnsetExists ensures that no value is present for Exists, not even an explicit nil
### GetMeterId

`func (o *UtilizationMeter) GetMeterId() string`

GetMeterId returns the MeterId field if non-nil, zero value otherwise.

### GetMeterIdOk

`func (o *UtilizationMeter) GetMeterIdOk() (*string, bool)`

GetMeterIdOk returns a tuple with the MeterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterId

`func (o *UtilizationMeter) SetMeterId(v string)`

SetMeterId sets MeterId field to given value.

### HasMeterId

`func (o *UtilizationMeter) HasMeterId() bool`

HasMeterId returns a boolean if a field has been set.

### GetMeterDesc

`func (o *UtilizationMeter) GetMeterDesc() string`

GetMeterDesc returns the MeterDesc field if non-nil, zero value otherwise.

### GetMeterDescOk

`func (o *UtilizationMeter) GetMeterDescOk() (*string, bool)`

GetMeterDescOk returns a tuple with the MeterDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeterDesc

`func (o *UtilizationMeter) SetMeterDesc(v string)`

SetMeterDesc sets MeterDesc field to given value.

### HasMeterDesc

`func (o *UtilizationMeter) HasMeterDesc() bool`

HasMeterDesc returns a boolean if a field has been set.

### GetRegion

`func (o *UtilizationMeter) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *UtilizationMeter) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *UtilizationMeter) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *UtilizationMeter) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetQuantity

`func (o *UtilizationMeter) GetQuantity() UtilizationMeterQuantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UtilizationMeter) GetQuantityOk() (*UtilizationMeterQuantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UtilizationMeter) SetQuantity(v UtilizationMeterQuantity)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *UtilizationMeter) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.


