import React, {useContext, useState} from 'react';
import {StyleSheet} from 'react-native';
import {useHistory} from 'react-router-native';
import {
  Container,
  Header,
  Content,
  Form,
  Item,
  Input,
  Label,
  Button,
  Text,
  Title,
  H1,
} from 'native-base';
import AuthContext from '../contexts/AuthContext';
import {signIn} from '../apis';

const styles = StyleSheet.create({
  h1: {margin: 15},
  button: {margin: 15, marginTop: 50},
});

const SignInScreen = () => {
  const auth = useContext(AuthContext);
  const history = useHistory();
  const [user, setUser] = useState('');
  const [pass, setPass] = useState('');

  const handleSubmit = () => {
    signIn(user, pass).then(data => {
      if (data) {
        auth.authenticate(() => {
          if (data.role === 'user') {
            history.push(`/${data.role}/${data.id}`);
          } else {
            history.push(`/${data.role}`);
          }
        });
      } else {
        // TODO: show error message
      }
    });
  };

  return (
    <Container>
      <Header>
        <Title>Login</Title>
      </Header>
      <Content>
        <H1 style={styles.h1}>Welcome to Review 360</H1>
        <Form>
          <Item fixedLabel>
            <Label>Username</Label>
            <Input
              autoCapitalize="none"
              value={user}
              onChangeText={text => setUser(text)}
            />
          </Item>
          <Item fixedLabel last>
            <Label>Password</Label>
            <Input
              secureTextEntry
              value={pass}
              onChangeText={text => setPass(text)}
            />
          </Item>
        </Form>
        <Button
          block
          style={styles.button}
          onPress={handleSubmit}
          onClick={handleSubmit}>
          <Text>Sign In</Text>
        </Button>
      </Content>
    </Container>
  );
};

export default SignInScreen;
