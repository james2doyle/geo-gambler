const storage = require('../app/storage');

module.exports = function() {
  return new Promise((resolve, reject) => {
    if (window.navigator.geolocation) {
      window.navigator.geolocation.watchPosition((position) => {
        // window.navigator.geolocation.getCurrentPosition((position) {
        const pos = {
          lat: position.coords.latitude,
          lng: position.coords.longitude
        };

        // save the user position in local storage
        storage.setLocation(pos);

        // send the position back
        resolve(pos);
      }, () => {
        reject('Error: The Geolocation service failed.');
      }, {
        enableHighAccuracy: false, // high accuraccy can be slow
        timeout: 5000,
        maximumAge: 0
      });
    } else {
      reject('Error: Your browser doesn\'t support geolocation.');
    }
  });
};