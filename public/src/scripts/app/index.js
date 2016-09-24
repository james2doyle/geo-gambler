function loadMap(locations) {
  return new Promise((resolve, reject) => {
    const mapEl = document.getElementById('map');
    const map = new google.maps.Map(mapEl, {
      // center: {
      //   lat: locations[0].latitude,
      //   lng: locations[0].longitude
      // },
      zoom: 15,
      scrollwheel: false
    });

    const infowindow = new google.maps.InfoWindow({});

    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(function(position) {
        const pos = {
          lat: position.coords.latitude,
          lng: position.coords.longitude
        };

        const myMarker = new google.maps.Marker({
          position: pos,
          icon: '/img/marker.svg',
          map: map
        });

        google.maps.event.addListener(myMarker, 'click', (function (m, i) {
          return function () {
            infowindow.setContent("Your Location");
            infowindow.open(map, m);
            map.panTo(m.getPosition());
            map.setZoom(17);
          };
        })(myMarker, 100));

        map.setCenter(pos);
        map.panTo(myMarker.getPosition());
        map.setZoom(17);
        document.body.classList.add('my-location-loaded');
      }, function() {
        handleLocationError(true, map.getCenter());
      });
    } else {
      // Browser doesn't support Geolocation
      handleLocationError(false, map.getCenter());
    }

    function handleLocationError(browserHasGeolocation, pos) {
      reject(browserHasGeolocation ?
        'Error: The Geolocation service failed.' :
        'Error: Your browser doesn\'t support geolocation.');
    }

    for (let i = 0; i < locations.length; i++) {
      let marker = new google.maps.Marker({
        position: new google.maps.LatLng(locations[i].latitude, locations[i].longitude),
        icon: '/img/dollar.svg',
        map: map
      }).addListener('click', function() {
        map.panTo(this.getPosition());
        infowindow.setContent(locations[i].title);
        infowindow.open(map, this);
      });
    }

    resolve(map);
  });
}

module.exports = loadMap;