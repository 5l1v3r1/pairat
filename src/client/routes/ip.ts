import { Router } from "express";
import axios from "axios";
import { url } from "../../../url.json";

const router = Router();

router.get("/ip", (_, res: any) => {
  axios
    .get(url)
    .then((r) => {
      let ipUrl = r.data[0]["Urlip"]; // get the ngrok /ip url
      axios
        .get(ipUrl)
        .then((r) => {
          res.json({ out: r.data }); // send the ip
        })
        .catch((err) => (err ? res.send("error") : null));
    })
    .catch((err) => (err ? res.send("error") : null));
});

module.exports = router;
