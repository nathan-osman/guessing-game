import React from 'react';

class Home extends React.Component {

  render() {
    return (
      <div className="container">
        <p>Welcome to the guessing game!</p>
        <p>
          Games that have been created are shown below.
          You can join one or create your own!
        </p>
      </div>
    );
  }
}

export default Home;
