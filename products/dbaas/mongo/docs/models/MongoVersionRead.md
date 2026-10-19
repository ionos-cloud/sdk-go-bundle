# MongoVersionRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the MongoVersion. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the MongoVersion. | |
|**Metadata** | **map[string]interface{}** |  | [readonly] |
|**Properties** | [**MongoVersion**](MongoVersion.md) |  | |

## Methods

### NewMongoVersionRead

`func NewMongoVersionRead(id string, type_ string, href string, metadata map[string]interface{}, properties MongoVersion, ) *MongoVersionRead`

NewMongoVersionRead instantiates a new MongoVersionRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoVersionReadWithDefaults

`func NewMongoVersionReadWithDefaults() *MongoVersionRead`

NewMongoVersionReadWithDefaults instantiates a new MongoVersionRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MongoVersionRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MongoVersionRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MongoVersionRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *MongoVersionRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MongoVersionRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MongoVersionRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *MongoVersionRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *MongoVersionRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *MongoVersionRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *MongoVersionRead) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *MongoVersionRead) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *MongoVersionRead) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *MongoVersionRead) GetProperties() MongoVersion`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *MongoVersionRead) GetPropertiesOk() (*MongoVersion, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *MongoVersionRead) SetProperties(v MongoVersion)`

SetProperties sets Properties field to given value.



