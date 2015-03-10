var app = angular.module('VHostApp', ['ui.router']);
// We need to configure stuff, right?
app.config(['$stateProvider', '$urlRouterProvider', function($stateProvider, $urlRouterProvider) {
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
            controllerAs: 'cl'
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
        });
}]);

app.controller('DockerInfoController', ['$http', function($http) {
    var dsc = this;
    $http.get('/dockerinfo')
        .success(function(data) {
            dsc.info = data;
        });
}]);

app.directive('spinner', function() {
    return {
        scope: {
            'while': '='
        },
        restrict: 'E',
        replace: 'true',
        template: '<div ng-show="while" class="loader"/></div>'
    };
});

app.filter('bytes', function() {
    return function(bytes, precision) {
        if (isNaN(parseFloat(bytes)) || !isFinite(bytes) || bytes == 0) return '-';
        if (typeof precision === 'undefined') precision = 1;
        var units = ['bytes', 'kB', 'MB', 'GB', 'TB', 'PB'],
            number = Math.floor(Math.log(bytes) / Math.log(1024));
        return (bytes / Math.pow(1024, Math.floor(number))).toFixed(precision) +  ' ' + units[number];
    }
});