import { Request, Response, NextFunction } from "express";
import { UploadService } from "../services/upload.service";
import { md5OfFile } from "../utils/md5";
import fs from "fs";

const service = new UploadService();

export class UploadController {
  static async uploadChunk(req: Request, res: Response, next: NextFunction) {
    try {
      const { filename, index, total } = req.body;

      if (!req.file) {
        return res.status(400).json({ error: "Missing chunk" });
      }

      const chunkIndex = parseInt(index);
      const totalChunks = parseInt(total);

      await service.saveChunk(filename, chunkIndex, req.file.buffer);

      const uploaded = await service.getUploadedChunks(filename);

      const isCompleted = uploaded.length === totalChunks;

      if (isCompleted) {
        const finalPath = await service.mergeChunks(filename, totalChunks);

        const finalBuffer = fs.readFileSync(finalPath);
        const chucksum = md5OfFile(finalBuffer);

        return res.json({
          message: "Upload complete",
          file: finalPath,
          md5: chucksum,
        });
      }
      return res.json({
        message: "Chunk uploaded",
        uploaded,
      });
    } catch (err) {
      next(err);
    }
  }

  static async resumeStatus(req: Request, res: Response, next: NextFunction) {
    try {
      const { filename } = req.params;

      const uploaded = await service.getUploadedChunks(filename);

      return res.json({
        uploaded,
        count: uploaded.length,
      });
    } catch (err) {
      next(err);
    }
  }
}
