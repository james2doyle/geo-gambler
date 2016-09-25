const storage = require('../app/storage');

module.exports = function() {
  const btnEvents = {
    handlePlayClick: function(id) {
      var pos = storage.getLocation();
      fetch(`https://geo.ohdoylerules.com/api/location/play/${id}?lat=${pos.lat}&long=${pos.lng}`, {
        method: 'post',
        headers: {
          'Accept': 'application/json',
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          number: 5
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