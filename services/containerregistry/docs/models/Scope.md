# Scope

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Actions** | **[]string** |  | |
|**Name** | **string** |  | |
|**Type** | **string** |  | |

## Methods

### NewScope

`func NewScope(actions []string, name string, type_ string, ) *Scope`

NewScope instantiates a new Scope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeWithDefaults

`func NewScopeWithDefaults() *Scope`

NewScopeWithDefaults instantiates a new Scope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *Scope) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *Scope) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *Scope) SetActions(v []string)`

SetActions sets Actions field to given value.


### SetActionsNil

`func (o *Scope) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *Scope) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetName

`func (o *Scope) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Scope) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Scope) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Scope) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Scope) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Scope) SetType(v string)`

SetType sets Type field to given value.



