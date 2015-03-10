app.controller('ContainerListController', ['$http', function($http) {
    var cl = this;
    cl.loading = true;
    $http.get('/containers/list')
        .success(function(data) {
            cl.containers = data;
            cl.loading = false;
        });
}]);

app.controller('ContainerInfoController', ['$http', '$stateParams', function($http, $stateParams) {
    var ci = this;
    ci.loading = true;
    $http.get('/containers/' + $stateParams.id + '/info')
        .success(function(data) {
            ci.info = data;
            ci.loading = false;
        });
}]);