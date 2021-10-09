import { Router } from "express";
import axios from "axios";
import { url } from "../../url.json"

const router = Router();

router.post("/command-sent", (req: any, res: any) => {
  const { command } = req.body;
  console.log(`input: ${command}`);
  axios
    .post(url, {
      command: command,
    })
    .then((r) => {
      console.log(`output: ${r.data}`);
      res.send(r.data);
    })
    .catch((err) => (err ? res.send("error") : null));
});

module.exports = router;
