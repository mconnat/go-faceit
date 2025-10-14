# RoundStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BestOf** | Pointer to **map[string]interface{}** |  | [optional] 
**CompetitionId** | Pointer to **map[string]interface{}** |  | [optional] 
**GameId** | Pointer to **map[string]interface{}** |  | [optional] 
**GameMode** | Pointer to **map[string]interface{}** |  | [optional] 
**MatchId** | Pointer to **map[string]interface{}** |  | [optional] 
**MatchRound** | Pointer to **map[string]interface{}** |  | [optional] 
**Played** | Pointer to **map[string]interface{}** |  | [optional] 
**RoundStats** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Teams** | Pointer to [**[]TeamStatsSimple**](TeamStatsSimple.md) |  | [optional] 

## Methods

### NewRoundStats

`func NewRoundStats() *RoundStats`

NewRoundStats instantiates a new RoundStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoundStatsWithDefaults

`func NewRoundStatsWithDefaults() *RoundStats`

NewRoundStatsWithDefaults instantiates a new RoundStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBestOf

`func (o *RoundStats) GetBestOf() map[string]interface{}`

GetBestOf returns the BestOf field if non-nil, zero value otherwise.

### GetBestOfOk

`func (o *RoundStats) GetBestOfOk() (*map[string]interface{}, bool)`

GetBestOfOk returns a tuple with the BestOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestOf

`func (o *RoundStats) SetBestOf(v map[string]interface{})`

SetBestOf sets BestOf field to given value.

### HasBestOf

`func (o *RoundStats) HasBestOf() bool`

HasBestOf returns a boolean if a field has been set.

### GetCompetitionId

`func (o *RoundStats) GetCompetitionId() map[string]interface{}`

GetCompetitionId returns the CompetitionId field if non-nil, zero value otherwise.

### GetCompetitionIdOk

`func (o *RoundStats) GetCompetitionIdOk() (*map[string]interface{}, bool)`

GetCompetitionIdOk returns a tuple with the CompetitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitionId

`func (o *RoundStats) SetCompetitionId(v map[string]interface{})`

SetCompetitionId sets CompetitionId field to given value.

### HasCompetitionId

`func (o *RoundStats) HasCompetitionId() bool`

HasCompetitionId returns a boolean if a field has been set.

### GetGameId

`func (o *RoundStats) GetGameId() map[string]interface{}`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *RoundStats) GetGameIdOk() (*map[string]interface{}, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *RoundStats) SetGameId(v map[string]interface{})`

SetGameId sets GameId field to given value.

### HasGameId

`func (o *RoundStats) HasGameId() bool`

HasGameId returns a boolean if a field has been set.

### GetGameMode

`func (o *RoundStats) GetGameMode() map[string]interface{}`

GetGameMode returns the GameMode field if non-nil, zero value otherwise.

### GetGameModeOk

`func (o *RoundStats) GetGameModeOk() (*map[string]interface{}, bool)`

GetGameModeOk returns a tuple with the GameMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameMode

`func (o *RoundStats) SetGameMode(v map[string]interface{})`

SetGameMode sets GameMode field to given value.

### HasGameMode

`func (o *RoundStats) HasGameMode() bool`

HasGameMode returns a boolean if a field has been set.

### GetMatchId

`func (o *RoundStats) GetMatchId() map[string]interface{}`

GetMatchId returns the MatchId field if non-nil, zero value otherwise.

### GetMatchIdOk

`func (o *RoundStats) GetMatchIdOk() (*map[string]interface{}, bool)`

GetMatchIdOk returns a tuple with the MatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchId

`func (o *RoundStats) SetMatchId(v map[string]interface{})`

SetMatchId sets MatchId field to given value.

### HasMatchId

`func (o *RoundStats) HasMatchId() bool`

HasMatchId returns a boolean if a field has been set.

### GetMatchRound

`func (o *RoundStats) GetMatchRound() map[string]interface{}`

GetMatchRound returns the MatchRound field if non-nil, zero value otherwise.

### GetMatchRoundOk

`func (o *RoundStats) GetMatchRoundOk() (*map[string]interface{}, bool)`

GetMatchRoundOk returns a tuple with the MatchRound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRound

`func (o *RoundStats) SetMatchRound(v map[string]interface{})`

SetMatchRound sets MatchRound field to given value.

### HasMatchRound

`func (o *RoundStats) HasMatchRound() bool`

HasMatchRound returns a boolean if a field has been set.

### GetPlayed

`func (o *RoundStats) GetPlayed() map[string]interface{}`

GetPlayed returns the Played field if non-nil, zero value otherwise.

### GetPlayedOk

`func (o *RoundStats) GetPlayedOk() (*map[string]interface{}, bool)`

GetPlayedOk returns a tuple with the Played field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayed

`func (o *RoundStats) SetPlayed(v map[string]interface{})`

SetPlayed sets Played field to given value.

### HasPlayed

`func (o *RoundStats) HasPlayed() bool`

HasPlayed returns a boolean if a field has been set.

### GetRoundStats

`func (o *RoundStats) GetRoundStats() map[string]map[string]interface{}`

GetRoundStats returns the RoundStats field if non-nil, zero value otherwise.

### GetRoundStatsOk

`func (o *RoundStats) GetRoundStatsOk() (*map[string]map[string]interface{}, bool)`

GetRoundStatsOk returns a tuple with the RoundStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundStats

`func (o *RoundStats) SetRoundStats(v map[string]map[string]interface{})`

SetRoundStats sets RoundStats field to given value.

### HasRoundStats

`func (o *RoundStats) HasRoundStats() bool`

HasRoundStats returns a boolean if a field has been set.

### GetTeams

`func (o *RoundStats) GetTeams() []TeamStatsSimple`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *RoundStats) GetTeamsOk() (*[]TeamStatsSimple, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *RoundStats) SetTeams(v []TeamStatsSimple)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *RoundStats) HasTeams() bool`

HasTeams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


