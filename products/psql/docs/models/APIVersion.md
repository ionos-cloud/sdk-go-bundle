# APIVersion

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Name** | Pointer to **string** |  | [optional] |
|**SwaggerUrl** | Pointer to **string** |  | [optional] |

## Methods

### NewAPIVersion

`func NewAPIVersion() *APIVersion`

NewAPIVersion instantiates a new APIVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIVersionWithDefaults

`func NewAPIVersionWithDefaults() *APIVersion`

NewAPIVersionWithDefaults instantiates a new APIVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *APIVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *APIVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *APIVersion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *APIVersion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSwaggerUrl

`func (o *APIVersion) GetSwaggerUrl() string`

GetSwaggerUrl returns the SwaggerUrl field if non-nil, zero value otherwise.

### GetSwaggerUrlOk

`func (o *APIVersion) GetSwaggerUrlOk() (*string, bool)`

GetSwaggerUrlOk returns a tuple with the SwaggerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwaggerUrl

`func (o *APIVersion) SetSwaggerUrl(v string)`

SetSwaggerUrl sets SwaggerUrl field to given value.

### HasSwaggerUrl

`func (o *APIVersion) HasSwaggerUrl() bool`

HasSwaggerUrl returns a boolean if a field has been set.


