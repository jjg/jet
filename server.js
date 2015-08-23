var config = require("./config.js");
var log = require("./jlog.js");
log.level = config.LOG_LEVEL;

log.message(log.INFO, "jet starting up...");
