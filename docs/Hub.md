# Hub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **string** |  | [optional] 
**BackgroundImage** | Pointer to **string** |  | [optional] 
**ChatRoomId** | Pointer to **string** |  | [optional] 
**CoverImage** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**FaceitUrl** | Pointer to **string** |  | [optional] 
**GameData** | Pointer to [**Game**](Game.md) |  | [optional] 
**GameId** | Pointer to **string** |  | [optional] 
**HubId** | Pointer to **string** |  | [optional] 
**JoinPermission** | Pointer to **string** |  | [optional] 
**MaxSkillLevel** | Pointer to **int64** |  | [optional] 
**MinSkillLevel** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrganizerData** | Pointer to [**Organizer**](Organizer.md) |  | [optional] 
**OrganizerId** | Pointer to **string** |  | [optional] 
**PlayersJoined** | Pointer to **int64** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**RuleId** | Pointer to **string** |  | [optional] 

## Methods

### NewHub

`func NewHub() *Hub`

NewHub instantiates a new Hub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHubWithDefaults

`func NewHubWithDefaults() *Hub`

NewHubWithDefaults instantiates a new Hub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *Hub) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *Hub) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *Hub) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *Hub) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetBackgroundImage

`func (o *Hub) GetBackgroundImage() string`

GetBackgroundImage returns the BackgroundImage field if non-nil, zero value otherwise.

### GetBackgroundImageOk

`func (o *Hub) GetBackgroundImageOk() (*string, bool)`

GetBackgroundImageOk returns a tuple with the BackgroundImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundImage

`func (o *Hub) SetBackgroundImage(v string)`

SetBackgroundImage sets BackgroundImage field to given value.

### HasBackgroundImage

`func (o *Hub) HasBackgroundImage() bool`

HasBackgroundImage returns a boolean if a field has been set.

### GetChatRoomId

`func (o *Hub) GetChatRoomId() string`

GetChatRoomId returns the ChatRoomId field if non-nil, zero value otherwise.

### GetChatRoomIdOk

`func (o *Hub) GetChatRoomIdOk() (*string, bool)`

GetChatRoomIdOk returns a tuple with the ChatRoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatRoomId

`func (o *Hub) SetChatRoomId(v string)`

SetChatRoomId sets ChatRoomId field to given value.

### HasChatRoomId

`func (o *Hub) HasChatRoomId() bool`

HasChatRoomId returns a boolean if a field has been set.

### GetCoverImage

`func (o *Hub) GetCoverImage() string`

GetCoverImage returns the CoverImage field if non-nil, zero value otherwise.

### GetCoverImageOk

`func (o *Hub) GetCoverImageOk() (*string, bool)`

GetCoverImageOk returns a tuple with the CoverImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImage

`func (o *Hub) SetCoverImage(v string)`

SetCoverImage sets CoverImage field to given value.

### HasCoverImage

`func (o *Hub) HasCoverImage() bool`

HasCoverImage returns a boolean if a field has been set.

### GetDescription

