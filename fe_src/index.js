import { TweenMax } from 'gsap';
mapboxgl.accessToken = 'pk.eyJ1IjoiYWRhcHQtYXNtdXNzZW4iLCJhIjoiY2pnN3h3dmUxMzB4aDJ5bW9vb2IzNmliMSJ9.Xp1FKOwzQ9s8Q9ZEUuwZeg';

let sources = [];
const addSource = (target, key, data) => {
  sources.push(key);
  target.addSource(key, { type: 'geojson', data });
  target.addLayer({
    "id": `layer_${key}`,
    "type": "circle",
    "source": key,
    "paint": {
      "circle-radius": 18,
      "circle-color": "#fff",
      "circle-opacity": 0.4
    }
  })
  setTimeout(removeSource, 5000, target);
}

const removeSource = (target) => {
  const key = sources.shift();
  target.removeLayer(`layer_${key}`);
  target.removeSource(key);
}

const getMarker = () => {
  let el = document.createElement('div');
  el.className = 'marker'
  return el;
}

const addFeatures = (geojson, target) => {
  geojson.features.forEach((feature) => {
    const markerEl = getMarker();
    const marker = new mapboxgl.Marker(markerEl)
      .setLngLat(feature.geometry.coordinates)
      .addTo(map);
    TweenMax.from(marker, 2, {alpha:0});
  })
}

const map = new mapboxgl.Map({
  container: 'map',
  style: 'mapbox://styles/mapbox/light-v9',
  center: [9.501785, 56.263920],
  zoom: 6,
  minZoom: 5,
  maxZoom: 12,
});

const setupWebSocket = () => {

  const web_socket = new WebSocket('ws://127.0.0.1:8888');

  web_socket.onopen = (event) => {

    web_socket.send("Here's some text client has sent to server upon socketd open event!");
  };

  web_socket.onmessage = (event) => {
    console.log(event.data);
  }
  console.log("hej");
}

const websocket_connection = (setupWebSocket)(); 

var points = {
  "type": "FeatureCollection",
  "features": [{
    "type" : "Feature",
    "geometry" : {
      "type": "Point",
      "coordinates": [10, 55]
    }
  },{
    "type" : "Feature",
    "geometry" : {
      "type": "Point",
      "coordinates": [10.5, 55.5]
    }    
  }]
}

map.on('load', function () {
  
  // addSource(map, 'cola', points);

  addFeatures(points, map);

});
