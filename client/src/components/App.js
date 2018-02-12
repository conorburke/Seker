import React, { Component } from 'react';
import { connect } from 'react-redux';
import { BrowserRouter, Route } from 'react-router-dom';

import './App.css';
import * as actions from '../actions';
import Landing from './Landing';
import Tools from './Tools';

class App extends Component {
  componentDidMount() {
    this.props.fetchTools();
  }

  render() {
    return (
      <div>
          <BrowserRouter>
              <div id='main'>
                  <Route exact path='/' component={Landing} />
                  <Route path='/tools' component={Tools} />
              </div>
          </BrowserRouter>
      </div>
    ); 
  }
}

export default connect(null, actions)(App);