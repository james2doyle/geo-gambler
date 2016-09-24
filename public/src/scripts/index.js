// require('./__modernizr.custom.js');

const loadMap = require('./app');
const updateUi = require('./ui');

window.addEventListener('load', () => {
  fetch('https://6c48ebbd.ngrok.io/location')
  .then(function(response) {
    return response.json();
  })
  .then(function(locations) {
    return updateUi(locations).then(loadMap);
  })
  .then(map => {
    window.btnEvents = require('./ui/window-events')();
    document.body.classList.add('loaded');
  }, err => {
    console.log(err);
  });
});
