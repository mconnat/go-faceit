# QueueBan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BanEnd** | Pointer to **time.Time** |  | [optional] 
**BanId** | Pointer to **string** |  | [optional] 
**BanStart** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**EntityId** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to **string** |  | [optional] 
**Expired** | Pointer to **bool** |  | [optional] 
**LastModified** | Pointer to **time.Time** |  | [optional] 
**Nickname** | Pointer to **string** |  | [optional] 
**OrganizerId** | Pointer to **string** |  | [optional] 
**QueueId** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewQueueBan

`func NewQueueBan() *QueueBan`

NewQueueBan instantiates a new QueueBan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueBanWithDefaults

`func NewQueueBanWithDefaults() *QueueBan`

NewQueueBanWithDefaults instantiates a new QueueBan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBanEnd

`func (o *QueueBan) GetBanEnd() time.Time`

GetBanEnd returns the BanEnd field if non-nil, zero value otherwise.

### GetBanEndOk

`func (o *QueueBan) GetBanEndOk() (*time.Time, bool)`

GetBanEndOk returns a tuple with the BanEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanEnd

`func (o *QueueBan) SetBanEnd(v time.Time)`

SetBanEnd sets BanEnd field to given value.

### HasBanEnd

`func (o *QueueBan) HasBanEnd() bool`

HasBanEnd returns a boolean if a field has been set.

### GetBanId

`func (o *QueueBan) GetBanId() string`

GetBanId returns the BanId field if non-nil, zero value otherwise.

### GetBanIdOk

`func (o *QueueBan) GetBanIdOk() (*string, bool)`

GetBanIdOk returns a tuple with the BanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanId

`func (o *QueueBan) SetBanId(v string)`

SetBanId sets BanId field to given value.

### HasBanId

`func (o *QueueBan) HasBanId() bool`

HasBanId returns a boolean if a field has been set.

### GetBanStart

`func (o *QueueBan) GetBanStart() time.Time`

GetBanStart returns the BanStart field if non-nil, zero value otherwise.

### GetBanStartOk

`func (o *QueueBan) GetBanStartOk() (*time.Time, bool)`

GetBanStartOk returns a tuple with the BanStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanStart

`func (o *QueueBan) SetBanStart(v time.Time)`

SetBanStart sets BanStart field to given value.

### HasBanStart

`func (o *QueueBan) HasBanStart() bool`

HasBanStart returns a boolean if a field has been set.

### GetCreatedAt

`func (o *QueueBan) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QueueBan) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QueueBan) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *QueueBan) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEntityId

`func (o *QueueBan) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *QueueBan) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *QueueBan) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *QueueBan) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityType

`func (o *QueueBan) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *QueueBan) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *QueueBan) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *QueueBan) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetExpired

`func (o *QueueBan) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *QueueBan) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *QueueBan) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *QueueBan) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetLastModified

`func (o *QueueBan) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *QueueBan) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *QueueBan) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *QueueBan) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetNickname

`func (o *QueueBan) GetNickname() string`

GetNickname returns the Nickname field if non-nil, zero value otherwise.

### GetNicknameOk

`func (o *QueueBan) GetNicknameOk() (*string, bool)`

GetNicknameOk returns a tuple with the Nickname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickname

`func (o *QueueBan) SetNickname(v string)`

SetNickname sets Nickname field to given value.

### HasNickname

`func (o *QueueBan) HasNickname() bool`

HasNickname returns a boolean if a field has been set.

### GetOrganizerId

`func (o *QueueBan) GetOrganizerId() string`

GetOrganizerId returns the OrganizerId field if non-nil, zero value otherwise.

### GetOrganizerIdOk

`func (o *QueueBan) GetOrganizerIdOk() (*string, bool)`

GetOrganizerIdOk returns a tuple with the OrganizerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizerId

`func (o *QueueBan) SetOrganizerId(v string)`

SetOrganizerId sets OrganizerId field to given value.

### HasOrganizerId

`func (o *QueueBan) HasOrganizerId() bool`

HasOrganizerId returns a boolean if a field has been set.

### GetQueueId

`func (o *QueueBan) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *QueueBan) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *QueueBan) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *QueueBan) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetReason

`func (o *QueueBan) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *QueueBan) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *QueueBan) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *QueueBan) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetType

`func (o *QueueBan) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *QueueBan) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *QueueBan) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *QueueBan) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *QueueBan) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *QueueBan) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *QueueBan) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *QueueBan) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetVersion

`func (o *QueueBan) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *QueueBan) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *QueueBan) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *QueueBan) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


