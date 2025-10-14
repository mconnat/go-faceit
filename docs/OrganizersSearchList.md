# OrganizersSearchList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int64** |  | [optional] 
**Items** | Pointer to [**[]OrganizerSearch**](OrganizerSearch.md) |  | [optional] 
**Start** | Pointer to **int64** |  | [optional] 

## Methods

### NewOrganizersSearchList

`func NewOrganizersSearchList() *OrganizersSearchList`

NewOrganizersSearchList instantiates a new OrganizersSearchList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizersSearchListWithDefaults

`func NewOrganizersSearchListWithDefaults() *OrganizersSearchList`

NewOrganizersSearchListWithDefaults instantiates a new OrganizersSearchList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *OrganizersSearchList) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *OrganizersSearchList) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *OrganizersSearchList) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *OrganizersSearchList) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetItems

`func (o *OrganizersSearchList) GetItems() []OrganizerSearch`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *OrganizersSearchList) GetItemsOk() (*[]OrganizerSearch, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *OrganizersSearchList) SetItems(v []OrganizerSearch)`

SetItems sets Items field to given value.

### HasItems

`func (o *OrganizersSearchList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetStart

`func (o *OrganizersSearchList) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *OrganizersSearchList) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *OrganizersSearchList) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *OrganizersSearchList) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


