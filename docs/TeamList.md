# TeamList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int64** |  | [optional] 
**Items** | [**[]Team**](Team.md) | The teams list. | [readonly] 
**Start** | Pointer to **int64** |  | [optional] 

## Methods

### NewTeamList

`func NewTeamList(items []Team, ) *TeamList`

NewTeamList instantiates a new TeamList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamListWithDefaults

`func NewTeamListWithDefaults() *TeamList`

NewTeamListWithDefaults instantiates a new TeamList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *TeamList) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *TeamList) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *TeamList) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *TeamList) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetItems

`func (o *TeamList) GetItems() []Team`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TeamList) GetItemsOk() (*[]Team, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TeamList) SetItems(v []Team)`

SetItems sets Items field to given value.


### GetStart

`func (o *TeamList) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *TeamList) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *TeamList) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *TeamList) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


