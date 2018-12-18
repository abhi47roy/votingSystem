(function () {
    'use strict';

    angular
        .module('app')
        .factory('TransactionService', Service);

    function Service($http, $q) {
        var service = {};
        service.getAllTransactions = getAllTransactions;
        service.addTransactionsRecords = addTransactionsRecords;
        
        return service;

        function getAllTransactions(username){
            var data = {
                username : username
            }
            return $http.get('/api/transaction/getAllTransactions?username=' + username).then(handleSuccess, handleError);
        }

        function addTransactionsRecords(username){
            var data = {
                username : username
            }
            return $http.post('/api/transaction/addTransactionsRecords',data).then(handleSuccess, handleError);
        }

        // private functions

        function handleSuccess(res) {
            return res.data;
        }

        function handleError(res) {
            return $q.reject(res.data);
        }
    }

})();
