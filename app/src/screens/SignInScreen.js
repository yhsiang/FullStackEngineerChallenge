import React, {useContext, useState, useEffect} from 'react';
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
import {Base64} from 'js-base64';

import AuthContext from '../contexts/AuthContext';
import {signIn} from '../apis';
import storage from '../storage';

const decode = token => {
  const [, payload] = token.split('.');
  const str = Base64.decode(payload);
  const data = JSON.parse(str);

  return data.id;
};

const styles = StyleSheet.create({
  h1: {margin: 15},
  button: {margin: 15, marginTop: 50},
});

const SignInScreen = () => {
  const auth = useContext(AuthContext);
  const history = useHistory();
  const [user, setUser] = useState('');
  const [pass, setPass] = useState('');

  useEffect(() => {
    if (!auth.isAuthenticated) {
      storage
        .load({key: 'token'})
        .then(token => {
          const userId = decode(token);
          auth.authenticate(token, () => {
            if (userId === 'admin') {
              history.push('/admin');
            } else if (userId.match(/user(\d+)/)) {
              const id = userId.match(/user(\d+)/)[1];
              history.push(`/user/${id}`);
            }
          });
        })
        .catch(() => {});
    }
  }, [auth, history]);

  const handleSubmit = () => {
    signIn(user, pass).then(data => {
      if (data) {
        storage.save({
          key: 'token',
          data: data.token,
        });
        const userId = decode(data.token);
        auth.authenticate(data.token, () => {
          if (userId === 'admin') {
            history.push('/admin');
          } else if (userId.match(/user(\d+)/)) {
            const id = userId.match(/user(\d+)/)[1];
            history.push(`/user/${id}`);
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
        <Button block style={styles.button} onPress={handleSubmit}>
          <Text>Sign In</Text>
        </Button>
      </Content>
    </Container>
  );
};

export default SignInScreen;
