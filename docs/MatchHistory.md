# MatchHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompetitionId** | Pointer to **string** |  | [optional] 
**CompetitionName** | Pointer to **string** |  | [optional] 
**CompetitionType** | Pointer to **string** |  | [optional] 
**FaceitUrl** | Pointer to **string** |  | [optional] 
**FinishedAt** | Pointer to **int64** |  | [optional] 
**GameId** | Pointer to **string** |  | [optional] 
**GameMode** | Pointer to **string** |  | [optional] 
**MatchId** | Pointer to **string** |  | [optional] 
**MatchType** | Pointer to **string** |  | [optional] 
**MaxPlayers** | Pointer to **int64** |  | [optional] 
**OrganizerId** | Pointer to **string** |  | [optional] 
**PlayingPlayers** | Pointer to **[]string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**MatchResult**](MatchResult.md) |  | [optional] 
**StartedAt** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Teams** | Pointer to [**map[string]HistoryFaction**](HistoryFaction.md) |  | [optional] 
**TeamsSize** | Pointer to **int64** |  | [optional] 

## Methods

### NewMatchHistory

`func NewMatchHistory() *MatchHistory`

NewMatchHistory instantiates a new MatchHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchHistoryWithDefaults

`func NewMatchHistoryWithDefaults() *MatchHistory`

NewMatchHistoryWithDefaults instantiates a new MatchHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompetitionId

`func (o *MatchHistory) GetCompetitionId() string`

GetCompetitionId returns the CompetitionId field if non-nil, zero value otherwise.

### GetCompetitionIdOk

`func (o *MatchHistory) GetCompetitionIdOk() (*string, bool)`

GetCompetitionIdOk returns a tuple with the CompetitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitionId

`func (o *MatchHistory) SetCompetitionId(v string)`

SetCompetitionId sets CompetitionId field to given value.

### HasCompetitionId

`func (o *MatchHistory) HasCompetitionId() bool`

HasCompetitionId returns a boolean if a field has been set.

### GetCompetitionName

`func (o *MatchHistory) GetCompetitionName() string`

GetCompetitionName returns the CompetitionName field if non-nil, zero value otherwise.

### GetCompetitionNameOk

`func (o *MatchHistory) GetCompetitionNameOk() (*string, bool)`

GetCompetitionNameOk returns a tuple with the CompetitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitionName

`func (o *MatchHistory) SetCompetitionName(v string)`

SetCompetitionName sets CompetitionName field to given value.

### HasCompetitionName

`func (o *MatchHistory) HasCompetitionName() bool`

HasCompetitionName returns a boolean if a field has been set.

### GetCompetitionType

`func (o *MatchHistory) GetCompetitionType() string`

GetCompetitionType returns the CompetitionType field if non-nil, zero value otherwise.

### GetCompetitionTypeOk

`func (o *MatchHistory) GetCompetitionTypeOk() (*string, bool)`

GetCompetitionTypeOk returns a tuple with the CompetitionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitionType

`func (o *MatchHistory) SetCompetitionType(v string)`

SetCompetitionType sets CompetitionType field to given value.

### HasCompetitionType

`func (o *MatchHistory) HasCompetitionType() bool`

HasCompetitionType returns a boolean if a field has been set.

### GetFaceitUrl

`func (o *MatchHistory) GetFaceitUrl() string`

GetFaceitUrl returns the FaceitUrl field if non-nil, zero value otherwise.

### GetFaceitUrlOk

`func (o *MatchHistory) GetFaceitUrlOk() (*string, bool)`

GetFaceitUrlOk returns a tuple with the FaceitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceitUrl

`func (o *MatchHistory) SetFaceitUrl(v string)`

SetFaceitUrl sets FaceitUrl field to given value.

### HasFaceitUrl

`func (o *MatchHistory) HasFaceitUrl() bool`

HasFaceitUrl returns a boolean if a field has been set.

### GetFinishedAt

`func (o *MatchHistory) GetFinishedAt() int64`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *MatchHistory) GetFinishedAtOk() (*int64, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *MatchHistory) SetFinishedAt(v int64)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *MatchHistory) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetGameId

`func (o *MatchHistory) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *MatchHistory) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *MatchHistory) SetGameId(v string)`

SetGameId sets GameId field to given value.

### HasGameId

`func (o *MatchHistory) HasGameId() bool`

HasGameId returns a boolean if a field has been set.

### GetGameMode

`func (o *MatchHistory) GetGameMode() string`

GetGameMode returns the GameMode field if non-nil, zero value otherwise.

### GetGameModeOk

