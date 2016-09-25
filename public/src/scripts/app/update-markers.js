const mapMarkerClick = require('./map-marker-click');

module.exports = function(data) {
  return new Promise((resolve, reject) => {
    // add all markers to the map and add click listeners
    data.locations.forEach((location) => {
      let marker = new google.maps.Marker({
        position: new google.maps.LatLng(location.latitude, location.longitude),
        icon: '/img/dollar.svg',
        map: data.map
      }).addListener('click', function() {
        // this is a marker
        mapMarkerClick.call(this, infowindow, location, data.map);
      });
    });

    resolve(data);
  });
};
