app.directive('containerControls', function() {
    return {
        scope: {
            containerId: '@'
        },
        restrict: 'E',
        replace: true,
        templateUrl: 'templates/container.controls.html',
        controller: ['$scope', '$http', function($scope, $http) {
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
        }]
    };
});