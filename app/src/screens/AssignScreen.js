import React, {useContext, useState} from 'react';
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
  Form,
  Item,
  Label,
  Input,
} from 'native-base';

import RouterPackage from '../Routing';
import {EmployeeList} from '../components';
import {
  SharedDataContext,
  addReviewerToState,
  removeReviewerFromState,
  DataProvider,
  updateEmployeeToState,
  removeEmployeeFromState,
} from '../contexts';
import {
  addReviewer,
  removeReviewer,
  updateEmployee,
  removeEmployee,
} from '../apis';

const {useHistory} = RouterPackage;
const styles = StyleSheet.create({
  header: {margin: 15},
  button: {margin: 15, marginTop: 20},
});

const AssignScreen = ({match, navigation}) => {
  const history = useHistory();
  const {state, dispatch} = useContext(SharedDataContext);
  const employeeID = +match.params.employee_id;
  const employee = state.find(em => em.id === employeeID);
  const reviewers = employee.reviewers.map(it => it.id);
  const [name, setName] = useState(employee.name);

  const updateName = () => {
    updateEmployee(employeeID, name).then(data => {
      dispatch(updateEmployeeToState(data));
    });
  };

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

  const deleteEmployee = () => {
    history.goBack();
    removeEmployee(employeeID).then(status => {
      if (status) {
        dispatch(removeEmployeeFromState(employeeID));
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
          <Form>
            <Item fixedLabel>
              <Label>Employee Name</Label>
              <Input
                autoCapitalize="none"
                value={name}
                onChangeText={text => setName(text)}
              />
            </Item>
          </Form>
          <Button block style={styles.button} onPress={updateName}>
            <Text>Update Name</Text>
          </Button>
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
          <Button block style={styles.button} onPress={deleteEmployee} danger>
            <Text>Delete Employee</Text>
          </Button>
        </Content>
      </Container>
    </DataProvider>
  );
};

export default AssignScreen;
