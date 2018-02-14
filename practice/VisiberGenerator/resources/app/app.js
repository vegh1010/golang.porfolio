angular.module('visiber', [
    'ui.bootstrap'
]);

angular.module('visiber').run(function($rootScope, $http){
    $rootScope.init = function() {
        // Init
        asticode.loader.init();
        asticode.modaler.init();
        asticode.notifier.init();

        // Wait for astilectron to be ready
        document.addEventListener('astilectron-ready', function() {
            $rootScope.initConfig();
        })
    };
    $rootScope.init();
});

angular.module('visiber').controller('appController', function ($rootScope, $scope, $http) {
    'use strict';

    $rootScope.user = {
        "name": "",
        "date": ""
    };
    $rootScope.printOutput = function(message) {
        asticode.notifier.info(message);
    };

    $rootScope.initConfig = function() {
        var initEventMessage = {
            "name": "init"
        };
        astilectron.sendMessage(initEventMessage, function(receiveMessage) {
            console.log(receiveMessage.payload);

            // Check error
            if (receiveMessage.name === "error") {
                $rootScope.printOutput("Error: " + receiveMessage.payload);
                return
            }
            $rootScope.printOutput(receiveMessage.payload);
        });
    };

    $rootScope.generateReport = function() {
        var initEventMessage = {
            "name": "generate_report",
            "payload": $rootScope.user
        };
        astilectron.sendMessage(initEventMessage, function(receiveMessage) {
            console.log(receiveMessage.payload);

            // Check error
            if (receiveMessage.name === "error") {
                $rootScope.printOutput("Error: " + receiveMessage.payload);
                return
            }
            $rootScope.printOutput(receiveMessage.payload);
        });
    };
});
