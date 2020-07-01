import {
  START_GAME,
  SET_QUESTION,
  SET_ANSWERS,
  PROCESS_GUESS,
  RESTART_GAME
} from '../constants/actionTypes';
import {
  WAITING_FOR_PLAYERS,
  WAITING_FOR_QUESTION,
  WAITING_FOR_ANSWERS,
  WAITING_FOR_GUESS
} from '../constants/gameStates';

const initialState = {
  state: WAITING_FOR_PLAYERS
};

export default (state = initialState, action) => {
  switch (action.type) {
    case START_GAME:
      return {
        ...state,
        playerSequence: action.payload.player_sequence,
        state: WAITING_FOR_QUESTION
      }
    case SET_QUESTION:
      return {
        ...state,
        question: action.payload.question,
        state: WAITING_FOR_ANSWERS
      }
    case SET_ANSWERS:
      return {
        ...state,
        answers: action.payload.answers,
        state: WAITING_FOR_GUESS
      }
    case PROCESS_GUESS:
      return {
        ...state,
        //...
      }
    case RESTART_GAME:
      return {
        ...state,
        //...
      }
    default:
      return state;
  }
};
