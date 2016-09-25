<!DOCTYPE html>
<html>
<head>
  <title>Geo Gambler</title>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width,initial-scale=1,minimum-scale=1,user-scalable=yes">
  <link rel="stylesheet" type="text/css" href="css/style.css">
  <script src="https://maps.googleapis.com/maps/api/js?key=AIzaSyCQ0NxmONBNjmeb9cJUkdIotSlUVClTYaw"></script>
</head>
<body>
  <div class="game-overlay">
    <button type="button" class="close-overlay" onclick="btnEvents.handleModalClose()">&times;</button>
    <form action="https://geo.ohdoylerules.com/api/location/play" method="post" accept-charset="utf-8">
      <h1>The Number Game</h1>
      <p>Welcome to the number game! Select a number below to guess what the number will be. Correct numbers pay out, and incorrect numbers debit.</p>
      <div class="split-group">
        <div class="wallet-group">
          <h3>Your Wallet</h3>
          <div><strong id="user-wallet"></strong> In Your Wallet.</div>
        </div>
        <div class="wallet-group">
          <h3>Location Wallet</h3>
          <div><strong id="location-wallet"></strong> Credits Available</div>
        </div>
      </div>
      <div id="game-message"></div>
      <div class="number-group">
        <input type="number" name="bet" id="bet" value="10" min="10" />
        <button type="button" class="btn default" onclick="btnEvents.increaseNumber('bet')">&plus;</button>
        <button type="button" class="btn default" onclick="btnEvents.decreaseNumber('bet')">&minus;</button>
      </div>
      <div class="number-group">
        <input type="number" name="number" id="number" value="1" min="1" max="10" />
        <button type="button" class="btn default" onclick="btnEvents.increaseNumber('number')">&plus;</button>
        <button type="button" class="btn default" onclick="btnEvents.decreaseNumber('number')">&minus;</button>
        <button type="button" id="submit-modal" class="btn success" onclick="btnEvents.handlePlayClick(this)" data-id="">Play!</button>
      </div>
    </form>
  </div>
  <div class="map-box">
    <div class="map-box-inner amount">Nearby Locations: <div id="location-count"></div></div>
    <div class="map-box-inner profile">Name: <div id="username">James</div></div>
    <div class="map-box-inner amount"><div id="credits">2000</div> Credits</div>
  </div>
  <div id="map" class="main-map"></div>
  <script src="js/script.js" type="text/javascript"></script>
</body>
</html>