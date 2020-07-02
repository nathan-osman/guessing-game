import { combineReducers } from 'redux';
import { SET_GAMES } from '../constants/actionTypes';
import game from './game';
import network from './network';

const lobby = (state = {}, action) => {
  switch (action.type) {
    case SET_GAMES:
      return {
        ...state,
        games: action.payload
      }
    default:
      return state;
  }
};

export default combineReducers({
  lobby,
  game,
  network
});
