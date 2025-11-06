import { Router } from "express";
import { graphStore } from "../graph/GraphStore";

const router = Router();

router.get("/", (_req, res) => {
  res.json(graphStore.getMap());
});

export default router;
