import {
  ADD_PLAYER,
  REMOVE_PLAYER
} from '../constants/actionTypes';

export default (state = {}, action) => {
  switch (action.type) {
    case ADD_PLAYER:
      return {
        ...state,
        [action.payload.guid]: action.payload
      }
    case REMOVE_PLAYER:
      state = Object.assign({}, state)
      delete state[action.payload.guid]
      return state
    default:
      return state;
  }
}
