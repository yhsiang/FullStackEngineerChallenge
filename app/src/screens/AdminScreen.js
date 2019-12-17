import React, {useContext, useState, useEffect} from 'react';
import {StyleSheet} from 'react-native';
import RouterPackage from '../Routing';
import {
  Container,
  Header,
  Content,
  Text,
  Form,
  Button,
  Item,
  Label,
  Input,
  Title,
  Left,
  Body,
  Right,
  Icon,
  Tabs,
  Tab,
  Textarea,
} from 'native-base';

import {
  AuthContext,
  SharedDataContext,
  addData,
  fetchInitial,
  PageContext,
} from '../contexts';
import {getEmployees, createEmployee, getReviews, createReview} from '../apis';
import {EmployeeList, ReviewList, DataPicker} from '../components';
import storage from '../storage';

const {useHistory} = RouterPackage;
const styles = StyleSheet.create({
  list: {margin: 15},
  button: {margin: 15, marginBottom: 50},
});

const AdminScreen = () => {
  const auth = useContext(AuthContext);
  const history = useHistory();
  const [name, setName] = useState('');
  const [reviews, setReviews] = useState([]);
  const [review, setReview] = useState('');
  const [reviewer, setReviewer] = useState(undefined);
  const [reviewee, setReviewee] = useState(undefined);
  const {state, dispatch} = useContext(SharedDataContext);
  const {page, setPage} = useContext(PageContext);

  useEffect(() => {
    getEmployees().then(data => dispatch(fetchInitial(data)));
    getReviews().then(data => setReviews(data));
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const addEmployee = () => {
    createEmployee(name).then(employee => {
      dispatch(addData(employee));
      setName('');
    });
  };

  const handleSignOut = () => {
    storage.remove({
      key: 'token',
    });
    auth.signout(() => {
      history.push('/');
    });
  };

  const handleAddReview = () => {
    createReview(reviewee, reviewer, review).then(data => {
      if (data.id) {
        setReviewee(undefined);
        setReviewer(undefined);
        setReview('');
        setReviews([...reviews, data]);
      }
    });
  };

  return (
    <Container>
      <Header hasTabs>
        <Left />
        <Body>
          <Title>Admin View</Title>
        </Body>
        <Right>
          <Button transparent onPress={handleSignOut}>
            <Icon name="log-out" />
          </Button>
        </Right>
      </Header>
      <Tabs page={page} onChangeTab={tab => setPage(tab.i)}>
        <Tab heading="Employees">
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
            <Button block style={styles.button} onPress={addEmployee}>
              <Text>Add</Text>
            </Button>
            <EmployeeList data={state} />
          </Content>
        </Tab>
        <Tab heading="Reviews">
          <Content>
            <Form>
              <Item picker fixedLabel>
                <Label>Reviewee</Label>
                <DataPicker
                  data={state}
                  selectedValue={reviewee}
                  onValueChange={id => setReviewee(id)}
                />
              </Item>
              <Item picker fixedLabel>
                <Label>Reviewer</Label>
                <DataPicker
                  data={state}
                  selectedValue={reviewer}
                  onValueChange={id => setReviewer(id)}
                />
              </Item>
              <Textarea
                style={styles.list}
                rowSpan={7}
                bordered
                placeholder="Please write down your review"
                onChangeText={text => setReview(text)}>
                {review}
              </Textarea>
            </Form>
            <Button block style={styles.button} onPress={handleAddReview}>
              <Text>Add Review</Text>
            </Button>
            <ReviewList data={reviews} />
          </Content>
        </Tab>
      </Tabs>
    </Container>
  );
};

export default AdminScreen;
