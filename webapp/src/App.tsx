import * as React from 'react';
import { Route, Switch, BrowserRouter } from "react-router-dom";
import { routes } from "./routes"
import Container from "./Container";
import './App.css';

const App = (): JSX.Element => {
  return (
    <BrowserRouter>
        <Container>
          <Switch>
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
  );
};

export default App;
