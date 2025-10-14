# QueueSimple

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **string** |  | [optional] 
**EntityType** | Pointer to **string** |  | [optional] 
**Game** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastModified** | Pointer to **time.Time** |  | [optional] 
**Open** | Pointer to **bool** |  | [optional] 
**OrganizerId** | Pointer to **string** |  | [optional] 
**QueueName** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewQueueSimple

`func NewQueueSimple() *QueueSimple`

NewQueueSimple instantiates a new QueueSimple object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueSimpleWithDefaults

`func NewQueueSimpleWithDefaults() *QueueSimple`

NewQueueSimpleWithDefaults instantiates a new QueueSimple object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *QueueSimple) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *QueueSimple) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *QueueSimple) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *QueueSimple) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityType

`func (o *QueueSimple) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *QueueSimple) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *QueueSimple) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *QueueSimple) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetGame

`func (o *QueueSimple) GetGame() string`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *QueueSimple) GetGameOk() (*string, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *QueueSimple) SetGame(v string)`

SetGame sets Game field to given value.

### HasGame

`func (o *QueueSimple) HasGame() bool`

HasGame returns a boolean if a field has been set.

### GetId

`func (o *QueueSimple) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QueueSimple) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QueueSimple) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QueueSimple) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastModified

`func (o *QueueSimple) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *QueueSimple) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *QueueSimple) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *QueueSimple) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetOpen

`func (o *QueueSimple) GetOpen() bool`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *QueueSimple) GetOpenOk() (*bool, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *QueueSimple) SetOpen(v bool)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *QueueSimple) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetOrganizerId

`func (o *QueueSimple) GetOrganizerId() string`

GetOrganizerId returns the OrganizerId field if non-nil, zero value otherwise.

### GetOrganizerIdOk

`func (o *QueueSimple) GetOrganizerIdOk() (*string, bool)`

GetOrganizerIdOk returns a tuple with the OrganizerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizerId

`func (o *QueueSimple) SetOrganizerId(v string)`

SetOrganizerId sets OrganizerId field to given value.

### HasOrganizerId

`func (o *QueueSimple) HasOrganizerId() bool`

HasOrganizerId returns a boolean if a field has been set.

### GetQueueName

`func (o *QueueSimple) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *QueueSimple) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *QueueSimple) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.

### HasQueueName

`func (o *QueueSimple) HasQueueName() bool`

HasQueueName returns a boolean if a field has been set.

### GetRegion

`func (o *QueueSimple) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *QueueSimple) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *QueueSimple) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *QueueSimple) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetState

`func (o *QueueSimple) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *QueueSimple) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *QueueSimple) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *QueueSimple) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


