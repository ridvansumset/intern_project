app.controller('productController', ['$scope','productService', '$routeParams', function($scope,productService,$routeParams) {
  $scope.products = productService.productResource.query();
  $scope.id = $routeParams.categoryId
  console.log($scope.id);
  // $routeParams.category_id = productService.productResource.query();
}]);
