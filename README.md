# GBF Bike #

GBF captains need a bike. `gbf-bike` helps captain find raid's room id from twitter.

## Usage ##

* Demo API Server: https://gbf-bike.herokuapp.com/

There are two ways to use gbf-bike. HTTP Get request and web socket.

### GET Request ###

The API will return "RAID" information on twitter in timeout seconds after get request.
By default timeout is 5 seconds.

* GET `https://gbf-bike.herokuapp.com/query`

```json
// RAIDs in 5 seconds on twitter.
[
  {
    "id": 863054895808561200,
    "level": "100",
    "roomId": "19A97526",
    "mobName": "バアル",
    "url": "https://t.co/mv65GoLxva"
  },
  {
    "id": 863054896903237600,
    "level": "75",
    "roomId": "10EC7A34",
    "mobName": "シュヴァリエ・マグナ",
    "url": "https://t.co/g9jPj10Whl"
  },
  {
    "id": 863054896936796200,
    "level": "75",
    "roomId": "63C1AF94",
    "mobName": "シュヴァリエ・マグナ",
    "url": "https://t.co/qWBsSrbsxc"
  },
  {
    "id": 863054898358665200,
    "level": "60",
    "roomId": "740AD1E2",
    "mobName": "青竜",
    "url": "https://t.co/8KiSakMvrO"
  },
  {
    "id": 863054899428245500,
    "level": "60",
    "roomId": "79302CD4",
    "mobName": "リヴァイアサン・マグナ",
    "url": "https://t.co/x6gYjzky3n"
  }
]
```

* Custom timeout: `https://gbf-bike.herokuapp.com/query?timeout=1`

### Web Socket ###

```html
<!doctype html>
<html lang="en">

<head>
  <meta charset="utf-8">
  <title>WebSocket</title>
</head>

<body>
  <p id="output"></p>

  <script>
    var loc = window.location;
    var uri = 'ws:';

    if (loc.protocol === 'https:') {
      uri = 'wss:';
    }
    uri = 'wss://gbf-bike.herokuapp.com/ws'

    ws = new WebSocket(uri)

    ws.onopen = function() {
      console.log('Connected')
    }

    ws.onmessage = function(evt) {
      var out = document.getElementById('output');
      out.innerHTML += evt.data + '<br>';
    }

    setInterval(function() {
      ws.send('Hello, Server!');
    }, 1000);
  </script>
</body>

</html>
```

## DIY ##

