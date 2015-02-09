var app = angular.module('VHostApp', ['mgcrea.ngStrap']); // 'mgcrea.ngStrap' if ngStrap

app.controller('ListController', function($http, $modal) {
    var lc = this;

    $http.get('/containers')
        .success(function(data) {
            lc.containers = data;
        });

    lc.getInfo = function(id) {
        $http.get('/container/' + id)
            .success(function (data) {
                $modal({
                    "title": "A response!",
                    "content": "<pre>" + JSON.stringify(data) + "</pre>",
                    "html": true,
                    "show": true
                });
            });
    }
});

app.controller('DockerInfoController', function($http) {
    var dsc = this;

    $http.get('/dockerinfo')
        .success(function(data) {
            dsc.info = data
        });
});