const url = "https://902a-181-58-226-188.ngrok.io/commands";

const rawResponse = await fetch(url, {
  method: "POST",
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
  body: JSON.stringify({
    command: `shutdown now`,
  }),
});

const content = await rawResponse.json();
console.log(content);
