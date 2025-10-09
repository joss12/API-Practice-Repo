import dotenv from "dotenv";
import app from "./server.js";

dotenv.config();
//Start the Server

console.log("booting server...");
const PORT = process.env.PORT || 3600;

app.use((err, req, res, next) => {
    console.error("->Server Error:", err);
    res.status(500).json({ error: "Internal Server Error" });
});

app.listen(PORT, () => {
    console.log(`->Server started on http://${PORT}`);
});
