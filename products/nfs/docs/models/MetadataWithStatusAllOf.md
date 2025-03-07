# MetadataWithStatusAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Status** | **string** | The status of the resource can be one of the following:  * &#x60;AVAILABLE&#x60; - The resource exists and is healthy. * &#x60;PROVISIONING&#x60; - The resource is being created or updated. * &#x60;DESTROYING&#x60; - A delete command was issued, and the resource is being deleted. * &#x60;FAILED&#x60; - The resource failed, with details provided in &#x60;statusMessage&#x60;.  | [readonly] |
|**StatusMessage** | Pointer to **string** | The message of the failure if the status is &#x60;FAILED&#x60;.  | [optional] [readonly] |

## Methods

### NewMetadataWithStatusAllOf

`func NewMetadataWithStatusAllOf(status string, ) *MetadataWithStatusAllOf`

NewMetadataWithStatusAllOf instantiates a new MetadataWithStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataWithStatusAllOfWithDefaults

`func NewMetadataWithStatusAllOfWithDefaults() *MetadataWithStatusAllOf`

NewMetadataWithStatusAllOfWithDefaults instantiates a new MetadataWithStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *MetadataWithStatusAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MetadataWithStatusAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MetadataWithStatusAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusMessage

`func (o *MetadataWithStatusAllOf) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *MetadataWithStatusAllOf) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *MetadataWithStatusAllOf) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *MetadataWithStatusAllOf) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.


