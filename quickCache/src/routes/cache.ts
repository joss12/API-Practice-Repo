import express from "express";
import * as cache from "../services/cacheService";

const router = express.Router();

router.post("/", (req, res) => {
  const { key, value, ttl } = req.body;
  if (!key || value === undefined)
    return res.status(400).json({ message: "key and value required" });
  cache.set(key, value, ttl);
  res.status(201).json({ message: "stored" });
});

router.get("/:key", (req, res) => {
  const value = cache.get(req.params.key);
  if (value === null)
    return res.status(400).json({ message: "Not found or expired" });
  res.json({ key: req.params.key, value });
});

router.delete("/:key", (req, res) => {
  cache.del(req.params.key);
  res.json({ message: "Delete" });
});

router.get("/stats/info", (req, res) => {
  res.json(cache.stats());
});
export default router;
