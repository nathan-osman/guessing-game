import {
  SET_GAMES,
  CLEAR_ERROR
} from "../constants/actionTypes";

export const clearError = () => ({ type: CLEAR_ERROR });

export const loadGames = () => ({
  type: SET_GAMES,
  url: '/api/games'
});
