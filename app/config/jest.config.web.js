const preset = require('react-native-web/jest-preset');
const path = require('path');

preset.setupFiles = [...preset.setupFiles, '<rootDir>/config/setup.web.js'];

module.exports = {
  rootDir: path.resolve(__dirname, '..'),
  ...preset,
  testMatch: ['**/__tests__/**/!(native.)*.js'],
  testRunner: 'jest-circus/runner',
  transformIgnorePatterns: [
    '/node_modules/(?!native-base|react-router-native)/',
  ],
};
