import { combineReducers } from 'redux';
import {
  SET_GAMES
} from '../constants/actionTypes';
import game from './game';

const lobby = (state = {}, action) => {
  switch (action.type) {
    case SET_GAMES:
      return {
        ...state,
        games: action.payload.games
      };
    default:
      return state;
  }
};

export default combineReducers({
  lobby,
  game,
  network
});
