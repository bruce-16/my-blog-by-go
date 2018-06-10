import React, {Component} from 'react';
import PropTypes from 'prop-types';
import {withStyles} from '@material-ui/core/styles';
import Avatar from '@material-ui/core/Avatar';
import Chip from '@material-ui/core/Chip';
import FaceIcon from '@material-ui/icons/Face';
import DoneIcon from '@material-ui/icons/Done';

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
    }
  }

  componentDidMount = () => {
    fetch('/get-labels')
      .then(data => data.json())
      .then(json => {
        if (json.status === 0) {
          this.renderChip(json.data)
        }
      })
      .catch(err => console.err("ERROR: ", err))
  }

  handleClick = (id) => {
    
  }

  renderChip = (data) => {
    if (Array.isArray(data)) {
      this.setState({
        chips: data.map(item => (
          <Chip
            label={item.Label}
            className={this.props.classes.chip}
            clickable
            key={JSON.stringify(item)}
            onClick={() => {
              this.handleClick(item.Id)
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
        {this.state.chips}
      </div>
    );
  }
}

Chips.propTypes = {
  classes: PropTypes.object.isRequired
};

export default withStyles(styles)(Chips);