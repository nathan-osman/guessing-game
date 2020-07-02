import {
  CLEAR_ERROR,
  SET_GAMES,
  CREATE_GAME
} from "../constants/actionTypes";

export const clearError = () => ({ type: CLEAR_ERROR });

export const loadGames = () => ({
  type: SET_GAMES,
  url: '/api/games',
  errorMessage: "Unable to load the list of games."
});

export const createGame = (name) => ({
  type: CREATE_GAME,
  url: '/api/create',
  data: {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ name })
  },
  errorMessage: "Unable to create a new game."
});
