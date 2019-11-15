import {applyMiddleware, combineReducers, compose, createStore} from "redux";
import thunk from "redux-thunk";
import {IUserState, userActionCreator, userStateReducer} from "./redux/states/UserState";
import {composeWithDevTools} from "redux-devtools-extension";

export type AppState = {
    user: IUserState
}

const composeEnhancers = composeWithDevTools({
});
export const store = createStore(
    combineReducers<AppState>({
        user: userStateReducer
    }),
    composeEnhancers(applyMiddleware(thunk)),
);


export const actionCreator = {
    userActionCreator
};

