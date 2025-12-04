import fs from "fs";
import path from "path";
import axios, { AxiosInstance } from "axios";
import FormData from "form-data";

export interface UploadProgress {
  uploadedBytes: number;
  totalBytes: number;
  uploadedChunks: number;
  totalChunks: number;
}

export interface UploadOptions {
  /**
   * Chunk size in bytes. Default: 5MB.
   */
  chunkSize?: number;
  /**
   * Remote filename to use. Default: basename(filePath).
   */
  filename?: string;
  /**
   * Progress callback fired after each successfully uploaded chunk.
   */
  onProgress?: (progress: UploadProgress) => void;
}

export interface UploadCompleteResult {
  filePath: string;
  md5: string;
}

interface ResumeStatusResponse {
  uploaded: number[];
  count: number;
}

interface ChunkUploadedResponse {
  message: "Chunk uploaded";
  uploaded: number[];
}

interface UploadFinishedResponse {
  message: "Upload complete";
  file: string;
  md5: string;
}

type ChunkResponse = ChunkUploadedResponse | UploadFinishedResponse;

/**
 * StreamVaultClient
 * TypeScript SDK for the StreamVault chunked upload API.
 *
 * Assumes the backend API exposes:
 *   POST /upload/chunk   (multipart/form-data)
 *   GET  /upload/resume/:filename
 */
export class StreamVaultClient {
  private http: AxiosInstance;
  private baseUrl: string;

  constructor(baseUrl: string) {
    this.baseUrl = baseUrl.replace(/\/+$/, ""); // remove trailing slash
    this.http = axios.create({
      baseURL: this.baseUrl,
      timeout: 60_000,
    });
  }

  /**
   * Returns the array of chunk indexes already uploaded for this filename.
   */
  async getResumeStatus(filename: string): Promise<number[]> {
    const res = await this.http.get<ResumeStatusResponse>(
      `/upload/resume/${encodeURIComponent(filename)}`,
    );
    // ensure sorted & unique
    return Array.from(new Set(res.data.uploaded)).sort((a, b) => a - b);
  }

  /**
   * Uploads a local file in chunks to the StreamVault backend.
   * Efficient: reads file chunk-by-chunk using fs, not loading the whole file into memory.
   */
  async uploadFile(
    filePath: string,
    options: UploadOptions = {},
  ): Promise<UploadCompleteResult> {
    const chunkSize = options.chunkSize ?? 5 * 1024 * 1024; // 5MB default
    const filename = options.filename ?? path.basename(filePath);

    const stat = await fs.promises.stat(filePath);
    const totalBytes = stat.size;

    if (totalBytes === 0) {
      throw new Error(`File is empty: ${filePath}`);
    }

    const totalChunks = Math.ceil(totalBytes / chunkSize);

    // Check which chunks already exist on the server (resume support)
    const alreadyUploaded = await this.getResumeStatus(filename);
    const uploadedSet = new Set(alreadyUploaded);

    let uploadedBytes = alreadyUploaded.length * chunkSize;
    if (uploadedBytes > totalBytes) {
      uploadedBytes = totalBytes; // clamp
    }

    let uploadedChunks = alreadyUploaded.length;
    let lastResult: UploadCompleteResult | null = null;

    const fd = await fs.promises.open(filePath, "r");

    try {
      for (let index = 0; index < totalChunks; index++) {
        if (uploadedSet.has(index)) {
          // Skip chunks that the server already has
          this.emitProgress(options, {
            uploadedBytes,
            totalBytes,
            uploadedChunks,
            totalChunks,
          });
          continue;
        }

        const start = index * chunkSize;
        const remaining = totalBytes - start;
        const currentChunkSize = Math.min(chunkSize, remaining);

        const buffer = Buffer.alloc(currentChunkSize);
        const { bytesRead } = await fd.read(buffer, 0, currentChunkSize, start);

        if (bytesRead !== currentChunkSize) {
          throw new Error(
            `Unexpected bytesRead: expected ${currentChunkSize}, got ${bytesRead}`,
          );
        }

        const chunkResponse = await this.uploadChunk(
          filename,
          index,
          totalChunks,
          buffer,
        );

        uploadedBytes += currentChunkSize;
        uploadedChunks += 1;

        this.emitProgress(options, {
          uploadedBytes,
          totalBytes,
          uploadedChunks,
          totalChunks,
        });

        if (chunkResponse.message === "Upload complete") {
          lastResult = {
            filePath: chunkResponse.file,
            md5: chunkResponse.md5,
          };
        }
      }
    } finally {
      await fd.close();
    }

    if (!lastResult) {
      throw new Error(
        "Upload finished but server did not return completion payload",
      );
    }

    return lastResult;
  }

  /**
   * Uploads a single chunk buffer.
   */
  private async uploadChunk(
    filename: string,
    index: number,
    totalChunks: number,
    chunkBuffer: Buffer,
  ): Promise<ChunkResponse> {
    const form = new FormData();
    // Name the part file purely for diagnostics; backend only cares about field name.
    form.append("chunk", chunkBuffer, {
      filename: `${filename}.part${index}`,
    });
    form.append("filename", filename);
    form.append("index", String(index));
    form.append("total", String(totalChunks));

    try {
      const res = await this.http.post<ChunkResponse>("/upload/chunk", form, {
        headers: form.getHeaders(),
        maxContentLength: Infinity,
        maxBodyLength: Infinity,
      });

      return res.data;
    } catch (err: any) {
      if (err.response) {
        // Server responded with non-2xx
        const status = err.response.status;
        const data = err.response.data;
        throw new Error(
          `Upload failed for chunk ${index}: HTTP ${status} - ${JSON.stringify(
            data,
          )}`,
        );
      }
      throw new Error(`Upload failed for chunk ${index}: ${String(err)}`);
    }
  }

  private emitProgress(options: UploadOptions, progress: UploadProgress): void {
    if (typeof options.onProgress === "function") {
      options.onProgress(progress);
    }
  }
}
