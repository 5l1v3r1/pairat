import express = require("express");
import { join } from "path";
const app = express();
const PORT = process.env.PORT || 8000;
app.use(express.urlencoded({ extended: false }));
app.use(express.json());

app.use(require("./routes/command"));
app.use(require("./routes/ip"));
app.use(require("./routes/os"));

app.use(express.static(join(__dirname, "public")));
app.listen(PORT, () => console.log(`client on http://127.0.0.1:${PORT}`));
