import {
  ADD_PLAYER,
  REMOVE_PLAYER
} from '../constants/actionTypes';

const initialState = {
  players: {}
};

export default (state = initialState, action) => {
  switch (action.type) {
    case ADD_PLAYER:
      return {
        ...state,
        players: {
          ...state.players,
          [action.payload.guid]: action.payload
        }
      }
    case REMOVE_PLAYER:
      state = Object.assign({}, state)
      delete state.players[action.payload.guid]
      return state
    default:
      return state;
  }
}
