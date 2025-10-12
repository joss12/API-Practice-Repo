const app = require("./server");
const http = require("http");

const PORT = process.env.PORT || 3000;

const server = http.createServer(app);

server
    .listen(PORT, () => {
        console.log(`->Server started on http://localhost:${PORT}`);
    })
    .on("error", (err) => {
        console.error(`Failed to start: ${err.message}`);
        process.exit(1);
    });

process.on("SIGINT", () => {
    console.log("\n Shutting down...");
    server.close(() => process.exit(0));
});
