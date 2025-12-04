import { Router } from "express";
import multer from "multer";
import { UploadController } from "../controllers/upload.controller";

const router = Router();
const upload = multer(); //memroy storage for streaming

router.post("/chunk", upload.single("chunk"), UploadController.uploadChunk);
router.get("/resume/:filename", UploadController.resumeStatus);

export default router;
