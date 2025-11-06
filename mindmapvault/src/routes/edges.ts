import { Router } from "express";
import { graphStore } from "../graph/GraphStore";

const router = Router();

router.post("/", (req, res) => {
  try {
    const { from, to } = req.body;
    graphStore.addEdge({ from, to });
    res.status(201).json({ message: "Edge added." });
  } catch (err: any) {
    res.status(400).json({ error: err.message });
  }
});

export default router;
