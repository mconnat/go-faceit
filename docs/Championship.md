# Championship

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnticheatRequired** | Pointer to **bool** |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**BackgroundImage** | Pointer to **string** |  | [optional] 
**ChampionshipId** | Pointer to **string** |  | [optional] 
**ChampionshipStart** | Pointer to **int64** |  | [optional] 
**CheckinClear** | Pointer to **int64** |  | [optional] 
**CheckinEnabled** | Pointer to **bool** |  | [optional] 
**CheckinStart** | Pointer to **int64** |  | [optional] 
**CoverImage** | Pointer to **string** |  | [optional] 
**CurrentSubscriptions** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FaceitUrl** | Pointer to **string** |  | [optional] 
**Featured** | Pointer to **bool** |  | [optional] 
**Full** | Pointer to **bool** |  | [optional] 
**GameData** | Pointer to [**Game**](Game.md) |  | [optional] 
**GameId** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Deprecated: use championship_id instead | [optional] 
**JoinChecks** | Pointer to [**JoinCheck**](JoinCheck.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrganizerData** | Pointer to [**Organizer**](Organizer.md) |  | [optional] 
**OrganizerId** | Pointer to **string** |  | [optional] 
**Prizes** | Pointer to [**[]Prize**](Prize.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**RulesId** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to [**map[string]ChampionshipSchedule**](ChampionshipSchedule.md) |  | [optional] 
**Screening** | Pointer to [**ChampionshipScreening**](ChampionshipScreening.md) |  | [optional] 
**SeedingStrategy** | Pointer to **string** |  | [optional] 
**Slots** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Stream** | Pointer to [**ChampionshipStream**](ChampionshipStream.md) |  | [optional] 
**SubscriptionEnd** | Pointer to **int64** |  | [optional] 
**SubscriptionStart** | Pointer to **int64** |  | [optional] 
**SubscriptionsLocked** | Pointer to **bool** |  | [optional] 
**SubstitutionConfiguration** | Pointer to [**SubstitutionConfiguration**](SubstitutionConfiguration.md) |  | [optional] 
**TotalGroups** | Pointer to **int64** |  | [optional] 
**TotalPrizes** | Pointer to **int64** |  | [optional] 
**TotalRounds** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewChampionship

`func NewChampionship() *Championship`

NewChampionship instantiates a new Championship object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChampionshipWithDefaults

`func NewChampionshipWithDefaults() *Championship`

NewChampionshipWithDefaults instantiates a new Championship object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnticheatRequired

`func (o *Championship) GetAnticheatRequired() bool`

GetAnticheatRequired returns the AnticheatRequired field if non-nil, zero value otherwise.

### GetAnticheatRequiredOk

`func (o *Championship) GetAnticheatRequiredOk() (*bool, bool)`

GetAnticheatRequiredOk returns a tuple with the AnticheatRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnticheatRequired

`func (o *Championship) SetAnticheatRequired(v bool)`

SetAnticheatRequired sets AnticheatRequired field to given value.

### HasAnticheatRequired

`func (o *Championship) HasAnticheatRequired() bool`

HasAnticheatRequired returns a boolean if a field has been set.

### GetAvatar

`func (o *Championship) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *Championship) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *Championship) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *Championship) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetBackgroundImage

`func (o *Championship) GetBackgroundImage() string`

GetBackgroundImage returns the BackgroundImage field if non-nil, zero value otherwise.

### GetBackgroundImageOk

`func (o *Championship) GetBackgroundImageOk() (*string, bool)`

GetBackgroundImageOk returns a tuple with the BackgroundImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundImage

`func (o *Championship) SetBackgroundImage(v string)`

SetBackgroundImage sets BackgroundImage field to given value.

### HasBackgroundImage

`func (o *Championship) HasBackgroundImage() bool`

HasBackgroundImage returns a boolean if a field has been set.

### GetChampionshipId

`func (o *Championship) GetChampionshipId() string`

GetChampionshipId returns the ChampionshipId field if non-nil, zero value otherwise.

### GetChampionshipIdOk

`func (o *Championship) GetChampionshipIdOk() (*string, bool)`

GetChampionshipIdOk returns a tuple with the ChampionshipId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChampionshipId

`func (o *Championship) SetChampionshipId(v string)`

SetChampionshipId sets ChampionshipId field to given value.

### HasChampionshipId

`func (o *Championship) HasChampionshipId() bool`

HasChampionshipId returns a boolean if a field has been set.

### GetChampionshipStart

`func (o *Championship) GetChampionshipStart() int64`

GetChampionshipStart returns the ChampionshipStart field if non-nil, zero value otherwise.

### GetChampionshipStartOk

`func (o *Championship) GetChampionshipStartOk() (*int64, bool)`

GetChampionshipStartOk returns a tuple with the ChampionshipStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChampionshipStart

`func (o *Championship) SetChampionshipStart(v int64)`

SetChampionshipStart sets ChampionshipStart field to given value.

### HasChampionshipStart

`func (o *Championship) HasChampionshipStart() bool`

HasChampionshipStart returns a boolean if a field has been set.

### GetCheckinClear

`func (o *Championship) GetCheckinClear() int64`

GetCheckinClear returns the CheckinClear field if non-nil, zero value otherwise.

### GetCheckinClearOk

`func (o *Championship) GetCheckinClearOk() (*int64, bool)`

GetCheckinClearOk returns a tuple with the CheckinClear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckinClear

`func (o *Championship) SetCheckinClear(v int64)`

SetCheckinClear sets CheckinClear field to given value.

### HasCheckinClear

`func (o *Championship) HasCheckinClear() bool`

HasCheckinClear returns a boolean if a field has been set.

### GetCheckinEnabled

`func (o *Championship) GetCheckinEnabled() bool`

GetCheckinEnabled returns the CheckinEnabled field if non-nil, zero value otherwise.

### GetCheckinEnabledOk

`func (o *Championship) GetCheckinEnabledOk() (*bool, bool)`

GetCheckinEnabledOk returns a tuple with the CheckinEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckinEnabled

`func (o *Championship) SetCheckinEnabled(v bool)`

SetCheckinEnabled sets CheckinEnabled field to given value.

### HasCheckinEnabled

`func (o *Championship) HasCheckinEnabled() bool`

HasCheckinEnabled returns a boolean if a field has been set.

### GetCheckinStart

`func (o *Championship) GetCheckinStart() int64`

GetCheckinStart returns the CheckinStart field if non-nil, zero value otherwise.

### GetCheckinStartOk

`func (o *Championship) GetCheckinStartOk() (*int64, bool)`

GetCheckinStartOk returns a tuple with the CheckinStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckinStart

`func (o *Championship) SetCheckinStart(v int64)`

SetCheckinStart sets CheckinStart field to given value.

### HasCheckinStart

`func (o *Championship) HasCheckinStart() bool`

HasCheckinStart returns a boolean if a field has been set.

### GetCoverImage

`func (o *Championship) GetCoverImage() string`

GetCoverImage returns the CoverImage field if non-nil, zero value otherwise.

### GetCoverImageOk

`func (o *Championship) GetCoverImageOk() (*string, bool)`

GetCoverImageOk returns a tuple with the CoverImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImage

`func (o *Championship) SetCoverImage(v string)`

SetCoverImage sets CoverImage field to given value.

### HasCoverImage

`func (o *Championship) HasCoverImage() bool`

HasCoverImage returns a boolean if a field has been set.

### GetCurrentSubscriptions

`func (o *Championship) GetCurrentSubscriptions() int64`

GetCurrentSubscriptions returns the CurrentSubscriptions field if non-nil, zero value otherwise.

### GetCurrentSubscriptionsOk

`func (o *Championship) GetCurrentSubscriptionsOk() (*int64, bool)`

GetCurrentSubscriptionsOk returns a tuple with the CurrentSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSubscriptions

`func (o *Championship) SetCurrentSubscriptions(v int64)`

SetCurrentSubscriptions sets CurrentSubscriptions field to given value.

### HasCurrentSubscriptions

`func (o *Championship) HasCurrentSubscriptions() bool`

HasCurrentSubscriptions returns a boolean if a field has been set.

### GetDescription

`func (o *Championship) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Championship) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Championship) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Championship) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFaceitUrl

