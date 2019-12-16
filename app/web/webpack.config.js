const path = require('path');
const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const rootDir = path.join(__dirname, '..');
const webpackEnv = process.env.NODE_ENV || 'development';

module.exports = {
  mode: webpackEnv,
  entry: {
    app: path.join(rootDir, './index.web.js'),
  },
  output: {
    path: path.resolve(rootDir, 'dist'),
    filename: 'app-[hash].bundle.js',
  },
  devtool: 'source-map',
  module: {
    rules: [
      {
        test: /\.(js)$/,
        include: [
          path.resolve('src'),
          path.resolve('node_modules/native-base-shoutem-theme'),
          path.resolve('node_modules/react-navigation'),
          path.resolve('node_modules/react-native-easy-grid'),
          path.resolve('node_modules/react-native-drawer'),
          path.resolve('node_modules/react-native-safe-area-view'),
          path.resolve('node_modules/react-native-vector-icons'),
          path.resolve('node_modules/react-native-keyboard-aware-scroll-view'),
          path.resolve('node_modules/react-native-web'),
          path.resolve('node_modules/react-native-tab-view'),
          path.resolve('node_modules/static-container'),
          path.resolve('node_modules/react-router-native'),
          path.resolve('node_modules/@react-native-community/async-storage'),
        ],
        loader: 'babel-loader',
      },
    ],
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: path.join(__dirname, './index.html'),
    }),
    new webpack.HotModuleReplacementPlugin(),
  ],
  resolve: {
    extensions: ['.web.js', '.js'],
    alias: Object.assign({
      'react-native$': 'react-native-web',
    }),
  },
};
