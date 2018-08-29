app.controller('productController', ['$scope', 'productService', 'categoryService', '$routeParams', '$modal',
function($scope, productService, categoryService, $routeParams,$modal) {
  $scope.products = productService.productResource.query({
    categoryId: $routeParams.categoryId
  });
  $scope.category = categoryService.categoryResource.get({
    categoryId: $routeParams.categoryId
  });

  $scope.product = {};
  $scope.showProductModal = function(product) {
    $scope.product = product;
    $scope.showModal();
  }
  var optionModal = $modal({
    scope: $scope,
    placement: 'center',
    templateUrl: 'assets/tpl/modals/option-modal.html',
    show: false
  });
  $scope.showModal = function() {
    optionModal.$promise.then(optionModal.show);
  };
  
  //  Cart Start Code
  $scope.cart = {
    products: [],
    totalAmountMethod: function() {
      optionModal.hide();
    },
    totalAmount: 0
  }
  // prod = {
  //   name:'asd',
  //   price: 5
  //   options: [{}, {
  //     name: 'aşsdl',
  //     choices: [{}, {}]
  //   }]
  // }
}]);
