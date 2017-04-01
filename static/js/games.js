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
          console.log($rootScope.gamedata);
          $rootScope.gamedata = data;
        }).error(logError);
    }
  };

  $scope.isHidden = function() {
    return !$rootScope.loggedin;
  }

  $scope.getFirstName = function(id) {
    return $rootScope.gamedata.players[id].firstname;
  }

  $scope.getLastName = function(id) {
    return $rootScope.gamedata.players[id].lastname;
  }

  $scope.getMail = function(id) {
    return $rootScope.gamedata.players[id].mail;
  }

  $scope.isPlayerInGame = function(id, game) {
    var arrayLength = game.players.length;
    for (var i = 0; i < arrayLength; i++) {
        if (game.players[i].id === id)
          return true;
    }
    return false;
  }

  $scope.addPlayerToGame = function(id, game) {
    $http.post('/addPlayerToGame', {playerid: id, gameid: game.id}).
      success(function() {
        refresh()
      }).error(logError);
  }

  $scope.removePlayerFromGame = function(id, game) {
    $http.post('/removePlayerFromGame', {playerid: id, gameid: game.id}).
      success(function() {
        refresh()
      }).error(logError);
  }

}
