function loadMap(locations) {
  return new Promise((resolve, reject) => {
    const mapEl = document.getElementById('map');
    const map = new google.maps.Map(mapEl, {
      zoom: 17,
      scrollwheel: false
    });

    const infowindow = new google.maps.InfoWindow({});

    const getUserLocation = require('../player/location');

    getUserLocation()
    .then((pos) => {
      const myMarker = new google.maps.Marker({
        position: pos,
        icon: '/img/marker.svg',
        map: map
      })

      myMarker.addListener('click', function() {
        infowindow.setContent("Your Location");
        infowindow.open(map, this);
        map.panTo(this.getPosition());
        map.setZoom(17);
      });

      map.panTo(myMarker.getPosition());
      document.body.classList.add('my-location-loaded');
    })
    .catch((err) => {
      console.error(err);
      reject(err);
    });

    resolve({
      map,
      locations,
      infowindow
    });
  });
}

module.exports = loadMap;