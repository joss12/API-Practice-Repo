import fs from "fs";
import path, { parse } from "path";
import fse from "fs-extra";
import { env } from "../config/env";

export class UploadService {
  private baseDir = env.UPLOAD_DIR;

  constructor() {
    fse.ensureDirSync(this.baseDir);
  }

  /**
   * Stores one chunk oon disk.
   */
  async saveChunk(filename: string, chunkIndex: number, chunkBuffer: Buffer) {
    const chunkDir = path.join(this.baseDir, filename);

    await fse.ensureDirSync(chunkDir);

    const chunkPath = path.join(chunkDir, `${chunkIndex}.part`);
    await fse.writeFile(chunkPath, chunkBuffer);
  }

  /**
   *Chunks how many chunks currently exist.
   */

  async getUploadedChunks(filename: string): Promise<number[]> {
    const chunkDir = path.join(this.baseDir, filename);

    if (!fs.existsSync(chunkDir)) return [];

    return fs
      .readdirSync(chunkDir)
      .map((f) => parseInt(f.replace(".part", ""), 10))
      .filter((n) => !isNaN(n))
      .sort((a, b) => a - b);
  }

  /**
   *Merges chunks into a final file.
   */

  async mergeChunks(filename: string, totalChunks: number): Promise<string> {
    const chunkDir = path.join(this.baseDir, filename);
    const finalPath = path.join(this.baseDir, `${filename}.final`);

    const writeStream = fs.createWriteStream(finalPath);

    for (let i = 0; i < totalChunks; i++) {
      const partPath = path.join(chunkDir, `${i}.ppart`);

      if (!fs.existsSync(partPath)) {
        throw new Error(`Missiing chunk: ${i}`);
      }

      const data = fs.readFileSync(partPath);
      writeStream.write(data);
    }

    writeStream.end();

    return finalPath;
  }
}
