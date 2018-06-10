import React, { Component } from 'react';

import Post from './Post'

class Home extends Component {
  render() {
    return (
      <div style={{ width: '100%' }}>
        <Post />
      </div>
    );
  }
}

export default Home;