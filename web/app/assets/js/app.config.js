app.config(['$routeProvider', '$locationProvider', function($routeProvider, $locationProvider) {
    $routeProvider
    .when("/", {
        templateUrl : "assets/tpl/components/categories.html"
    })
    .when("/categories/:categoryId", {
        templateUrl : "assets/tpl/components/products.html"
    })
    .otherwise({
			redirectTo: '/'
		});
    $locationProvider.hashPrefix('');
}]);
