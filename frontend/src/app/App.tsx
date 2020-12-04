import { useEffect } from "react";
import { connect, Provider } from "react-redux";
import {
  BrowserRouter, Route, Switch
} from "react-router-dom";
import { start } from "src/app/actions";
import { NotFound } from 'src/app/NotFound';
import { Feedback } from 'src/feedback/Feedback';
import { Home } from 'src/home/Home';
import { createStore, RootState } from 'src/root/model';
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
      <BrowserRouter>
        <ConnectedApp/>
      </BrowserRouter>
    </Provider>
  );
}
