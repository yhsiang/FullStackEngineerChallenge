const preset = require('react-native/jest-preset');
const path = require('path');

module.exports = {
  rootDir: path.resolve(__dirname, '..'),
  ...preset,
  testRunner: 'jest-circus/runner',
  testMatch: ['**/__tests__/**/!(web.)*.js'],
  transformIgnorePatterns: [
    '/node_modules/(?!native-base|react-router-native)/',
  ],
};
