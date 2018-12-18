(function () {
    'use strict';

    angular
        .module('app', ['ui.router'])
        .config(config)
        .run(run);

        function config($stateProvider, $urlRouterProvider) {
            // default route
            $urlRouterProvider.otherwise("/main");
    
            $stateProvider.state('main', {
                    url: '/main',
                    templateUrl: 'Dashboard/dashboard.html',
                    controller: 'Dashboard.DashboardController',
                    controllerAs: 'vm',
                    data: { pageTitle: 'Dashboard' }
                });
        }
    
        function run($http, $rootScope, $window) {
            // add JWT token as default auth header
            $http.defaults.headers.common['Authorization'] = 'Bearer ' + $window.jwtToken;
    
            // update active tab on state change
            $rootScope.$on('$stateChangeSuccess', function (event, toState, toParams, fromState, fromParams) {
                $rootScope.$on('$stateChangeSuccess', function(event, toState, toParams) {
                    if (toState.name === 'main') {
                            $state.go('main', {}, {});
                    }
                    
                });
            });
        }

    // manually bootstrap angular after the JWT token is retrieved from the server
    $(function () {
        // get JWT token from server
        $.get('/app/token', function (token) {
            window.jwtToken = token;

            angular.bootstrap(document, ['app']);
        });
    });
})();