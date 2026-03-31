# PostgresVersionRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the PostgresVersion. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the PostgresVersion. | |
|**Metadata** | **map[string]interface{}** |  | [readonly] |
|**Properties** | [**PostgresVersion**](PostgresVersion.md) |  | |

## Methods

### NewPostgresVersionRead

`func NewPostgresVersionRead(id string, type_ string, href string, metadata map[string]interface{}, properties PostgresVersion, ) *PostgresVersionRead`

NewPostgresVersionRead instantiates a new PostgresVersionRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresVersionReadWithDefaults

`func NewPostgresVersionReadWithDefaults() *PostgresVersionRead`

NewPostgresVersionReadWithDefaults instantiates a new PostgresVersionRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostgresVersionRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostgresVersionRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostgresVersionRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PostgresVersionRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostgresVersionRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostgresVersionRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *PostgresVersionRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PostgresVersionRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PostgresVersionRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *PostgresVersionRead) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PostgresVersionRead) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PostgresVersionRead) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *PostgresVersionRead) GetProperties() PostgresVersion`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *PostgresVersionRead) GetPropertiesOk() (*PostgresVersion, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *PostgresVersionRead) SetProperties(v PostgresVersion)`

SetProperties sets Properties field to given value.



