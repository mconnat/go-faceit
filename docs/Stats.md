# Stats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rating** | Pointer to **int64** |  | [optional] 
**SkillLevel** | Pointer to [**StatsSkillLevel**](StatsSkillLevel.md) |  | [optional] 
**WinProbability** | Pointer to **float64** |  | [optional] 

## Methods

### NewStats

`func NewStats() *Stats`

NewStats instantiates a new Stats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsWithDefaults

`func NewStatsWithDefaults() *Stats`

NewStatsWithDefaults instantiates a new Stats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRating

`func (o *Stats) GetRating() int64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *Stats) GetRatingOk() (*int64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *Stats) SetRating(v int64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *Stats) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetSkillLevel

`func (o *Stats) GetSkillLevel() StatsSkillLevel`

GetSkillLevel returns the SkillLevel field if non-nil, zero value otherwise.

### GetSkillLevelOk

`func (o *Stats) GetSkillLevelOk() (*StatsSkillLevel, bool)`

GetSkillLevelOk returns a tuple with the SkillLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillLevel

`func (o *Stats) SetSkillLevel(v StatsSkillLevel)`

SetSkillLevel sets SkillLevel field to given value.

### HasSkillLevel

`func (o *Stats) HasSkillLevel() bool`

HasSkillLevel returns a boolean if a field has been set.

### GetWinProbability

`func (o *Stats) GetWinProbability() float64`

GetWinProbability returns the WinProbability field if non-nil, zero value otherwise.

### GetWinProbabilityOk

`func (o *Stats) GetWinProbabilityOk() (*float64, bool)`

GetWinProbabilityOk returns a tuple with the WinProbability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinProbability

`func (o *Stats) SetWinProbability(v float64)`

SetWinProbability sets WinProbability field to given value.

### HasWinProbability

`func (o *Stats) HasWinProbability() bool`

HasWinProbability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


