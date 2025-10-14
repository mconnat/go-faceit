# HubsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **int64** |  | [optional] 
**Items** | Pointer to [**[]Hub**](Hub.md) |  | [optional] 
**Start** | Pointer to **int64** |  | [optional] 

## Methods

### NewHubsList

`func NewHubsList() *HubsList`

NewHubsList instantiates a new HubsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubsListWithDefaults

`func NewHubsListWithDefaults() *HubsList`

NewHubsListWithDefaults instantiates a new HubsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *HubsList) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *HubsList) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *HubsList) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *HubsList) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetItems

`func (o *HubsList) GetItems() []Hub`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *HubsList) GetItemsOk() (*[]Hub, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *HubsList) SetItems(v []Hub)`

SetItems sets Items field to given value.

### HasItems

`func (o *HubsList) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetStart

`func (o *HubsList) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *HubsList) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *HubsList) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *HubsList) HasStart() bool`

HasStart returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


