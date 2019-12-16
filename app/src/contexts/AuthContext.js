import React from 'react';
import api from '../apis/api';

export const auth = {
  isAuthenticated: false,
  token: null,
  authenticate(token, cb) {
    this.isAuthenticated = true;
    this.token = token;
    api.defaults.headers.Authorization = `Bearer ${token}`;

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
