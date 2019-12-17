import api from './api';

export const getEmployees = async () => {
  const response = await api.get('/admin/employees');

  if (response.data.status && response.data.data) {
    return response.data.data;
  }
  // TODO: handle error here
  return [];
};

export const createEmployee = async name => {
  const response = await api.post(
    '/admin/employees',
    {name},
    {
      headers: {
        'Content-Type': 'application/json',
      },
    },
  );

  if (response.data.status && response.data.data) {
    return response.data.data;
  }
  // TODO: handle error here
  return {};
};

export const updateEmployee = async (id, name) => {
  const response = await api.put(
    `/admin/employees/${id}`,
    {name},
    {
      headers: {
        'Content-Type': 'application/json',
      },
    },
  );

  if (response.data.status && response.data.data) {
    return response.data.data;
  }
  // TODO: handle error here
  return {};
};

export const addReviewer = async (reviewee, reviewer) => {
  const response = await api.post(
    '/admin/reviewers/add',
    {reviewee, reviewer},
    {
      headers: {
        'Content-Type': 'application/json',
      },
    },
  );

  return response.data.status;
};

export const removeReviewer = async (reviewee, reviewer) => {
  const response = await api.post(
    '/admin/reviewers/remove',
    {reviewee, reviewer},
    {
      headers: {
        'Content-Type': 'application/json',
      },
    },
  );

  return response.data.status;
};
