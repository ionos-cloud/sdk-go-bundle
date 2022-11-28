# ApiInfoDto

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Href** | Pointer to **string** | The API entry point. | [optional] [readonly] |
|**Name** | Pointer to **string** | The API name. | [optional] [readonly] |
|**Version** | Pointer to **string** | The API version. | [optional] [readonly] |

## Methods

### NewApiInfoDto

`func NewApiInfoDto() *ApiInfoDto`

NewApiInfoDto instantiates a new ApiInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiInfoDtoWithDefaults

`func NewApiInfoDtoWithDefaults() *ApiInfoDto`

NewApiInfoDtoWithDefaults instantiates a new ApiInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *ApiInfoDto) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ApiInfoDto) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ApiInfoDto) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ApiInfoDto) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetName

`func (o *ApiInfoDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiInfoDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiInfoDto) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiInfoDto) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVersion

`func (o *ApiInfoDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApiInfoDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApiInfoDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApiInfoDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


