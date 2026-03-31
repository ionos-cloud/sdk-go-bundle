# PostgresClusterConnection

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**DatacenterId** | **string** | The datacenter to connect your instance to. | |
|**LanId** | **string** | The numeric LAN ID to connect your instance to. | |
|**PrimaryInstanceAddress** | **string** | Assigns the IP address and netmask to the cluster&#39;s primary instance. Note the following unavailable IP range: 10.208.0.0/12  | |

## Methods

### NewPostgresClusterConnection

`func NewPostgresClusterConnection(datacenterId string, lanId string, primaryInstanceAddress string, ) *PostgresClusterConnection`

NewPostgresClusterConnection instantiates a new PostgresClusterConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresClusterConnectionWithDefaults

`func NewPostgresClusterConnectionWithDefaults() *PostgresClusterConnection`

NewPostgresClusterConnectionWithDefaults instantiates a new PostgresClusterConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenterId

`func (o *PostgresClusterConnection) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *PostgresClusterConnection) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *PostgresClusterConnection) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.


### GetLanId

`func (o *PostgresClusterConnection) GetLanId() string`

GetLanId returns the LanId field if non-nil, zero value otherwise.

### GetLanIdOk

`func (o *PostgresClusterConnection) GetLanIdOk() (*string, bool)`

GetLanIdOk returns a tuple with the LanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanId

`func (o *PostgresClusterConnection) SetLanId(v string)`

SetLanId sets LanId field to given value.


### GetPrimaryInstanceAddress

`func (o *PostgresClusterConnection) GetPrimaryInstanceAddress() string`

GetPrimaryInstanceAddress returns the PrimaryInstanceAddress field if non-nil, zero value otherwise.

### GetPrimaryInstanceAddressOk

`func (o *PostgresClusterConnection) GetPrimaryInstanceAddressOk() (*string, bool)`

GetPrimaryInstanceAddressOk returns a tuple with the PrimaryInstanceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryInstanceAddress

`func (o *PostgresClusterConnection) SetPrimaryInstanceAddress(v string)`

SetPrimaryInstanceAddress sets PrimaryInstanceAddress field to given value.



