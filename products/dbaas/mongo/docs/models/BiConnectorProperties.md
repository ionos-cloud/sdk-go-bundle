# BiConnectorProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Enabled** | Pointer to **bool** | The MongoDB Connector for Business Intelligence allows you to query a MongoDB database using SQL commands to aid in data analysis.  | [optional] [default to false]|
|**Host** | Pointer to **string** | The host where this new BI Connector is installed.  | [optional] [readonly] |
|**Port** | Pointer to **string** | Port number used when connecting to this new BI Connector.  | [optional] [readonly] |

## Methods

### NewBiConnectorProperties

`func NewBiConnectorProperties() *BiConnectorProperties`

NewBiConnectorProperties instantiates a new BiConnectorProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBiConnectorPropertiesWithDefaults

`func NewBiConnectorPropertiesWithDefaults() *BiConnectorProperties`

NewBiConnectorPropertiesWithDefaults instantiates a new BiConnectorProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *BiConnectorProperties) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BiConnectorProperties) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BiConnectorProperties) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BiConnectorProperties) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHost

`func (o *BiConnectorProperties) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *BiConnectorProperties) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *BiConnectorProperties) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *BiConnectorProperties) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *BiConnectorProperties) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *BiConnectorProperties) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *BiConnectorProperties) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *BiConnectorProperties) HasPort() bool`

HasPort returns a boolean if a field has been set.


