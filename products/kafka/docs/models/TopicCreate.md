# TopicCreate

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Metadata** | Pointer to **map[string]interface{}** | Metadata | [optional] |
|**Properties** | [**Topic**](Topic.md) |  | |

## Methods

### NewTopicCreate

`func NewTopicCreate(properties Topic, ) *TopicCreate`

NewTopicCreate instantiates a new TopicCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicCreateWithDefaults

`func NewTopicCreateWithDefaults() *TopicCreate`

NewTopicCreateWithDefaults instantiates a new TopicCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *TopicCreate) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TopicCreate) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TopicCreate) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TopicCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProperties

`func (o *TopicCreate) GetProperties() Topic`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TopicCreate) GetPropertiesOk() (*Topic, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TopicCreate) SetProperties(v Topic)`

SetProperties sets Properties field to given value.



