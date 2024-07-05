# ConnectionPooler

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Enabled** | Pointer to **bool** |  | [optional] |
|**PoolMode** | Pointer to [**PoolMode**](PoolMode.md) |  | [optional] |

## Methods

### NewConnectionPooler

`func NewConnectionPooler() *ConnectionPooler`

NewConnectionPooler instantiates a new ConnectionPooler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionPoolerWithDefaults

`func NewConnectionPoolerWithDefaults() *ConnectionPooler`

NewConnectionPoolerWithDefaults instantiates a new ConnectionPooler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ConnectionPooler) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ConnectionPooler) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ConnectionPooler) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ConnectionPooler) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPoolMode

`func (o *ConnectionPooler) GetPoolMode() PoolMode`

GetPoolMode returns the PoolMode field if non-nil, zero value otherwise.

### GetPoolModeOk

`func (o *ConnectionPooler) GetPoolModeOk() (*PoolMode, bool)`

GetPoolModeOk returns a tuple with the PoolMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMode

`func (o *ConnectionPooler) SetPoolMode(v PoolMode)`

SetPoolMode sets PoolMode field to given value.

### HasPoolMode

`func (o *ConnectionPooler) HasPoolMode() bool`

HasPoolMode returns a boolean if a field has been set.


