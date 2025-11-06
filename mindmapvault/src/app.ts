import express from "express";
import nodesRoutes from "./routes/nodes";
import edgesRoute from "./routes/edges";
import mapRoute from "./routes/map";
import viewRoute from "./routes/view";

import path from "path";

const app = express();
app.use(express.json());

app.use("/nodes", nodesRoutes);
app.use("/edges", edgesRoute);
app.use("/map", mapRoute);
app.use("/view", viewRoute);

app.use(express.static(path.join(__dirname, "../public")));

export default app;
