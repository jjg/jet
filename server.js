var http = require("http");
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
			// check cache for request hash
			if(cache[req_hash]){
				log.message(log.INFO, "Cache HIT");
				// handle RANGE correctly
				if(req.headers.range){
					log.message(log.INFO, "RANGE request: " + req.headers.range);
					var range_string = req.headers.range;
					var range_parts = range_string.substring(range_string.indexOf("=")+1).split("-");
					var range_begin = range_parts[0];
					var range_end = cache[req_hash].content_length - 1;
					if(range_parts[1] && range_parts[1].length > 0){
						range_end = range_parts[1];
					}
					log.message(log.DEBUG, "Range begin: " + range_begin);
					log.message(log.DEBUG, "Range end: " + range_end);
					// set the headers and status
					res.setHeader("Content-Range", "bytes " + range_begin + "-" + range_end + "/" +  cache[req_hash].content_length);
					res.setHeader("Content-Length", cache[req_hash].content_length);
					res.setHeader("Accept-Ranges", "bytes");
					res.statusCode = 206;
					//  return only the requested bytes from the cache
					res.write(cache[req_hash].data.slice(range_begin, range_end));
					res.end();
					log.message(log.INFO, req.method + " request complete");
				} else {
					// return response headers from cache metadata
					res.setHeader("Content-Length", cache[req_hash].content_length);
					res.setHeader("Content-Type", cache[req_hash].content_type);
					// return data from cache
					res.write(cache[req_hash].data);
					res.end();
					log.message(log.INFO, req.method + " request complete");
				}
			} else {
				log.message(log.INFO, "Cache MISS");
				// create new cache entry
				cache[req_hash] = {};
				// relay request to origin server
				var origin_req_options = {
					hostname: config.ORIGIN_SERVER_HOST,	// set via config, alternatively req.hostname,
					port: config.ORIGIN_SERVER_PORT,		// set via config, alternatively req.port,
					path: req.url,
					//headers: req.headers,					// todo: include headers, just not RANGE
					method: "GET"
				};
				log.message(log.DEBUG, "Origin request options: " + JSON.stringify(origin_req_options));
				var origin_req = http.request(origin_req_options, function(origin_res){
					log.message(log.DEBUG, "Origin server request status: " + origin_res.statusCode);
					// todo: if origin status isn't good, end request
					// initialize cache entry metadata using origin server response headers
					cache[req_hash].content_length = origin_res.headers["content-length"];
					cache[req_hash].content_type = origin_res.headers["content-type"];
					cache[req_hash].data = new Buffer("");

					// write headers from origin to client
					res.setHeader("Content-Length", cache[req_hash].content_length);
					res.setHeader("Content-Type", cache[req_hash].content_type);
					origin_res.on("data", function(chunk){
						//log.message(log.DEBUG, "Received " + chunk.length + " bytes from origin server");	
						// write bytes from origin server to client request
						res.write(chunk);
						// append bytes from origin server to cache
						cache[req_hash].data = new Buffer.concat([cache[req_hash].data, chunk]);
						// todo: use a message to trigger res.write() vs tightly bound like above? 
					});
					origin_res.on("end", function(){
						log.message(log.INFO, "Origin server response ended");
						// todo: close origin server connection?
						// todo: close client request connection
						res.end();
						log.message(log.INFO, req.method + " request complete");
					});
					origin_res.on("error", function(error){
						log.message(log.ERROR, "Origin server response error: " + error);
						res.statusCode = 500;
						res.end();
						log.message(log.INFO, req.method + " request comlete");
					});
				});
				origin_req.on("error", function(error){
					log.message(log.ERROR, "Origin server request error: " + error);
				});
				origin_req.end();
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
