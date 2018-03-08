
var container = require('./controller.js');


module.exports = function (app) {

  app.get('/get_data/:id', function (req, res) {
    container.get_data(req, res);
  });

  app.get('/set_data/:holder', function (req, res) {
    container.set_data(req, res);
  });

}



