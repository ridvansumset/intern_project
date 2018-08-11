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

  $scope.new_category = {
      name : '',
      list_order: 5
    }
  $scope.addItem = function () {
    $scope.errortext = "";
    if ($scope.new_category.name != '') { // categories'nin içinde olup olmadığını kontrol ediyor,
      new categoryService.categoryResource($scope.new_category).$save(function(createdCategory){
        $scope.categories.push(createdCategory);
      });
      $scope.new_category.name = '';
    } else {
      $scope.errortext = "Boş bırakmayın a";
    }
  }


  $scope.removeItem = function(i, c) {
    categoryService.categoryResource.delete({
      categoryId: c.id
    });
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
