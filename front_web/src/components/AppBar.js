import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import {withRouter} from "react-router-dom";

const styles = {
  root: {
    flexGrow: 1,
  },
  flex: {
    flex: 1,
  },
  menuButton: {
    marginLeft: -12,
    marginRight: 20,
  },
};

function ButtonAppBar(props) {
  const { classes, rightButton, history } = props;

  const handleClick = () => {
    history.push("/")
  }

  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <Typography variant="title" color="inherit" className={classes.flex} onClick={handleClick}>
            Zachrey's blog
          </Typography>
          {rightButton}
        </Toolbar>
      </AppBar>
    </div>
  );
}

ButtonAppBar.propTypes = {
  classes: PropTypes.object.isRequired,
  rightButton: PropTypes.arrayOf(PropTypes.element)
};

export default withRouter(withStyles(styles)(ButtonAppBar));