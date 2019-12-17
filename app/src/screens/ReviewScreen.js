import React, {useEffect, useState} from 'react';
import {StyleSheet} from 'react-native';
import RouterPackage from '../Routing';
import {
  Container,
  Header,
  Content,
  Button,
  Title,
  Left,
  Body,
  Right,
  Icon,
  Form,
  Textarea,
  H1,
  Text,
} from 'native-base';

import {getEmployee, createReview, getReview, updateReview} from '../apis';
const {useHistory} = RouterPackage;
const styles = StyleSheet.create({
  field: {marginTop: 30},
  textArea: {padding: 10},
});

const ReviewScreen = ({match: {path, params}}) => {
  const [name, setName] = useState('');
  const [content, setContent] = useState('');
  let reviewee, reviewer, reviewID;
  const isEdit = path.match(/^\/edit/);
  if (isEdit) {
    reviewID = +params.review_id;
  } else {
    reviewee = +params.reviewee_id;
    reviewer = +params.reviewer_id;
  }

  const history = useHistory();
  useEffect(() => {
    if (isEdit) {
      getReview(reviewID).then(data => setContent(data.content));
    } else {
      getEmployee(reviewee).then(data => setName(data.name));
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const handleSubmit = () => {
    if (isEdit) {
      updateReview(reviewID, content).then(data => {
        setContent(data.content);
      });
    } else {
      createReview(reviewee, reviewer, content).then(() => {
        setContent('');
        history.goBack('/');
      });
    }
  };

  return (
    <Container>
      <Header>
        <Left>
          <Button transparent onPress={() => history.goBack()}>
            <Icon name="arrow-back" />
          </Button>
        </Left>
        <Body>
          <Title>Feedback</Title>
        </Body>
        <Right />
      </Header>
      <Content padder>
        <H1>{isEdit ? 'Update Review' : `Review ${name}`}</H1>
        <Form style={styles.field}>
          <Textarea
            rowSpan={7}
            bordered
            placeholder="Please write down your review"
            onChangeText={text => setContent(text)}>
            {content}
          </Textarea>
        </Form>
        <Button block style={styles.field} onPress={handleSubmit}>
          <Text>{isEdit ? 'Update Review' : 'Submit Feedback'}</Text>
        </Button>
      </Content>
    </Container>
  );
};

export default ReviewScreen;
