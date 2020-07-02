import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { loadGames } from '../actions';
import { Link } from 'react-router-dom';

class Home extends React.Component {

  static propTypes = {
    games: PropTypes.array
  }

  componentDidMount() {
    this.props.loadGames();
  }

  render() {
    const games = this.props.payload || [];
    return (
      <div className="container">
        <p>Welcome to the guessing game!</p>
        <p>
          Games that have been created are shown below.
          You can join one or create your own!
        </p>
        {this.props.games &&
          this.props.games.map(game => {
            return (
              <div>{game.name}</div>
            );
          })
        }
        <div className="buttons">
          <Link className="button" to="/create">Create</Link>
        </div>
      </div>
    );
  }
}

export default connect(
  (state) => ({
    ...state.lobby
  }),
  { loadGames }
)(Home);
