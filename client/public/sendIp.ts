const json = fetch("/ip");

json
  .then((r) => {
    return r.json();
  })
  .then((jsonData) => {
    document.getElementById(
      "ip"
    )!.innerHTML = `ip: ${jsonData.out.query}\ncountry: ${jsonData.out.country}\ncity: ${jsonData.out.city}`;
  });
