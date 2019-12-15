import api from './api';
import {Cookie} from 'tough-cookie';
import {Platform} from 'react-native';

export const signIn = async (user, pass) => {
  const response = await api.post('/signIn', `user=${user}&pass=${pass}`);
  const cookies = response.headers['set-cookie'].map(Cookie.parse);

  if (Platform.OS !== 'web') {
    api.defaults.headers.Cookie = `${cookies[0].key}=${cookies[0].value}`;
  }
  if (response.data.status && response.data.data) {
    return {role: 'user', id: response.data.data.id};
  }
  if (response.data.status) {
    return {role: 'admin'};
  }

  return null;
};

export const signOut = async (user, pass) => {
  const response = await api.post('/signOut');
  return response.data.status;
};
