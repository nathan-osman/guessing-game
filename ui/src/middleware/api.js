import {
  ASYNC_START,
  ASYNC_END,
  SET_ERROR
} from "../constants/actionTypes";

export default store => next => action => {
  let url = action['url'];
  if (typeof url === 'undefined') {
    return next(action);
  }
  const prefix = process.env.REACT_APP_API_PREFIX;
  if (typeof prefix !== 'undefined') {
    url = `${prefix}${url}`;
  }
  store.dispatch({ type: ASYNC_START });
  fetch(url)
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
