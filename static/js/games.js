SignCtrl.$inject = ['$scope', '$http', '$rootScope'];

function initGameData(data) {
  for(var gameid in data.games) {
    var game = data.games[gameid]
    var arrayLength = game.players.length;
    for (var i = 0; i < arrayLength; i++) {
      if (game.players[i].id === data.currentplayerid) {
        game.players[i].newComment = game.players[i].comment
      }
    }
  }
  for(var teamid in data.teams) {
    var team = data.teams[teamid]
    team.newName = team.name;
    var arrayLength = team.players.length;
    for (var i = 0; i < arrayLength; i++) {
      if (team.players[i].id === data.currentplayerid) {
        team.players[i].newComment = team.players[i].comment
      }
    }
  }
}

function GameCtrl($scope, $http, $rootScope) {
  $rootScope.gamedata = {};

  var logError = function(data, status) {
    console.log('code '+status+': '+data);
  };

  var refresh = function() {
    if($rootScope.loggedin){
      $http.get('/games/'+$rootScope.mail+'?rnd='+new Date().getTime()).
        success(function(data) {
          $rootScope.gamedata = data;
          initGameData($rootScope.gamedata);
          console.log($rootScope.gamedata);
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

  $scope.getPlayerInGame = function(id, game) {
    var arrayLength = game.players.length;
    for (var i = 0; i < arrayLength; i++) {
      if (game.players[i].id === id)
      return game.players[i];
    }
    return null;
  }

  $scope.getTeam = function(teamid) {
    return $rootScope.gamedata.teams[teamid];
  }

  $scope.isPlayerInTeam = function(id, teamid) {
    var team = $scope.getTeam(teamid);
    var arrayLength = team.players.length;
    for (var i = 0; i < arrayLength; i++) {
      if (team.players[i].id === id)
      return true;
    }
    return false;
  }

  $scope.getPlayerInTeam = function(id, teamid) {
    var team = $scope.getTeam(teamid);
    var arrayLength = team.players.length;
    for (var i = 0; i < arrayLength; i++) {
      if (team.players[i].id === id)
      return team.players[i];
    }
    return null;
  }


  $scope.addPlayerToGame = function(id, game) {
    $http.post('/addPlayerToGame?rnd='+new Date().getTime(), {playerid: id, gameid: game.id}).
      success(function() {
        refresh()
      }).error(logError);
  }

  $scope.removePlayerFromGame = function(id, game) {
    $http.post('/removePlayerFromGame?rnd='+new Date().getTime(), {playerid: id, gameid: game.id}).
      success(function() {
        refresh()
      }).error(logError);
  }

  $scope.submitPlayerGameComment = function(id, game) {
    $http.post('/submitPlayerGameComment?rnd='+new Date().getTime(), {playerid: id, gameid: game.id, comment: $scope.getPlayerInGame(id, game).newComment}).
      success(function() {
        refresh()
      }).error(logError);
  }

  $scope.addTeamToGame = function(teamname, managerid, game) {
    if(teamname == undefined || teamname === "") {
      //todo error handling
      return;
    }
    $http.post('/addTeamToGame?rnd='+new Date().getTime(), {teamname: teamname, managerid: managerid, gameid: game.id}).
      success(function() {
        refresh()
      }).error(logError);
  }

  $scope.removeTeamFromGame = function(teamid, managerid, game) {
    $http.post('/removeTeamFromGame?rnd='+new Date().getTime(), {teamid: teamid, managerid: managerid, gameid: game.id}).
      success(function() {
        refresh()
      }).error(logError);
  }

  $scope.changeTeamName = function(teamid, teamname, managerid) {
    if(teamname == undefined || teamname === "") {
      //todo error handling
      return;
    }
    $http.post('/changeTeamName?rnd='+new Date().getTime(), {teamid: teamid, teamname: teamname, managerid: managerid}).
      success(function() {
        refresh()
      }).error(logError);
  }

  $scope.addPlayerToTeam = function(id, teamid) {
    $http.post('/addPlayerToTeam?rnd='+new Date().getTime(), {playerid: id, teamid: teamid}).
      success(function() {
        refresh()
      }).error(logError);
  }

  $scope.removePlayerFromTeam = function(id, teamid) {
    $http.post('/removePlayerFromTeam?rnd='+new Date().getTime(), {playerid: id, teamid: teamid}).
      success(function() {
        refresh()
      }).error(logError);
  }

  $scope.submitPlayerTeamComment = function(id, teamid) {
    $http.post('/submitPlayerTeamComment?rnd='+new Date().getTime(), {playerid: id, teamid: teamid, comment: $scope.getPlayerInTeam(id, teamid).newComment}).
      success(function() {
        refresh()
      }).error(logError);
  }


}
