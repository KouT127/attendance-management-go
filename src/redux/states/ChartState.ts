import {Action, AnyAction, Dispatch} from "redux";
import {ThunkAction} from "redux-thunk";
import {actionCreator, AppState} from "../../store";
import firebase from "firebase";

export type UserPayload = {
    userState: IUserState
};


export interface IUserState {
    id: string
    name: string | null
    email: string
    imageUrl: string | null
}

export const initialState: IUserState = {
    id: '',
    name: '',
    email: '',
    imageUrl: ''
};

//Actionの定義
//Action-Creatorの定義store
//Reducerの定義
export interface LoadedUserAction extends Action {
    type: "LOADED_USER";
    payload: UserPayload;
}

export const loadedUser = (payload: UserPayload): LoadedUserAction => {
    return {
        type: "LOADED_USER",
        payload
    };
};

export const userStateReducer = (state: IUserState = initialState, action: LoadedUserAction) => {
    switch (action.type) {
        case "LOADED_USER": {
            const user = action.payload.userState;
            return {...state, user};
        }
        default:
            return state;
    }
};

//Thunk-Actionの定義
export const observeAuth = (payload: void): ThunkAction<void, AppState, any, AnyAction> => (dispatch: Dispatch) => {
    firebase.auth().onAuthStateChanged((user) => {
        if (!user || !user.uid || !user.email) {
            // implement error
            return;
        }
        const payload: UserPayload = {
            userState: {
                id: user.uid,
                name: user.displayName,
                email: user.email,
                imageUrl: user.photoURL
            }
        };
        dispatch(userActionCreator.loadedUser(payload))
    });
};

//ドメイン毎にまとめる。
export const userActionCreator = {
    loadedUser,
    observeAuth
};