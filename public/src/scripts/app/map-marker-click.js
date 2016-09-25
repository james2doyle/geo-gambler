function markerTemplate(location) {
  return `<div class="on-map-marker">
    <h4>${location.title}</h4>
    <div><small>3M away</small></div>
    <div><small>${location.wallet} Credit Jackpot</small></div>
    <div><small>Expires in 22 mins</small></div>
    <div><button type="button" onclick="btnEvents.handlePlayClick(${location.id})">Play!</button></div>
  </div>`;
}

module.exports = function (infowindow, location, map) {
  map.panTo(this.getPosition());
  infowindow.setContent(markerTemplate(location));
  infowindow.open(map, this);
};