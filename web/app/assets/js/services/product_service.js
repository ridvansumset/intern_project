app.factory('productService', ['$resource', function($resource) {
  var productResource = $resource(
    'http://localhost:1323/categories/:categoryId/products',
    {
      categoryId: '@id'
    },
    {
      'query': {
        method: 'GET',
        isArray: true
      }
    }
  )

  return {
    productResource: productResource
  }

}]);
