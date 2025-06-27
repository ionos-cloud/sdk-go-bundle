# MetadataWithStatusAllOf

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**State** | **string** | The status of the object. The status can be: * &#x60;AVAILABLE&#x60; - resource exists and is healthy. * &#x60;PROVISIONING&#x60; - resource is being created or updated. * &#x60;DESTROYING&#x60; - delete command was issued, the resource is being deleted. * &#x60;FAILED&#x60; - resource failed, details in &#x60;failureMessage&#x60;.  | [readonly] |

## Methods

### NewMetadataWithStatusAllOf

`func NewMetadataWithStatusAllOf(state string, ) *MetadataWithStatusAllOf`

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



