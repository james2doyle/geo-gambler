module.exports = function (infowindow, location, map, marker) {
  infowindow.setContent(location.title);
  infowindow.open(map, marker);
  map.panTo(marker.getPosition());
  map.setZoom(17);
};