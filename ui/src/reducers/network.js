import {
  ASYNC_START,
  ASYNC_END,
  SET_ERROR,
  CLEAR_ERROR
} from '../constants/actionTypes';

export default (state = {}, action) => {
  switch (action.type) {
    case ASYNC_START:
      return {
        ...state,
        isLoading: true
      };
    case ASYNC_END:
      return {
        ...state,
        isLoading: false
      };
    case SET_ERROR:
      return {
        ...state,
        errorMessage: action.message
      };
    case CLEAR_ERROR:
      return {
        ...state,
        errorMessage: null
      };
    default:
      return state;
  };
};