`func (o *Hub) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Hub) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Hub) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Hub) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFaceitUrl

`func (o *Hub) GetFaceitUrl() string`

GetFaceitUrl returns the FaceitUrl field if non-nil, zero value otherwise.

### GetFaceitUrlOk

`func (o *Hub) GetFaceitUrlOk() (*string, bool)`

GetFaceitUrlOk returns a tuple with the FaceitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceitUrl

`func (o *Hub) SetFaceitUrl(v string)`

SetFaceitUrl sets FaceitUrl field to given value.

### HasFaceitUrl

`func (o *Hub) HasFaceitUrl() bool`

HasFaceitUrl returns a boolean if a field has been set.

### GetGameData

`func (o *Hub) GetGameData() Game`

GetGameData returns the GameData field if non-nil, zero value otherwise.

### GetGameDataOk

`func (o *Hub) GetGameDataOk() (*Game, bool)`

GetGameDataOk returns a tuple with the GameData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameData

`func (o *Hub) SetGameData(v Game)`

SetGameData sets GameData field to given value.

### HasGameData

`func (o *Hub) HasGameData() bool`

HasGameData returns a boolean if a field has been set.

### GetGameId

`func (o *Hub) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *Hub) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *Hub) SetGameId(v string)`

SetGameId sets GameId field to given value.

### HasGameId

`func (o *Hub) HasGameId() bool`

HasGameId returns a boolean if a field has been set.

### GetHubId

`func (o *Hub) GetHubId() string`

GetHubId returns the HubId field if non-nil, zero value otherwise.

### GetHubIdOk

`func (o *Hub) GetHubIdOk() (*string, bool)`

GetHubIdOk returns a tuple with the HubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubId

`func (o *Hub) SetHubId(v string)`

SetHubId sets HubId field to given value.

### HasHubId

`func (o *Hub) HasHubId() bool`

HasHubId returns a boolean if a field has been set.

### GetJoinPermission

`func (o *Hub) GetJoinPermission() string`

GetJoinPermission returns the JoinPermission field if non-nil, zero value otherwise.

### GetJoinPermissionOk

`func (o *Hub) GetJoinPermissionOk() (*string, bool)`

GetJoinPermissionOk returns a tuple with the JoinPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinPermission

`func (o *Hub) SetJoinPermission(v string)`

SetJoinPermission sets JoinPermission field to given value.

### HasJoinPermission

`func (o *Hub) HasJoinPermission() bool`

HasJoinPermission returns a boolean if a field has been set.

### GetMaxSkillLevel

`func (o *Hub) GetMaxSkillLevel() int64`

GetMaxSkillLevel returns the MaxSkillLevel field if non-nil, zero value otherwise.

### GetMaxSkillLevelOk

`func (o *Hub) GetMaxSkillLevelOk() (*int64, bool)`

GetMaxSkillLevelOk returns a tuple with the MaxSkillLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSkillLevel

`func (o *Hub) SetMaxSkillLevel(v int64)`

SetMaxSkillLevel sets MaxSkillLevel field to given value.

### HasMaxSkillLevel

`func (o *Hub) HasMaxSkillLevel() bool`

HasMaxSkillLevel returns a boolean if a field has been set.

### GetMinSkillLevel

`func (o *Hub) GetMinSkillLevel() int64`

GetMinSkillLevel returns the MinSkillLevel field if non-nil, zero value otherwise.

### GetMinSkillLevelOk

`func (o *Hub) GetMinSkillLevelOk() (*int64, bool)`

GetMinSkillLevelOk returns a tuple with the MinSkillLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSkillLevel

`func (o *Hub) SetMinSkillLevel(v int64)`

SetMinSkillLevel sets MinSkillLevel field to given value.

### HasMinSkillLevel

`func (o *Hub) HasMinSkillLevel() bool`

HasMinSkillLevel returns a boolean if a field has been set.

### GetName

`func (o *Hub) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Hub) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Hub) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Hub) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizerData

`func (o *Hub) GetOrganizerData() Organizer`

GetOrganizerData returns the OrganizerData field if non-nil, zero value otherwise.

### GetOrganizerDataOk

`func (o *Hub) GetOrganizerDataOk() (*Organizer, bool)`

GetOrganizerDataOk returns a tuple with the OrganizerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizerData

`func (o *Hub) SetOrganizerData(v Organizer)`

SetOrganizerData sets OrganizerData field to given value.

### HasOrganizerData

`func (o *Hub) HasOrganizerData() bool`

HasOrganizerData returns a boolean if a field has been set.

### GetOrganizerId

`func (o *Hub) GetOrganizerId() string`

GetOrganizerId returns the OrganizerId field if non-nil, zero value otherwise.

### GetOrganizerIdOk

`func (o *Hub) GetOrganizerIdOk() (*string, bool)`

GetOrganizerIdOk returns a tuple with the OrganizerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizerId

`func (o *Hub) SetOrganizerId(v string)`

SetOrganizerId sets OrganizerId field to given value.

### HasOrganizerId

`func (o *Hub) HasOrganizerId() bool`

HasOrganizerId returns a boolean if a field has been set.

### GetPlayersJoined

`func (o *Hub) GetPlayersJoined() int64`

GetPlayersJoined returns the PlayersJoined field if non-nil, zero value otherwise.

### GetPlayersJoinedOk

`func (o *Hub) GetPlayersJoinedOk() (*int64, bool)`

GetPlayersJoinedOk returns a tuple with the PlayersJoined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayersJoined

`func (o *Hub) SetPlayersJoined(v int64)`

SetPlayersJoined sets PlayersJoined field to given value.

### HasPlayersJoined

`func (o *Hub) HasPlayersJoined() bool`

HasPlayersJoined returns a boolean if a field has been set.

### GetRegion

`func (o *Hub) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Hub) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Hub) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Hub) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRuleId

`func (o *Hub) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *Hub) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *Hub) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *Hub) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


