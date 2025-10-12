const express = require("express");
const morgan = require("morgan");
const helmet = require("helmet");
const cors = require("cors");
const path = require("path");

const rateLimiter = require("./middleware/rateLimiter");
const submitRoute = require("./routes/submit");
const authRoutes = require("./routes/auth");
const submissionRoutes = require("./routes/submit");

require("dotenv").config();

const app = express();

app.use(express.json());
app.use(express.static(path.join(__dirname, "..", "public")));
app.use(morgan("dev"));
app.use(helmet());
app.use(cors());

app.use("/auth", authRoutes);

app.use(rateLimiter);

app.use("/submit", submitRoute);
app.use("/submissions", submissionRoutes);

module.exports = app;
