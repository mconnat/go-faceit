# Season

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | Pointer to **time.Time** | The end date of the season. | [optional] 
**Number** | Pointer to **int64** | The season number. | [optional] 
**PlacementMatchCount** | Pointer to **int64** | The placement match mount. | [optional] 
**StartDate** | Pointer to **time.Time** | The start date of the season. | [optional] 

## Methods

### NewSeason

`func NewSeason() *Season`

NewSeason instantiates a new Season object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeasonWithDefaults

`func NewSeasonWithDefaults() *Season`

NewSeasonWithDefaults instantiates a new Season object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *Season) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Season) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Season) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Season) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetNumber

`func (o *Season) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Season) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Season) SetNumber(v int64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Season) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPlacementMatchCount

`func (o *Season) GetPlacementMatchCount() int64`

GetPlacementMatchCount returns the PlacementMatchCount field if non-nil, zero value otherwise.

### GetPlacementMatchCountOk

`func (o *Season) GetPlacementMatchCountOk() (*int64, bool)`

GetPlacementMatchCountOk returns a tuple with the PlacementMatchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementMatchCount

`func (o *Season) SetPlacementMatchCount(v int64)`

SetPlacementMatchCount sets PlacementMatchCount field to given value.

### HasPlacementMatchCount

`func (o *Season) HasPlacementMatchCount() bool`

HasPlacementMatchCount returns a boolean if a field has been set.

### GetStartDate

`func (o *Season) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Season) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Season) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *Season) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


