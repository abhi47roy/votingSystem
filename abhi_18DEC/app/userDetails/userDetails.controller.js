(function() {
  "use strict";

  angular
    .module("app")
    .controller("UserDetails.UserDetailsController", Controller);

  function Controller($rootScope,$scope, $http, $window, UserService, UtilityService,FlashService) {
    var vm = this;
    vm.clearUserSession = clearUserSession;
    vm.user = null;
    initController();

    function initController() {
      var user;
      
      UserService.GetCurrent().then(function(currentUser) {
              user = currentUser;
              vm.user = user;
              UtilityService.CurrentUser = user;
              
          })
          .catch(function(error) {
              FlashService.Error(error);
          });
       
  }

    function clearUserSession() {
      //clears user session on click on logout button
      if (vm.user.role == "ADMIN") {
        $rootScope.logout = false;
      }
      sessionStorage.clear();
      $http.defaults.headers.common.Authorization = "";
      localStorage.setItem("logged_in", false);
      $http.get("/login").then(function() {
        location.href = "/login";
      });
    }
  }
})();
