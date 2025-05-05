# ClusterLogsInstancesMessages

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Time** | Pointer to [**time.Time**](time.Time.md) |  | [optional] |
|**Message** | Pointer to **string** |  | [optional] |

## Methods

### NewClusterLogsInstancesMessages

`func NewClusterLogsInstancesMessages() *ClusterLogsInstancesMessages`

NewClusterLogsInstancesMessages instantiates a new ClusterLogsInstancesMessages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLogsInstancesMessagesWithDefaults

`func NewClusterLogsInstancesMessagesWithDefaults() *ClusterLogsInstancesMessages`

NewClusterLogsInstancesMessagesWithDefaults instantiates a new ClusterLogsInstancesMessages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *ClusterLogsInstancesMessages) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ClusterLogsInstancesMessages) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ClusterLogsInstancesMessages) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ClusterLogsInstancesMessages) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetMessage

`func (o *ClusterLogsInstancesMessages) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterLogsInstancesMessages) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterLogsInstancesMessages) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClusterLogsInstancesMessages) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


