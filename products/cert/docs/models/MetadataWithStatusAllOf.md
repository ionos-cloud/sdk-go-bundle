# MetadataWithStatusAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**State** | **string** | The resource state. | |
|**Message** | **string** | A human readable message describing the current state. In case of an error, the message will contain a detailed error message.  | |

## Methods

### NewMetadataWithStatusAllOf

`func NewMetadataWithStatusAllOf(state string, message string, ) *MetadataWithStatusAllOf`

NewMetadataWithStatusAllOf instantiates a new MetadataWithStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithStatusAllOfWithDefaults

`func NewMetadataWithStatusAllOfWithDefaults() *MetadataWithStatusAllOf`

NewMetadataWithStatusAllOfWithDefaults instantiates a new MetadataWithStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *MetadataWithStatusAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MetadataWithStatusAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MetadataWithStatusAllOf) SetState(v string)`

SetState sets State field to given value.


### GetMessage

`func (o *MetadataWithStatusAllOf) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MetadataWithStatusAllOf) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MetadataWithStatusAllOf) SetMessage(v string)`

SetMessage sets Message field to given value.



