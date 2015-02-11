var app = angular.module('VHostApp', ['ui.router']);
// We need to configure stuff, right?
app.config(function($stateProvider, $urlRouterProvider) {
    // For unmatched URLs, redirect to /
    $urlRouterProvider.otherwise("/containers/list");
    // Now we set up our routes
    $stateProvider
        .state("containers", {
            abstract: true,
            url: "/containers",
            templateUrl: "templates/container.html"
        })
        .state("containers.list", {
            url: "/list",
            templateUrl: "templates/container.list.html",
            controller: 'ContainerListController',
            controllerAs: 'lc'
        })
        .state("containers.info", {
            url: "/info/{id}",
            templateUrl: "templates/container.info.html",
            controller: 'ContainerInfoController',
            controllerAs: 'ci'
        })
        .state("images", {
            abstract: true,
            url: "/images",
            templateUrl: "templates/image.html"
        })
        .state("images.list", {
            url: "/list",
            templateUrl: "templates/image.list.html",
            controller: "ImageListController",
            controllerAs: "il"
        })
        .state("images.info", {
            url: "/info/{id}",
            templateUrl: "templates/image.info.html",
            controller: "ImageInfoController",
            controllerAs: "ii"
        })
});


// Controller for the list of containers
app.controller('ContainerListController', function($http) {
    var cl = this;

    $http.get('/containers')
        .success(function(data) {
            cl.containers = data;
        });
});

app.controller('ContainerInfoController', function($http, $stateParams) {
    var ci = this;
    $http.get('/container/' + $stateParams.id)
        .success(function(data) {
            ci.info = data;
        });
});

app.controller('ImageListController', function($http) {
    var il = this;
    $http.get('/images')
        .success(function(data) {
            il.images = data;
        });
});

app.controller('ImageInfoController', function($http, $stateParams) {
    var ii = this;
    $http.get('/image/' + $stateParams.id)
        .success(function(data) {
            ii.info = data;
        });
});

app.controller('DockerInfoController', function($http) {
    var dsc = this;

    $http.get('/dockerinfo')
        .success(function(data) {
            dsc.info = data;
        });
});