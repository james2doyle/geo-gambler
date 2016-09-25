const storage = require('../app/storage');

function markerTemplate(location) {
  return `<div class="on-map-marker">
    <h4>${location.title}</h4>
    <div><small>${location.distance}m away</small></div>
    <div><small>${location.wallet} Credit Jackpot</small></div>
    <div><small>Expires in 22 mins</small></div>
    <div><button type="button" onclick="btnEvents.handlePlayClick(${location.id})">Play!</button></div>
  </div>`;
}

module.exports = function (infowindow, location, map) {
  var pos = storage.getLocation();
  fetch(`https://geo.ohdoylerules.com/api/location/${location.id}?lat=${pos.lat}&long=${pos.lng}`)
  .then((response) => {
    return response.json();
  }).then((res) => {
    console.log(res);
    map.panTo(this.getPosition());
    infowindow.setContent(markerTemplate(location));
    infowindow.open(map, this);
  }).catch((err) => {
    console.error(err);
  });
};