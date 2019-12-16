import React from 'react';
import RouterPackage from '../Routing';
import {
  Text,
  List,
  ListItem,
  Left,
  Right,
  Icon,
  CheckBox,
  Body,
} from 'native-base';

const {useHistory} = RouterPackage;

const EmployeeList = ({
  data,
  checkable = false,
  reviewers = [],
  addReviewer,
  removeReviewer,
  reviewer = null,
}) => {
  const history = useHistory();
  if (data.length === 0) {
    return <Text>No Data</Text>;
  }

  if (checkable) {
    return (
      <List>
        {data.map(d => {
          const checked = reviewers.indexOf(d.id) !== -1;
          return (
            <ListItem key={d.id}>
              <CheckBox
                checked={checked}
                onPress={() =>
                  checked ? removeReviewer(d.id) : addReviewer(d.id)
                }
              />
              <Body>
                <Text>{d.name}</Text>
              </Body>
            </ListItem>
          );
        })}
      </List>
    );
  }

  return (
    <List>
      {data.map(d => {
        return (
          <ListItem
            key={d.id}
            button
            onPress={() => {
              if (reviewer && d.review_id > 0) {
                history.push(`/edit/${d.review_id}`);
              } else if (reviewer && d.review_id === 0) {
                history.push(`/review/${d.id}/${reviewer}`);
              } else {
                history.push(`/assign/${d.id}`);
              }
            }}>
            <Left>
              <Text>{d.name}</Text>
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

export default EmployeeList;
