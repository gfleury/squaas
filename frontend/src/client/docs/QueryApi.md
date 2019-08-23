# DBqueryBench.QueryApi

All URIs are relative to *https://dbqueryBench/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**addQuery**](QueryApi.md#addQuery) | **POST** /query | Add a new query to the queue
[**approveQuery**](QueryApi.md#approveQuery) | **POST** /query/approve/{queryId} | Approve a query in the queue
[**deleteQuery**](QueryApi.md#deleteQuery) | **DELETE** /query/approve/{queryId} | Deletes a query
[**findQueryByOwner**](QueryApi.md#findQueryByOwner) | **GET** /query/findByOwner | Finds Query by Owner
[**findQueryByStatus**](QueryApi.md#findQueryByStatus) | **GET** /query/findByStatus | Finds Query by status
[**getQueryById**](QueryApi.md#getQueryById) | **GET** /query/{queryId} | Find query by ID
[**updateQuery**](QueryApi.md#updateQuery) | **PUT** /query | Update an existing query


<a name="addQuery"></a>
# **addQuery**
> addQuery(body)

Add a new query to the queue



### Example
```javascript
var DBqueryBench = require('d_bquery_bench');

var apiInstance = new DBqueryBench.QueryApi();

var body = new DBqueryBench.Query(); // Query | Query that needs to be queued


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.addQuery(body, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Query**](Query.md)| Query that needs to be queued | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="approveQuery"></a>
# **approveQuery**
> approveQuery(queryId, status)

Approve a query in the queue



### Example
```javascript
var DBqueryBench = require('d_bquery_bench');

var apiInstance = new DBqueryBench.QueryApi();

var queryId = 789; // Number | ID of query that needs to be updated

var status = "status_example"; // String | Updated status of the query


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.approveQuery(queryId, status, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryId** | **Number**| ID of query that needs to be updated | 
 **status** | **String**| Updated status of the query | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="deleteQuery"></a>
# **deleteQuery**
> deleteQuery(queryId)

Deletes a query



### Example
```javascript
var DBqueryBench = require('d_bquery_bench');

var apiInstance = new DBqueryBench.QueryApi();

var queryId = 789; // Number | Query id to delete


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.deleteQuery(queryId, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryId** | **Number**| Query id to delete | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="findQueryByOwner"></a>
# **findQueryByOwner**
> [Query] findQueryByOwner(owner)

Finds Query by Owner

Muliple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

### Example
```javascript
var DBqueryBench = require('d_bquery_bench');

var apiInstance = new DBqueryBench.QueryApi();

var owner = ["owner_example"]; // [String] | Owner to filter by


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.findQueryByOwner(owner, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | [**[String]**](String.md)| Owner to filter by | 

### Return type

[**[Query]**](Query.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="findQueryByStatus"></a>
# **findQueryByStatus**
> [Query] findQueryByStatus(status)

Finds Query by status

Multiple status values can be provided with comma separated strings

### Example
```javascript
var DBqueryBench = require('d_bquery_bench');

var apiInstance = new DBqueryBench.QueryApi();

var status = ["status_example"]; // [String] | Status values that need to be considered for filter


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.findQueryByStatus(status, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | [**[String]**](String.md)| Status values that need to be considered for filter | 

### Return type

[**[Query]**](Query.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="getQueryById"></a>
# **getQueryById**
> Query getQueryById(queryId)

Find query by ID

Returns a single query

### Example
```javascript
var DBqueryBench = require('d_bquery_bench');

var apiInstance = new DBqueryBench.QueryApi();

var queryId = 789; // Number | ID of query to return


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getQueryById(queryId, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **queryId** | **Number**| ID of query to return | 

### Return type

[**Query**](Query.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="updateQuery"></a>
# **updateQuery**
> updateQuery(body)

Update an existing query



### Example
```javascript
var DBqueryBench = require('d_bquery_bench');

var apiInstance = new DBqueryBench.QueryApi();

var body = new DBqueryBench.Query(); // Query | Query that needs to be updated


var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully.');
  }
};
apiInstance.updateQuery(body, callback);
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Query**](Query.md)| Query that needs to be updated | 

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

