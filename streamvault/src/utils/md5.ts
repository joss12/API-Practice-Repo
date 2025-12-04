import crypto from "crypto";

export function md5OfFile(buffer: Buffer): string {
  return crypto.createHash("md5").update(buffer).digest("hex");
}
