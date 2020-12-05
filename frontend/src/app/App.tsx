import { useEffect } from "react";
import { useSelector } from "react-redux";
import {
  BrowserRouter, Route, Switch
} from "react-router-dom";
import { NotFound } from 'src/app/NotFound';
import { Feedback } from 'src/feedback/Feedback';
import { Home } from 'src/home/Home';
import { RootState } from 'src/root/model';
import { useStart } from "./actions";
import './App.css';

export const App: React.FC = () => {
  const loading = useSelector((state: RootState) => state.app.loading);
  const start = useStart();

  useEffect(() => {
    start();
  });

  if (loading) {
    return <div></div>
  }

  return (
    <BrowserRouter>
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
    </BrowserRouter>
  );
}
