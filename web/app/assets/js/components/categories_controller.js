app.controller('categoryController', ['$scope', 'categoryService', function($scope,categoryService) {
  $scope.categories = categoryService.categoryResource.query();

  var categoiess = $scope.categories;
  $scope.convert = "";
  $scope.$watch("searchCategory.name", function(newValue, oldValue) {
    // if ( newValue !== oldValue) {
      $scope.convert = "";
      for (var i = 0; i < categoiess.length; i++) {
        if (newValue.includes(categoiess[i])) {
          $scope.convert = "var";
        } else {
          $scope.convert = "Yook";
        }
      }
      console.log(oldValue + ":" + newValue);
    // }
  }, true);

}]);
