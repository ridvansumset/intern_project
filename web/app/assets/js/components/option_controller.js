app.controller('optionController', ['$scope', 'optionService', 'choiceService', '$routeParams',
function($scope, optionService, choiceService, $routeParams) {
  $scope.options = [];
  optionService.optionResource.query({
    categoryId: $routeParams.categoryId,
    productId: $scope.product.id
  }, function(options) {
    options.forEach(function(option, index) {
      choiceService.choiceResource.query({
        categoryId: $routeParams.categoryId,
        productId: $scope.product.id,
        optionId: option.id
      }, function(choices) {
        choices.forEach(function(choice) {
          choice.added = false;
        });
        option.choices = choices;
        $scope.options.push(option);
      });
    });
  });

  // Cart Start Code
  // $scope.selected_options = [];
  $scope.selected_choices = [];
  var add = 0;

  $scope.addChoiceToArray = function (choice) {
    $scope.selected_choices.push(choice.price);
    for (var i = 0; i < $scope.selected_choices.length; i++) {
      add += $scope.selected_choices[i];
      $scope.totalAmount = add;
    }
    console.log($scope.selected_choices);
  }
  $scope.removeChoiceToArray = function (i,c) {
    $scope.selected_choices.splice(i, 1);
    for (var i = 0; i < $scope.selected_choices.length; i++) {
      add -= $scope.selected_choices[i];
      $scope.totalAmount = add;
    }
    console.log($scope.selected_choices);
  }
}]);
