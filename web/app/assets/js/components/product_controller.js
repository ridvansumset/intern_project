app.controller('productController', ['$scope','productService', '$routeParams', function($scope,productService,$routeParams) {
  $scope.products = productService.productResource.query();
  console.log($routeParams.categoryId);
  // $routeParams.category_id = productService.productResource.query();
}]);
