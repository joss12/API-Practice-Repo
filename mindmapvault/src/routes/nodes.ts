import { Router } from "express";
import { graphStore } from "../graph/GraphStore";

const router = Router();

router.post("/", (req, res) => {
  try {
    const { id, label } = req.body;
    graphStore.addNode({ id, label });
    res.status(201).json({ message: "Node added." });
  } catch (err: any) {
    res.status(400).json({ error: err.message });
  }
});

export default router;
