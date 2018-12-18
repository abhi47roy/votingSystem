var config = require('config.json');
var express = require('express');
var router = express.Router();
var transactionService = require('services/transaction.service');

// routes

router.get('/getAllTransactions', getAllTransactions);
router.post('/addTransactionsRecords', addTransactionsRecords);

module.exports = router;

function getAllTransactions(req, res) {
    transactionService.getAllTransactions(req.query.username)
    .then(function(response) {
        res.send(response);
    })
    .catch(function(err) {
        res.status(400).send(err);
    });
}


function addTransactionsRecords(req, res) {
    transactionService.addTransactionsRecords(req)
    .then(function(response) {
        res.send(response);
    })
    .catch(function(err) {
        res.status(400).send(err);
    });
}


