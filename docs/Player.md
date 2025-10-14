# Player

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivatedAt** | Pointer to **time.Time** |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CoverFeaturedImage** | Pointer to **string** | Deprecated: no more in use | [optional] 
**CoverImage** | Pointer to **string** |  | [optional] 
**FaceitUrl** | Pointer to **string** |  | [optional] 
**FriendsIds** | Pointer to **[]string** |  | [optional] 
**Games** | Pointer to [**map[string]GameDetail**](GameDetail.md) |  | [optional] 
**Infractions** | Pointer to **map[string]interface{}** | Deprecated: no more in use | [optional] 
**MembershipType** | Pointer to **string** | Deprecated: use memberships instead | [optional] 
**Memberships** | Pointer to **[]string** |  | [optional] 
**NewSteamId** | Pointer to **string** |  | [optional] 
**Nickname** | Pointer to **string** |  | [optional] 
**Platforms** | Pointer to **map[string]string** |  | [optional] 
**PlayerId** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**UserSettings**](UserSettings.md) |  | [optional] 
**SteamId64** | Pointer to **string** |  | [optional] 
**SteamNickname** | Pointer to **string** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 

## Methods

### NewPlayer

`func NewPlayer() *Player`

NewPlayer instantiates a new Player object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerWithDefaults

`func NewPlayerWithDefaults() *Player`

NewPlayerWithDefaults instantiates a new Player object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivatedAt

`func (o *Player) GetActivatedAt() time.Time`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *Player) GetActivatedAtOk() (*time.Time, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedAt

`func (o *Player) SetActivatedAt(v time.Time)`

SetActivatedAt sets ActivatedAt field to given value.

### HasActivatedAt

`func (o *Player) HasActivatedAt() bool`

HasActivatedAt returns a boolean if a field has been set.

### GetAvatar

`func (o *Player) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *Player) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *Player) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *Player) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetCountry

`func (o *Player) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Player) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Player) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Player) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCoverFeaturedImage

`func (o *Player) GetCoverFeaturedImage() string`

GetCoverFeaturedImage returns the CoverFeaturedImage field if non-nil, zero value otherwise.

### GetCoverFeaturedImageOk

`func (o *Player) GetCoverFeaturedImageOk() (*string, bool)`

GetCoverFeaturedImageOk returns a tuple with the CoverFeaturedImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverFeaturedImage

`func (o *Player) SetCoverFeaturedImage(v string)`

SetCoverFeaturedImage sets CoverFeaturedImage field to given value.

### HasCoverFeaturedImage

`func (o *Player) HasCoverFeaturedImage() bool`

HasCoverFeaturedImage returns a boolean if a field has been set.

### GetCoverImage

`func (o *Player) GetCoverImage() string`

GetCoverImage returns the CoverImage field if non-nil, zero value otherwise.

### GetCoverImageOk

`func (o *Player) GetCoverImageOk() (*string, bool)`

GetCoverImageOk returns a tuple with the CoverImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImage

`func (o *Player) SetCoverImage(v string)`

SetCoverImage sets CoverImage field to given value.

### HasCoverImage

`func (o *Player) HasCoverImage() bool`

HasCoverImage returns a boolean if a field has been set.

### GetFaceitUrl

`func (o *Player) GetFaceitUrl() string`

GetFaceitUrl returns the FaceitUrl field if non-nil, zero value otherwise.

### GetFaceitUrlOk

`func (o *Player) GetFaceitUrlOk() (*string, bool)`

GetFaceitUrlOk returns a tuple with the FaceitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceitUrl

`func (o *Player) SetFaceitUrl(v string)`

SetFaceitUrl sets FaceitUrl field to given value.

### HasFaceitUrl

`func (o *Player) HasFaceitUrl() bool`

HasFaceitUrl returns a boolean if a field has been set.

### GetFriendsIds

`func (o *Player) GetFriendsIds() []string`

GetFriendsIds returns the FriendsIds field if non-nil, zero value otherwise.

### GetFriendsIdsOk

`func (o *Player) GetFriendsIdsOk() (*[]string, bool)`

GetFriendsIdsOk returns a tuple with the FriendsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendsIds

`func (o *Player) SetFriendsIds(v []string)`

SetFriendsIds sets FriendsIds field to given value.

### HasFriendsIds

`func (o *Player) HasFriendsIds() bool`

HasFriendsIds returns a boolean if a field has been set.

### GetGames

`func (o *Player) GetGames() map[string]GameDetail`

GetGames returns the Games field if non-nil, zero value otherwise.

### GetGamesOk

`func (o *Player) GetGamesOk() (*map[string]GameDetail, bool)`

