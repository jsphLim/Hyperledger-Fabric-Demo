'use strict';

var app = angular.module('application', []);
app.controller('appController', function ($scope, appFactory) {


	$scope.Get = function () {

		var key = $scope.container_key;

		appFactory.Get(key, function (data) {
			$scope.query_res = data;
			console.log(data)
		});
	}


	$scope.Set = function () {

		appFactory.Set($scope.holder, function (data) {
			$scope.change_holder = data;
			console.log(data)

		});
	}


});

app.factory('appFactory', function ($http) {

	var factory = {};


	factory.Get = function (key, callback) {
		$http.get('/get_data/' + key).success(function (output) {
			callback(output)
		});
	}



	factory.Set = function (data, callback) {

		var changes = data.id + "-" + data.content;

		$http.get('/set_data/' + changes).success(function (output) {
			callback(output)
		});
	}

	return factory;
});


