import { Router } from "express";
import axios from "axios";
import { url } from "../../../url.json";

const router = Router();

router.get("/ip/os", (_, res: any) => {
  axios
    .get(url)
    .then((r) => {
      let OsUrl = r.data[0]["Urlos"]; // get the ngrok /ip/os url
      axios
        .get(OsUrl)
        .then((r) => {
          res.json({ out: r.data }); // send the os
        })
        .catch((err) => (err ? res.send("error") : null));
    })
    .catch((err) => (err ? res.send("error") : null));
});

module.exports = router;
