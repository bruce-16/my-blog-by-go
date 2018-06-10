import React, { Component } from 'react';

import Post from './Post'

class Home extends Component {

  constructor(props) {
    super(props);
    this.state = {
      posts: [],
    }
  }
  
  componentDidMount = async () => {
    let data = await fetch("/get-posts")
    let json = await data.json()
    if (json.status === 0) {
      this.setState({
        posts: json.data.map(item => (
          <Post
            title={item.Title}
            date={item.CreateTime}
            amount={item.TextAmount}
            key={JSON.stringify(item)}
            id={item.Id}
          />
        ))
      })
    }
  }

  render() {
    return (
      <div style={{ width: '100%' }}>
        {this.state.posts}
      </div>
    );
  }
}

export default Home;