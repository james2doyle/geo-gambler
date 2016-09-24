// require('./__modernizr.custom.js');

const loadMap = require('./app');

window.addEventListener('load', () => {
  fetch('https://6c48ebbd.ngrok.io/location')
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
