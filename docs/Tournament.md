# Tournament

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnticheatRequired** | Pointer to **bool** |  | [optional] 
**BestOf** | Pointer to **map[string]interface{}** |  | [optional] 
**CalculateElo** | Pointer to **bool** |  | [optional] 
**CompetitionId** | Pointer to **string** | DEPRECATED: use tournament_id instead | [optional] 
**CoverImage** | Pointer to **string** |  | [optional] 
**Custom** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FaceitUrl** | Pointer to **string** |  | [optional] 
**FeaturedImage** | Pointer to **string** |  | [optional] 
**GameData** | Pointer to [**Game**](Game.md) |  | [optional] 
**GameId** | Pointer to **string** |  | [optional] 
**InviteType** | Pointer to **string** |  | [optional] 
**MatchType** | Pointer to **string** |  | [optional] 
**MaxSkill** | Pointer to **int64** |  | [optional] 
**MembershipType** | Pointer to **string** |  | [optional] 
**MinSkill** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NumberOfPlayers** | Pointer to **int64** |  | [optional] 
**NumberOfPlayersCheckedin** | Pointer to **int64** |  | [optional] 
**NumberOfPlayersJoined** | Pointer to **int64** |  | [optional] 
**NumberOfPlayersParticipants** | Pointer to **int64** |  | [optional] 
**OrganizerData** | Pointer to [**Organizer**](Organizer.md) |  | [optional] 
**OrganizerId** | Pointer to **string** |  | [optional] 
**PrizeType** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Rounds** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Rule** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SubstitutesAllowed** | Pointer to **int64** |  | [optional] 
**SubstitutionsAllowed** | Pointer to **int64** |  | [optional] 
**TeamSize** | Pointer to **int64** |  | [optional] 
**TotalPrize** | Pointer to **map[string]interface{}** |  | [optional] 
**TournamentId** | Pointer to **string** |  | [optional] 
**Voting** | Pointer to **map[string]interface{}** |  | [optional] 
**WhitelistCountries** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTournament

`func NewTournament() *Tournament`

NewTournament instantiates a new Tournament object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTournamentWithDefaults

`func NewTournamentWithDefaults() *Tournament`

NewTournamentWithDefaults instantiates a new Tournament object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnticheatRequired

`func (o *Tournament) GetAnticheatRequired() bool`

GetAnticheatRequired returns the AnticheatRequired field if non-nil, zero value otherwise.

### GetAnticheatRequiredOk

`func (o *Tournament) GetAnticheatRequiredOk() (*bool, bool)`

GetAnticheatRequiredOk returns a tuple with the AnticheatRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnticheatRequired

`func (o *Tournament) SetAnticheatRequired(v bool)`

SetAnticheatRequired sets AnticheatRequired field to given value.

### HasAnticheatRequired

`func (o *Tournament) HasAnticheatRequired() bool`

HasAnticheatRequired returns a boolean if a field has been set.

### GetBestOf

`func (o *Tournament) GetBestOf() map[string]interface{}`

GetBestOf returns the BestOf field if non-nil, zero value otherwise.

### GetBestOfOk

`func (o *Tournament) GetBestOfOk() (*map[string]interface{}, bool)`

GetBestOfOk returns a tuple with the BestOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestOf

`func (o *Tournament) SetBestOf(v map[string]interface{})`

SetBestOf sets BestOf field to given value.

### HasBestOf

`func (o *Tournament) HasBestOf() bool`

HasBestOf returns a boolean if a field has been set.

### GetCalculateElo

`func (o *Tournament) GetCalculateElo() bool`

GetCalculateElo returns the CalculateElo field if non-nil, zero value otherwise.

### GetCalculateEloOk

`func (o *Tournament) GetCalculateEloOk() (*bool, bool)`

GetCalculateEloOk returns a tuple with the CalculateElo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalculateElo

`func (o *Tournament) SetCalculateElo(v bool)`

SetCalculateElo sets CalculateElo field to given value.

### HasCalculateElo

`func (o *Tournament) HasCalculateElo() bool`

HasCalculateElo returns a boolean if a field has been set.

### GetCompetitionId

