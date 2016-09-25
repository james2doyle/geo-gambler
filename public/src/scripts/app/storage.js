module.exports.setLocation = function (pos) {
  window.localStorage.setItem('playerLocation', JSON.stringify(pos));
};

module.exports.getLocation = function () {
  const pos = window.localStorage.getItem('playerLocation');
  return JSON.parse(pos);
};

module.exports.setApiKey = function (key) {
  window.localStorage.setItem('playerApiKey', key);
};

module.exports.getApiKey = function () {
  return window.localStorage.getItem('playerApiKey');
};

module.exports.setUsername = function (name) {
  window.localStorage.setItem('playerUsername', name);
};

module.exports.getUsername = function () {
  return window.localStorage.getItem('playerUsername');
};
