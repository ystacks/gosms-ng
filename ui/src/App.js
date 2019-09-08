import React, { Component } from 'react';
import { Route, Switch } from 'react-router-dom';
import { Content } from 'carbon-components-react/lib/components/UIShell';
import './App.scss';
import UIHeader from './components/Header';
import LandingPage from './contents/LandingPage';
import SMSPage from './contents/SMSPage';

class App extends Component {
  render() {
    return (
      <div className="App">
        <UIHeader />
        <Content>
          <Switch>
            <Route exact path="/" component={SMSPage} />
            <Route path="/landing" component={LandingPage} />
          </Switch>
        </Content>
      </div>
    );
  }
}

export default App;
