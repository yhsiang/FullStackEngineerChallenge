import React, {useContext, useState, useEffect} from 'react';
// import {StyleSheet} from 'react-native';
import {useHistory} from 'react-router-native';
import {
  Container,
  Header,
  Content,
  Text,
  Button,
  Title,
  Left,
  Body,
  Right,
  Icon,
  // Form,
  // Textarea,
} from 'native-base';

import {signOut, getEmployee} from '../apis';
import {AuthContext} from '../contexts';
import EmployeeList from '../components/EmployeeList';

// const styles = StyleSheet.create({
//   list: {margin: 15},
//   button: {margin: 15, marginBottom: 50},
// });

const UserScreen = ({match}) => {
  const auth = useContext(AuthContext);
  const history = useHistory();
  // const {state, dispatch} = useContext(SharedDataContext);
  const id = match.params.id;
  const [reviewees, setReviewees] = useState([]);

  // const [name, setName] = useState('');
  // const {state, dispatch} = useContext(SharedDataContext);

  useEffect(() => {
    getEmployee(id).then(data => setReviewees(data.reviewees));
  }, [id]);

  const handleSignOut = () => {
    signOut().then(() => {
      auth.signout(() => {
        history.push('/');
      });
    });
  };
  return (
    <Container>
      <Header>
        <Left />
        <Body>
          <Title>Reviewee</Title>
        </Body>
        <Right>
          <Button transparent onPress={handleSignOut} onClick={handleSignOut}>
            <Icon name="log-out" />
          </Button>
        </Right>
      </Header>
      <Content padder>
        <Text>Choose one and write down your feedback.</Text>
        <EmployeeList data={reviewees} reviewer={id} />
      </Content>
    </Container>
  );
};

export default UserScreen;
