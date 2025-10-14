# SpacePlayer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivatedAt** | Pointer to **time.Time** |  | [optional] 
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**BirthDate** | Pointer to [**BirthDate**](BirthDate.md) |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CoverFeaturedImage** | Pointer to **string** | Deprecated: no more in use | [optional] 
**CoverImage** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**FaceitUrl** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**FriendsIds** | Pointer to **[]string** |  | [optional] 
**Games** | Pointer to [**map[string]GameDetail**](GameDetail.md) |  | [optional] 
**Infractions** | Pointer to **map[string]interface{}** | Deprecated: no more in use | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
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

### NewSpacePlayer

`func NewSpacePlayer() *SpacePlayer`

NewSpacePlayer instantiates a new SpacePlayer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpacePlayerWithDefaults

`func NewSpacePlayerWithDefaults() *SpacePlayer`

NewSpacePlayerWithDefaults instantiates a new SpacePlayer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivatedAt

`func (o *SpacePlayer) GetActivatedAt() time.Time`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *SpacePlayer) GetActivatedAtOk() (*time.Time, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedAt

`func (o *SpacePlayer) SetActivatedAt(v time.Time)`

SetActivatedAt sets ActivatedAt field to given value.

### HasActivatedAt

`func (o *SpacePlayer) HasActivatedAt() bool`

HasActivatedAt returns a boolean if a field has been set.

### GetAddress

`func (o *SpacePlayer) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SpacePlayer) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SpacePlayer) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SpacePlayer) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAvatar

`func (o *SpacePlayer) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *SpacePlayer) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *SpacePlayer) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *SpacePlayer) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetBirthDate

`func (o *SpacePlayer) GetBirthDate() BirthDate`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *SpacePlayer) GetBirthDateOk() (*BirthDate, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *SpacePlayer) SetBirthDate(v BirthDate)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *SpacePlayer) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### GetCountry

`func (o *SpacePlayer) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SpacePlayer) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SpacePlayer) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SpacePlayer) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCoverFeaturedImage

`func (o *SpacePlayer) GetCoverFeaturedImage() string`

GetCoverFeaturedImage returns the CoverFeaturedImage field if non-nil, zero value otherwise.

### GetCoverFeaturedImageOk

`func (o *SpacePlayer) GetCoverFeaturedImageOk() (*string, bool)`

GetCoverFeaturedImageOk returns a tuple with the CoverFeaturedImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverFeaturedImage

`func (o *SpacePlayer) SetCoverFeaturedImage(v string)`

SetCoverFeaturedImage sets CoverFeaturedImage field to given value.

### HasCoverFeaturedImage

`func (o *SpacePlayer) HasCoverFeaturedImage() bool`

HasCoverFeaturedImage returns a boolean if a field has been set.

### GetCoverImage

`func (o *SpacePlayer) GetCoverImage() string`

GetCoverImage returns the CoverImage field if non-nil, zero value otherwise.

### GetCoverImageOk

`func (o *SpacePlayer) GetCoverImageOk() (*string, bool)`

GetCoverImageOk returns a tuple with the CoverImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImage

`func (o *SpacePlayer) SetCoverImage(v string)`

SetCoverImage sets CoverImage field to given value.

### HasCoverImage

`func (o *SpacePlayer) HasCoverImage() bool`

HasCoverImage returns a boolean if a field has been set.

### GetEmail

`func (o *SpacePlayer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SpacePlayer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SpacePlayer) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SpacePlayer) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFaceitUrl

`func (o *SpacePlayer) GetFaceitUrl() string`

GetFaceitUrl returns the FaceitUrl field if non-nil, zero value otherwise.

### GetFaceitUrlOk

`func (o *SpacePlayer) GetFaceitUrlOk() (*string, bool)`

GetFaceitUrlOk returns a tuple with the FaceitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceitUrl

`func (o *SpacePlayer) SetFaceitUrl(v string)`

SetFaceitUrl sets FaceitUrl field to given value.

### HasFaceitUrl

`func (o *SpacePlayer) HasFaceitUrl() bool`

HasFaceitUrl returns a boolean if a field has been set.

### GetFirstName

`func (o *SpacePlayer) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SpacePlayer) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SpacePlayer) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *SpacePlayer) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetFriendsIds

`func (o *SpacePlayer) GetFriendsIds() []string`

GetFriendsIds returns the FriendsIds field if non-nil, zero value otherwise.

### GetFriendsIdsOk

`func (o *SpacePlayer) GetFriendsIdsOk() (*[]string, bool)`

GetFriendsIdsOk returns a tuple with the FriendsIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendsIds

`func (o *SpacePlayer) SetFriendsIds(v []string)`

SetFriendsIds sets FriendsIds field to given value.

### HasFriendsIds

`func (o *SpacePlayer) HasFriendsIds() bool`

HasFriendsIds returns a boolean if a field has been set.

### GetGames

`func (o *SpacePlayer) GetGames() map[string]GameDetail`

GetGames returns the Games field if non-nil, zero value otherwise.

### GetGamesOk

`func (o *SpacePlayer) GetGamesOk() (*map[string]GameDetail, bool)`

GetGamesOk returns a tuple with the Games field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGames

`func (o *SpacePlayer) SetGames(v map[string]GameDetail)`

SetGames sets Games field to given value.

### HasGames

`func (o *SpacePlayer) HasGames() bool`

HasGames returns a boolean if a field has been set.

### GetInfractions

`func (o *SpacePlayer) GetInfractions() map[string]interface{}`

GetInfractions returns the Infractions field if non-nil, zero value otherwise.

### GetInfractionsOk

`func (o *SpacePlayer) GetInfractionsOk() (*map[string]interface{}, bool)`

GetInfractionsOk returns a tuple with the Infractions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfractions

`func (o *SpacePlayer) SetInfractions(v map[string]interface{})`

SetInfractions sets Infractions field to given value.

### HasInfractions

`func (o *SpacePlayer) HasInfractions() bool`

HasInfractions returns a boolean if a field has been set.

### GetLastName

`func (o *SpacePlayer) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *SpacePlayer) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *SpacePlayer) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *SpacePlayer) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetMembershipType

