import express from "express";
import {
    loginUser,
    promoteUserToAdmin,
    registerUser,
} from "../controllers/auth.controller.js";
import { protect } from "../middlewares/auth.js";
import { authorizationRoles } from "../middlewares/authorizeRoles.js";

const router = express.Router();

// ✅ Add this just for debugging from browser
router.get("/auth", (req, res) => {
    res.status(200).json({ message: "Auth Route is Working" });
});

// POST api/auth/register
router.post("/register", registerUser);
router.post("/login", loginUser);

//Test routes for protection
router.get("/me", protect, (req, res) => {
    res.json({ message: "You are authenticated", user: req.user });
});

//Test route for roles
router.delete(
    "/user/:id",
    protect,
    authorizationRoles("ADMIN"),
    async (req, res) => {
        //only admins reach here
        res.json({ message: `Delete user ${req.params.id}` });
    },
);

//To promote user to admin.
router.put(
    "/promote/:id",
    protect,
    authorizationRoles("ADMIN"),
    promoteUserToAdmin,
);

export default router;