`func (o *Tournament) GetCompetitionId() string`

GetCompetitionId returns the CompetitionId field if non-nil, zero value otherwise.

### GetCompetitionIdOk

`func (o *Tournament) GetCompetitionIdOk() (*string, bool)`

GetCompetitionIdOk returns a tuple with the CompetitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompetitionId

`func (o *Tournament) SetCompetitionId(v string)`

SetCompetitionId sets CompetitionId field to given value.

### HasCompetitionId

`func (o *Tournament) HasCompetitionId() bool`

HasCompetitionId returns a boolean if a field has been set.

### GetCoverImage

`func (o *Tournament) GetCoverImage() string`

GetCoverImage returns the CoverImage field if non-nil, zero value otherwise.

### GetCoverImageOk

`func (o *Tournament) GetCoverImageOk() (*string, bool)`

GetCoverImageOk returns a tuple with the CoverImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImage

`func (o *Tournament) SetCoverImage(v string)`

SetCoverImage sets CoverImage field to given value.

### HasCoverImage

`func (o *Tournament) HasCoverImage() bool`

HasCoverImage returns a boolean if a field has been set.

### GetCustom

`func (o *Tournament) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *Tournament) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *Tournament) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *Tournament) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetDescription

`func (o *Tournament) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Tournament) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Tournament) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Tournament) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFaceitUrl

`func (o *Tournament) GetFaceitUrl() string`

GetFaceitUrl returns the FaceitUrl field if non-nil, zero value otherwise.

### GetFaceitUrlOk

`func (o *Tournament) GetFaceitUrlOk() (*string, bool)`

GetFaceitUrlOk returns a tuple with the FaceitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceitUrl

`func (o *Tournament) SetFaceitUrl(v string)`

SetFaceitUrl sets FaceitUrl field to given value.

### HasFaceitUrl

`func (o *Tournament) HasFaceitUrl() bool`

HasFaceitUrl returns a boolean if a field has been set.

### GetFeaturedImage

`func (o *Tournament) GetFeaturedImage() string`

GetFeaturedImage returns the FeaturedImage field if non-nil, zero value otherwise.

### GetFeaturedImageOk

`func (o *Tournament) GetFeaturedImageOk() (*string, bool)`

GetFeaturedImageOk returns a tuple with the FeaturedImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedImage

`func (o *Tournament) SetFeaturedImage(v string)`

SetFeaturedImage sets FeaturedImage field to given value.

### HasFeaturedImage

`func (o *Tournament) HasFeaturedImage() bool`

HasFeaturedImage returns a boolean if a field has been set.

### GetGameData

`func (o *Tournament) GetGameData() Game`

GetGameData returns the GameData field if non-nil, zero value otherwise.

### GetGameDataOk

`func (o *Tournament) GetGameDataOk() (*Game, bool)`

GetGameDataOk returns a tuple with the GameData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameData

`func (o *Tournament) SetGameData(v Game)`

SetGameData sets GameData field to given value.

### HasGameData

`func (o *Tournament) HasGameData() bool`

HasGameData returns a boolean if a field has been set.

### GetGameId

`func (o *Tournament) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *Tournament) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *Tournament) SetGameId(v string)`

SetGameId sets GameId field to given value.

### HasGameId

`func (o *Tournament) HasGameId() bool`

HasGameId returns a boolean if a field has been set.

### GetInviteType

`func (o *Tournament) GetInviteType() string`

GetInviteType returns the InviteType field if non-nil, zero value otherwise.

### GetInviteTypeOk

`func (o *Tournament) GetInviteTypeOk() (*string, bool)`

GetInviteTypeOk returns a tuple with the InviteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteType

`func (o *Tournament) SetInviteType(v string)`

SetInviteType sets InviteType field to given value.

### HasInviteType

`func (o *Tournament) HasInviteType() bool`

HasInviteType returns a boolean if a field has been set.

### GetMatchType

