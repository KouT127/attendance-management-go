import {applyMiddleware, combineReducers, createStore} from "redux";
import thunk from "redux-thunk";
import {IUserState, userActionCreator, userStateReducer} from "./redux/states/UserState";
import {composeWithDevTools} from "redux-devtools-extension";
import {applicationActionCreator, applicationStateReducer, IApplicationState} from "./redux/states/ApplicationState";

export type AppState = {
    application: IApplicationState
    user: IUserState
}

const composeEnhancers = composeWithDevTools({});
export const store = createStore(
    combineReducers<AppState>({
        application: applicationStateReducer,
        user: userStateReducer
    }),
    composeEnhancers(applyMiddleware(thunk)),
);


export const actionCreator = {
    userActionCreator,
    applicationActionCreator
};

