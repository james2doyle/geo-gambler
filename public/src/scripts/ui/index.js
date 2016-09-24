function updateUi(locations) {
  return new Promise((resolve, reject) => {
    const username = document.getElementById('username');
    const credits = document.getElementById('credits');
    const locationCount = document.getElementById('location-count');
    locationCount.innerHTML = locations.length;
    resolve(locations);
  });
}

module.exports = updateUi;