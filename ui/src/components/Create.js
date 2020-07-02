import React from 'react';
import { connect } from 'react-redux';
import { createGame } from '../actions';
import { withRouter } from 'react-router-dom';

class Create extends React.Component {

  constructor(props) {
    super(props);
    this.state = { value: '' };
  }

  handleChange = (event) => {
    this.setState({ value: event.target.value });
  }

  createGame = () => {
    this.props.createGame(this.state.value)
  }

  goBack = () => {
    this.props.history.goBack();
  }

  render() {
    return (
      <div className="container">
        <p>Enter a name for your game so that others can easily find it:</p>
        <input
          className="input"
          type="text"
          name="name"
          value={this.state.value}
          onChange={this.handleChange}
          autoFocus />
        <div className="buttons">
          <button className="button" onClick={this.createGame}>Create</button>
          <button className="button" onClick={this.goBack}>
            Back
          </button>
        </div>
      </div>
    );
  }
}

export default connect(null, { createGame })(withRouter(Create));
