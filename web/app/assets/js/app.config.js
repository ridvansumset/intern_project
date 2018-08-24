app
.config(['$routeProvider', '$locationProvider', function($routeProvider, $locationProvider) {
    $routeProvider
    .when("/", {
        templateUrl : "assets/tpl/components/categories.html"
    })
    .when("/categories/:categoryId", {
        templateUrl : "assets/tpl/components/products.html"
    })
    .when("/categories/:categoryId/products", {
        templateUrl : "assets/tpl/components/options.html"
    })
    .when("/admin", {
        templateUrl : "assets/tpl/components/admin.html"
    })
    .otherwise({
			redirectTo: '/'
		});
    $locationProvider.hashPrefix('');
}])

.config(function($modalProvider) {
  angular.extend($modalProvider.defaults, {
    animation: 'am-flip-x',
    placement: 'center'
  });
})
