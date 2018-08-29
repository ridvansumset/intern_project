app.controller('mainController', ['$scope', '$localStorage', function($scope, $localStorage) {
  $scope.localStorage = $localStorage;

  $scope.init = function() {
    $localStorage.cart = {
        products: [],
        totalAmount: 0
    }
  }
  // console.log($scope.localStorage.cart);
}])