GetGamesOk returns a tuple with the Games field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGames

`func (o *Player) SetGames(v map[string]GameDetail)`

SetGames sets Games field to given value.

### HasGames

`func (o *Player) HasGames() bool`

HasGames returns a boolean if a field has been set.

### GetInfractions

`func (o *Player) GetInfractions() map[string]interface{}`

GetInfractions returns the Infractions field if non-nil, zero value otherwise.

### GetInfractionsOk

`func (o *Player) GetInfractionsOk() (*map[string]interface{}, bool)`

GetInfractionsOk returns a tuple with the Infractions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfractions

`func (o *Player) SetInfractions(v map[string]interface{})`

SetInfractions sets Infractions field to given value.

### HasInfractions

`func (o *Player) HasInfractions() bool`

HasInfractions returns a boolean if a field has been set.

### GetMembershipType

`func (o *Player) GetMembershipType() string`

GetMembershipType returns the MembershipType field if non-nil, zero value otherwise.

### GetMembershipTypeOk

`func (o *Player) GetMembershipTypeOk() (*string, bool)`

GetMembershipTypeOk returns a tuple with the MembershipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipType

`func (o *Player) SetMembershipType(v string)`

SetMembershipType sets MembershipType field to given value.

### HasMembershipType

`func (o *Player) HasMembershipType() bool`

HasMembershipType returns a boolean if a field has been set.

### GetMemberships

`func (o *Player) GetMemberships() []string`

GetMemberships returns the Memberships field if non-nil, zero value otherwise.

### GetMembershipsOk

`func (o *Player) GetMembershipsOk() (*[]string, bool)`

GetMembershipsOk returns a tuple with the Memberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberships

`func (o *Player) SetMemberships(v []string)`

SetMemberships sets Memberships field to given value.

### HasMemberships

`func (o *Player) HasMemberships() bool`

HasMemberships returns a boolean if a field has been set.

### GetNewSteamId

`func (o *Player) GetNewSteamId() string`

GetNewSteamId returns the NewSteamId field if non-nil, zero value otherwise.

### GetNewSteamIdOk

`func (o *Player) GetNewSteamIdOk() (*string, bool)`

GetNewSteamIdOk returns a tuple with the NewSteamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSteamId

`func (o *Player) SetNewSteamId(v string)`

SetNewSteamId sets NewSteamId field to given value.

### HasNewSteamId

`func (o *Player) HasNewSteamId() bool`

HasNewSteamId returns a boolean if a field has been set.

### GetNickname

`func (o *Player) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *Player) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *Player) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *Player) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetPlatforms

`func (o *Player) GetPlatforms() map[string]string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *Player) GetPlatformsOk() (*map[string]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *Player) SetPlatforms(v map[string]string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *Player) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetPlayerId

`func (o *Player) GetPlayerId() string`

GetPlayerId returns the PlayerId field if non-nil, zero value otherwise.

### GetPlayerIdOk

`func (o *Player) GetPlayerIdOk() (*string, bool)`

GetPlayerIdOk returns a tuple with the PlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerId

`func (o *Player) SetPlayerId(v string)`

SetPlayerId sets PlayerId field to given value.

### HasPlayerId

`func (o *Player) HasPlayerId() bool`

HasPlayerId returns a boolean if a field has been set.

### GetSettings

`func (o *Player) GetSettings() UserSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *Player) GetSettingsOk() (*UserSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *Player) SetSettings(v UserSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *Player) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSteamId64

`func (o *Player) GetSteamId64() string`

GetSteamId64 returns the SteamId64 field if non-nil, zero value otherwise.

### GetSteamId64Ok

`func (o *Player) GetSteamId64Ok() (*string, bool)`

GetSteamId64Ok returns a tuple with the SteamId64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamId64

`func (o *Player) SetSteamId64(v string)`

SetSteamId64 sets SteamId64 field to given value.

### HasSteamId64

`func (o *Player) HasSteamId64() bool`

HasSteamId64 returns a boolean if a field has been set.

### GetSteamNickname

`func (o *Player) GetSteamNickname() string`

GetSteamNickname returns the SteamNickname field if non-nil, zero value otherwise.

### GetSteamNicknameOk

`func (o *Player) GetSteamNicknameOk() (*string, bool)`

GetSteamNicknameOk returns a tuple with the SteamNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamNickname

`func (o *Player) SetSteamNickname(v string)`

SetSteamNickname sets SteamNickname field to given value.

### HasSteamNickname

`func (o *Player) HasSteamNickname() bool`

HasSteamNickname returns a boolean if a field has been set.

### GetVerified

`func (o *Player) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *Player) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *Player) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *Player) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


