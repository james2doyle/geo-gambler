const storage = require('../app/storage');

function getLocation(id) {
  var pos = storage.getLocation();
  return fetch(`https://geo.ohdoylerules.com/api/location/${id}?lat=${pos.lat}&long=${pos.lng}`);
}

function renderGameMessage(status, message) {
  return (status) ? `<div class="message win-message">${message}</div>` : `<div class="message loss-message">${message}</div>`;
}

module.exports = function(id) {
  document.getElementById('submit-modal').setAttribute('data-id', id);
  getLocation(id)
  .then((response) => {
    return response.json();
  })
  .then((location) => {
    document.getElementById('user-wallet').innerHTML = storage.getWallet();
    document.getElementById('location-wallet').innerHTML = location.wallet;
    document.getElementById('game-message').innerHTML = renderGameMessage(true, 'Game message here');
  })
  .catch(function(err) {
    console.error(err);
  });
};
