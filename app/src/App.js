import React from 'react';
import {View, StyleSheet} from 'react-native';
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
} from 'native-base';

const styles = StyleSheet.create({
  buttonContainer: {
    padding: 16,
  },
});

const App = () => {
  return (
    <Container>
      <Header />
      <Content>
        <Form>
          <Item fixedLabel>
            <Label>Username</Label>
            <Input />
          </Item>
          <Item fixedLabel last>
            <Label>Password</Label>
            <Input />
          </Item>
        </Form>
        <View style={styles.buttonContainer}>
          <Button>
            <Text>Success</Text>
          </Button>
        </View>
      </Content>
    </Container>
  );
};

export default App;