`func (o *SpacePlayer) GetMembershipType() string`

GetMembershipType returns the MembershipType field if non-nil, zero value otherwise.

### GetMembershipTypeOk

`func (o *SpacePlayer) GetMembershipTypeOk() (*string, bool)`

GetMembershipTypeOk returns a tuple with the MembershipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipType

`func (o *SpacePlayer) SetMembershipType(v string)`

SetMembershipType sets MembershipType field to given value.

### HasMembershipType

`func (o *SpacePlayer) HasMembershipType() bool`

HasMembershipType returns a boolean if a field has been set.

### GetMemberships

`func (o *SpacePlayer) GetMemberships() []string`

GetMemberships returns the Memberships field if non-nil, zero value otherwise.

### GetMembershipsOk

`func (o *SpacePlayer) GetMembershipsOk() (*[]string, bool)`

GetMembershipsOk returns a tuple with the Memberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberships

`func (o *SpacePlayer) SetMemberships(v []string)`

SetMemberships sets Memberships field to given value.

### HasMemberships

`func (o *SpacePlayer) HasMemberships() bool`

HasMemberships returns a boolean if a field has been set.

### GetNewSteamId

`func (o *SpacePlayer) GetNewSteamId() string`

GetNewSteamId returns the NewSteamId field if non-nil, zero value otherwise.

### GetNewSteamIdOk

`func (o *SpacePlayer) GetNewSteamIdOk() (*string, bool)`

GetNewSteamIdOk returns a tuple with the NewSteamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSteamId

`func (o *SpacePlayer) SetNewSteamId(v string)`

SetNewSteamId sets NewSteamId field to given value.

### HasNewSteamId

`func (o *SpacePlayer) HasNewSteamId() bool`

HasNewSteamId returns a boolean if a field has been set.

### GetNickname

`func (o *SpacePlayer) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *SpacePlayer) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *SpacePlayer) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *SpacePlayer) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetPlatforms

`func (o *SpacePlayer) GetPlatforms() map[string]string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *SpacePlayer) GetPlatformsOk() (*map[string]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *SpacePlayer) SetPlatforms(v map[string]string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *SpacePlayer) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetPlayerId

`func (o *SpacePlayer) GetPlayerId() string`

GetPlayerId returns the PlayerId field if non-nil, zero value otherwise.

### GetPlayerIdOk

`func (o *SpacePlayer) GetPlayerIdOk() (*string, bool)`

GetPlayerIdOk returns a tuple with the PlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerId

`func (o *SpacePlayer) SetPlayerId(v string)`

SetPlayerId sets PlayerId field to given value.

### HasPlayerId

`func (o *SpacePlayer) HasPlayerId() bool`

HasPlayerId returns a boolean if a field has been set.

### GetSettings

`func (o *SpacePlayer) GetSettings() UserSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *SpacePlayer) GetSettingsOk() (*UserSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *SpacePlayer) SetSettings(v UserSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *SpacePlayer) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSteamId64

`func (o *SpacePlayer) GetSteamId64() string`

GetSteamId64 returns the SteamId64 field if non-nil, zero value otherwise.

### GetSteamId64Ok

`func (o *SpacePlayer) GetSteamId64Ok() (*string, bool)`

GetSteamId64Ok returns a tuple with the SteamId64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamId64

`func (o *SpacePlayer) SetSteamId64(v string)`

SetSteamId64 sets SteamId64 field to given value.

### HasSteamId64

`func (o *SpacePlayer) HasSteamId64() bool`

HasSteamId64 returns a boolean if a field has been set.

### GetSteamNickname

`func (o *SpacePlayer) GetSteamNickname() string`

GetSteamNickname returns the SteamNickname field if non-nil, zero value otherwise.

### GetSteamNicknameOk

`func (o *SpacePlayer) GetSteamNicknameOk() (*string, bool)`

GetSteamNicknameOk returns a tuple with the SteamNickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteamNickname

`func (o *SpacePlayer) SetSteamNickname(v string)`

SetSteamNickname sets SteamNickname field to given value.

### HasSteamNickname

`func (o *SpacePlayer) HasSteamNickname() bool`

HasSteamNickname returns a boolean if a field has been set.

### GetVerified

`func (o *SpacePlayer) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *SpacePlayer) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *SpacePlayer) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *SpacePlayer) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


