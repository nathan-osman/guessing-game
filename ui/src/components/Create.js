import React from 'react';
import { withRouter } from 'react-router-dom';

class Create extends React.Component {

  render() {
    return (
      <div className="container">
        <p>Enter a name for your game so that others can easily find it:</p>
        <input className="input" type="text" name="name" autoFocus />
        <div className="buttons">
          <button className="button">Create</button>
          <button className="button" onClick={this.props.history.goBack}>
            Back
          </button>
        </div>
      </div>
    );
  }
}

export default withRouter(Create);