`func (o *Championship) GetFaceitUrl() string`

GetFaceitUrl returns the FaceitUrl field if non-nil, zero value otherwise.

### GetFaceitUrlOk

`func (o *Championship) GetFaceitUrlOk() (*string, bool)`

GetFaceitUrlOk returns a tuple with the FaceitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceitUrl

`func (o *Championship) SetFaceitUrl(v string)`

SetFaceitUrl sets FaceitUrl field to given value.

### HasFaceitUrl

`func (o *Championship) HasFaceitUrl() bool`

HasFaceitUrl returns a boolean if a field has been set.

### GetFeatured

`func (o *Championship) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *Championship) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *Championship) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *Championship) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetFull

`func (o *Championship) GetFull() bool`

GetFull returns the Full field if non-nil, zero value otherwise.

### GetFullOk

`func (o *Championship) GetFullOk() (*bool, bool)`

GetFullOk returns a tuple with the Full field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFull

`func (o *Championship) SetFull(v bool)`

SetFull sets Full field to given value.

### HasFull

`func (o *Championship) HasFull() bool`

HasFull returns a boolean if a field has been set.

### GetGameData

`func (o *Championship) GetGameData() Game`

GetGameData returns the GameData field if non-nil, zero value otherwise.

### GetGameDataOk

`func (o *Championship) GetGameDataOk() (*Game, bool)`

GetGameDataOk returns a tuple with the GameData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameData

`func (o *Championship) SetGameData(v Game)`

SetGameData sets GameData field to given value.

### HasGameData

`func (o *Championship) HasGameData() bool`

HasGameData returns a boolean if a field has been set.

### GetGameId

`func (o *Championship) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *Championship) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *Championship) SetGameId(v string)`

SetGameId sets GameId field to given value.

### HasGameId

`func (o *Championship) HasGameId() bool`

HasGameId returns a boolean if a field has been set.

### GetId

`func (o *Championship) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Championship) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Championship) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Championship) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJoinChecks

