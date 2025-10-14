# TournamentSimple

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnticheatRequired** | Pointer to **bool** |  | [optional] 
**Custom** | Pointer to **bool** |  | [optional] 
**FaceitUrl** | Pointer to **string** |  | [optional] 
**FeaturedImage** | Pointer to **string** |  | [optional] 
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
**OrganizerId** | Pointer to **string** |  | [optional] 
**PrizeType** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SubscriptionsCount** | Pointer to **int64** |  | [optional] 
**TeamSize** | Pointer to **int64** |  | [optional] 
**TotalPrize** | Pointer to **map[string]interface{}** |  | [optional] 
**TournamentId** | Pointer to **string** |  | [optional] 
**WhitelistCountries** | Pointer to **[]string** |  | [optional] 

## Methods

### NewTournamentSimple

`func NewTournamentSimple() *TournamentSimple`

NewTournamentSimple instantiates a new TournamentSimple object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTournamentSimpleWithDefaults

`func NewTournamentSimpleWithDefaults() *TournamentSimple`

NewTournamentSimpleWithDefaults instantiates a new TournamentSimple object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnticheatRequired

`func (o *TournamentSimple) GetAnticheatRequired() bool`

GetAnticheatRequired returns the AnticheatRequired field if non-nil, zero value otherwise.

### GetAnticheatRequiredOk

`func (o *TournamentSimple) GetAnticheatRequiredOk() (*bool, bool)`

GetAnticheatRequiredOk returns a tuple with the AnticheatRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnticheatRequired

`func (o *TournamentSimple) SetAnticheatRequired(v bool)`

SetAnticheatRequired sets AnticheatRequired field to given value.

### HasAnticheatRequired

`func (o *TournamentSimple) HasAnticheatRequired() bool`

HasAnticheatRequired returns a boolean if a field has been set.

### GetCustom

`func (o *TournamentSimple) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *TournamentSimple) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *TournamentSimple) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *TournamentSimple) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetFaceitUrl

`func (o *TournamentSimple) GetFaceitUrl() string`

GetFaceitUrl returns the FaceitUrl field if non-nil, zero value otherwise.

### GetFaceitUrlOk

`func (o *TournamentSimple) GetFaceitUrlOk() (*string, bool)`

GetFaceitUrlOk returns a tuple with the FaceitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceitUrl

`func (o *TournamentSimple) SetFaceitUrl(v string)`

SetFaceitUrl sets FaceitUrl field to given value.

### HasFaceitUrl

`func (o *TournamentSimple) HasFaceitUrl() bool`

HasFaceitUrl returns a boolean if a field has been set.

### GetFeaturedImage

`func (o *TournamentSimple) GetFeaturedImage() string`

GetFeaturedImage returns the FeaturedImage field if non-nil, zero value otherwise.

### GetFeaturedImageOk

`func (o *TournamentSimple) GetFeaturedImageOk() (*string, bool)`

GetFeaturedImageOk returns a tuple with the FeaturedImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeaturedImage

`func (o *TournamentSimple) SetFeaturedImage(v string)`

SetFeaturedImage sets FeaturedImage field to given value.

### HasFeaturedImage

`func (o *TournamentSimple) HasFeaturedImage() bool`

HasFeaturedImage returns a boolean if a field has been set.

### GetGameId

`func (o *TournamentSimple) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *TournamentSimple) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *TournamentSimple) SetGameId(v string)`

SetGameId sets GameId field to given value.

### HasGameId

`func (o *TournamentSimple) HasGameId() bool`

HasGameId returns a boolean if a field has been set.

### GetInviteType

`func (o *TournamentSimple) GetInviteType() string`

GetInviteType returns the InviteType field if non-nil, zero value otherwise.

### GetInviteTypeOk

`func (o *TournamentSimple) GetInviteTypeOk() (*string, bool)`

GetInviteTypeOk returns a tuple with the InviteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteType

`func (o *TournamentSimple) SetInviteType(v string)`

SetInviteType sets InviteType field to given value.

### HasInviteType

`func (o *TournamentSimple) HasInviteType() bool`

HasInviteType returns a boolean if a field has been set.

### GetMatchType

`func (o *TournamentSimple) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *TournamentSimple) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *TournamentSimple) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.

### HasMatchType

`func (o *TournamentSimple) HasMatchType() bool`

HasMatchType returns a boolean if a field has been set.

### GetMaxSkill

`func (o *TournamentSimple) GetMaxSkill() int64`

GetMaxSkill returns the MaxSkill field if non-nil, zero value otherwise.

### GetMaxSkillOk

`func (o *TournamentSimple) GetMaxSkillOk() (*int64, bool)`

GetMaxSkillOk returns a tuple with the MaxSkill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSkill

`func (o *TournamentSimple) SetMaxSkill(v int64)`

SetMaxSkill sets MaxSkill field to given value.

### HasMaxSkill

