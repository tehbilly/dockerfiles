app.controller('ImageListController', ['$http', function($http) {
    var il = this;
    il.loading = true;
    il.TotalSize = 0;
    $http.get('/images/list')
        .success(function(data) {
            il.images = data;
            il.loading = false;
            for (var i = 0; i < il.images.length; i++) {
                il.TotalSize = il.TotalSize + il.images[i].VirtualSize;
            }
        });
}]);

app.controller('ImageInfoController', ['$http', '$stateParams', function($http, $stateParams) {
    var ii = this;
    ii.loading = true;
    $http.get('/images/' + $stateParams.id + '/info')
        .success(function(data) {
            ii.info = data;
            ii.loading = false;
        });
}]);