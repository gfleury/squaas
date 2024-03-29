/**
 * DBqueryBench
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0.0
 * Contact: apiteam@swagger.io
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.8
 *
 * Do not edit the class manually.
 *
 */

(function (root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['../ApiClient', '../model/Server'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/Server'));
  } else {
    // Browser globals (root is window)
    if (!root.DBqueryBench) {
      root.DBqueryBench = {};
    }
    root.DBqueryBench.DatabasesApi = factory(root.DBqueryBench.ApiClient, root.DBqueryBench.Server);
  }
}(this, function (ApiClient, Server) {
  'use strict';

  /**
   * Databases service.
   * @module api/DatabasesApi
   * @version 1.0.0
   */

  /**
   * Constructs a new DatabasesApi. 
   * @alias module:api/DatabasesApi
   * @class
   * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function (apiClient) {
    this.apiClient = apiClient || ApiClient.instance;


    /**
     * Callback function to receive the result of the getDatabases operation.
     * @callback module:api/DatabasesApi~getDatabasesCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/Server>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get list of databases
     * @param {module:api/DatabasesApi~getDatabasesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/Server>}
     */
    this.getDatabases = function (callback) {
      var postBody = null;


      var pathParams = {
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = [];
      var contentTypes = [];
      var accepts = ['application/json'];
      var returnType = [Server];

      return this.apiClient.callApi(
        '/databases', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
  };

  return exports;
}));
