# PostgresVersionList

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Data** | Pointer to [**[]PostgresVersionListData**](PostgresVersionListData.md) |  | [optional] |

## Methods

### NewPostgresVersionList

`func NewPostgresVersionList() *PostgresVersionList`

NewPostgresVersionList instantiates a new PostgresVersionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresVersionListWithDefaults

`func NewPostgresVersionListWithDefaults() *PostgresVersionList`

NewPostgresVersionListWithDefaults instantiates a new PostgresVersionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PostgresVersionList) GetData() []PostgresVersionListData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PostgresVersionList) GetDataOk() (*[]PostgresVersionListData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PostgresVersionList) SetData(v []PostgresVersionListData)`

SetData sets Data field to given value.

### HasData

`func (o *PostgresVersionList) HasData() bool`

HasData returns a boolean if a field has been set.


