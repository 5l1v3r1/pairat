import { Router } from "express";
import axios from "axios";
import { urlip } from "../../url.json";

const router = Router();

router.get("/ip", (_, res: any) => {
  axios
    .get(urlip)
    .then((r) => {
      res.json({ out: r.data });
    })
    .catch((err) => (err ? res.send("error") : null));
});

module.exports = router;
