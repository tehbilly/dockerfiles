var app = angular.module('VHostApp', ['ui.router']);
// We need to configure stuff, right?
app.config(function($stateProvider, $urlRouterProvider, $httpProvider) {
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
});

app.controller('ContainerListController', ['$http', '$log', function($http, $log) {
    var cl = this;
    cl.loading = true;
    $http.get('/containers/list')
        .success(function(data) {
            cl.containers = data;
            cl.loading = false;
        });
}]);

app.controller('ContainerInfoController', function($http, $stateParams) {
    var ci = this;
    ci.loading = true;
    $http.get('/containers/' + $stateParams.id + '/info')
        .success(function(data) {
            ci.info = data;
            ci.loading = false;
        });
});

app.controller('ImageListController', function($http) {
    var il = this;
    il.loading = true;
    $http.get('/images/list')
        .success(function(data) {
            il.images = data;
            il.loading = false;
        });
});

app.controller('ImageInfoController', function($http, $stateParams) {
    var ii = this;
    ii.loading = true;
    $http.get('/images/' + $stateParams.id + '/info')
        .success(function(data) {
            ii.info = data;
            ii.loading = false;
        });
});

app.controller('DockerInfoController', function($http) {
    var dsc = this;
    $http.get('/dockerinfo')
        .success(function(data) {
            dsc.info = data;
        });
});

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

app.directive('containerControls', function() {
    return {
        scope: {
            containerId: '@'
        },
        restrict: 'E',
        replace: true,
        templateUrl: 'templates/container.controls.html',
        controller: function($scope, $http) {
            var defaultAlert = {
                show: false
            };

            $scope.alert = defaultAlert;

            $scope.clearMessage = function() {
                $scope.alert = defaultAlert;
            };

            var containerAction = function(action) {
                $http.get('/containers/' + $scope.containerId + '/' + action)
                    .success(function(data) {
                        console.log(action + ' response: ' + data);
                        $scope.alert = {
                            show: true,
                            class: 'alert-success',
                            message: data
                        };
                    })
                    .error(function(data) {
                        $scope.alert = {
                            show: true,
                            class: 'alert-warning',
                            message: data
                        };
                    });
            };

            $scope.startContainer = function() {
                containerAction('start');
            };
            $scope.stopContainer = function() {
                containerAction('stop');
            };
            $scope.restartContainer = function() {
                containerAction('restart');
            };
            $scope.killContainer = function() {
                containerAction('kill');
            };
        }
    };
});