`func (o *Tournament) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *Tournament) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *Tournament) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.

### HasMatchType

`func (o *Tournament) HasMatchType() bool`

HasMatchType returns a boolean if a field has been set.

### GetMaxSkill

`func (o *Tournament) GetMaxSkill() int64`

GetMaxSkill returns the MaxSkill field if non-nil, zero value otherwise.

### GetMaxSkillOk

`func (o *Tournament) GetMaxSkillOk() (*int64, bool)`

GetMaxSkillOk returns a tuple with the MaxSkill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSkill

`func (o *Tournament) SetMaxSkill(v int64)`

SetMaxSkill sets MaxSkill field to given value.

### HasMaxSkill

`func (o *Tournament) HasMaxSkill() bool`

HasMaxSkill returns a boolean if a field has been set.

### GetMembershipType

`func (o *Tournament) GetMembershipType() string`

GetMembershipType returns the MembershipType field if non-nil, zero value otherwise.

### GetMembershipTypeOk

`func (o *Tournament) GetMembershipTypeOk() (*string, bool)`

GetMembershipTypeOk returns a tuple with the MembershipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipType

`func (o *Tournament) SetMembershipType(v string)`

SetMembershipType sets MembershipType field to given value.

### HasMembershipType

`func (o *Tournament) HasMembershipType() bool`

HasMembershipType returns a boolean if a field has been set.

### GetMinSkill

`func (o *Tournament) GetMinSkill() int64`

GetMinSkill returns the MinSkill field if non-nil, zero value otherwise.

### GetMinSkillOk

`func (o *Tournament) GetMinSkillOk() (*int64, bool)`

GetMinSkillOk returns a tuple with the MinSkill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSkill

`func (o *Tournament) SetMinSkill(v int64)`

SetMinSkill sets MinSkill field to given value.

### HasMinSkill

`func (o *Tournament) HasMinSkill() bool`

HasMinSkill returns a boolean if a field has been set.

### GetName

`func (o *Tournament) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tournament) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tournament) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Tournament) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberOfPlayers

`func (o *Tournament) GetNumberOfPlayers() int64`

GetNumberOfPlayers returns the NumberOfPlayers field if non-nil, zero value otherwise.

### GetNumberOfPlayersOk

`func (o *Tournament) GetNumberOfPlayersOk() (*int64, bool)`

GetNumberOfPlayersOk returns a tuple with the NumberOfPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPlayers

`func (o *Tournament) SetNumberOfPlayers(v int64)`

SetNumberOfPlayers sets NumberOfPlayers field to given value.

### HasNumberOfPlayers

`func (o *Tournament) HasNumberOfPlayers() bool`

HasNumberOfPlayers returns a boolean if a field has been set.

### GetNumberOfPlayersCheckedin

`func (o *Tournament) GetNumberOfPlayersCheckedin() int64`

GetNumberOfPlayersCheckedin returns the NumberOfPlayersCheckedin field if non-nil, zero value otherwise.

### GetNumberOfPlayersCheckedinOk

`func (o *Tournament) GetNumberOfPlayersCheckedinOk() (*int64, bool)`

GetNumberOfPlayersCheckedinOk returns a tuple with the NumberOfPlayersCheckedin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPlayersCheckedin

`func (o *Tournament) SetNumberOfPlayersCheckedin(v int64)`

SetNumberOfPlayersCheckedin sets NumberOfPlayersCheckedin field to given value.

### HasNumberOfPlayersCheckedin

`func (o *Tournament) HasNumberOfPlayersCheckedin() bool`

HasNumberOfPlayersCheckedin returns a boolean if a field has been set.

### GetNumberOfPlayersJoined

`func (o *Tournament) GetNumberOfPlayersJoined() int64`

GetNumberOfPlayersJoined returns the NumberOfPlayersJoined field if non-nil, zero value otherwise.

### GetNumberOfPlayersJoinedOk

`func (o *Tournament) GetNumberOfPlayersJoinedOk() (*int64, bool)`

GetNumberOfPlayersJoinedOk returns a tuple with the NumberOfPlayersJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPlayersJoined

`func (o *Tournament) SetNumberOfPlayersJoined(v int64)`

SetNumberOfPlayersJoined sets NumberOfPlayersJoined field to given value.

### HasNumberOfPlayersJoined

`func (o *Tournament) HasNumberOfPlayersJoined() bool`

