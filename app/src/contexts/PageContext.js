import React, {useState} from 'react';

const init = {page: 0, setPage: () => {}};

const PageContext = React.createContext(init);

export const PageProvider = props => {
  const [page, setPage] = useState(init.page);

  return (
    <PageContext.Provider value={{page, setPage}}>
      {props.children}
    </PageContext.Provider>
  );
};

export default PageContext;
