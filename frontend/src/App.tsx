import React from 'react';
import './App.css';
import {SignIn} from './SignIn';
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";

function App() {
  return (
    <Router>
      <Switch>
        <Route path='/'>
          <SignIn/>
        </Route>
      </Switch>
    </Router>
  );
}

export default App;