`func (o *TournamentSimple) HasMaxSkill() bool`

HasMaxSkill returns a boolean if a field has been set.

### GetMembershipType

`func (o *TournamentSimple) GetMembershipType() string`

GetMembershipType returns the MembershipType field if non-nil, zero value otherwise.

### GetMembershipTypeOk

`func (o *TournamentSimple) GetMembershipTypeOk() (*string, bool)`

GetMembershipTypeOk returns a tuple with the MembershipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipType

`func (o *TournamentSimple) SetMembershipType(v string)`

SetMembershipType sets MembershipType field to given value.

### HasMembershipType

`func (o *TournamentSimple) HasMembershipType() bool`

HasMembershipType returns a boolean if a field has been set.

### GetMinSkill

`func (o *TournamentSimple) GetMinSkill() int64`

GetMinSkill returns the MinSkill field if non-nil, zero value otherwise.

### GetMinSkillOk

`func (o *TournamentSimple) GetMinSkillOk() (*int64, bool)`

GetMinSkillOk returns a tuple with the MinSkill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSkill

`func (o *TournamentSimple) SetMinSkill(v int64)`

SetMinSkill sets MinSkill field to given value.

### HasMinSkill

`func (o *TournamentSimple) HasMinSkill() bool`

HasMinSkill returns a boolean if a field has been set.

### GetName

`func (o *TournamentSimple) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TournamentSimple) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TournamentSimple) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TournamentSimple) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNumberOfPlayers

`func (o *TournamentSimple) GetNumberOfPlayers() int64`

GetNumberOfPlayers returns the NumberOfPlayers field if non-nil, zero value otherwise.

### GetNumberOfPlayersOk

`func (o *TournamentSimple) GetNumberOfPlayersOk() (*int64, bool)`

GetNumberOfPlayersOk returns a tuple with the NumberOfPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPlayers

`func (o *TournamentSimple) SetNumberOfPlayers(v int64)`

SetNumberOfPlayers sets NumberOfPlayers field to given value.

### HasNumberOfPlayers

`func (o *TournamentSimple) HasNumberOfPlayers() bool`

HasNumberOfPlayers returns a boolean if a field has been set.

### GetNumberOfPlayersCheckedin

`func (o *TournamentSimple) GetNumberOfPlayersCheckedin() int64`

GetNumberOfPlayersCheckedin returns the NumberOfPlayersCheckedin field if non-nil, zero value otherwise.

### GetNumberOfPlayersCheckedinOk

`func (o *TournamentSimple) GetNumberOfPlayersCheckedinOk() (*int64, bool)`

GetNumberOfPlayersCheckedinOk returns a tuple with the NumberOfPlayersCheckedin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPlayersCheckedin

`func (o *TournamentSimple) SetNumberOfPlayersCheckedin(v int64)`

SetNumberOfPlayersCheckedin sets NumberOfPlayersCheckedin field to given value.

### HasNumberOfPlayersCheckedin

`func (o *TournamentSimple) HasNumberOfPlayersCheckedin() bool`

HasNumberOfPlayersCheckedin returns a boolean if a field has been set.

### GetNumberOfPlayersJoined

`func (o *TournamentSimple) GetNumberOfPlayersJoined() int64`

GetNumberOfPlayersJoined returns the NumberOfPlayersJoined field if non-nil, zero value otherwise.

### GetNumberOfPlayersJoinedOk

`func (o *TournamentSimple) GetNumberOfPlayersJoinedOk() (*int64, bool)`

GetNumberOfPlayersJoinedOk returns a tuple with the NumberOfPlayersJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPlayersJoined

`func (o *TournamentSimple) SetNumberOfPlayersJoined(v int64)`

SetNumberOfPlayersJoined sets NumberOfPlayersJoined field to given value.

### HasNumberOfPlayersJoined

`func (o *TournamentSimple) HasNumberOfPlayersJoined() bool`

HasNumberOfPlayersJoined returns a boolean if a field has been set.

### GetNumberOfPlayersParticipants

`func (o *TournamentSimple) GetNumberOfPlayersParticipants() int64`

GetNumberOfPlayersParticipants returns the NumberOfPlayersParticipants field if non-nil, zero value otherwise.

### GetNumberOfPlayersParticipantsOk

`func (o *TournamentSimple) GetNumberOfPlayersParticipantsOk() (*int64, bool)`

GetNumberOfPlayersParticipantsOk returns a tuple with the NumberOfPlayersParticipants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPlayersParticipants

`func (o *TournamentSimple) SetNumberOfPlayersParticipants(v int64)`

SetNumberOfPlayersParticipants sets NumberOfPlayersParticipants field to given value.

### HasNumberOfPlayersParticipants

`func (o *TournamentSimple) HasNumberOfPlayersParticipants() bool`

HasNumberOfPlayersParticipants returns a boolean if a field has been set.

### GetOrganizerId

`func (o *TournamentSimple) GetOrganizerId() string`

GetOrganizerId returns the OrganizerId field if non-nil, zero value otherwise.

### GetOrganizerIdOk

`func (o *TournamentSimple) GetOrganizerIdOk() (*string, bool)`

GetOrganizerIdOk returns a tuple with the OrganizerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizerId

`func (o *TournamentSimple) SetOrganizerId(v string)`

SetOrganizerId sets OrganizerId field to given value.

### HasOrganizerId

`func (o *TournamentSimple) HasOrganizerId() bool`

HasOrganizerId returns a boolean if a field has been set.

### GetPrizeType

`func (o *TournamentSimple) GetPrizeType() string`

GetPrizeType returns the PrizeType field if non-nil, zero value otherwise.

### GetPrizeTypeOk

`func (o *TournamentSimple) GetPrizeTypeOk() (*string, bool)`

GetPrizeTypeOk returns a tuple with the PrizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrizeType

`func (o *TournamentSimple) SetPrizeType(v string)`

SetPrizeType sets PrizeType field to given value.

### HasPrizeType

`func (o *TournamentSimple) HasPrizeType() bool`

HasPrizeType returns a boolean if a field has been set.

### GetRegion

`func (o *TournamentSimple) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TournamentSimple) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TournamentSimple) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *TournamentSimple) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStartedAt

