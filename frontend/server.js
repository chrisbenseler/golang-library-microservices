"use strict";
const express = require("express");
const compression = require("compression");

const _port = 3000;
const _app_folder = process.env.APPFOLDER || './dist/frontend/';

const app = express();
app.use(compression());


app.get('/frontend', (req, res) => {
    res.json({ ok: 'ok'})
});

// ---- SERVE STATIC FILES ---- //
app.get('*.*', express.static(_app_folder, {maxAge: '1y'}));

// ---- SERVE APLICATION PATHS ---- //
app.all('*', function (req, res) {
    res.status(200).sendFile(`/`, {root: _app_folder});
});

// ---- START UP THE NODE SERVER  ----
app.listen(_port, function () {
    console.log("Node Express server for " + app.name + " listening on http://localhost:" + _port);
});