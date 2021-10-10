import { Router } from "express";
import axios from "axios";
import { urlip } from "../../url.json";

const router = Router();

interface CuloUrl {
  Url: string;
}

router.get("/ip", (_, res: any) => {
  axios
    .get(urlip)
    .then((r) => {
      res.json({ out: r.data });
    })
    .catch((err) => (err ? res.send("error") : null));
  axios
    .get("http://127.0.0.1:1323/ngrok")
    .then((r) => {
      console.log(r.data);
    })
    .catch((err) => (err ? console.log("error") : null));
});

module.exports = router;
