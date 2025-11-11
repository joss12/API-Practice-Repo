import express from "express";
import dotenv from "dotenv";
import appRoute from "./routes/router";

dotenv.config();
const app = express();
app.use(express.json());

app.use("/api", appRoute);

export default app;
