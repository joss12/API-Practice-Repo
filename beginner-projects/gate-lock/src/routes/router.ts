import express from "express";
import { ipWhiteList } from "../middleware/ip";
import { verifyToken } from "../middleware/jwt";
import jwt from "jsonwebtoken";

const router = express.Router();

router.get("/secure", ipWhiteList, verifyToken, (req, res) => {
  res.json({ message: "Access Granted", user: (req as any).user });
});

router.post("/login", (_req, res) => {
  const token = jwt.sign(
    { userId: "abc123" },
    process.env.JWT_SECRET as string,
    { expiresIn: "1h" },
  );
  res.json({ token });
});

export default router;
