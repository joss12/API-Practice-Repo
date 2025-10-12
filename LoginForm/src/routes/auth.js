const express = require("express");
const router = express.Router();
const auth = require("../controllers/authCtrl");

router.post("/register", auth.Register);
router.post("/login", auth.Login);

module.exports = router;
