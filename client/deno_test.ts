const data = JSON.parse(Deno.readTextFileSync('../url.json'));

const rawResponse = await fetch(data.url, {
  method: "POST",
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
  body: JSON.stringify({
    command: `neofetch`,
  }),
});

const content = await rawResponse.json();
console.log(content);
