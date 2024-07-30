# ResourceStatus

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Status** | **string** | The current status of the resource. The status can be:  * &#x60;AVAILABLE&#x60; - resource exists and is healthy. * &#x60;PROVISIONING&#x60; - resource is being created or updated. * &#x60;DESTROYING&#x60; - delete command was issued, the resource is being deleted. * &#x60;FAILED&#x60;: - resource failed, details in &#x60;statusMessage&#x60;.  | [readonly] |
|**StatusMessage** | Pointer to **string** | The message of the failure if the status is &#x60;FAILED&#x60;.  | [optional] [readonly] |

## Methods

### NewResourceStatus

`func NewResourceStatus(status string, ) *ResourceStatus`

NewResourceStatus instantiates a new ResourceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceStatusWithDefaults

`func NewResourceStatusWithDefaults() *ResourceStatus`

NewResourceStatusWithDefaults instantiates a new ResourceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ResourceStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *ResourceStatus) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ResourceStatus) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ResourceStatus) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ResourceStatus) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.


