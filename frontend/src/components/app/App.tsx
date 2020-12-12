import React, { useEffect } from "react";
import {
  BrowserRouter, Route, Switch
} from "react-router-dom";
import { About } from "src/components/about/About";
import { useStart } from "src/components/app/actions";
import { Feedback } from 'src/components/feedback/Feedback';
import { Header } from "src/components/header/Header";
import { Home } from 'src/components/home/Home';
import { NotFound } from 'src/components/notfound/NotFound';
import { useSelector } from "src/root/model";

export const App: React.FC = () => {
  const loading = useSelector(state => state.app.loading);
  const start = useStart();

  useEffect(() => {
    start();
  }, [start]);

  if (loading) {
    return <div data-testid="loading"/>
  }

  return (
    <BrowserRouter>
      <Header/>
      <Switch>
        <Route exact path='/'>
          <Home/>
        </Route>
        <Route path='/about'>
          <About/>
        </Route>
        <Route path='/feedback/:key'>
          <Feedback/>
        </Route>
        <Route path='*'>
          <NotFound/>
        </Route>
      </Switch>
    </BrowserRouter>
  );
}
