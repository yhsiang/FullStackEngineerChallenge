import api from './api';

export const signIn = async (user, pass) => {
  const response = await api.post('/signIn', `user=${user}&pass=${pass}`);
  if (response.data.token) {
    return response.data;
  }

  return null;
};

export const signOut = async (user, pass) => {
  const response = await api.post('/signOut');
  return response.data.status;
};
