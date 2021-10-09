const url = "http://127.0.0.1:1323/commands";

const rawResponse = await fetch(url, {
  method: "POST",
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
  body: JSON.stringify({
    command: `echo hello world`,
  }),
});
const content = await rawResponse.json();
console.log(content);
