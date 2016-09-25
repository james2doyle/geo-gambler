const storage = require('../app/storage');

function markerTemplate(location) {
  const walkTime = Math.round((location.distance * 2) / 60);
  let distance = `Less than 1 minute away.`;
  if (walkTime > 1) {
    distance = `${walkTime} minutes away.`;
  }
  return `<div class="on-map-marker">
    <h4>${location.title}</h4>
    <div><small>${distance}</small></div>
    <div><small>${location.wallet} Credits Jackpot</small></div>
    <div><small>Odds: <em>1 in 10</em></small></div>
    <div>
      <input type="number" name="number" id="number" value="1" min="1" max="10" />
      <button type="button" onclick="btnEvents.handlePlayClick(${location.id})">Play!</button>
    </div>
  </div>`;
}

module.exports = function (infowindow, location, map) {
  var pos = storage.getLocation();
  fetch(`https://geo.ohdoylerules.com/api/location/${location.id}?lat=${pos.lat}&long=${pos.lng}`)
  .then((response) => {
    return response.json();
  }).then((res) => {
    map.panTo(this.getPosition());
    infowindow.setContent(markerTemplate(res));
    infowindow.open(map, this);
  }).catch((err) => {
    console.error(err);
  });
};