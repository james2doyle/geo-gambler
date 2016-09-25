const storage = require('../app/storage');

function getLocation(id) {
  var pos = storage.getLocation();
  return fetch(`https://geo.ohdoylerules.com/api/location/${id}?lat=${pos.lat}&long=${pos.lng}`);
}

function renderGameMessage(won, message) {
  return (won) ? `<div class="message win-message">${message}</div>` : `<div class="message loss-message">${message}</div>`;
}

module.exports = function(data, ajax = false) {
  // zero out the message
  document.getElementById('game-message').innerHTML = '';
  if (ajax) {
    document.getElementById('submit-modal').setAttribute('data-id', data.location.id);
    storage.setWallet(data.user.credit);
    document.getElementById('user-wallet').innerHTML = data.user.credit;
    document.getElementById('location-wallet').innerHTML = data.location.wallet;
    document.getElementById('game-message').innerHTML = renderGameMessage(data.won, data.detail);
    return;
  }
  document.getElementById('submit-modal').setAttribute('data-id', data);
  getLocation(data)
  .then((response) => {
    return response.json();
  })
  .then((location) => {
    document.getElementById('user-wallet').innerHTML = storage.getWallet();
    document.getElementById('location-wallet').innerHTML = location.wallet;
  })
  .catch(function(err) {
    console.error(err);
  });
};
