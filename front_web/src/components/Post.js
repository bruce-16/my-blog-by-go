import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Paper from '@material-ui/core/Paper';
import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import { withRouter } from 'react-router-dom'

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
  const { classes, id, title, date, amount, history} = props;
  let postDate = new Date(date * 1000)
  postDate = `${postDate.getFullYear()}年${String(postDate.getMonth() + 1).padStart(2, '0')}月${String(postDate.getDate()).padStart(2, '0')}日`
  const handleClick = () => {
    history.push("/article", {id})
  }
  return (
    <Button variant="outlined" className={classes.button} onClick={handleClick}>
      <Paper className={classes.root} elevation={4}>
        <Typography variant="headline" component="h3" style={{ textAlign: 'start'}}>
          {title}
        </Typography>
        <Typography component="p" style={{ textAlign: 'start', marginTop: 20 }}>
          {postDate}
        </Typography>
        <Typography component="p" style={{ textAlign: 'end'}}>
          字数：{amount}
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

export default withRouter(withStyles(styles)(PaperSheet));