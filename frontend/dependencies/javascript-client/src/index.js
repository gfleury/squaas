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

(function (factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['./ApiClient', './model/Query', './model/QueryApprovals', './model/QueryResult', './model/Server', './model/User', './api/DatabasesApi', './api/QueryApi'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('./ApiClient'), require('./model/Query'), require('./model/QueryApprovals'), require('./model/QueryResult'), require('./model/Server'), require('./model/User'), require('./api/DatabasesApi'), require('./api/QueryApi'));
  }
}(function (ApiClient, Query, QueryApprovals, QueryResult, Server, User, DatabasesApi, QueryApi) {
  'use strict';

  /**
   * .<br>
   * The <code>index</code> module provides access to constructors for all the classes which comprise the public API.
   * <p>
   * An AMD (recommended!) or CommonJS application will generally do something equivalent to the following:
   * <pre>
   * var DBqueryBench = require('index'); // See note below*.
   * var xxxSvc = new DBqueryBench.XxxApi(); // Allocate the API class we're going to use.
   * var yyyModel = new DBqueryBench.Yyy(); // Construct a model instance.
   * yyyModel.someProperty = 'someValue';
   * ...
   * var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
   * ...
   * </pre>
   * <em>*NOTE: For a top-level AMD script, use require(['index'], function(){...})
   * and put the application logic within the callback function.</em>
   * </p>
   * <p>
   * A non-AMD browser application (discouraged) might do something like this:
   * <pre>
   * var xxxSvc = new DBqueryBench.XxxApi(); // Allocate the API class we're going to use.
   * var yyy = new DBqueryBench.Yyy(); // Construct a model instance.
   * yyyModel.someProperty = 'someValue';
   * ...
   * var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
   * ...
   * </pre>
   * </p>
   * @module index
   * @version 1.0.0
   */
  var exports = {
    /**
     * The ApiClient constructor.
     * @property {module:ApiClient}
     */
    ApiClient: ApiClient,
    /**
     * The Query model constructor.
     * @property {module:model/Query}
     */
    Query: Query,
    /**
     * The QueryApprovals model constructor.
     * @property {module:model/QueryApprovals}
     */
    QueryApprovals: QueryApprovals,
    /**
     * The QueryResult model constructor.
     * @property {module:model/QueryResult}
     */
    QueryResult: QueryResult,
    /**
     * The Server model constructor.
     * @property {module:model/Server}
     */
    Server: Server,
    /**
     * The User model constructor.
     * @property {module:model/User}
     */
    User: User,
    /**
     * The DatabasesApi service constructor.
     * @property {module:api/DatabasesApi}
     */
    DatabasesApi: DatabasesApi,
    /**
     * The QueryApi service constructor.
     * @property {module:api/QueryApi}
     */
    QueryApi: QueryApi
  };

  return exports;
}));
