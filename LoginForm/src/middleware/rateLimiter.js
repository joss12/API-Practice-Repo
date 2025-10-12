const rateLimit = require("express-rate-limit");

module.exports = rateLimit({
    windowMs: 60 * 1000, //1 minute
    max: 5, //5 request per minutes per IP
    message: "Too many submission from this IP, please try again later",
});
