import { Router } from "express";
import path from "path";

const router = Router();

router.get("/", (_req, res) => {
  res.sendFile(path.resolve(__dirname, "../../public/graph.html"));
});

export default router;
