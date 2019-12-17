import React from 'react';
import {StyleSheet} from 'react-native';
import {Picker, Icon} from 'native-base';

const styles = StyleSheet.create({
  placeholder: {color: '#bfc6ea'},
});
const DataPicker = ({data, selectedValue, onValueChange}) => {
  return (
    <Picker
      mode="dropdown"
      iosIcon={<Icon name="arrow-down" />}
      style={{width: undefined}}
      placeholder="Select one employee"
      placeholderStyle={styles.placeholder}
      placeholderIconColor="#007aff"
      selectedValue={selectedValue}
      onValueChange={onValueChange}>
      {data.map(d => (
        <Picker.Item key={d.id} label={d.name} value={d.id} />
      ))}
    </Picker>
  );
};

export default DataPicker;
