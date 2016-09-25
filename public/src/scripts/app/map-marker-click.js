const storage = require('../app/storage');
const renderMarkerTemplate = require('../ui/render-marker-template');

module.exports = function (infowindow, location, map) {
  var pos = storage.getLocation();
  fetch(`https://geo.ohdoylerules.com/api/location/${location.id}?lat=${pos.lat}&long=${pos.lng}`)
  .then((response) => {
    return response.json();
  }).then((res) => {
    map.panTo(this.getPosition());
    infowindow.setContent(renderMarkerTemplate(res));
    infowindow.open(map, this);
  }).catch((err) => {
    console.error(err);
  });
};