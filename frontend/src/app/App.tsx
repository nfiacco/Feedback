import { start } from "app/actions";
import { createStore, RootState } from 'app/model';
import { NotFound } from 'app/NotFound';
import { Feedback } from 'feedback/Feedback';
import { Home } from 'home/Home';
import { useEffect } from "react";
import { connect, Provider } from "react-redux";
import {
  BrowserRouter as Router,

  Route, Switch
} from "react-router-dom";
import './App.css';

interface AppProps {
  loading: boolean,
  start: () => void;
}

const AppImpl: React.FC<AppProps> = props => {
  useEffect(() => {
    props.start();
  });

  if (props.loading) {
    return <div>Loading!</div>
  }

  return (
    <Switch>
      <Route exact path='/'>
        <Home/>
      </Route>
      <Route path='/:id'>
        <Feedback/>
      </Route>
      <Route path='/'>
        <NotFound/>
      </Route>
    </Switch>
  );
}

const Connector = connect(
  (state: RootState) => ({
    loading: state.app.loading,
  }),
  {
    start: start,
  }
);
const ConnectedApp = Connector(AppImpl);

export function App() {
  const store = createStore();

  return (
    <Provider store={store}>
      <Router>
        <ConnectedApp/>
      </Router>
    </Provider>
  );
}
