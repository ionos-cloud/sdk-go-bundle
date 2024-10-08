# Purl

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Type** | **string** | The affected package type | |
|**Name** | **string** | The affected package name | |
|**Version** | **string** | The affected package version | |

## Methods

### NewPurl

`func NewPurl(type_ string, name string, version string, ) *Purl`

NewPurl instantiates a new Purl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurlWithDefaults

`func NewPurlWithDefaults() *Purl`

NewPurlWithDefaults instantiates a new Purl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Purl) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Purl) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Purl) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *Purl) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Purl) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Purl) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *Purl) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Purl) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Purl) SetVersion(v string)`

SetVersion sets Version field to given value.