`func (o *MatchHistory) GetGameModeOk() (*string, bool)`

GetGameModeOk returns a tuple with the GameMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameMode

`func (o *MatchHistory) SetGameMode(v string)`

SetGameMode sets GameMode field to given value.

### HasGameMode

`func (o *MatchHistory) HasGameMode() bool`

HasGameMode returns a boolean if a field has been set.

### GetMatchId

`func (o *MatchHistory) GetMatchId() string`

GetMatchId returns the MatchId field if non-nil, zero value otherwise.

### GetMatchIdOk

`func (o *MatchHistory) GetMatchIdOk() (*string, bool)`

GetMatchIdOk returns a tuple with the MatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchId

`func (o *MatchHistory) SetMatchId(v string)`

SetMatchId sets MatchId field to given value.

### HasMatchId

`func (o *MatchHistory) HasMatchId() bool`

HasMatchId returns a boolean if a field has been set.

### GetMatchType

`func (o *MatchHistory) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *MatchHistory) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *MatchHistory) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.

### HasMatchType

`func (o *MatchHistory) HasMatchType() bool`

HasMatchType returns a boolean if a field has been set.

### GetMaxPlayers

`func (o *MatchHistory) GetMaxPlayers() int64`

GetMaxPlayers returns the MaxPlayers field if non-nil, zero value otherwise.

### GetMaxPlayersOk

`func (o *MatchHistory) GetMaxPlayersOk() (*int64, bool)`

GetMaxPlayersOk returns a tuple with the MaxPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPlayers

`func (o *MatchHistory) SetMaxPlayers(v int64)`

SetMaxPlayers sets MaxPlayers field to given value.

### HasMaxPlayers

`func (o *MatchHistory) HasMaxPlayers() bool`

HasMaxPlayers returns a boolean if a field has been set.

### GetOrganizerId

`func (o *MatchHistory) GetOrganizerId() string`

GetOrganizerId returns the OrganizerId field if non-nil, zero value otherwise.

### GetOrganizerIdOk

`func (o *MatchHistory) GetOrganizerIdOk() (*string, bool)`

GetOrganizerIdOk returns a tuple with the OrganizerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizerId

`func (o *MatchHistory) SetOrganizerId(v string)`

SetOrganizerId sets OrganizerId field to given value.

### HasOrganizerId

`func (o *MatchHistory) HasOrganizerId() bool`

HasOrganizerId returns a boolean if a field has been set.

### GetPlayingPlayers

`func (o *MatchHistory) GetPlayingPlayers() []string`

GetPlayingPlayers returns the PlayingPlayers field if non-nil, zero value otherwise.

### GetPlayingPlayersOk

`func (o *MatchHistory) GetPlayingPlayersOk() (*[]string, bool)`

GetPlayingPlayersOk returns a tuple with the PlayingPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayingPlayers

`func (o *MatchHistory) SetPlayingPlayers(v []string)`

SetPlayingPlayers sets PlayingPlayers field to given value.

### HasPlayingPlayers

`func (o *MatchHistory) HasPlayingPlayers() bool`

HasPlayingPlayers returns a boolean if a field has been set.

### GetRegion

`func (o *MatchHistory) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MatchHistory) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MatchHistory) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MatchHistory) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetResults

`func (o *MatchHistory) GetResults() MatchResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *MatchHistory) GetResultsOk() (*MatchResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *MatchHistory) SetResults(v MatchResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *MatchHistory) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetStartedAt

`func (o *MatchHistory) GetStartedAt() int64`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *MatchHistory) GetStartedAtOk() (*int64, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *MatchHistory) SetStartedAt(v int64)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *MatchHistory) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *MatchHistory) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MatchHistory) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MatchHistory) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MatchHistory) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTeams

`func (o *MatchHistory) GetTeams() map[string]HistoryFaction`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *MatchHistory) GetTeamsOk() (*map[string]HistoryFaction, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *MatchHistory) SetTeams(v map[string]HistoryFaction)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *MatchHistory) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetTeamsSize

`func (o *MatchHistory) GetTeamsSize() int64`

GetTeamsSize returns the TeamsSize field if non-nil, zero value otherwise.

### GetTeamsSizeOk

`func (o *MatchHistory) GetTeamsSizeOk() (*int64, bool)`

GetTeamsSizeOk returns a tuple with the TeamsSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsSize

`func (o *MatchHistory) SetTeamsSize(v int64)`

SetTeamsSize sets TeamsSize field to given value.

### HasTeamsSize

`func (o *MatchHistory) HasTeamsSize() bool`

HasTeamsSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


