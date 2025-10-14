# Match

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BestOf** | Pointer to **int64** |  | [optional] 
**BroadcastStartTime** | Pointer to **int64** |  | [optional] 
**BroadcastStartTimeLabel** | Pointer to **string** |  | [optional] 
**CalculateElo** | Pointer to **bool** |  | [optional] 
**ChatRoomId** | Pointer to **string** |  | [optional] 
**CompetitionId** | Pointer to **string** |  | [optional] 
**CompetitionName** | Pointer to **string** |  | [optional] 
**CompetitionType** | Pointer to **string** |  | [optional] 
**ConfiguredAt** | Pointer to **int64** |  | [optional] 
**DemoUrl** | Pointer to **[]string** |  | [optional] 
**DetailedResults** | Pointer to [**[]DetailedMatchResult**](DetailedMatchResult.md) |  | [optional] 
**FaceitUrl** | Pointer to **string** |  | [optional] 
**FinishedAt** | Pointer to **int64** |  | [optional] 
**Game** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **int64** |  | [optional] 
**MatchId** | Pointer to **string** |  | [optional] 
**OrganizerId** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Results** | Pointer to [**MatchResult**](MatchResult.md) |  | [optional] 
**Round** | Pointer to **int64** |  | [optional] 
**ScheduledAt** | Pointer to **int64** |  | [optional] 
**StartedAt** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Teams** | Pointer to [**map[string]Faction**](Faction.md) |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 
**Voting** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewMatch

`func NewMatch() *Match`

NewMatch instantiates a new Match object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchWithDefaults

`func NewMatchWithDefaults() *Match`

NewMatchWithDefaults instantiates a new Match object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBestOf

`func (o *Match) GetBestOf() int64`

GetBestOf returns the BestOf field if non-nil, zero value otherwise.

### GetBestOfOk

`func (o *Match) GetBestOfOk() (*int64, bool)`

GetBestOfOk returns a tuple with the BestOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestOf

`func (o *Match) SetBestOf(v int64)`

SetBestOf sets BestOf field to given value.

### HasBestOf

`func (o *Match) HasBestOf() bool`

HasBestOf returns a boolean if a field has been set.

### GetBroadcastStartTime

`func (o *Match) GetBroadcastStartTime() int64`

GetBroadcastStartTime returns the BroadcastStartTime field if non-nil, zero value otherwise.

### GetBroadcastStartTimeOk

`func (o *Match) GetBroadcastStartTimeOk() (*int64, bool)`

GetBroadcastStartTimeOk returns a tuple with the BroadcastStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastStartTime

`func (o *Match) SetBroadcastStartTime(v int64)`

SetBroadcastStartTime sets BroadcastStartTime field to given value.

### HasBroadcastStartTime

`func (o *Match) HasBroadcastStartTime() bool`

HasBroadcastStartTime returns a boolean if a field has been set.

### GetBroadcastStartTimeLabel

`func (o *Match) GetBroadcastStartTimeLabel() string`

GetBroadcastStartTimeLabel returns the BroadcastStartTimeLabel field if non-nil, zero value otherwise.

### GetBroadcastStartTimeLabelOk

`func (o *Match) GetBroadcastStartTimeLabelOk() (*string, bool)`

GetBroadcastStartTimeLabelOk returns a tuple with the BroadcastStartTimeLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastStartTimeLabel

`func (o *Match) SetBroadcastStartTimeLabel(v string)`

SetBroadcastStartTimeLabel sets BroadcastStartTimeLabel field to given value.

### HasBroadcastStartTimeLabel

`func (o *Match) HasBroadcastStartTimeLabel() bool`

HasBroadcastStartTimeLabel returns a boolean if a field has been set.

### GetCalculateElo

`func (o *Match) GetCalculateElo() bool`

GetCalculateElo returns the CalculateElo field if non-nil, zero value otherwise.

### GetCalculateEloOk

`func (o *Match) GetCalculateEloOk() (*bool, bool)`

GetCalculateEloOk returns a tuple with the CalculateElo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculateElo

`func (o *Match) SetCalculateElo(v bool)`

SetCalculateElo sets CalculateElo field to given value.

### HasCalculateElo

`func (o *Match) HasCalculateElo() bool`

HasCalculateElo returns a boolean if a field has been set.

