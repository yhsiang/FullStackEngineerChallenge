import mockAsyncStorage from '@react-native-community/async-storage/jest/async-storage-mock';

/* eslint-disable no-undef */
jest.mock('@react-native-community/async-storage', () => mockAsyncStorage);
jest.mock('react-router-native', () => ({
  useHistory: () => ({
    push: jest.fn(),
  }),
}));
