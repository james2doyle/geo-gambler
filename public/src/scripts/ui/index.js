const storage = require('../app/storage');

function updateUserDetails(locations, res) {
  const username = document.getElementById('username');
  username.innerHTML = res.name;
  const credits = document.getElementById('credits');
  credits.innerHTML = res.credit;
  const locationCount = document.getElementById('location-count');
  locationCount.innerHTML = locations.length;
}

function updateUi(locations) {
  return new Promise((resolve, reject) => {
    const id = storage.getApiKey();
    if (id) {
      fetch(`https://geo.ohdoylerules.com/api/user/me?auth=${id}`)
      .then((response) => {
        return response.json();
      }).then((res) => {
        updateUserDetails(locations, res);
        resolve(locations);
      }).catch((err) => {
        console.error(err);
        window.localStorage.clear();
      });
      return;
    }
    let username = storage.getUsername();
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
      updateUserDetails(locations, res);
      resolve(locations);
    }).catch((err) => {
      reject(err);
    });
  });
}

module.exports = updateUi;