HasNumberOfPlayersJoined returns a boolean if a field has been set.

### GetNumberOfPlayersParticipants

`func (o *Tournament) GetNumberOfPlayersParticipants() int64`

GetNumberOfPlayersParticipants returns the NumberOfPlayersParticipants field if non-nil, zero value otherwise.

### GetNumberOfPlayersParticipantsOk

`func (o *Tournament) GetNumberOfPlayersParticipantsOk() (*int64, bool)`

GetNumberOfPlayersParticipantsOk returns a tuple with the NumberOfPlayersParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPlayersParticipants

`func (o *Tournament) SetNumberOfPlayersParticipants(v int64)`

SetNumberOfPlayersParticipants sets NumberOfPlayersParticipants field to given value.

### HasNumberOfPlayersParticipants

`func (o *Tournament) HasNumberOfPlayersParticipants() bool`

HasNumberOfPlayersParticipants returns a boolean if a field has been set.

### GetOrganizerData

`func (o *Tournament) GetOrganizerData() Organizer`

GetOrganizerData returns the OrganizerData field if non-nil, zero value otherwise.

### GetOrganizerDataOk

`func (o *Tournament) GetOrganizerDataOk() (*Organizer, bool)`

GetOrganizerDataOk returns a tuple with the OrganizerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizerData

`func (o *Tournament) SetOrganizerData(v Organizer)`

SetOrganizerData sets OrganizerData field to given value.

### HasOrganizerData

`func (o *Tournament) HasOrganizerData() bool`

HasOrganizerData returns a boolean if a field has been set.

### GetOrganizerId

`func (o *Tournament) GetOrganizerId() string`

GetOrganizerId returns the OrganizerId field if non-nil, zero value otherwise.

### GetOrganizerIdOk

`func (o *Tournament) GetOrganizerIdOk() (*string, bool)`

GetOrganizerIdOk returns a tuple with the OrganizerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizerId

`func (o *Tournament) SetOrganizerId(v string)`

SetOrganizerId sets OrganizerId field to given value.

### HasOrganizerId

`func (o *Tournament) HasOrganizerId() bool`

HasOrganizerId returns a boolean if a field has been set.

### GetPrizeType

`func (o *Tournament) GetPrizeType() string`

GetPrizeType returns the PrizeType field if non-nil, zero value otherwise.

### GetPrizeTypeOk

`func (o *Tournament) GetPrizeTypeOk() (*string, bool)`

GetPrizeTypeOk returns a tuple with the PrizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrizeType

`func (o *Tournament) SetPrizeType(v string)`

SetPrizeType sets PrizeType field to given value.

### HasPrizeType

`func (o *Tournament) HasPrizeType() bool`

HasPrizeType returns a boolean if a field has been set.

### GetRegion

`func (o *Tournament) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Tournament) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Tournament) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Tournament) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRounds

`func (o *Tournament) GetRounds() []map[string]interface{}`

GetRounds returns the Rounds field if non-nil, zero value otherwise.

### GetRoundsOk

`func (o *Tournament) GetRoundsOk() (*[]map[string]interface{}, bool)`

GetRoundsOk returns a tuple with the Rounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRounds

`func (o *Tournament) SetRounds(v []map[string]interface{})`

SetRounds sets Rounds field to given value.

### HasRounds

`func (o *Tournament) HasRounds() bool`

HasRounds returns a boolean if a field has been set.

### GetRule

`func (o *Tournament) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *Tournament) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *Tournament) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *Tournament) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetStartedAt

`func (o *Tournament) GetStartedAt() int64`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *Tournament) GetStartedAtOk() (*int64, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *Tournament) SetStartedAt(v int64)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *Tournament) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *Tournament) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Tournament) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Tournament) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Tournament) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubstitutesAllowed

`func (o *Tournament) GetSubstitutesAllowed() int64`

GetSubstitutesAllowed returns the SubstitutesAllowed field if non-nil, zero value otherwise.

### GetSubstitutesAllowedOk

`func (o *Tournament) GetSubstitutesAllowedOk() (*int64, bool)`

GetSubstitutesAllowedOk returns a tuple with the SubstitutesAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstitutesAllowed

