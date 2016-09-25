module.exports.setLocation = function (pos) {
  window.localStorage.setItem('playerLocation', JSON.stringify(pos));
};

module.exports.getLocation = function () {
  const pos = window.localStorage.getItem('playerLocation');
  return JSON.parse(pos);
};