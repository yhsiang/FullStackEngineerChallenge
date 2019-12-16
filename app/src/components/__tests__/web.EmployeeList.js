import React from 'react';
import {render} from '@testing-library/react';
import '@testing-library/jest-dom/extend-expect';
import EmployeeList from '../EmployeeList';

test('should be render properly', () => {
  const data = [
    {id: 1, name: 'test1'},
    {id: 1, name: 'test2'},
    {id: 1, name: 'test3'},
  ];
  const test4 = 'test4';
  const {queryByText, getByText} = render(<EmployeeList data={data} />);

  expect(queryByText(test4)).toBeNull();

  expect(getByText(data[0].name)).toBeInTheDocument();
});
