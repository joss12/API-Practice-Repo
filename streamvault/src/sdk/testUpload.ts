import { StreamVaultClient } from "./streamvaultClient";

async function main() {
  const client = new StreamVaultClient("http://localhost:3000");

  const result = await client.uploadFile("/home/egs/Downloads/Part1.mp4", {
    chunkSize: 5 * 1024 * 1024,
    onProgress: (p) => {
      const percent = ((p.uploadedBytes / p.totalBytes) * 100).toFixed(2);
      console.log(
        `Chunks ${p.uploadedChunks}/${p.totalChunks} - ${p.uploadedBytes}/${p.totalBytes} bytes (${percent}%)`,
      );
    },
  });

  console.log("Upload complete:", result);
}

main().catch((err) => {
  console.log(err);
  process.exit(1);
});
