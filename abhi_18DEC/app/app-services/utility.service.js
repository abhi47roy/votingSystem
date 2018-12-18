(function () {
    'use strict';

    angular
        .module('app')
        .factory('UtilityService', Service);

    function Service($rootScope) {
        var service = {};
        service.CurrentUser = {};
        return service;
        
    }

})();