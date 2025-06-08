const express = require('express');
const app = express();

app.get('/temperature', (req, res) => {
    const location = req.query.location || 'unknown';
    const temperature = +(Math.random() * (30 - 16) + 16).toFixed(2); // 16..30°C
    res.json({ location, temperature });
});

app.listen(8081, () => {
    console.log('Temperature-API running on port 8081');
});
