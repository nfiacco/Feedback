import { useEffect } from "react";
import {
  BrowserRouter, Route, Switch
} from "react-router-dom";
import { Feedback } from 'src/components/feedback/Feedback';
import { Home } from 'src/components/home/Home';
import { NotFound } from 'src/components/NotFound';
import { useSelector } from "src/root/model";
import { useStart } from "./actions";

export const App: React.FC = () => {
  const loading = useSelector(state => state.app.loading);
  const start = useStart();

  useEffect(() => {
    start();
  });

  if (loading) {
    return <div/>
  }

  return (
    <BrowserRouter>
      <Switch>
      <Route exact path='/'>
        <Home/>
      </Route>
      <Route path='/feedback/:id'>
        <Feedback/>
      </Route>
      <Route path='/'>
        <NotFound/>
      </Route>
    </Switch>
    </BrowserRouter>
  );
}
