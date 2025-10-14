# UserSimple

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Avatar** | Pointer to **string** | The Avatar of a user | [optional] 
**Country** | Pointer to **string** | The Country of a user | [optional] 
**FaceitUrl** | Pointer to **string** | The FaceitUrl of a user | [optional] 
**MembershipType** | Pointer to **string** | Deprecated: use memberships instead | [optional] 
**Memberships** | Pointer to **[]string** | The Memberships of a user | [optional] 
**Nickname** | Pointer to **string** | The Nickname of a user | [optional] 
**SkillLevel** | Pointer to **int64** | The SkillLevel of a user | [optional] 
**UserId** | Pointer to **string** | The ID of a user | [optional] 

## Methods

### NewUserSimple

`func NewUserSimple() *UserSimple`

NewUserSimple instantiates a new UserSimple object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSimpleWithDefaults

`func NewUserSimpleWithDefaults() *UserSimple`

NewUserSimpleWithDefaults instantiates a new UserSimple object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatar

`func (o *UserSimple) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *UserSimple) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *UserSimple) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *UserSimple) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetCountry

`func (o *UserSimple) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UserSimple) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UserSimple) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UserSimple) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetFaceitUrl

`func (o *UserSimple) GetFaceitUrl() string`

GetFaceitUrl returns the FaceitUrl field if non-nil, zero value otherwise.

### GetFaceitUrlOk

`func (o *UserSimple) GetFaceitUrlOk() (*string, bool)`

GetFaceitUrlOk returns a tuple with the FaceitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaceitUrl

`func (o *UserSimple) SetFaceitUrl(v string)`

SetFaceitUrl sets FaceitUrl field to given value.

### HasFaceitUrl

`func (o *UserSimple) HasFaceitUrl() bool`

HasFaceitUrl returns a boolean if a field has been set.

### GetMembershipType

`func (o *UserSimple) GetMembershipType() string`

GetMembershipType returns the MembershipType field if non-nil, zero value otherwise.

### GetMembershipTypeOk

`func (o *UserSimple) GetMembershipTypeOk() (*string, bool)`

GetMembershipTypeOk returns a tuple with the MembershipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipType

`func (o *UserSimple) SetMembershipType(v string)`

SetMembershipType sets MembershipType field to given value.

### HasMembershipType

`func (o *UserSimple) HasMembershipType() bool`

HasMembershipType returns a boolean if a field has been set.

### GetMemberships

`func (o *UserSimple) GetMemberships() []string`

GetMemberships returns the Memberships field if non-nil, zero value otherwise.

### GetMembershipsOk

`func (o *UserSimple) GetMembershipsOk() (*[]string, bool)`

GetMembershipsOk returns a tuple with the Memberships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberships

`func (o *UserSimple) SetMemberships(v []string)`

SetMemberships sets Memberships field to given value.

### HasMemberships

`func (o *UserSimple) HasMemberships() bool`

HasMemberships returns a boolean if a field has been set.

### GetNickname

`func (o *UserSimple) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *UserSimple) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *UserSimple) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *UserSimple) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetSkillLevel

`func (o *UserSimple) GetSkillLevel() int64`

GetSkillLevel returns the SkillLevel field if non-nil, zero value otherwise.

### GetSkillLevelOk

`func (o *UserSimple) GetSkillLevelOk() (*int64, bool)`

GetSkillLevelOk returns a tuple with the SkillLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillLevel

`func (o *UserSimple) SetSkillLevel(v int64)`

SetSkillLevel sets SkillLevel field to given value.

### HasSkillLevel

`func (o *UserSimple) HasSkillLevel() bool`

HasSkillLevel returns a boolean if a field has been set.

### GetUserId

`func (o *UserSimple) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserSimple) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserSimple) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *UserSimple) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


