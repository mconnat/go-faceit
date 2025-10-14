# DetailedMatchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AscScore** | Pointer to **bool** | True if the scores should be interpreted ascending (lower score wins) | [optional] 
**Factions** | Pointer to [**map[string]FactionResult**](FactionResult.md) | The factions of the match | [optional] 
**Winner** | Pointer to **string** | The winner faction | [optional] 

## Methods

### NewDetailedMatchResult

`func NewDetailedMatchResult() *DetailedMatchResult`

NewDetailedMatchResult instantiates a new DetailedMatchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedMatchResultWithDefaults

`func NewDetailedMatchResultWithDefaults() *DetailedMatchResult`

NewDetailedMatchResultWithDefaults instantiates a new DetailedMatchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAscScore

`func (o *DetailedMatchResult) GetAscScore() bool`

GetAscScore returns the AscScore field if non-nil, zero value otherwise.

### GetAscScoreOk

`func (o *DetailedMatchResult) GetAscScoreOk() (*bool, bool)`

GetAscScoreOk returns a tuple with the AscScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAscScore

`func (o *DetailedMatchResult) SetAscScore(v bool)`

SetAscScore sets AscScore field to given value.

### HasAscScore

`func (o *DetailedMatchResult) HasAscScore() bool`

HasAscScore returns a boolean if a field has been set.

### GetFactions

`func (o *DetailedMatchResult) GetFactions() map[string]FactionResult`

GetFactions returns the Factions field if non-nil, zero value otherwise.

### GetFactionsOk

`func (o *DetailedMatchResult) GetFactionsOk() (*map[string]FactionResult, bool)`

GetFactionsOk returns a tuple with the Factions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactions

`func (o *DetailedMatchResult) SetFactions(v map[string]FactionResult)`

SetFactions sets Factions field to given value.

### HasFactions

`func (o *DetailedMatchResult) HasFactions() bool`

HasFactions returns a boolean if a field has been set.

### GetWinner

`func (o *DetailedMatchResult) GetWinner() string`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *DetailedMatchResult) GetWinnerOk() (*string, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *DetailedMatchResult) SetWinner(v string)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *DetailedMatchResult) HasWinner() bool`

HasWinner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


