const json = fetch("/ip");
const os = fetch("/ip/os");
const command = fetch("/command-sent");

json
  .then((r) => {
    return r.json();
  })
  .then((jsonData) => {
    document.getElementById(
      "ip"
    )!.innerHTML = `ip: ${jsonData.out.query}\ncountry: ${jsonData.out.country}\ncity: ${jsonData.out.city}`;
  });

os.then((r) => {
  return r.json();
}).then((jsonData) => {
  document.getElementById(
    "os"
  )!.innerHTML = `Operative System: ${jsonData.out}`;
});

interface api {
  out: string;
}
function sendAndReceive() {
  const command: HTMLInputElement = document.getElementById(
    "command"
  ) as HTMLInputElement;
  fetch("/command-sent", {
    method: "POST",
    body: JSON.stringify({ command: command.value }),
    headers: {
      "content-type": "application/json",
    },
  })
    .then((r) => r.json())
    .then(
      (d: api) =>
        (document.getElementById("o")!.innerHTML = d.out.replace(
          /</g,
          "<span><</span>"
        ))
    );
}
