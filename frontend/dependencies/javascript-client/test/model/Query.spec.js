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

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.DBqueryBench);
  }
}(this, function(expect, DBqueryBench) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new DBqueryBench.Query();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('Query', function() {
    it('should create an instance of Query', function() {
      // uncomment below and update the code to test Query
      //var instance = new DBqueryBench.Query();
      //expect(instance).to.be.a(DBqueryBench.Query);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instance = new DBqueryBench.Query();
      //expect(instance).to.be();
    });

    it('should have the property ticketid (base name: "ticketid")', function() {
      // uncomment below and update the code to test the property ticketid
      //var instance = new DBqueryBench.Query();
      //expect(instance).to.be();
    });

    it('should have the property approvals (base name: "approvals")', function() {
      // uncomment below and update the code to test the property approvals
      //var instance = new DBqueryBench.Query();
      //expect(instance).to.be();
    });

    it('should have the property owner (base name: "owner")', function() {
      // uncomment below and update the code to test the property owner
      //var instance = new DBqueryBench.Query();
      //expect(instance).to.be();
    });

    it('should have the property query (base name: "query")', function() {
      // uncomment below and update the code to test the property query
      //var instance = new DBqueryBench.Query();
      //expect(instance).to.be();
    });

    it('should have the property servername (base name: "servername")', function() {
      // uncomment below and update the code to test the property servername
      //var instance = new DBqueryBench.Query();
      //expect(instance).to.be();
    });

    it('should have the property hasselect (base name: "hasselect")', function() {
      // uncomment below and update the code to test the property hasselect
      //var instance = new DBqueryBench.Query();
      //expect(instance).to.be();
    });

    it('should have the property hasdelete (base name: "hasdelete")', function() {
      // uncomment below and update the code to test the property hasdelete
      //var instance = new DBqueryBench.Query();
      //expect(instance).to.be();
    });

    it('should have the property hasinsert (base name: "hasinsert")', function() {
      // uncomment below and update the code to test the property hasinsert
      //var instance = new DBqueryBench.Query();
      //expect(instance).to.be();
    });

    it('should have the property hasupdate (base name: "hasupdate")', function() {
      // uncomment below and update the code to test the property hasupdate
      //var instance = new DBqueryBench.Query();
      //expect(instance).to.be();
    });

    it('should have the property hastransaction (base name: "hastransaction")', function() {
      // uncomment below and update the code to test the property hastransaction
      //var instance = new DBqueryBench.Query();
      //expect(instance).to.be();
    });

    it('should have the property status (base name: "status")', function() {
      // uncomment below and update the code to test the property status
      //var instance = new DBqueryBench.Query();
      //expect(instance).to.be();
    });

  });

}));
