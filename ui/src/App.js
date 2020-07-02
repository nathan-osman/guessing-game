import React from 'react';
import {
  BrowserRouter as Router,
  Route,
  Switch
} from 'react-router-dom';
import Create from './components/Create';
import Header from './components/Header';
import Home from './components/Home';
import Network from './components/Network';

class App extends React.Component {

  render() {
    return (
      <Router>
        <Network />
        <Header />
        <Switch>

          {/* Home page showing the list of games */}
          <Route exact path="/">
            <Home />
          </Route>

          {/* Create new game page */}
          <Route exact path="/create">
            <Create />
          </Route>

        </Switch>
      </Router>
    )
  }
}

export default App;
