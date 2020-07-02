import {
  ASYNC_START,
  ASYNC_END,
  SET_ERROR
} from "../constants/actionTypes";

const API_HOST = process.env.REACT_APP_API_HOST;
const transformURL = (typeof API_HOST !== 'undefined')
  ? u => `http://${API_HOST}${u}`
  : u => u;

export default store => next => action => {
  const { url, data } = action['url'];
  if (typeof url === 'undefined') {
    return next(action);
  }
  store.dispatch({ type: ASYNC_START });
  fetch(transformURL(url), data)
    .then(response => {
      if (!response.ok) {
        throw Error(response.statusText);
      }
      return response;
    })
    .then(response => response.json())
    .then(json => {
      next({
        ...action,
        payload: json
      });
    })
    .catch(error => {
      store.dispatch({
        type: SET_ERROR,
        message: "Unable to retrieve the list of active games"
      });
    })
    .finally(() => {
      store.dispatch({ type: ASYNC_END });
    });
};
