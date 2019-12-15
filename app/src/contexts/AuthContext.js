import React from 'react';

export const auth = {
  isAuthenticated: false,
  authenticate(cb) {
    this.isAuthenticated = true;
    if (cb) {
      setTimeout(cb, 100);
    }
  },
  signout(cb) {
    this.isAuthenticated = false;
    if (cb) {
      setTimeout(cb, 100);
    }
  },
};

const AuthContext = React.createContext(auth);

export default AuthContext;
