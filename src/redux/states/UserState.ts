import {Action, AnyAction, Dispatch} from "redux";
import {ThunkAction} from "redux-thunk";
import {actionCreator, AppState} from "../../store";
import {firebaseApp} from "../../lib/firebase";


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
export type UserPayload = {
    initialLoaded: boolean
    userState: IUserState
};

export interface LoadedUserAction extends Action {
    type: "LOADED_USER";
    payload: UserPayload;
}

const loadedUser = (payload: UserPayload): LoadedUserAction => {
    return {
        type: "LOADED_USER",
        payload
    };
};

//Thunk-Actionの定義
const observeAuth = (payload: void): ThunkAction<void, AppState, any, AnyAction> => (dispatch: Dispatch) => {
    console.log('observe');
    firebaseApp.auth().onAuthStateChanged((user) => {
        if (!user || !user.uid || !user.email) {
            dispatch(actionCreator.applicationActionCreator.loadedApplication());
            return;
        }
        const payload: UserPayload = {
            initialLoaded: true,
            userState: {
                id: user.uid,
                name: user.displayName,
                email: user.email,
                imageUrl: user.photoURL
            }
        };
        dispatch(actionCreator.userActionCreator.loadedUser(payload));
        dispatch(actionCreator.applicationActionCreator.loadedApplication());
    });
};

export const userStateReducer = (state: IUserState = initialState, action: LoadedUserAction) => {
    switch (action.type) {
        case "LOADED_USER": {
            const user = action.payload.userState;
            return {...state, ...user};
        }
        default:
            return state;
    }
};

//ドメイン毎にまとめる。
export const userActionCreator = {
    loadedUser,
    observeAuth
};
