# DeleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | Pointer to **string** | The image ID of an image that was deleted | [optional] 
**Untagged** | Pointer to **string** | The image ID of an image that was untagged | [optional] 

## Methods

### NewDeleteResponse

`func NewDeleteResponse() *DeleteResponse`

NewDeleteResponse instantiates a new DeleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteResponseWithDefaults

`func NewDeleteResponseWithDefaults() *DeleteResponse`

NewDeleteResponseWithDefaults instantiates a new DeleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *DeleteResponse) GetDeleted() string`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *DeleteResponse) GetDeletedOk() (*string, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *DeleteResponse) SetDeleted(v string)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *DeleteResponse) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetUntagged

`func (o *DeleteResponse) GetUntagged() string`

GetUntagged returns the Untagged field if non-nil, zero value otherwise.

### GetUntaggedOk

`func (o *DeleteResponse) GetUntaggedOk() (*string, bool)`

GetUntaggedOk returns a tuple with the Untagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagged

`func (o *DeleteResponse) SetUntagged(v string)`

SetUntagged sets Untagged field to given value.

### HasUntagged

`func (o *DeleteResponse) HasUntagged() bool`

HasUntagged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


