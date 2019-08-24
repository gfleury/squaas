# DBqueryBench.Query

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **Number** |  | [optional] 
**ticketid** | **String** |  | 
**approvals** | [**[User]**](User.md) |  | [optional] 
**owner** | [**User**](User.md) |  | [optional] 
**query** | **String** |  | 
**hasselect** | **Boolean** |  | [optional] 
**hasdelete** | **Boolean** |  | [optional] 
**hasinsert** | **Boolean** |  | [optional] 
**hasupdate** | **Boolean** |  | [optional] 
**hastransaction** | **Boolean** |  | [optional] 
**status** | **String** | query status in the store | [optional] 


<a name="StatusEnum"></a>
## Enum: StatusEnum


* `done` (value: `"done"`)

* `pending` (value: `"pending"`)

* `approved` (value: `"approved"`)

* `running` (value: `"running"`)

* `failed` (value: `"failed"`)




