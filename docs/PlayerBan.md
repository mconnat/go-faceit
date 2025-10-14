# PlayerBan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndsAt** | Pointer to **time.Time** |  | [optional] 
**Game** | Pointer to **string** |  | [optional] 
**Nickname** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**StartsAt** | Pointer to **time.Time** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewPlayerBan

`func NewPlayerBan() *PlayerBan`

NewPlayerBan instantiates a new PlayerBan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerBanWithDefaults

`func NewPlayerBanWithDefaults() *PlayerBan`

NewPlayerBanWithDefaults instantiates a new PlayerBan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndsAt

`func (o *PlayerBan) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *PlayerBan) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *PlayerBan) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *PlayerBan) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### GetGame

`func (o *PlayerBan) GetGame() string`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *PlayerBan) GetGameOk() (*string, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *PlayerBan) SetGame(v string)`

SetGame sets Game field to given value.

### HasGame

`func (o *PlayerBan) HasGame() bool`

HasGame returns a boolean if a field has been set.

### GetNickname

`func (o *PlayerBan) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *PlayerBan) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *PlayerBan) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *PlayerBan) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetReason

`func (o *PlayerBan) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PlayerBan) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PlayerBan) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PlayerBan) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStartsAt

`func (o *PlayerBan) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *PlayerBan) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *PlayerBan) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *PlayerBan) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetType

`func (o *PlayerBan) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PlayerBan) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PlayerBan) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PlayerBan) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *PlayerBan) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PlayerBan) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PlayerBan) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PlayerBan) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


