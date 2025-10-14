# MatchStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rounds** | Pointer to [**[]RoundStats**](RoundStats.md) |  | [optional] 

## Methods

### NewMatchStats

`func NewMatchStats() *MatchStats`

NewMatchStats instantiates a new MatchStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchStatsWithDefaults

`func NewMatchStatsWithDefaults() *MatchStats`

NewMatchStatsWithDefaults instantiates a new MatchStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRounds

`func (o *MatchStats) GetRounds() []RoundStats`

GetRounds returns the Rounds field if non-nil, zero value otherwise.

### GetRoundsOk

`func (o *MatchStats) GetRoundsOk() (*[]RoundStats, bool)`

GetRoundsOk returns a tuple with the Rounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRounds

`func (o *MatchStats) SetRounds(v []RoundStats)`

SetRounds sets Rounds field to given value.

### HasRounds

`func (o *MatchStats) HasRounds() bool`

HasRounds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


