import React from 'react';
import PropTypes from 'prop-types';
import classNames from 'classnames';
import { connect } from 'react-redux';
import { clearError } from '../actions';

class Network extends React.Component {

  static propTypes = {
    isLoading: PropTypes.bool,
    errorMessage: PropTypes.string
  }

  render() {
    const { isLoading, errorMessage } = this.props;
    const divClass = classNames(
      'overlay',
      { visible: isLoading || errorMessage }
    );
    return (
      <div className={divClass}>
        {isLoading &&
          <div className="spinner"></div>
        }
        {errorMessage &&
          <div className="error">
            <div className="header">Error</div>
            {errorMessage}
            <div className="buttons">
              <button className="button" onClick={this.props.clearError}>
                OK
              </button>
            </div>
          </div>
        }
      </div >
    );
  }
}

export default connect(
  (state) => ({
    ...state.network
  }),
  { clearError }
)(Network);
