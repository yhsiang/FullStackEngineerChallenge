import React, {useContext} from 'react';
import {StyleSheet} from 'react-native';
import {
  Container,
  Header,
  Content,
  Text,
  Left,
  Button,
  Icon,
  Body,
  Title,
  Right,
} from 'native-base';

import RouterPackage from '../Routing';
import {EmployeeList} from '../components';
import {
  SharedDataContext,
  addReviewerToState,
  removeReviewerFromState,
  DataProvider,
} from '../contexts';
import {addReviewer, removeReviewer} from '../apis';

const {useHistory} = RouterPackage;
const styles = StyleSheet.create({
  header: {margin: 15},
});

const AssignScreen = ({match, navigation}) => {
  const history = useHistory();
  const {state, dispatch} = useContext(SharedDataContext);
  const employeeID = +match.params.employee_id;
  const employee = state.find(em => em.id === employeeID);
  const reviewers = employee.reviewers.map(it => it.id);

  const handleAdd = reviewer => {
    addReviewer(employeeID, reviewer).then(status => {
      if (status) {
        dispatch(addReviewerToState({reviewee: employeeID, reviewer}));
      }
    });
  };
  const handleRemove = reviewer => {
    removeReviewer(employeeID, reviewer).then(status => {
      if (status) {
        dispatch(removeReviewerFromState({reviewee: employeeID, reviewer}));
      }
    });
  };

  return (
    <DataProvider>
      <Container>
        <Header>
          <Left>
            <Button transparent onPress={() => history.goBack()}>
              <Icon name="arrow-back" />
            </Button>
          </Left>
          <Body>
            <Title>Reviewers</Title>
          </Body>
          <Right />
        </Header>
        <Content>
          <Text style={styles.header}>{`Please choose ${
            employee.name
          }'s reviewer`}</Text>
          <EmployeeList
            checkable
            data={state.filter(em => em.id !== employeeID)}
            reviewers={reviewers}
            addReviewer={handleAdd}
            removeReviewer={handleRemove}
          />
        </Content>
      </Container>
    </DataProvider>
  );
};

export default AssignScreen;