### GetChatRoomId

`func (o *Match) GetChatRoomId() string`

GetChatRoomId returns the ChatRoomId field if non-nil, zero value otherwise.

### GetChatRoomIdOk

`func (o *Match) GetChatRoomIdOk() (*string, bool)`

GetChatRoomIdOk returns a tuple with the ChatRoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatRoomId

`func (o *Match) SetChatRoomId(v string)`

SetChatRoomId sets ChatRoomId field to given value.

### HasChatRoomId

`func (o *Match) HasChatRoomId() bool`

HasChatRoomId returns a boolean if a field has been set.

### GetCompetitionId

`func (o *Match) GetCompetitionId() string`

GetCompetitionId returns the CompetitionId field if non-nil, zero value otherwise.

### GetCompetitionIdOk

`func (o *Match) GetCompetitionIdOk() (*string, bool)`

GetCompetitionIdOk returns a tuple with the CompetitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitionId

`func (o *Match) SetCompetitionId(v string)`

SetCompetitionId sets CompetitionId field to given value.

### HasCompetitionId

`func (o *Match) HasCompetitionId() bool`

HasCompetitionId returns a boolean if a field has been set.

### GetCompetitionName

`func (o *Match) GetCompetitionName() string`

GetCompetitionName returns the CompetitionName field if non-nil, zero value otherwise.

### GetCompetitionNameOk

`func (o *Match) GetCompetitionNameOk() (*string, bool)`

GetCompetitionNameOk returns a tuple with the CompetitionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitionName

`func (o *Match) SetCompetitionName(v string)`

SetCompetitionName sets CompetitionName field to given value.

### HasCompetitionName

`func (o *Match) HasCompetitionName() bool`

HasCompetitionName returns a boolean if a field has been set.

### GetCompetitionType

`func (o *Match) GetCompetitionType() string`

GetCompetitionType returns the CompetitionType field if non-nil, zero value otherwise.

### GetCompetitionTypeOk

`func (o *Match) GetCompetitionTypeOk() (*string, bool)`

GetCompetitionTypeOk returns a tuple with the CompetitionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitionType

`func (o *Match) SetCompetitionType(v string)`

SetCompetitionType sets CompetitionType field to given value.

### HasCompetitionType

`func (o *Match) HasCompetitionType() bool`

HasCompetitionType returns a boolean if a field has been set.

### GetConfiguredAt

`func (o *Match) GetConfiguredAt() int64`

GetConfiguredAt returns the ConfiguredAt field if non-nil, zero value otherwise.

### GetConfiguredAtOk

`func (o *Match) GetConfiguredAtOk() (*int64, bool)`

GetConfiguredAtOk returns a tuple with the ConfiguredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredAt

`func (o *Match) SetConfiguredAt(v int64)`

SetConfiguredAt sets ConfiguredAt field to given value.

### HasConfiguredAt

`func (o *Match) HasConfiguredAt() bool`

HasConfiguredAt returns a boolean if a field has been set.

### GetDemoUrl

`func (o *Match) GetDemoUrl() []string`

GetDemoUrl returns the DemoUrl field if non-nil, zero value otherwise.

### GetDemoUrlOk

`func (o *Match) GetDemoUrlOk() (*[]string, bool)`

GetDemoUrlOk returns a tuple with the DemoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemoUrl

`func (o *Match) SetDemoUrl(v []string)`

SetDemoUrl sets DemoUrl field to given value.

### HasDemoUrl

`func (o *Match) HasDemoUrl() bool`

HasDemoUrl returns a boolean if a field has been set.

### GetDetailedResults

`func (o *Match) GetDetailedResults() []DetailedMatchResult`

GetDetailedResults returns the DetailedResults field if non-nil, zero value otherwise.

### GetDetailedResultsOk

`func (o *Match) GetDetailedResultsOk() (*[]DetailedMatchResult, bool)`

GetDetailedResultsOk returns a tuple with the DetailedResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedResults

`func (o *Match) SetDetailedResults(v []DetailedMatchResult)`

SetDetailedResults sets DetailedResults field to given value.

### HasDetailedResults

`func (o *Match) HasDetailedResults() bool`

HasDetailedResults returns a boolean if a field has been set.

### GetFaceitUrl

