const express = require('express');
const app = express();

function getSensorMapping(location, sensorId) {
// If no location is provided, use a default based on sensor ID
    if (!location) {
        switch (sensorId) {
            case "1": location = "Living Room"; break;
            case "2": location = "Bedroom"; break;
            case "3": location = "Kitchen"; break;
            default:  location = "Unknown";
        }
    }
// If no sensor ID is provided, generate one based on location
    if (!sensorId) {
        switch (location) {
            case "Living Room": sensorId = "1"; break;
            case "Bedroom": sensorId = "2"; break;
            case "Kitchen": sensorId = "3"; break;
            default: sensorId = "0";
        }
    }
    return { sensorId, location };
}

// /temperature?location=...
app.get('/temperature', (req, res) => {
    let { location, sensorId } = req.query;
    ({ sensorId, location } = getSensorMapping(location, sensorId));
    const value = +(Math.random() * (30-16) + 16).toFixed(2);
    res.json({ sensorId, location, value });
});

// /temperature/:sensorId
app.get('/temperature/:sensorId', (req, res) => {
    const sensorId = req.params.sensorId;
    const { location } = getSensorMapping(null, sensorId);
    const value = +(Math.random() * (30-16) + 16).toFixed(2);
    res.json({ sensorId, location, value });
});

app.listen(8081, () => {
    console.log('Temperature-API running on port 8081');
});