const express = require("express");
const router = express.Router();
const fs = require("fs");
const path = require("path");
const { verifyToken, requireRole } = require("../middleware/auth");
const {
    submitMiddleware,
    mySubmit,
    deleteUser,
} = require("../middleware/submit");

const SUB_PATH = path.join(__dirname, "..", "..", "data.json");

//Admin routes
router.get("/", verifyToken, requireRole("admin"), (req, res) => {
    if (!fs.existsSync(SUB_PATH)) return res.json([]);
    const raw = fs.readFileSync(SUB_PATH, "utf-8");
    const data = raw.trim() ? JSON.parse(raw) : [];
    res.json(data);
});

router.post("/", submitMiddleware);
router.get("/my-submissions", verifyToken, requireRole("user"), mySubmit);
router.delete("/delete/:id", verifyToken, requireRole("admin"), deleteUser);
module.exports = router;
