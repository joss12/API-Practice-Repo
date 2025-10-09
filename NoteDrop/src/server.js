import express from "express";
import cors from "cors";
import dotenv from "dotenv";
import { PrismaClient } from "@prisma/client";

//Import routes
import authRoutes from "./routes/auth.routes.js";
import noteRoutes from "./routes/note.routes.js";

// Loads .env variable
dotenv.config();

//Init app + Prisma client
const app = express();
const prisma = new PrismaClient();

//Middleware
app.use(cors());
app.use(express.json());

//Health check
app.get("/", (req, res) => {
    res.json({ message: "NoteDrop API is alive" });
});

//Api roputes
app.use("/api/auth/", authRoutes);
app.use("/api/notes/", noteRoutes);

export default app;
