const storage = require('../app/storage');

function updateUi(locations) {
  return new Promise((resolve, reject) => {
    var username = storage.getUsername();
    if (!username) {
      username = prompt('Enter a username to continue:');
      storage.setUsername(username);
    }
    fetch('https://geo.ohdoylerules.com/api/user', {
      method: 'post',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        name: username
      })
    }).then((res) => {
      return res.json();
    }).then((res) => {
      storage.setApiKey(res.id);
      const username = document.getElementById('username');
      username.innerHTML = res.name;
      const credits = document.getElementById('credits');
      credits.innerHTML = res.credit;
      const locationCount = document.getElementById('location-count');
      locationCount.innerHTML = locations.length;
      resolve(locations);
    }).catch((err) => {
      reject(err);
    });
  });
}

module.exports = updateUi;