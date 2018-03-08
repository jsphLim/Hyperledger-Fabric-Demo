
var express = require('express');
var path = require('path');

var app = express();

require('./routes.js')(app);

app.use(express.static(path.join(__dirname, './client')));

var port = process.env.PORT || 7775;

app.listen(port, function () {
  console.log("Serving on port: " + port);
});

