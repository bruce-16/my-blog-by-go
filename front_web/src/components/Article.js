import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';
import withRouter from 'react-router-dom/withRouter';

const styles = theme => ({
  root: theme.mixins.gutters({
    paddingTop: 16,
    paddingBottom: 16,
    marginTop: theme.spacing.unit * 3,
  }),
});

class PaperSheet extends Component {
  constructor(props) {
    super(props);
    this.state = {
      html: "",
    }
  }
  componentDidMount = async () => {
    const { location } = this.props;
    const data = await fetch("/get-html-str/" + location.state.id)
    const json = await data.json()
    if (json.status === 0) {
      this.setState({
        html: json.data,
      })
    }
  } 
  render() {
    const { classes } = this.props;
    return (
      <div>
        <Paper className={classes.root} elevation={4}>
          <div
            dangerouslySetInnerHTML={{ __html: this.state.html }}
            style={{ padding: '2% 25%' }}>
          </div>
        </Paper>
      </div>
    );
  }
}

export default withRouter(withStyles(styles)(PaperSheet));