`func (o *Tournament) SetSubstitutesAllowed(v int64)`

SetSubstitutesAllowed sets SubstitutesAllowed field to given value.

### HasSubstitutesAllowed

`func (o *Tournament) HasSubstitutesAllowed() bool`

HasSubstitutesAllowed returns a boolean if a field has been set.

### GetSubstitutionsAllowed

`func (o *Tournament) GetSubstitutionsAllowed() int64`

GetSubstitutionsAllowed returns the SubstitutionsAllowed field if non-nil, zero value otherwise.

### GetSubstitutionsAllowedOk

`func (o *Tournament) GetSubstitutionsAllowedOk() (*int64, bool)`

GetSubstitutionsAllowedOk returns a tuple with the SubstitutionsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstitutionsAllowed

`func (o *Tournament) SetSubstitutionsAllowed(v int64)`

SetSubstitutionsAllowed sets SubstitutionsAllowed field to given value.

### HasSubstitutionsAllowed

`func (o *Tournament) HasSubstitutionsAllowed() bool`

HasSubstitutionsAllowed returns a boolean if a field has been set.

### GetTeamSize

`func (o *Tournament) GetTeamSize() int64`

GetTeamSize returns the TeamSize field if non-nil, zero value otherwise.

### GetTeamSizeOk

`func (o *Tournament) GetTeamSizeOk() (*int64, bool)`

GetTeamSizeOk returns a tuple with the TeamSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamSize

`func (o *Tournament) SetTeamSize(v int64)`

SetTeamSize sets TeamSize field to given value.

### HasTeamSize

`func (o *Tournament) HasTeamSize() bool`

HasTeamSize returns a boolean if a field has been set.

### GetTotalPrize

`func (o *Tournament) GetTotalPrize() map[string]interface{}`

GetTotalPrize returns the TotalPrize field if non-nil, zero value otherwise.

### GetTotalPrizeOk

`func (o *Tournament) GetTotalPrizeOk() (*map[string]interface{}, bool)`

GetTotalPrizeOk returns a tuple with the TotalPrize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrize

`func (o *Tournament) SetTotalPrize(v map[string]interface{})`

SetTotalPrize sets TotalPrize field to given value.

### HasTotalPrize

`func (o *Tournament) HasTotalPrize() bool`

HasTotalPrize returns a boolean if a field has been set.

### GetTournamentId

`func (o *Tournament) GetTournamentId() string`

GetTournamentId returns the TournamentId field if non-nil, zero value otherwise.

### GetTournamentIdOk

`func (o *Tournament) GetTournamentIdOk() (*string, bool)`

GetTournamentIdOk returns a tuple with the TournamentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTournamentId

`func (o *Tournament) SetTournamentId(v string)`

SetTournamentId sets TournamentId field to given value.

### HasTournamentId

`func (o *Tournament) HasTournamentId() bool`

HasTournamentId returns a boolean if a field has been set.

### GetVoting

`func (o *Tournament) GetVoting() map[string]interface{}`

GetVoting returns the Voting field if non-nil, zero value otherwise.

### GetVotingOk

`func (o *Tournament) GetVotingOk() (*map[string]interface{}, bool)`

GetVotingOk returns a tuple with the Voting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoting

`func (o *Tournament) SetVoting(v map[string]interface{})`

SetVoting sets Voting field to given value.

### HasVoting

`func (o *Tournament) HasVoting() bool`

HasVoting returns a boolean if a field has been set.

### GetWhitelistCountries

`func (o *Tournament) GetWhitelistCountries() []string`

GetWhitelistCountries returns the WhitelistCountries field if non-nil, zero value otherwise.

### GetWhitelistCountriesOk

`func (o *Tournament) GetWhitelistCountriesOk() (*[]string, bool)`

GetWhitelistCountriesOk returns a tuple with the WhitelistCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistCountries

`func (o *Tournament) SetWhitelistCountries(v []string)`

SetWhitelistCountries sets WhitelistCountries field to given value.

### HasWhitelistCountries

`func (o *Tournament) HasWhitelistCountries() bool`

HasWhitelistCountries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


