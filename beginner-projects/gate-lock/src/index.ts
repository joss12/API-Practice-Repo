import app from "./server";

const PORT = process.env.PORT;

app.listen(PORT, () => {
  console.log(`Server started on http://localhost:${PORT}`);
});
