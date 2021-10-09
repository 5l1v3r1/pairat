const url = "http://127.0.0.1:1323/command"
const Router = require('express')
const axios = require('axios')

const router = Router();

router.post("/command-sent", async (req, res) => {
  const  command  = req.body;
  console.log(`input: ${command}`);

  axios
    .post(url, {
      command: command,
    })
    .then((r) => {
        console.log(`output: ${r.data}`);
        res.send(r.data)
    }).catch((err) => err ? res.send("error") : null);
});

module.exports = router;
