import express from "express";
import cors from "cors";
import uploadRoutes from "./routes/upload.routes";
import { errorHandler } from "./middleware/errorHandler";

const app = express();

app.use(cors());
app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.use("/upload", uploadRoutes);
app.use(errorHandler);

export default app;