`func (o *Championship) GetJoinChecks() JoinCheck`

GetJoinChecks returns the JoinChecks field if non-nil, zero value otherwise.

### GetJoinChecksOk

`func (o *Championship) GetJoinChecksOk() (*JoinCheck, bool)`

GetJoinChecksOk returns a tuple with the JoinChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinChecks

`func (o *Championship) SetJoinChecks(v JoinCheck)`

SetJoinChecks sets JoinChecks field to given value.

### HasJoinChecks

`func (o *Championship) HasJoinChecks() bool`

HasJoinChecks returns a boolean if a field has been set.

### GetName

`func (o *Championship) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Championship) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Championship) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Championship) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizerData

`func (o *Championship) GetOrganizerData() Organizer`

GetOrganizerData returns the OrganizerData field if non-nil, zero value otherwise.

### GetOrganizerDataOk

`func (o *Championship) GetOrganizerDataOk() (*Organizer, bool)`

GetOrganizerDataOk returns a tuple with the OrganizerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizerData

`func (o *Championship) SetOrganizerData(v Organizer)`

SetOrganizerData sets OrganizerData field to given value.

### HasOrganizerData

`func (o *Championship) HasOrganizerData() bool`

HasOrganizerData returns a boolean if a field has been set.

### GetOrganizerId

`func (o *Championship) GetOrganizerId() string`

GetOrganizerId returns the OrganizerId field if non-nil, zero value otherwise.

### GetOrganizerIdOk

`func (o *Championship) GetOrganizerIdOk() (*string, bool)`

GetOrganizerIdOk returns a tuple with the OrganizerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizerId

`func (o *Championship) SetOrganizerId(v string)`

SetOrganizerId sets OrganizerId field to given value.

### HasOrganizerId

`func (o *Championship) HasOrganizerId() bool`

HasOrganizerId returns a boolean if a field has been set.

### GetPrizes

`func (o *Championship) GetPrizes() []Prize`

GetPrizes returns the Prizes field if non-nil, zero value otherwise.

### GetPrizesOk

`func (o *Championship) GetPrizesOk() (*[]Prize, bool)`

GetPrizesOk returns a tuple with the Prizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrizes

`func (o *Championship) SetPrizes(v []Prize)`

SetPrizes sets Prizes field to given value.

### HasPrizes

`func (o *Championship) HasPrizes() bool`

HasPrizes returns a boolean if a field has been set.

### GetRegion

`func (o *Championship) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Championship) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Championship) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Championship) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRulesId

