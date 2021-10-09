const url = "http://localhost:1323/commands";

const rawResponse = await fetch(url, {
  method: "POST",
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
  body: JSON.stringify({
    command: `cowsay cum`,
  }),
});

const content = await rawResponse.json();
console.log(content);
