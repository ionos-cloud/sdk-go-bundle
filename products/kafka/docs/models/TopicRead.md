# TopicRead

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | **string** | The ID (UUID) of the Topic. | |
|**Type** | **string** | The type of the resource. | |
|**Href** | **string** | The URL of the Topic. | |
|**Metadata** | [**ResourceMetadata**](ResourceMetadata.md) |  | |
|**Properties** | [**Topic**](Topic.md) |  | |

## Methods

### NewTopicRead

`func NewTopicRead(id string, type_ string, href string, metadata ResourceMetadata, properties Topic, ) *TopicRead`

NewTopicRead instantiates a new TopicRead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopicReadWithDefaults

`func NewTopicReadWithDefaults() *TopicRead`

NewTopicReadWithDefaults instantiates a new TopicRead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TopicRead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TopicRead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TopicRead) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *TopicRead) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TopicRead) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TopicRead) SetType(v string)`

SetType sets Type field to given value.


### GetHref

`func (o *TopicRead) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *TopicRead) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *TopicRead) SetHref(v string)`

SetHref sets Href field to given value.


### GetMetadata

`func (o *TopicRead) GetMetadata() ResourceMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TopicRead) GetMetadataOk() (*ResourceMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TopicRead) SetMetadata(v ResourceMetadata)`

SetMetadata sets Metadata field to given value.


### GetProperties

`func (o *TopicRead) GetProperties() Topic`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TopicRead) GetPropertiesOk() (*Topic, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TopicRead) SetProperties(v Topic)`

SetProperties sets Properties field to given value.