`func (o *Match) GetFaceitUrl() string`

GetFaceitUrl returns the FaceitUrl field if non-nil, zero value otherwise.

### GetFaceitUrlOk

`func (o *Match) GetFaceitUrlOk() (*string, bool)`

GetFaceitUrlOk returns a tuple with the FaceitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceitUrl

`func (o *Match) SetFaceitUrl(v string)`

SetFaceitUrl sets FaceitUrl field to given value.

### HasFaceitUrl

`func (o *Match) HasFaceitUrl() bool`

HasFaceitUrl returns a boolean if a field has been set.

### GetFinishedAt

`func (o *Match) GetFinishedAt() int64`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *Match) GetFinishedAtOk() (*int64, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *Match) SetFinishedAt(v int64)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *Match) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetGame

`func (o *Match) GetGame() string`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *Match) GetGameOk() (*string, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *Match) SetGame(v string)`

SetGame sets Game field to given value.

### HasGame

`func (o *Match) HasGame() bool`

HasGame returns a boolean if a field has been set.

### GetGroup

`func (o *Match) GetGroup() int64`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Match) GetGroupOk() (*int64, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Match) SetGroup(v int64)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Match) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetMatchId

`func (o *Match) GetMatchId() string`

GetMatchId returns the MatchId field if non-nil, zero value otherwise.

### GetMatchIdOk

`func (o *Match) GetMatchIdOk() (*string, bool)`

GetMatchIdOk returns a tuple with the MatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchId

`func (o *Match) SetMatchId(v string)`

SetMatchId sets MatchId field to given value.

### HasMatchId

`func (o *Match) HasMatchId() bool`

HasMatchId returns a boolean if a field has been set.

### GetOrganizerId

`func (o *Match) GetOrganizerId() string`

GetOrganizerId returns the OrganizerId field if non-nil, zero value otherwise.

### GetOrganizerIdOk

`func (o *Match) GetOrganizerIdOk() (*string, bool)`

GetOrganizerIdOk returns a tuple with the OrganizerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizerId

`func (o *Match) SetOrganizerId(v string)`

SetOrganizerId sets OrganizerId field to given value.

### HasOrganizerId

`func (o *Match) HasOrganizerId() bool`

HasOrganizerId returns a boolean if a field has been set.

### GetRegion

`func (o *Match) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Match) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Match) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Match) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetResults

`func (o *Match) GetResults() MatchResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *Match) GetResultsOk() (*MatchResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *Match) SetResults(v MatchResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *Match) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetRound

`func (o *Match) GetRound() int64`

GetRound returns the Round field if non-nil, zero value otherwise.

### GetRoundOk

`func (o *Match) GetRoundOk() (*int64, bool)`

GetRoundOk returns a tuple with the Round field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRound

`func (o *Match) SetRound(v int64)`

SetRound sets Round field to given value.

### HasRound

`func (o *Match) HasRound() bool`

HasRound returns a boolean if a field has been set.

### GetScheduledAt

`func (o *Match) GetScheduledAt() int64`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *Match) GetScheduledAtOk() (*int64, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *Match) SetScheduledAt(v int64)`

SetScheduledAt sets ScheduledAt field to given value.

### HasScheduledAt

`func (o *Match) HasScheduledAt() bool`

HasScheduledAt returns a boolean if a field has been set.

### GetStartedAt

`func (o *Match) GetStartedAt() int64`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Match) GetStartedAtOk() (*int64, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Match) SetStartedAt(v int64)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Match) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *Match) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Match) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Match) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Match) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTeams

`func (o *Match) GetTeams() map[string]Faction`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *Match) GetTeamsOk() (*map[string]Faction, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *Match) SetTeams(v map[string]Faction)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *Match) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetVersion

`func (o *Match) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Match) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Match) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Match) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVoting

`func (o *Match) GetVoting() map[string]interface{}`

GetVoting returns the Voting field if non-nil, zero value otherwise.

### GetVotingOk

`func (o *Match) GetVotingOk() (*map[string]interface{}, bool)`

GetVotingOk returns a tuple with the Voting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoting

`func (o *Match) SetVoting(v map[string]interface{})`

SetVoting sets Voting field to given value.

### HasVoting

`func (o *Match) HasVoting() bool`

HasVoting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


