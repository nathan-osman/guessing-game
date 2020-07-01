import { combineReducers } from 'redux';
import appReducer from './app';
import gameReducer from './game';
import playersReducer from './players';

export default combineReducers({
  appReducer,
  gameReducer,
  playersReducer
});