`func (o *Championship) GetRulesId() string`

GetRulesId returns the RulesId field if non-nil, zero value otherwise.

### GetRulesIdOk

`func (o *Championship) GetRulesIdOk() (*string, bool)`

GetRulesIdOk returns a tuple with the RulesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesId

`func (o *Championship) SetRulesId(v string)`

SetRulesId sets RulesId field to given value.

### HasRulesId

`func (o *Championship) HasRulesId() bool`

HasRulesId returns a boolean if a field has been set.

### GetSchedule

`func (o *Championship) GetSchedule() map[string]ChampionshipSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Championship) GetScheduleOk() (*map[string]ChampionshipSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Championship) SetSchedule(v map[string]ChampionshipSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *Championship) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetScreening

`func (o *Championship) GetScreening() ChampionshipScreening`

GetScreening returns the Screening field if non-nil, zero value otherwise.

### GetScreeningOk

`func (o *Championship) GetScreeningOk() (*ChampionshipScreening, bool)`

GetScreeningOk returns a tuple with the Screening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreening

`func (o *Championship) SetScreening(v ChampionshipScreening)`

SetScreening sets Screening field to given value.

### HasScreening

`func (o *Championship) HasScreening() bool`

HasScreening returns a boolean if a field has been set.

### GetSeedingStrategy

`func (o *Championship) GetSeedingStrategy() string`

GetSeedingStrategy returns the SeedingStrategy field if non-nil, zero value otherwise.

### GetSeedingStrategyOk

`func (o *Championship) GetSeedingStrategyOk() (*string, bool)`

GetSeedingStrategyOk returns a tuple with the SeedingStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedingStrategy

`func (o *Championship) SetSeedingStrategy(v string)`

SetSeedingStrategy sets SeedingStrategy field to given value.

### HasSeedingStrategy

`func (o *Championship) HasSeedingStrategy() bool`

HasSeedingStrategy returns a boolean if a field has been set.

### GetSlots

`func (o *Championship) GetSlots() int64`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *Championship) GetSlotsOk() (*int64, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *Championship) SetSlots(v int64)`

SetSlots sets Slots field to given value.

### HasSlots

`func (o *Championship) HasSlots() bool`

HasSlots returns a boolean if a field has been set.

### GetStatus

`func (o *Championship) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Championship) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Championship) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Championship) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStream

`func (o *Championship) GetStream() ChampionshipStream`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *Championship) GetStreamOk() (*ChampionshipStream, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *Championship) SetStream(v ChampionshipStream)`

SetStream sets Stream field to given value.

### HasStream

