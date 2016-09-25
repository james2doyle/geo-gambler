const storage = require('../app/storage');

module.exports = function() {
  const btnEvents = {
    handlePlayClick: function(id) {
      var pos = storage.getLocation();
      var auth = storage.getApiKey();
      fetch(`https://geo.ohdoylerules.com/api/location/play/${id}?lat=${pos.lat}&long=${pos.lng}&auth=${auth}`, {
        method: 'post',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          number: parseInt(document.getElementById('number').value, 10)
        })
      }).then((response) => {
        return response.json();
      }).then((res) => {
        console.log(res);
      }).catch(function(err) {
        console.error(err);
      });
    }
  };

  return btnEvents;
};