# PostgresVersionReadListAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | ID of the list of PostgresVersion resources. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the list of PostgresVersion resources. | |
|**Items** | Pointer to [**[]PostgresVersionRead**](PostgresVersionRead.md) | The list of PostgresVersion resources. | [optional] |

## Methods

### NewPostgresVersionReadListAllOf

`func NewPostgresVersionReadListAllOf(id string, type_ string, href string, ) *PostgresVersionReadListAllOf`

NewPostgresVersionReadListAllOf instantiates a new PostgresVersionReadListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresVersionReadListAllOfWithDefaults

`func NewPostgresVersionReadListAllOfWithDefaults() *PostgresVersionReadListAllOf`

NewPostgresVersionReadListAllOfWithDefaults instantiates a new PostgresVersionReadListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostgresVersionReadListAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostgresVersionReadListAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostgresVersionReadListAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *PostgresVersionReadListAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostgresVersionReadListAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostgresVersionReadListAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *PostgresVersionReadListAllOf) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *PostgresVersionReadListAllOf) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *PostgresVersionReadListAllOf) SetHref(v string)`

SetHref sets Href field to given value.


### GetItems

`func (o *PostgresVersionReadListAllOf) GetItems() []PostgresVersionRead`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PostgresVersionReadListAllOf) GetItemsOk() (*[]PostgresVersionRead, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PostgresVersionReadListAllOf) SetItems(v []PostgresVersionRead)`

SetItems sets Items field to given value.

### HasItems

`func (o *PostgresVersionReadListAllOf) HasItems() bool`

HasItems returns a boolean if a field has been set.


