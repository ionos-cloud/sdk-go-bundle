# PutBucketWebsiteRequest

## Properties

|Name | Type | Description | Notes|
|------------ | ------------- | ------------- | -------------|
|**ErrorDocument** | Pointer to [**ErrorDocument**](ErrorDocument.md) |  | [optional] |
|**IndexDocument** | Pointer to [**IndexDocument**](IndexDocument.md) |  | [optional] |
|**RedirectAllRequestsTo** | Pointer to [**RedirectAllRequestsTo**](RedirectAllRequestsTo.md) |  | [optional] |
|**RoutingRules** | Pointer to [**[]RoutingRule**](RoutingRule.md) |  | [optional] |

## Methods

### NewPutBucketWebsiteRequest

`func NewPutBucketWebsiteRequest() *PutBucketWebsiteRequest`

NewPutBucketWebsiteRequest instantiates a new PutBucketWebsiteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutBucketWebsiteRequestWithDefaults

`func NewPutBucketWebsiteRequestWithDefaults() *PutBucketWebsiteRequest`

NewPutBucketWebsiteRequestWithDefaults instantiates a new PutBucketWebsiteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorDocument

`func (o *PutBucketWebsiteRequest) GetErrorDocument() ErrorDocument`

GetErrorDocument returns the ErrorDocument field if non-nil, zero value otherwise.

### GetErrorDocumentOk

`func (o *PutBucketWebsiteRequest) GetErrorDocumentOk() (*ErrorDocument, bool)`

GetErrorDocumentOk returns a tuple with the ErrorDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDocument

`func (o *PutBucketWebsiteRequest) SetErrorDocument(v ErrorDocument)`

SetErrorDocument sets ErrorDocument field to given value.

### HasErrorDocument

`func (o *PutBucketWebsiteRequest) HasErrorDocument() bool`

HasErrorDocument returns a boolean if a field has been set.

### GetIndexDocument

`func (o *PutBucketWebsiteRequest) GetIndexDocument() IndexDocument`

GetIndexDocument returns the IndexDocument field if non-nil, zero value otherwise.

### GetIndexDocumentOk

`func (o *PutBucketWebsiteRequest) GetIndexDocumentOk() (*IndexDocument, bool)`

GetIndexDocumentOk returns a tuple with the IndexDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexDocument

`func (o *PutBucketWebsiteRequest) SetIndexDocument(v IndexDocument)`

SetIndexDocument sets IndexDocument field to given value.

### HasIndexDocument

`func (o *PutBucketWebsiteRequest) HasIndexDocument() bool`

HasIndexDocument returns a boolean if a field has been set.

### GetRedirectAllRequestsTo

`func (o *PutBucketWebsiteRequest) GetRedirectAllRequestsTo() RedirectAllRequestsTo`

GetRedirectAllRequestsTo returns the RedirectAllRequestsTo field if non-nil, zero value otherwise.

### GetRedirectAllRequestsToOk

`func (o *PutBucketWebsiteRequest) GetRedirectAllRequestsToOk() (*RedirectAllRequestsTo, bool)`

GetRedirectAllRequestsToOk returns a tuple with the RedirectAllRequestsTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectAllRequestsTo

`func (o *PutBucketWebsiteRequest) SetRedirectAllRequestsTo(v RedirectAllRequestsTo)`

SetRedirectAllRequestsTo sets RedirectAllRequestsTo field to given value.

### HasRedirectAllRequestsTo

`func (o *PutBucketWebsiteRequest) HasRedirectAllRequestsTo() bool`

HasRedirectAllRequestsTo returns a boolean if a field has been set.

### GetRoutingRules

`func (o *PutBucketWebsiteRequest) GetRoutingRules() []RoutingRule`

GetRoutingRules returns the RoutingRules field if non-nil, zero value otherwise.

### GetRoutingRulesOk

`func (o *PutBucketWebsiteRequest) GetRoutingRulesOk() (*[]RoutingRule, bool)`

GetRoutingRulesOk returns a tuple with the RoutingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingRules

`func (o *PutBucketWebsiteRequest) SetRoutingRules(v []RoutingRule)`

SetRoutingRules sets RoutingRules field to given value.

### HasRoutingRules

`func (o *PutBucketWebsiteRequest) HasRoutingRules() bool`

HasRoutingRules returns a boolean if a field has been set.