`func (o *Championship) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetSubscriptionEnd

`func (o *Championship) GetSubscriptionEnd() int64`

GetSubscriptionEnd returns the SubscriptionEnd field if non-nil, zero value otherwise.

### GetSubscriptionEndOk

`func (o *Championship) GetSubscriptionEndOk() (*int64, bool)`

GetSubscriptionEndOk returns a tuple with the SubscriptionEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionEnd

`func (o *Championship) SetSubscriptionEnd(v int64)`

SetSubscriptionEnd sets SubscriptionEnd field to given value.

### HasSubscriptionEnd

`func (o *Championship) HasSubscriptionEnd() bool`

HasSubscriptionEnd returns a boolean if a field has been set.

### GetSubscriptionStart

`func (o *Championship) GetSubscriptionStart() int64`

GetSubscriptionStart returns the SubscriptionStart field if non-nil, zero value otherwise.

### GetSubscriptionStartOk

`func (o *Championship) GetSubscriptionStartOk() (*int64, bool)`

GetSubscriptionStartOk returns a tuple with the SubscriptionStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionStart

`func (o *Championship) SetSubscriptionStart(v int64)`

SetSubscriptionStart sets SubscriptionStart field to given value.

### HasSubscriptionStart

`func (o *Championship) HasSubscriptionStart() bool`

HasSubscriptionStart returns a boolean if a field has been set.

### GetSubscriptionsLocked

`func (o *Championship) GetSubscriptionsLocked() bool`

GetSubscriptionsLocked returns the SubscriptionsLocked field if non-nil, zero value otherwise.

### GetSubscriptionsLockedOk

`func (o *Championship) GetSubscriptionsLockedOk() (*bool, bool)`

GetSubscriptionsLockedOk returns a tuple with the SubscriptionsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionsLocked

`func (o *Championship) SetSubscriptionsLocked(v bool)`

SetSubscriptionsLocked sets SubscriptionsLocked field to given value.

### HasSubscriptionsLocked

`func (o *Championship) HasSubscriptionsLocked() bool`

HasSubscriptionsLocked returns a boolean if a field has been set.

### GetSubstitutionConfiguration

`func (o *Championship) GetSubstitutionConfiguration() SubstitutionConfiguration`

GetSubstitutionConfiguration returns the SubstitutionConfiguration field if non-nil, zero value otherwise.

### GetSubstitutionConfigurationOk

`func (o *Championship) GetSubstitutionConfigurationOk() (*SubstitutionConfiguration, bool)`

GetSubstitutionConfigurationOk returns a tuple with the SubstitutionConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubstitutionConfiguration

`func (o *Championship) SetSubstitutionConfiguration(v SubstitutionConfiguration)`

SetSubstitutionConfiguration sets SubstitutionConfiguration field to given value.

### HasSubstitutionConfiguration

`func (o *Championship) HasSubstitutionConfiguration() bool`

HasSubstitutionConfiguration returns a boolean if a field has been set.

### GetTotalGroups

`func (o *Championship) GetTotalGroups() int64`

GetTotalGroups returns the TotalGroups field if non-nil, zero value otherwise.

### GetTotalGroupsOk

`func (o *Championship) GetTotalGroupsOk() (*int64, bool)`

GetTotalGroupsOk returns a tuple with the TotalGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGroups

`func (o *Championship) SetTotalGroups(v int64)`

SetTotalGroups sets TotalGroups field to given value.

### HasTotalGroups

`func (o *Championship) HasTotalGroups() bool`

HasTotalGroups returns a boolean if a field has been set.

### GetTotalPrizes

`func (o *Championship) GetTotalPrizes() int64`

GetTotalPrizes returns the TotalPrizes field if non-nil, zero value otherwise.

### GetTotalPrizesOk

`func (o *Championship) GetTotalPrizesOk() (*int64, bool)`

GetTotalPrizesOk returns a tuple with the TotalPrizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPrizes

`func (o *Championship) SetTotalPrizes(v int64)`

SetTotalPrizes sets TotalPrizes field to given value.

### HasTotalPrizes

`func (o *Championship) HasTotalPrizes() bool`

HasTotalPrizes returns a boolean if a field has been set.

### GetTotalRounds

`func (o *Championship) GetTotalRounds() int64`

GetTotalRounds returns the TotalRounds field if non-nil, zero value otherwise.

### GetTotalRoundsOk

`func (o *Championship) GetTotalRoundsOk() (*int64, bool)`

GetTotalRoundsOk returns a tuple with the TotalRounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRounds

`func (o *Championship) SetTotalRounds(v int64)`

SetTotalRounds sets TotalRounds field to given value.

### HasTotalRounds

`func (o *Championship) HasTotalRounds() bool`

HasTotalRounds returns a boolean if a field has been set.

### GetType

`func (o *Championship) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Championship) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Championship) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Championship) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


