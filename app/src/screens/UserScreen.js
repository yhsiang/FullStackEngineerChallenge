import React, {useContext, useState, useEffect} from 'react';
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
} from 'native-base';

import {getEmployee} from '../apis';
import {AuthContext} from '../contexts';
import EmployeeList from '../components/EmployeeList';
import storage from '../storage';

const UserScreen = ({match}) => {
  const auth = useContext(AuthContext);
  const history = useHistory();
  const id = match.params.id;
  const [reviewees, setReviewees] = useState([]);

  useEffect(() => {
    getEmployee(id).then(data => setReviewees(data.reviewees));
  }, [id]);

  const handleSignOut = () => {
    storage.remove({
      key: 'token',
    });
    auth.signout(() => {
      history.push('/');
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
