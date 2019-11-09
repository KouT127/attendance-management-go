import { applyMiddleware, combineReducers, createStore } from "redux";
import thunk from "redux-thunk";

export type AppState = {}


export const store = createStore(
  combineReducers<AppState>({}),
  applyMiddleware(thunk)
);


export const actionCreator = {};

