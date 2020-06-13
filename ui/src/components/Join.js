import React from 'react';

/**
 * Join presents the user with an input for entering their name
 */
class Join extends React.Component {

  render() {
    return (
      <div>
        <p>You are about to join a game.</p>
        <p>Please enter your name below to continue.</p>
        <input type="text" value={this.state.name} />
        <button type="submit">Join</button>
      </div>
    )
  }
}

export default Join;
