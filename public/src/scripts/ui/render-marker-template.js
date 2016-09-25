module.exports = function(location) {
  const walkTime = Math.round((location.distance * 2) / 60);
  let distance = `Less than 1 minute away.`;
  if (walkTime > 1) {
    distance = `${walkTime} minutes away.`;
  }
  const button = (location.can_play) ? `<div><button type="button" class="btn" onclick="btnEvents.handleModalOpen(${location.id})">Play!</button></div>` : `<div class="color-danger">Location out of range.</div>`;
  return `<div id="on-map-maker" class="on-map-marker">
    <h4>${location.title}</h4>
    <div><small>${distance}</small></div>
    <div><small>${location.wallet} Credits Jackpot</small></div>
    <div><small>Odds: <em>1 in 10</em></small></div>
    ${button}
  </div>`;
};
