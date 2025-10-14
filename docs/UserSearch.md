# UserSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Games** | Pointer to [**[]GameUserSearch**](GameUserSearch.md) |  | [optional] 
**Nickname** | Pointer to **string** |  | [optional] 
**PlayerId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserSearch

`func NewUserSearch() *UserSearch`

NewUserSearch instantiates a new UserSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSearchWithDefaults

`func NewUserSearchWithDefaults() *UserSearch`

NewUserSearchWithDefaults instantiates a new UserSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *UserSearch) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UserSearch) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UserSearch) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UserSearch) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetCountry

`func (o *UserSearch) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UserSearch) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UserSearch) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UserSearch) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetGames

`func (o *UserSearch) GetGames() []GameUserSearch`

GetGames returns the Games field if non-nil, zero value otherwise.

### GetGamesOk

`func (o *UserSearch) GetGamesOk() (*[]GameUserSearch, bool)`

GetGamesOk returns a tuple with the Games field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGames

`func (o *UserSearch) SetGames(v []GameUserSearch)`

SetGames sets Games field to given value.

### HasGames

`func (o *UserSearch) HasGames() bool`

HasGames returns a boolean if a field has been set.

### GetNickname

`func (o *UserSearch) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *UserSearch) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *UserSearch) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *UserSearch) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetPlayerId

`func (o *UserSearch) GetPlayerId() string`

GetPlayerId returns the PlayerId field if non-nil, zero value otherwise.

### GetPlayerIdOk

`func (o *UserSearch) GetPlayerIdOk() (*string, bool)`

GetPlayerIdOk returns a tuple with the PlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerId

`func (o *UserSearch) SetPlayerId(v string)`

SetPlayerId sets PlayerId field to given value.

### HasPlayerId

`func (o *UserSearch) HasPlayerId() bool`

HasPlayerId returns a boolean if a field has been set.

### GetStatus

`func (o *UserSearch) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UserSearch) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UserSearch) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UserSearch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVerified

`func (o *UserSearch) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *UserSearch) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *UserSearch) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *UserSearch) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


