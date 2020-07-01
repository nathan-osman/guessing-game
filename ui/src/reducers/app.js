import {
  SET_GAMES
} from '../constants/actionTypes';

const initialState = {
  games: []
};

export default (state = initialState, action) => {
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
