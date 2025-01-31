# ResourceState

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**State** | **string** | State of the resource. Resource states: &#x60;AVAILABLE&#x60;: There are no pending modification requests for this item. &#x60;BUSY&#x60;: There is at least one modification request pending and all following requests will be queued. &#x60;DEPLOYING&#x60;: The resource is being created. &#x60;FAILED&#x60;: The creation of the resource failed. &#x60;UPDATING&#x60;: The resource is being updated. &#x60;FAILED_UPDATING&#x60;: An update to the resource was not successful. &#x60;DESTROYING&#x60;: A delete command was issued, and the resource is being deleted.  | |
|**Message** | Pointer to **string** | A human readable message describing the current state. In case of an error, the message will contain a detailed error message.  | [optional] |

## Methods

### NewResourceState

`func NewResourceState(state string, ) *ResourceState`

NewResourceState instantiates a new ResourceState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceStateWithDefaults

`func NewResourceStateWithDefaults() *ResourceState`

NewResourceStateWithDefaults instantiates a new ResourceState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *ResourceState) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ResourceState) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ResourceState) SetState(v string)`

SetState sets State field to given value.


### GetMessage

`func (o *ResourceState) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ResourceState) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ResourceState) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ResourceState) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


