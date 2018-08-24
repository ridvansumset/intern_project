app.controller('productController', ['$scope','productService' function($scope,productService) {
  $scope.categories = categoryService.categoryResource.query();


}]);
