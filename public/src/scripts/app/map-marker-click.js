module.exports = function (infowindow, location, map) {
  map.panTo(this.getPosition());
  infowindow.setContent(location.title);
  infowindow.open(map, this);
};