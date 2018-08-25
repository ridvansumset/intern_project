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
      },
      'update': {
        method: 'PUT'
      }
    }
  )

  return {
    productResource: productResource
  }
}])
