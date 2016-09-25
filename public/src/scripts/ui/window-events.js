const storage = require('../app/storage');
const renderMarkerTemplate = require('../ui/render-marker-template');
const updateModal = require('../ui/update-modal');

module.exports = function() {
  const btnEvents = {
    increaseNumber: function() {
      let el = document.getElementById('number');
      let number = parseInt(el.value, 10);
      if (number < el.max) {
        // make a string so the input updates
        el.value = (number + 1) + "";
      }
    },

    decreaseNumber: function() {
      let el = document.getElementById('number');
      let number = parseInt(el.value, 10);
      if (number > el.min) {
        // make a string so the input updates
        el.value = (number - 1) + "";
      }
    },

    handleModalOpen: function(id) {
      document.body.classList.add('game-overlay-on');
      updateModal(id);
    },

    handleModalClose: function() {
      document.body.classList.remove('game-overlay-on');
    },

    handlePlayClick: function(el) {
      const id = el.getAttribute('data-id');
      const pos = storage.getLocation();
      const auth = storage.getApiKey();
      const number = parseInt(document.getElementById('number').value, 10);
      fetch(`https://geo.ohdoylerules.com/api/location/play/${id}?lat=${pos.lat}&long=${pos.lng}&auth=${auth}`, {
        method: 'post',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          number,
        })
      }).then((response) => {
        return response.json();
      }).then((res) => {
        updateModal(res, true);
        document.getElementById('on-map-maker').innerHTML = renderMarkerTemplate(res.location);
      }).catch(function(err) {
        console.error(err);
      });
    }
  };

  return btnEvents;
};