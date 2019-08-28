# DBqueryBench.DatabasesApi

All URIs are relative to *https://dbqueryBench/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**getDatabases**](DatabasesApi.md#getDatabases) | **GET** /databases | Get list of databases


<a name="getDatabases"></a>
# **getDatabases**
> [Server] getDatabases()

Get list of databases

### Example
```javascript
var DBqueryBench = require('d_bquery_bench');

var apiInstance = new DBqueryBench.DatabasesApi();

var callback = function(error, data, response) {
  if (error) {
    console.error(error);
  } else {
    console.log('API called successfully. Returned data: ' + data);
  }
};
apiInstance.getDatabases(callback);
```

### Parameters
This endpoint does not need any parameter.

### Return type

[**[Server]**](Server.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

