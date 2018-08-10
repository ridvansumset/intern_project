var app = angular.module('myApp', ['ngResource']);

app.factory('categoryService', ['$resource', function($resource) {
  var categoryResource = $resource(
    'http://localhost:1323/categories/:categoryId',
    {
      categoryId: '@id'
    },
    {
      'query': {
        method: 'GET',
        isArray: true
      }
    }
  )

  return {
    categoryResource: categoryResource
  }

}]);

app.controller('categoryController', ['$scope', 'categoryService', function($scope, categoryService) {

  // $scope.category = categoryService.category.get({
  //   categoryId: '1'
  // });
  // console.log($scope.category);
  $scope.categories = categoryService.categoryResource.query();

  var category_request_obj = {
    name      : 'Tatlılar',
    list_order: 5
  }
  $scope.new_category = new categoryService.categoryResource(category_request_obj).$save();
  // $scope.addItem = function () {
  //   $scope.errortext = "";
  //   if (!$scope.new_category) {return;}
  //   if ($scope.categories.indexOf($scope.new_category) == -1) {
  //     $scope.categories.push($scope.new_category);
  //   } else {
  //     $scope.errortext = "The item is already in your shopping list.";
  //   }
  // }
  $scope.removeItem = function(i, c) {
    // categoryService.category.delete({
    //   categoryId: i.toString()
    // });
    $scope.categories.splice(i, 1);
  }

  // $scope.categories = [];
  // $http({
  //   method : 'GET',
  //   url    : 'http://localhost:1323/categories'
  // }).then(function mySuccess(response) {
  //   $scope.categories = response.data;
  //   console.log($scope.categories);
  // }, function myError(err) {
  //     $scope.error = err.statusText;
  //   });
}]);
