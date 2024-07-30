# ClusterLogsInstances

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** | The name of the MongoDB instance. | [optional] |
|**Messages** | Pointer to [**[]ClusterLogsInstancesMessages**](ClusterLogsInstancesMessages.md) |  | [optional] |

## Methods

### NewClusterLogsInstances

`func NewClusterLogsInstances() *ClusterLogsInstances`

NewClusterLogsInstances instantiates a new ClusterLogsInstances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLogsInstancesWithDefaults

`func NewClusterLogsInstancesWithDefaults() *ClusterLogsInstances`

NewClusterLogsInstancesWithDefaults instantiates a new ClusterLogsInstances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterLogsInstances) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterLogsInstances) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterLogsInstances) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterLogsInstances) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMessages

`func (o *ClusterLogsInstances) GetMessages() []ClusterLogsInstancesMessages`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *ClusterLogsInstances) GetMessagesOk() (*[]ClusterLogsInstancesMessages, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *ClusterLogsInstances) SetMessages(v []ClusterLogsInstancesMessages)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *ClusterLogsInstances) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


