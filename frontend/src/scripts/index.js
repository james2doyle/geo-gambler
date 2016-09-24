// require('./__modernizr.custom.js');

const loadMap = require('./app');

window.coords = [{
  title: 'Location 1',
  latitude: 49.281355,
  longitude: -123.116686,
}, {
  title: 'Location 2',
  latitude: 49.283623,
  longitude: -123.103991,
}];

const mapEl = document.getElementById('map');

window.addEventListener('load', () => {
  loadMap(mapEl, coords)
  .then(map => {
    // body...
  }, err => {
    // body...
  });
});
