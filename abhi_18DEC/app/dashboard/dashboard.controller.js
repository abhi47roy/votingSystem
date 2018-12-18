(function () {
    'use strict';

    angular
        .module('app')
        .controller('Dashboard.DashboardController', Controller);

    function Controller($scope,UserService,UtilityService) {
        var vm = this;

        vm.user = null;

        initController();

        function initController() {
            var user;
            // get current user
            UserService.GetCurrent().then(function (currentUser) {
              user = currentUser;
              vm.user = user;
              UtilityService.CurrentUser = user;
              $scope.$broadcast('mydata', vm.user);
            });
        }
    }

})();