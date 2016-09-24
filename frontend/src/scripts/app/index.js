function loadMap(mapEl, locations) {
  return new Promise(function(resolve, reject) {
    var map = new google.maps.Map(mapEl, {
      center: { lat: locations[0].latitude, lng: locations[0].longitude },
      zoom: 15,
      scrollwheel: false
    });

    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(function(position) {
        var pos = {
          lat: position.coords.latitude,
          lng: position.coords.longitude
        };

        var myMarker = new google.maps.Marker({
          position: pos,
          icon: '/img/marker.png',
          map: map
        });

        var myLocationInfoWindow = new google.maps.InfoWindow({});

        google.maps.event.addListener(myMarker, 'click', (function (m, i) {
          return function () {
            myLocationInfoWindow.setContent("Your Location");
            myLocationInfoWindow.open(map, m);
            map.panTo(m.getPosition());
            map.setZoom(17);
          };
        })(myMarker, 100));

        map.setCenter(pos);
        map.panTo(myMarker.getPosition());
        map.setZoom(17);
      }, function() {
        handleLocationError(true, map.getCenter());
      });
    } else {
      // Browser doesn't support Geolocation
      handleLocationError(false, map.getCenter());
    }

    function handleLocationError(browserHasGeolocation, pos) {
      alert(browserHasGeolocation ?
        'Error: The Geolocation service failed.' :
        'Error: Your browser doesn\'t support geolocation.');
    }

    var infowindow = new google.maps.InfoWindow({});
    var marker, i;

    for (i = 0; i < locations.length; i++) {
      marker = new google.maps.Marker({
        position: new google.maps.LatLng(locations[i].latitude, locations[i].longitude),
        map: map
      });

      google.maps.event.addListener(marker, 'click', (function (marker, i) {
        return function () {
          infowindow.setContent(locations[i].title);
          infowindow.open(map, marker);
          map.panTo(marker.getPosition());
          map.setZoom(17);
        };
      })(marker, i));
    }

    resolve(map);
  });
}

module.exports = loadMap;