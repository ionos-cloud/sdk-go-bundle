# ProfilesGet200Response

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**Companies** | Pointer to [**[]Metadata**](Metadata.md) |  | [optional] |

## Methods

### NewProfilesGet200Response

`func NewProfilesGet200Response() *ProfilesGet200Response`

NewProfilesGet200Response instantiates a new ProfilesGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfilesGet200ResponseWithDefaults

`func NewProfilesGet200ResponseWithDefaults() *ProfilesGet200Response`

NewProfilesGet200ResponseWithDefaults instantiates a new ProfilesGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompanies

`func (o *ProfilesGet200Response) GetCompanies() []Metadata`

GetCompanies returns the Companies field if non-nil, zero value otherwise.

### GetCompaniesOk

`func (o *ProfilesGet200Response) GetCompaniesOk() (*[]Metadata, bool)`

GetCompaniesOk returns a tuple with the Companies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanies

`func (o *ProfilesGet200Response) SetCompanies(v []Metadata)`

SetCompanies sets Companies field to given value.

### HasCompanies

`func (o *ProfilesGet200Response) HasCompanies() bool`

HasCompanies returns a boolean if a field has been set.


