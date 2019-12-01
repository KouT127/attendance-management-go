import {Action} from "redux";


export interface IUserState {
    id: string
    name: string | null
    email: string
    imageUrl: string | null
    shouldEdit: boolean
}

export const initialState: IUserState = {
    id: '',
    name: '',
    email: '',
    imageUrl: '',
    shouldEdit: true
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
    loadedUser
};
