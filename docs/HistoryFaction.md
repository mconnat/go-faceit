# HistoryFaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **string** |  | [optional] 
**Nickname** | Pointer to **string** |  | [optional] 
**Players** | Pointer to [**[]MatchHistoryPlayer**](MatchHistoryPlayer.md) |  | [optional] 
**TeamId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewHistoryFaction

`func NewHistoryFaction() *HistoryFaction`

NewHistoryFaction instantiates a new HistoryFaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoryFactionWithDefaults

`func NewHistoryFactionWithDefaults() *HistoryFaction`

NewHistoryFactionWithDefaults instantiates a new HistoryFaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *HistoryFaction) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *HistoryFaction) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *HistoryFaction) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *HistoryFaction) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetNickname

`func (o *HistoryFaction) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *HistoryFaction) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *HistoryFaction) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *HistoryFaction) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetPlayers

`func (o *HistoryFaction) GetPlayers() []MatchHistoryPlayer`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *HistoryFaction) GetPlayersOk() (*[]MatchHistoryPlayer, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *HistoryFaction) SetPlayers(v []MatchHistoryPlayer)`

SetPlayers sets Players field to given value.

### HasPlayers

`func (o *HistoryFaction) HasPlayers() bool`

HasPlayers returns a boolean if a field has been set.

### GetTeamId

`func (o *HistoryFaction) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *HistoryFaction) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *HistoryFaction) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.

### HasTeamId

`func (o *HistoryFaction) HasTeamId() bool`

HasTeamId returns a boolean if a field has been set.

### GetType

`func (o *HistoryFaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HistoryFaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HistoryFaction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HistoryFaction) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


