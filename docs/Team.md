# Team

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **string** | The Avatar of a team | [optional] 
**ChatRoomId** | Pointer to **string** | The ChatID of a team | [optional] 
**CoverImage** | Pointer to **string** | The CoverImage of a team | [optional] 
**Description** | Pointer to **string** | The Description of a team | [optional] 
**Facebook** | Pointer to **string** | The Facebook of a team | [optional] 
**FaceitUrl** | Pointer to **string** | The FaceitUrl of a team | [optional] 
**Game** | Pointer to **string** | The Game of a team | [optional] 
**Leader** | Pointer to **string** | The Leader of a team | [optional] 
**Members** | Pointer to [**[]UserSimple**](UserSimple.md) | The Members of a team  Can be empty if not supported | [optional] 
**Name** | Pointer to **string** | The Name of a team | [optional] 
**Nickname** | Pointer to **string** | The Nickname of a team | [optional] 
**TeamId** | Pointer to **string** | The ID of a team | [optional] 
**TeamType** | Pointer to **string** | The TeamType of a team | [optional] 
**Twitter** | Pointer to **string** | The Twitter of a team | [optional] 
**Website** | Pointer to **string** | The Website of a team | [optional] 
**Youtube** | Pointer to **string** | The Youtube of a team | [optional] 

## Methods

### NewTeam

`func NewTeam() *Team`

NewTeam instantiates a new Team object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamWithDefaults

`func NewTeamWithDefaults() *Team`

NewTeamWithDefaults instantiates a new Team object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *Team) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *Team) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *Team) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *Team) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetChatRoomId

`func (o *Team) GetChatRoomId() string`

GetChatRoomId returns the ChatRoomId field if non-nil, zero value otherwise.

### GetChatRoomIdOk

`func (o *Team) GetChatRoomIdOk() (*string, bool)`

GetChatRoomIdOk returns a tuple with the ChatRoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatRoomId

`func (o *Team) SetChatRoomId(v string)`

SetChatRoomId sets ChatRoomId field to given value.

### HasChatRoomId

`func (o *Team) HasChatRoomId() bool`

HasChatRoomId returns a boolean if a field has been set.

### GetCoverImage

`func (o *Team) GetCoverImage() string`

GetCoverImage returns the CoverImage field if non-nil, zero value otherwise.

### GetCoverImageOk

`func (o *Team) GetCoverImageOk() (*string, bool)`

GetCoverImageOk returns a tuple with the CoverImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImage

`func (o *Team) SetCoverImage(v string)`

SetCoverImage sets CoverImage field to given value.

### HasCoverImage

`func (o *Team) HasCoverImage() bool`

HasCoverImage returns a boolean if a field has been set.

### GetDescription

`func (o *Team) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Team) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Team) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Team) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFacebook

`func (o *Team) GetFacebook() string`

GetFacebook returns the Facebook field if non-nil, zero value otherwise.

### GetFacebookOk

`func (o *Team) GetFacebookOk() (*string, bool)`

GetFacebookOk returns a tuple with the Facebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebook

`func (o *Team) SetFacebook(v string)`

SetFacebook sets Facebook field to given value.

### HasFacebook

`func (o *Team) HasFacebook() bool`

HasFacebook returns a boolean if a field has been set.

### GetFaceitUrl

`func (o *Team) GetFaceitUrl() string`

GetFaceitUrl returns the FaceitUrl field if non-nil, zero value otherwise.

### GetFaceitUrlOk

`func (o *Team) GetFaceitUrlOk() (*string, bool)`

GetFaceitUrlOk returns a tuple with the FaceitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceitUrl

`func (o *Team) SetFaceitUrl(v string)`

SetFaceitUrl sets FaceitUrl field to given value.

### HasFaceitUrl

`func (o *Team) HasFaceitUrl() bool`

HasFaceitUrl returns a boolean if a field has been set.

### GetGame

`func (o *Team) GetGame() string`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *Team) GetGameOk() (*string, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *Team) SetGame(v string)`

SetGame sets Game field to given value.

### HasGame

`func (o *Team) HasGame() bool`

HasGame returns a boolean if a field has been set.

### GetLeader

`func (o *Team) GetLeader() string`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *Team) GetLeaderOk() (*string, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *Team) SetLeader(v string)`

SetLeader sets Leader field to given value.

### HasLeader

`func (o *Team) HasLeader() bool`

HasLeader returns a boolean if a field has been set.

### GetMembers

`func (o *Team) GetMembers() []UserSimple`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Team) GetMembersOk() (*[]UserSimple, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Team) SetMembers(v []UserSimple)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Team) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetName

`func (o *Team) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Team) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Team) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Team) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNickname

`func (o *Team) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *Team) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *Team) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *Team) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetTeamId

`func (o *Team) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *Team) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *Team) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *Team) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetTeamType

`func (o *Team) GetTeamType() string`

GetTeamType returns the TeamType field if non-nil, zero value otherwise.

### GetTeamTypeOk

`func (o *Team) GetTeamTypeOk() (*string, bool)`

GetTeamTypeOk returns a tuple with the TeamType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamType

`func (o *Team) SetTeamType(v string)`

SetTeamType sets TeamType field to given value.

### HasTeamType

`func (o *Team) HasTeamType() bool`

HasTeamType returns a boolean if a field has been set.

### GetTwitter

`func (o *Team) GetTwitter() string`

GetTwitter returns the Twitter field if non-nil, zero value otherwise.

### GetTwitterOk

`func (o *Team) GetTwitterOk() (*string, bool)`

GetTwitterOk returns a tuple with the Twitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitter

`func (o *Team) SetTwitter(v string)`

SetTwitter sets Twitter field to given value.

### HasTwitter

`func (o *Team) HasTwitter() bool`

HasTwitter returns a boolean if a field has been set.

### GetWebsite

`func (o *Team) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *Team) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *Team) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *Team) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetYoutube

`func (o *Team) GetYoutube() string`

GetYoutube returns the Youtube field if non-nil, zero value otherwise.

### GetYoutubeOk

`func (o *Team) GetYoutubeOk() (*string, bool)`

GetYoutubeOk returns a tuple with the Youtube field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYoutube

`func (o *Team) SetYoutube(v string)`

SetYoutube sets Youtube field to given value.

### HasYoutube

`func (o *Team) HasYoutube() bool`

HasYoutube returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


