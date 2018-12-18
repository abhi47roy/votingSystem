(function () {
    'use strict';

    angular
        .module('app')
        .controller('TransactionDetails.TransactionDetailsController', Controller);

    function Controller($rootScope,$scope,$http,$window, UserService, UtilityService,TransactionService ,FlashService){
        var vm = this;
        vm.userDetails= UtilityService.CurrentUser;
        vm.user = vm.userDetails;
        vm.addRecords = addRecords;
        vm.transactionsList = [];
        
        initController();

        function initController() {
         vm.username = vm.userDetails;
         
         TransactionService.getAllTransactions(vm.username.username).then(function (transactionsRecord) {
            //JSON.parse(transactionsRecord);
            console.log(transactionsRecord);
            });
        }


        function addRecords(){
            console.log("Record insertion called",vm.userDetails);
            TransactionService.addTransactionsRecords(vm.userDetails.username).then(function (transactionsRecord) {
                console.log("Records entered ",transactionsRecord);
                });
        }
        
    }    


})();