import api from './api';

export const getEmployee = async id => {
  const response = await api.get(`/employees/${id}`);

  if (response.data.status && response.data.data) {
    return response.data.data;
  }
  // TODO: handle error here
  return {};
};

export const createReview = async (reviewee, reviewer, content) => {
  const response = await api.post('/reviews', {
    reviewee,
    reviewer,
    content,
  });

  if (response.data.status && response.data.data) {
    return response.data.data;
  }
  // TODO: handle error here
  return {};
};

export const getReview = async reviewID => {
  const response = await api.get(`/reviews/${reviewID}`);

  if (response.data.status && response.data.data) {
    return response.data.data;
  }
  // TODO: handle error here
  return {};
};

export const updateReview = async (id, content) => {
  const response = await api.put('/reviews', {
    id,
    content,
  });

  if (response.data.status && response.data.data) {
    return response.data.data;
  }
  // TODO: handle error here
  return {};
};
