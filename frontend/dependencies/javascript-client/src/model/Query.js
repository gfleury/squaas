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
    define(['../ApiClient', '../model/QueryApprovals', '../model/User'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('./QueryApprovals'), require('./User'));
  } else {
    // Browser globals (root is window)
    if (!root.DBqueryBench) {
      root.DBqueryBench = {};
    }
    root.DBqueryBench.Query = factory(root.DBqueryBench.ApiClient, root.DBqueryBench.QueryApprovals, root.DBqueryBench.User);
  }
}(this, function (ApiClient, QueryApprovals, User) {
  'use strict';




  /**
   * The Query model module.
   * @module model/Query
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>Query</code>.
   * @alias module:model/Query
   * @class
   * @param ticketid {String} 
   * @param query {String} 
   */
  var exports = function (ticketid, query) {
    var _this = this;


    _this['ticketid'] = ticketid;


    _this['query'] = query;







  };

  /**
   * Constructs a <code>Query</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/Query} obj Optional instance to populate.
   * @return {module:model/Query} The populated <code>Query</code> instance.
   */
  exports.constructFromObject = function (data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('id')) {
        obj['id'] = ApiClient.convertToType(data['id'], 'String');
      }
      if (data.hasOwnProperty('ticketid')) {
        obj['ticketid'] = ApiClient.convertToType(data['ticketid'], 'String');
      }
      if (data.hasOwnProperty('approvals')) {
        obj['approvals'] = ApiClient.convertToType(data['approvals'], [QueryApprovals]);
      }
      if (data.hasOwnProperty('owner')) {
        obj['owner'] = User.constructFromObject(data['owner']);
      }
      if (data.hasOwnProperty('query')) {
        obj['query'] = ApiClient.convertToType(data['query'], 'String');
      }
      if (data.hasOwnProperty('servername')) {
        obj['servername'] = ApiClient.convertToType(data['servername'], 'String');
      }
      if (data.hasOwnProperty('hasselect')) {
        obj['hasselect'] = ApiClient.convertToType(data['hasselect'], 'Boolean');
      }
      if (data.hasOwnProperty('hasdelete')) {
        obj['hasdelete'] = ApiClient.convertToType(data['hasdelete'], 'Boolean');
      }
      if (data.hasOwnProperty('hasinsert')) {
        obj['hasinsert'] = ApiClient.convertToType(data['hasinsert'], 'Boolean');
      }
      if (data.hasOwnProperty('hasupdate')) {
        obj['hasupdate'] = ApiClient.convertToType(data['hasupdate'], 'Boolean');
      }
      if (data.hasOwnProperty('hastransaction')) {
        obj['hastransaction'] = ApiClient.convertToType(data['hastransaction'], 'Boolean');
      }
      if (data.hasOwnProperty('status')) {
        obj['status'] = ApiClient.convertToType(data['status'], 'String');
      }
      if (data.hasOwnProperty('createdAt')) {
        obj['createdAt'] = ApiClient.convertToType(data['createdAt'], 'String');
      }
      if (data.hasOwnProperty('updatedAt')) {
        obj['updatedAt'] = ApiClient.convertToType(data['updatedAt'], 'String');
      }
    }
    return obj;
  }

  /**
   * @member {Number} id
   */
  exports.prototype['id'] = undefined;
  /**
   * @member {String} ticketid
   */
  exports.prototype['ticketid'] = undefined;
  /**
   * @member {Array.<module:model/QueryApprovals>} approvals
   */
  exports.prototype['approvals'] = undefined;
  /**
   * @member {module:model/User} owner
   */
  exports.prototype['owner'] = undefined;
  /**
   * @member {String} query
   */
  exports.prototype['query'] = undefined;
  /**
   * @member {String} servername
   */
  exports.prototype['servername'] = undefined;
  /**
   * @member {Boolean} hasselect
   */
  exports.prototype['hasselect'] = undefined;
  /**
   * @member {Boolean} hasdelete
   */
  exports.prototype['hasdelete'] = undefined;
  /**
   * @member {Boolean} hasinsert
   */
  exports.prototype['hasinsert'] = undefined;
  /**
   * @member {Boolean} hasupdate
   */
  exports.prototype['hasupdate'] = undefined;
  /**
   * @member {Boolean} hastransaction
   */
  exports.prototype['hastransaction'] = undefined;
  /**
   * query status in the store
   * @member {module:model/Query.StatusEnum} status
   */
  exports.prototype['status'] = undefined;
  /**
   * query status in the store
   * @member {module:model/Query.StatusEnum} status
   */
  exports.prototype['createdAt'] = undefined;
  /**
   * query status in the store
   * @member {module:model/Query.StatusEnum} status
   */
  exports.prototype['updatedAt'] = undefined;

  /**
   * Allowed values for the <code>status</code> property.
   * @enum {String}
   * @readonly
   */
  exports.StatusEnum = {
    /**
     * value: "done"
     * @const
     */
    "done": "done",
    /**
     * value: "pending"
     * @const
     */
    "pending": "pending",
    /**
     * value: "approved"
     * @const
     */
    "approved": "approved",
    /**
     * value: "running"
     * @const
     */
    "running": "running",
    /**
     * value: "failed"
     * @const
     */
    "failed": "failed"
  };


  return exports;
}));