`func (o *TournamentSimple) GetStartedAt() int64`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *TournamentSimple) GetStartedAtOk() (*int64, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *TournamentSimple) SetStartedAt(v int64)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *TournamentSimple) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *TournamentSimple) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TournamentSimple) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TournamentSimple) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TournamentSimple) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubscriptionsCount

`func (o *TournamentSimple) GetSubscriptionsCount() int64`

GetSubscriptionsCount returns the SubscriptionsCount field if non-nil, zero value otherwise.

### GetSubscriptionsCountOk

`func (o *TournamentSimple) GetSubscriptionsCountOk() (*int64, bool)`

GetSubscriptionsCountOk returns a tuple with the SubscriptionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionsCount

`func (o *TournamentSimple) SetSubscriptionsCount(v int64)`

SetSubscriptionsCount sets SubscriptionsCount field to given value.

### HasSubscriptionsCount

`func (o *TournamentSimple) HasSubscriptionsCount() bool`

HasSubscriptionsCount returns a boolean if a field has been set.

### GetTeamSize

`func (o *TournamentSimple) GetTeamSize() int64`

GetTeamSize returns the TeamSize field if non-nil, zero value otherwise.

### GetTeamSizeOk

`func (o *TournamentSimple) GetTeamSizeOk() (*int64, bool)`

GetTeamSizeOk returns a tuple with the TeamSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamSize

`func (o *TournamentSimple) SetTeamSize(v int64)`

SetTeamSize sets TeamSize field to given value.

### HasTeamSize

`func (o *TournamentSimple) HasTeamSize() bool`

HasTeamSize returns a boolean if a field has been set.

### GetTotalPrize

`func (o *TournamentSimple) GetTotalPrize() map[string]interface{}`

GetTotalPrize returns the TotalPrize field if non-nil, zero value otherwise.

### GetTotalPrizeOk

`func (o *TournamentSimple) GetTotalPrizeOk() (*map[string]interface{}, bool)`

GetTotalPrizeOk returns a tuple with the TotalPrize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrize

`func (o *TournamentSimple) SetTotalPrize(v map[string]interface{})`

SetTotalPrize sets TotalPrize field to given value.

### HasTotalPrize

`func (o *TournamentSimple) HasTotalPrize() bool`

HasTotalPrize returns a boolean if a field has been set.

### GetTournamentId

`func (o *TournamentSimple) GetTournamentId() string`

GetTournamentId returns the TournamentId field if non-nil, zero value otherwise.

### GetTournamentIdOk

`func (o *TournamentSimple) GetTournamentIdOk() (*string, bool)`

GetTournamentIdOk returns a tuple with the TournamentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTournamentId

`func (o *TournamentSimple) SetTournamentId(v string)`

SetTournamentId sets TournamentId field to given value.

### HasTournamentId

`func (o *TournamentSimple) HasTournamentId() bool`

HasTournamentId returns a boolean if a field has been set.

### GetWhitelistCountries

`func (o *TournamentSimple) GetWhitelistCountries() []string`

GetWhitelistCountries returns the WhitelistCountries field if non-nil, zero value otherwise.

### GetWhitelistCountriesOk

`func (o *TournamentSimple) GetWhitelistCountriesOk() (*[]string, bool)`

GetWhitelistCountriesOk returns a tuple with the WhitelistCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistCountries

`func (o *TournamentSimple) SetWhitelistCountries(v []string)`

SetWhitelistCountries sets WhitelistCountries field to given value.

### HasWhitelistCountries

`func (o *TournamentSimple) HasWhitelistCountries() bool`

HasWhitelistCountries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


