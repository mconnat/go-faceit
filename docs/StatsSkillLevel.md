# StatsSkillLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Average** | Pointer to **int64** |  | [optional] 
**Range** | Pointer to [**StatsSkillLevelRange**](StatsSkillLevelRange.md) |  | [optional] 

## Methods

### NewStatsSkillLevel

`func NewStatsSkillLevel() *StatsSkillLevel`

NewStatsSkillLevel instantiates a new StatsSkillLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsSkillLevelWithDefaults

`func NewStatsSkillLevelWithDefaults() *StatsSkillLevel`

NewStatsSkillLevelWithDefaults instantiates a new StatsSkillLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAverage

`func (o *StatsSkillLevel) GetAverage() int64`

GetAverage returns the Average field if non-nil, zero value otherwise.

### GetAverageOk

`func (o *StatsSkillLevel) GetAverageOk() (*int64, bool)`

GetAverageOk returns a tuple with the Average field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverage

`func (o *StatsSkillLevel) SetAverage(v int64)`

SetAverage sets Average field to given value.

### HasAverage

`func (o *StatsSkillLevel) HasAverage() bool`

HasAverage returns a boolean if a field has been set.

### GetRange

`func (o *StatsSkillLevel) GetRange() StatsSkillLevelRange`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *StatsSkillLevel) GetRangeOk() (*StatsSkillLevelRange, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *StatsSkillLevel) SetRange(v StatsSkillLevelRange)`

SetRange sets Range field to given value.

### HasRange

`func (o *StatsSkillLevel) HasRange() bool`

HasRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


