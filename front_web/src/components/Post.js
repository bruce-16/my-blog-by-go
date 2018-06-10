import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';

const styles = theme => ({
  root: theme.mixins.gutters({
    width: '100%',
    paddingTop: 16,
    paddingBottom: 16,
    // 
  }),
  button: {
    marginTop: theme.spacing.unit * 2,
    dispaly: 'block',
    width: '100%',
    border: 'none',
  },
});

function PaperSheet(props) {
  const { classes } = props;
  return (
    <Button variant="outlined" className={classes.button}>
      <Paper className={classes.root} elevation={4}>
        <Typography variant="headline" component="h3" style={{ textAlign: 'start'}}>
          第一篇博客~
        </Typography>
        <Typography component="p" style={{ textAlign: 'start'}}>
          2018年05月27日
        </Typography>
        <Typography component="p" style={{ textAlign: 'end'}}>
          字数：200
        </Typography>
      </Paper>
    </Button>
  );
}

PaperSheet.propTypes = {
  classes: PropTypes.object.isRequired,
  title: PropTypes.string.isRequired,
  date: PropTypes.string.isRequired,
  amount: PropTypes.string.isRequired,
};

export default withStyles(styles)(PaperSheet);