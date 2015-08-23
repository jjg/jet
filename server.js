var http = require("http");
var https = require("https");
var crypto = require("crypto");
var url = require("url");
var RQ = require("rq.js");
var config = require("./config.js");
var log = require("jlog.js");
log.level = config.LOG_LEVEL;

http.createServer(function(req, res){

	// todo: set response headers necissary for CORS support
	res.setHeader("Access-Control-Allow-Methods","GET,OPTIONS");
	res.setHeader("Access-Control-Allow-Headers","*");
	res.setHeader("Access-Control-Allow-Origin","*");

	// determine HTTP method
	log.message(log.INFO, "Processing " + req.method + " request");
	switch(req.method){
		case "OPTIONS":
			// handle OPTIONS request
			res.writeHead();
			res.end();
			break;
		case "GET":
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
			break;
		default:
			// todo: reject all other HTTP method requests
			break;
	}
}).listen(config.SERVER_PORT);

log.message(log.INFO, "Starting jet on port " + config.SERVER_PORT + ".");
