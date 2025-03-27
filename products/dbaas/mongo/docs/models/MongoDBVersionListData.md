# MongoDBVersionListData

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The MongoDB version of your cluster. | [optional] |

## Methods

### NewMongoDBVersionListData

`func NewMongoDBVersionListData() *MongoDBVersionListData`

NewMongoDBVersionListData instantiates a new MongoDBVersionListData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMongoDBVersionListDataWithDefaults

`func NewMongoDBVersionListDataWithDefaults() *MongoDBVersionListData`

NewMongoDBVersionListDataWithDefaults instantiates a new MongoDBVersionListData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MongoDBVersionListData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MongoDBVersionListData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MongoDBVersionListData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MongoDBVersionListData) HasName() bool`

HasName returns a boolean if a field has been set.


