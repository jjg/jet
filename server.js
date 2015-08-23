var https = require("https");
var crypto = require("crypto");
var url = require("url");
var RQ = require("rq.js");
var config = require("./config.js");
var log = require("jlog.js");
log.level = config.LOG_LEVEL;

log.message(log.INFO, "jet starting up...");

// todo: determine HTTP method

// todo: handle OPTIONS request

// todo: reject all other HTTP method requests

// todo: handle GET request

// todo: generate hash of relevant request parameters (URL, headers, etc.)

// todo: check cache for request hash

// todo: create new cache entry

// todo: relay request to origin server
//	do not forwart RANGE request parameters

// todo: initialize cache entry metadata using origin server response headers

// todo: update cache entry metadata:
//  measure inbound datarate from origin server
//  and throttle client response to maintain
//  buffer underrun

// todo: write bytes from origin server to client request

// todo: write bytes from origin server to cache

// todo: close origin server connection

// todo: close client request connection
