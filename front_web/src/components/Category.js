import React, {Component} from 'react';
import PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';
import Chip from '@material-ui/core/Chip';


import ShowPosts from './ShowPosts'
import Post from './Post'


const styles = theme => ({
  root: {
    display: 'flex',
    justifyContent: 'center',
    flexWrap: 'wrap',
    margin: '8% 5%'
  },
  chip: {
    margin: theme.spacing.unit
  }
});


class Chips extends Component {

  constructor(props) {
    super(props)
    this.state = {
      chips: [],
      tips: "点击分类查看相应文章",
      posts: [],
      isShow: false,
      disable: true,
    }
  }

  componentDidMount = () => {
    fetch('/get-categoies')
      .then(data => data.json())
      .then(json => {
        if (json.status === 0) {
          this.renderChip(json.data)
        }
      })
      .catch(err => console.err("ERROR: ", err))
  }

  handleClick = (id, label) => {
    fetch('/get-posts-by-category/' + id)
      .then(data => data.json())
      .then(json => {
        if (json.status === 0) {
          this.setState({
            posts: json.data.map(item => (
              <Post
                title={item.Title}
                date={item.CreateTime}
                amount={item.TextAmount}
                key={JSON.stringify(item)}
                id={item.PostId}
              />
            )),
            tips: label,
            isShow: true,
            disable: false,
          })
        }
      })
      .catch(err => console.err("ERROR: ", err))
  }

  renderChip = (data) => {
    if (Array.isArray(data)) {
      this.setState({
        chips: data.map(item => (
          <Chip
            label={item.Category}
            className={this.props.classes.chip}
            clickable
            key={JSON.stringify(item)}
            onClick={() => {
              this.handleClick(item.Id, item.Category)
            }}
          />
        ))
      })
    }
  }

  render() {
    const {classes} = this.props;
    return (
      <div className={classes.root}>
        <div>{this.state.chips}</div>
        <ShowPosts tips={this.state.tips} disable={this.state.disable} isShow={this.state.isShow} onChange={this.onChange}>
        { this.state.posts}
        </ShowPosts>
      </div>
    );
  }
}

Chips.propTypes = {
  classes: PropTypes.object.isRequired
};

export default withStyles(styles)(Chips);