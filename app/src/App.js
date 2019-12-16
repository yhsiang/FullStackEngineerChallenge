import React, {useContext} from 'react';
import RouterPackage, {Router} from './Routing';

import {
  SignInScreen,
  AdminScreen,
  AssignScreen,
  UserScreen,
  ReviewScreen,
} from './screens';
import {AuthContext, auth, DataProvider} from './contexts';

const {Route, Redirect} = RouterPackage;
const App = () => {
  return (
    <Router>
      <AuthContext.Provider value={auth}>
        <DataProvider>
          <Route exact path="/" component={SignInScreen} />
          <PrivateRoute path="/user/:id" component={UserScreen} />
          <PrivateRoute path="/edit/:review_id" component={ReviewScreen} />
          <PrivateRoute
            path="/review/:reviewee_id/:reviewer_id"
            component={ReviewScreen}
          />
          <PrivateRoute path="/admin" component={AdminScreen} />
          <PrivateRoute path="/assign/:employee_id" component={AssignScreen} />
        </DataProvider>
      </AuthContext.Provider>
    </Router>
  );
};

const PrivateRoute = ({component: Component, ...rest}) => {
  const authValue = useContext(AuthContext);
  return (
    <Route
      {...rest}
      render={props =>
        authValue.isAuthenticated ? (
          <Component {...props} />
        ) : (
          <Redirect
            to={{
              pathname: '/',
              state: {from: props.location},
            }}
          />
        )
      }
    />
  );
};

export default App;
