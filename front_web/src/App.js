import React, { Component } from 'react';
import Button from '@material-ui/core/Button';
import { BrowserRouter as Router, Route, Link } from "react-router-dom";

import './App.css';

import AppBar from './components/AppBar'
import Home from './components/Home'
import Label from './components/Label'
import Category from './components/Category'



class App extends Component {
  buttons = [
    <Button color="inherit"><Link to="/label" style={{ textDecoration: 'none', color: '#fff' }}>标签</Link></Button>,
    <Button color="inherit"><Link to="/category" style={{ textDecoration: 'none', color: '#fff' }}>分类</Link></Button>
  ]
  render() {
    return (
      <Router>
        <div>
          <AppBar rightButton={this.buttons}/>
          <Route exact path="/" component={Home} />
          <Route exact path="/label" component={Label} />
          <Route exact path="/category" component={Category} />
        </div>
      </Router>
    );
  }
}

export default App;
