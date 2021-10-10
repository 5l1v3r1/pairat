import { Router } from "express";
import axios from "axios";
import { urlos } from"../../../url.json"

const router = Router();

router.get("/ip/os", (_, res: any) => {
  axios
    .get(urlos)
    .then((r) => {
      res.json({ out: r.data });
    })
    .catch((err) => (err ? res.send("error") : null));
});

module.exports = router;
