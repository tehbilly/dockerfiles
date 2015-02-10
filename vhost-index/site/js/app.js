var app = angular.module('VHostApp', ['ui.router']);
// We need to configure stuff, right?
app.config(function($stateProvider, $urlRouterProvider) {
    // For unmatched URLs, redirect to /
    $urlRouterProvider.otherwise("/");
    // Now we set up our routes
    $stateProvider
        .state("container-list", {
            url: "/",
            templateUrl: "templates/container-list.html",
            controller: 'ListController',
            controllerAs: 'lc'
        })
        .state("container-info", {
            url: "/container/{id}",
            templateUrl: "templates/container-info.html",
            controller: 'ContainerInfoController',
            controllerAs: 'ci'
        })
});


// Controller for the list of containers
app.controller('ListController', function($http) {
    var lc = this;

    $http.get('/containers')
        .success(function(data) {
            lc.containers = data;
        });
});

app.controller('ContainerInfoController', function($http, $stateParams) {
    var ci = this
    $http.get('/container/' + $stateParams.id)
        .success(function(data) {
            ci.info = data
        });
});

app.controller('DockerInfoController', function($http) {
    var dsc = this;

    $http.get('/dockerinfo')
        .success(function(data) {
            dsc.info = data
        });
});