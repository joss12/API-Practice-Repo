const fs = require("fs");
const path = require("path");

//const file = "video.mp4";
const file = "/home/egs/Downloads/Part1.mp4";
const chunkSize = 1024 * 1024 * 5; // 5MB chunks

if (!fs.existsSync("chunks")) {
  fs.mkdirSync("chunks");
}

const data = fs.readFileSync(file);
const chunks = Math.ceil(data.length / chunkSize);

for (let i = 0; i < chunks; i++) {
  const start = i * chunkSize;
  const end = Math.min(start + chunkSize, data.length);
  const chunk = data.slice(start, end);

  fs.writeFileSync(path.join("chunks", `chunk_${i}.bin`), chunk);
}

console.log("chunks created:", chunks);
