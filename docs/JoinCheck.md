# JoinCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedTeamTypes** | Pointer to **[]string** |  | [optional] 
**BlacklistGeoCountries** | Pointer to **[]string** |  | [optional] 
**JoinPolicy** | Pointer to **string** |  | [optional] 
**MaxSkillLevel** | Pointer to **int64** |  | [optional] 
**MembershipType** | Pointer to **string** |  | [optional] 
**MinSkillLevel** | Pointer to **int64** |  | [optional] 
**WhitelistGeoCountries** | Pointer to **[]string** |  | [optional] 
**WhitelistGeoCountriesMinPlayers** | Pointer to **int64** |  | [optional] 

## Methods

### NewJoinCheck

`func NewJoinCheck() *JoinCheck`

NewJoinCheck instantiates a new JoinCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJoinCheckWithDefaults

`func NewJoinCheckWithDefaults() *JoinCheck`

NewJoinCheckWithDefaults instantiates a new JoinCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedTeamTypes

`func (o *JoinCheck) GetAllowedTeamTypes() []string`

GetAllowedTeamTypes returns the AllowedTeamTypes field if non-nil, zero value otherwise.

### GetAllowedTeamTypesOk

`func (o *JoinCheck) GetAllowedTeamTypesOk() (*[]string, bool)`

GetAllowedTeamTypesOk returns a tuple with the AllowedTeamTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedTeamTypes

`func (o *JoinCheck) SetAllowedTeamTypes(v []string)`

SetAllowedTeamTypes sets AllowedTeamTypes field to given value.

### HasAllowedTeamTypes

`func (o *JoinCheck) HasAllowedTeamTypes() bool`

HasAllowedTeamTypes returns a boolean if a field has been set.

### GetBlacklistGeoCountries

`func (o *JoinCheck) GetBlacklistGeoCountries() []string`

GetBlacklistGeoCountries returns the BlacklistGeoCountries field if non-nil, zero value otherwise.

### GetBlacklistGeoCountriesOk

`func (o *JoinCheck) GetBlacklistGeoCountriesOk() (*[]string, bool)`

GetBlacklistGeoCountriesOk returns a tuple with the BlacklistGeoCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistGeoCountries

`func (o *JoinCheck) SetBlacklistGeoCountries(v []string)`

SetBlacklistGeoCountries sets BlacklistGeoCountries field to given value.

### HasBlacklistGeoCountries

`func (o *JoinCheck) HasBlacklistGeoCountries() bool`

HasBlacklistGeoCountries returns a boolean if a field has been set.

### GetJoinPolicy

`func (o *JoinCheck) GetJoinPolicy() string`

GetJoinPolicy returns the JoinPolicy field if non-nil, zero value otherwise.

### GetJoinPolicyOk

`func (o *JoinCheck) GetJoinPolicyOk() (*string, bool)`

GetJoinPolicyOk returns a tuple with the JoinPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinPolicy

`func (o *JoinCheck) SetJoinPolicy(v string)`

SetJoinPolicy sets JoinPolicy field to given value.

### HasJoinPolicy

`func (o *JoinCheck) HasJoinPolicy() bool`

HasJoinPolicy returns a boolean if a field has been set.

### GetMaxSkillLevel

`func (o *JoinCheck) GetMaxSkillLevel() int64`

GetMaxSkillLevel returns the MaxSkillLevel field if non-nil, zero value otherwise.

### GetMaxSkillLevelOk

`func (o *JoinCheck) GetMaxSkillLevelOk() (*int64, bool)`

GetMaxSkillLevelOk returns a tuple with the MaxSkillLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSkillLevel

`func (o *JoinCheck) SetMaxSkillLevel(v int64)`

SetMaxSkillLevel sets MaxSkillLevel field to given value.

### HasMaxSkillLevel

`func (o *JoinCheck) HasMaxSkillLevel() bool`

HasMaxSkillLevel returns a boolean if a field has been set.

### GetMembershipType

`func (o *JoinCheck) GetMembershipType() string`

GetMembershipType returns the MembershipType field if non-nil, zero value otherwise.

### GetMembershipTypeOk

`func (o *JoinCheck) GetMembershipTypeOk() (*string, bool)`

GetMembershipTypeOk returns a tuple with the MembershipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembershipType

`func (o *JoinCheck) SetMembershipType(v string)`

SetMembershipType sets MembershipType field to given value.

### HasMembershipType

`func (o *JoinCheck) HasMembershipType() bool`

HasMembershipType returns a boolean if a field has been set.

### GetMinSkillLevel

`func (o *JoinCheck) GetMinSkillLevel() int64`

GetMinSkillLevel returns the MinSkillLevel field if non-nil, zero value otherwise.

### GetMinSkillLevelOk

`func (o *JoinCheck) GetMinSkillLevelOk() (*int64, bool)`

GetMinSkillLevelOk returns a tuple with the MinSkillLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSkillLevel

`func (o *JoinCheck) SetMinSkillLevel(v int64)`

SetMinSkillLevel sets MinSkillLevel field to given value.

### HasMinSkillLevel

`func (o *JoinCheck) HasMinSkillLevel() bool`

HasMinSkillLevel returns a boolean if a field has been set.

### GetWhitelistGeoCountries

`func (o *JoinCheck) GetWhitelistGeoCountries() []string`

GetWhitelistGeoCountries returns the WhitelistGeoCountries field if non-nil, zero value otherwise.

### GetWhitelistGeoCountriesOk

`func (o *JoinCheck) GetWhitelistGeoCountriesOk() (*[]string, bool)`

GetWhitelistGeoCountriesOk returns a tuple with the WhitelistGeoCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistGeoCountries

`func (o *JoinCheck) SetWhitelistGeoCountries(v []string)`

SetWhitelistGeoCountries sets WhitelistGeoCountries field to given value.

### HasWhitelistGeoCountries

`func (o *JoinCheck) HasWhitelistGeoCountries() bool`

HasWhitelistGeoCountries returns a boolean if a field has been set.

### GetWhitelistGeoCountriesMinPlayers

`func (o *JoinCheck) GetWhitelistGeoCountriesMinPlayers() int64`

GetWhitelistGeoCountriesMinPlayers returns the WhitelistGeoCountriesMinPlayers field if non-nil, zero value otherwise.

### GetWhitelistGeoCountriesMinPlayersOk

`func (o *JoinCheck) GetWhitelistGeoCountriesMinPlayersOk() (*int64, bool)`

GetWhitelistGeoCountriesMinPlayersOk returns a tuple with the WhitelistGeoCountriesMinPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelistGeoCountriesMinPlayers

`func (o *JoinCheck) SetWhitelistGeoCountriesMinPlayers(v int64)`

SetWhitelistGeoCountriesMinPlayers sets WhitelistGeoCountriesMinPlayers field to given value.

### HasWhitelistGeoCountriesMinPlayers

`func (o *JoinCheck) HasWhitelistGeoCountriesMinPlayers() bool`

HasWhitelistGeoCountriesMinPlayers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


