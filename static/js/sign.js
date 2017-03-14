function SignCtrl($scope, $http) {
  $scope.user = "";
  $scope.password = "";
  $scope.loggedin = false;
  $scope.errormessage = "";

  var logError = function(data, status) {
    console.log('code '+status+': '+data);
  };

  var refresh = function() {
    return $http.get('/').
      success(function() {

      }).error(logError);
  };

  $scope.signin = function() {
    console.log('trying to sign in: '+$scope.user+','+$scope.password);
    $scope.errormessage = "";

    /*
    $http.post('/signin/'+$scope.user, {password: $scope.password}).
      success(function() {
        $http.get('/signin/'+$scope.user).
          success(function(data) {
            console.log('sign in process completed');
            $scope.user = data.user;
            $scope.haserror = data.haserror;
            $scope.error = data.error;
          }).error(logError);
      }).error(logError);
    $scope.password = "";
    */
    $http.post('/signin', {user: $scope.user, password: $scope.password}).
      success(function() {
        $scope.loggedin = true;
      }).error(function(error, status) {
        switch(status) {
          case 403:
              $scope.errormessage = "Invalid username or password";
              break;
          case 404:
              $scope.errormessage = "Missing username";
              break;
          case 405:
              $scope.errormessage = "Missing password";
              break;
          default:
              $scope.errormessage = "Unknown error, status: "+status;
        }
      });

    //reset password in all cases
    $scope.password = "";
  };

  $scope.signup = function() {
    console.log('trying to sign up: '+$scope.user);
    $scope.errormessage = "";

    $http.post('/signup', {user: $scope.user}).
      success(function() {
      }).error(function(error, status) {
        switch(status) {
          case 403:
              $scope.errormessage = "Invalid username";
              break;
          case 404:
              $scope.errormessage = "Missing username";
              break;
          default:
              $scope.errormessage = "Unknown error, status: "+status;
        }
      });

    //reset password in all cases
    $scope.password = "";

  };

  refresh().then(function() { $scope.loggedin = false; });
}
