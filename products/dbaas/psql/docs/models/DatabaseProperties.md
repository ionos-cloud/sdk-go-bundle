# DatabaseProperties

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | **string** | The databasename of a given database. | |
|**Owner** | **string** | The name of the role owning a given database. | |

## Methods

### NewDatabaseProperties

`func NewDatabaseProperties(name string, owner string, ) *DatabaseProperties`

NewDatabaseProperties instantiates a new DatabaseProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatabasePropertiesWithDefaults

`func NewDatabasePropertiesWithDefaults() *DatabaseProperties`

NewDatabasePropertiesWithDefaults instantiates a new DatabaseProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatabaseProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatabaseProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatabaseProperties) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *DatabaseProperties) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DatabaseProperties) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DatabaseProperties) SetOwner(v string)`

SetOwner sets Owner field to given value.



