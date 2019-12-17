import React from 'react';
import RouterPackage from '../Routing';
import {Text, List, ListItem, Left, Right, Icon} from 'native-base';

const {useHistory} = RouterPackage;

const ReviewList = ({data}) => {
  const history = useHistory();
  if (data.length === 0) {
    return <Text>No Data</Text>;
  }

  return (
    <List>
      {data.map(d => {
        return (
          <ListItem
            key={d.id}
            button
            onPress={() => {
              history.push(`/edit/${d.id}`);
            }}>
            <Left>
              <Text>{d.content}</Text>
            </Left>
            <Right>
              <Icon name="arrow-forward" />
            </Right>
          </ListItem>
        );
      })}
    </List>
  );
};

export default ReviewList;
