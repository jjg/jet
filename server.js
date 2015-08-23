var http = require("http");
var https = require("https");
var crypto = require("crypto");
var url = require("url");
var RQ = require("rq.js");
var config = require("./config.js");
var log = require("jlog.js");
log.level = config.LOG_LEVEL;

// array to hold in-memory cache
var cache = {};

http.createServer(function(req, res){

	// set response headers necissary for CORS support
	res.setHeader("Access-Control-Allow-Methods","GET,OPTIONS");
	res.setHeader("Access-Control-Allow-Headers","*");
	res.setHeader("Access-Control-Allow-Origin","*");

	// determine HTTP method
	log.message(log.INFO, "Processing " + req.method + " request");
	switch(req.method){
		case "OPTIONS":
			// handle OPTIONS request
			res.end();
			break;
		case "GET":
			// handle GET request

			// generate hash of relevant request parameters (URL, headers, etc.)
			var shasum = crypto.createHash("sha1");
			shasum.update(req.url 
				+ req.headers["x-mms-src"]
				+ req.headers["x-mms-dst"]
				+ req.headers["x-mms-fmt"]
				+ req.headers["x-mms-meta"]
				+ req.headers["x-mms-auth"]
				+ req.headers["x-mms-scale-width"]
				+ req.headers["x-mms-crop"]
			);
			var req_hash = shasum.digest("hex");
			log.message(log.DEBUG, "Request hash: " + req_hash);

			// todo: check cache for request hash
			if(cache[req_hash]){
				log.message(log.INFO, "Request found in cache");
				// todo: return response headers from cache metadata
				// todo: return data from cache
				res.end(cache[req_hash].data);
				log.message(log.INFO, req.method + " request complete");
			} else {
				// todo: create new cache entry
				cache[req_hash] = {};

				// todo: relay request to origin server
				//	do not forwart RANGE request parameters

				// todo: initialize cache entry metadata using origin server response headers

				// todo: update cache entry metadata:
				//  measure inbound datarate from origin server
				//  and throttle client response to maintain
				//  buffer underrun
	
				// todo: write bytes from origin server to client request
				res.write("foo");
	
				// todo: write bytes from origin server to cache
				cache[req_hash].data = "foo";
	
				// todo: close origin server connection
	
				// todo: close client request connection
				res.end();
				log.message(log.INFO, req.method + " request complete");
			}
			break;
		default:
			// reject all other HTTP method requests
			res.statusCode = 405;	// METHOD NOT ALLOWED
			res.end();
			break;
	}
}).listen(config.SERVER_PORT);

log.message(log.INFO, "Starting jet on port " + config.SERVER_PORT + ".");
