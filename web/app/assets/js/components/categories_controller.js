app.controller('categoryController', ['$scope', 'categoryService', function($scope,categoryService) {
  $scope.categories = categoryService.categoryResource.query();
  $scope.deneme = "KAMİL KAPLAN";
}])
