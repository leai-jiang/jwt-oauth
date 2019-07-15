import * as React  from 'react';
import { Route, Redirect, Switch, BrowserRouter } from "react-router-dom";
import { routes } from "./routes"
import Container from "./Container";
import reducer from "./context/reducer";
import GlobalStore from "./context/GlobalStore";
import './App.css';

const { useContext, useReducer } = React;

const App = (): JSX.Element => {
  const store = useContext(GlobalStore);
  const [state, dispatch] = useReducer(reducer, store);

  return (
    <GlobalStore.Provider value={{ state, dispatch }}>
      <BrowserRouter>
          <Container>
            <Switch>
              <Redirect path="/" exact to="/home" />
              {routes.map(({ key, path, exact = true, component }) => (
                <Route
                  key={key}
                  path={path}
                  exact={exact}
                  component={component}
                />
              ))}
            </Switch>
          </Container>
      </BrowserRouter>
    </GlobalStore.Provider>
  );
};

export default App;
