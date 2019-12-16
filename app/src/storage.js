import Storage from 'react-native-storage';
import {Platform} from 'react-native';

let AsyncStorage;
if (Platform.OS !== 'web') {
  AsyncStorage = require('@react-native-community/async-storage').default;
} else {
  AsyncStorage = window.localStorage;
}

const storage = new Storage({
  storageBackend: AsyncStorage,
});

export default storage;
