# ReferenceById

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Id** | Pointer to **int64** | The Contract ID | [optional] |
|**Type** | Pointer to **string** | A string indicating the type | [optional] |
|**Href** | Pointer to **string** | A URL reference | [optional] |

## Methods

### NewReferenceById

`func NewReferenceById() *ReferenceById`

NewReferenceById instantiates a new ReferenceById object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceByIdWithDefaults

`func NewReferenceByIdWithDefaults() *ReferenceById`

NewReferenceByIdWithDefaults instantiates a new ReferenceById object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReferenceById) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReferenceById) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReferenceById) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ReferenceById) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ReferenceById) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReferenceById) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReferenceById) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReferenceById) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHref

`func (o *ReferenceById) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *ReferenceById) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *ReferenceById) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *ReferenceById) HasHref() bool`

HasHref returns a boolean if a field has been set.


