// require('./__modernizr.custom.js');

const loadMap = require('./app');

const mapEl = document.getElementById('map');

window.addEventListener('load', () => {
  fetch('http://6c48ebbd.ngrok.io/location')
  .then(function(response) {
    return response.json();
  })
  .then(loadMap)
  .then(map => {
    console.log('loaded');
  }, err => {
    console.log(err);
  });
});
