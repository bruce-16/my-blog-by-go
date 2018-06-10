import React, { Component } from 'react';
import Button from '@material-ui/core/Button';
import { BrowserRouter as Router, Route, Link } from "react-router-dom";

import './App.css';

import AppBar from './components/AppBar'
import Home from './components/Home'
import Label from './components/Label'
import Category from './components/Category'
import Article from './components/Article'



class App extends Component {
  buttons = [
    <Button key={1} color="inherit"><Link to="/label" style={{ textDecoration: 'none', color: '#fff' }}>标签</Link></Button>,
    <Button key={2} color="inherit"><Link to="/category" style={{ textDecoration: 'none', color: '#fff' }}>分类</Link></Button>
  ]
  render() {
    return (
      <Router>
        <div>
          <AppBar rightButton={this.buttons}/>
          <Route exact path="/" component={Home} />
          <Route exact path="/label" component={Label} />
          <Route exact path="/category" component={Category} />
          <Route exact path="/article" component={Article} />
        </div>
      </Router>
    );
  }
}

export default App;
