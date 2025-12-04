import app from "./app";
import { env } from "./config/env";
import fse from "fs-extra";

fse.ensureDirSync(env.UPLOAD_DIR);

app.listen(env.PORT, () => {
  console.log(`StreamVault running on http://localhost:${env.PORT}`);
});
