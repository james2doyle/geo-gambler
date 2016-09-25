// require('./__modernizr.custom.js');

const loadMap = require('./app');
const updateMarkers = require('./app/update-markers');
const updateUi = require('./ui');

window.addEventListener('load', () => {
  fetch('https://6c48ebbd.ngrok.io/location')
  .then((response) => {
    return response.json();
  })
  .then((locations) => {
    return updateUi(locations)
    .then(loadMap) // resolves the google map and locations as an object
    .then(updateMarkers); // resolves the same as the above
  })
  .then(data => {
    window.btnEvents = require('./ui/window-events')();
    document.body.classList.add('loaded');
  })
  .catch(err => {
    console.log(err);
  });
});
