app.controller('productController', ['$scope','productService', '$routeParams', function($scope,productService,$routeParams) {
  $routeParams.category_id = productService.productResource.query();
}]);
