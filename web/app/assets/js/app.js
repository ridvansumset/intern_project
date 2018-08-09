var app = angular.module("myApp", []);

app.controller("categoryController", ['$scope', '$http', function($scope, $http){
  $scope.categories = [];
  var s = "string";
  console.log(s);
  $http({
    method : "GET",
    url    : "http://localhost:1323/categories"
  }).then(function mySuccess(response) {
    $scope.categories = response.data;
    console.log($scope.categories);
  }, function myError(err) {
      $scope.error = err.statusText;
    });

  $scope.removeItem = function(i, c) {
    // $scope.errortext = "";
    $scope.categories.splice(i, 1);
  }

}]);
