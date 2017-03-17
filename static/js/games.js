SignCtrl.$inject = ['$scope', '$http', '$rootScope'];

function GameCtrl($scope, $http, $rootScope) {
  $rootScope.gamedata = {};

  var logError = function(data, status) {
    console.log('code '+status+': '+data);
  };

  var refresh = function() {
    if($rootScope.loggedin){
      $http.get('/games/'+$rootScope.mail).
        success(function(data) {
          $rootScope.gamedata = data
        }).error(logError);
    }
  };

  $scope.isHidden = function() {
    return !$rootScope.loggedin;
  }
}
