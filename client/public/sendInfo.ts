const json = fetch("/ip");
const os = fetch("/ip/os");

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
  document.getElementById("os")!.innerHTML = `Operative System: ${jsonData.out}`;
});
