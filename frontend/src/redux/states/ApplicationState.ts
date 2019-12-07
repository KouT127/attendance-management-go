import {Action} from "redux";

export interface IApplicationState {
    initialLoaded: boolean
}

export const initialState: IApplicationState = {
    initialLoaded: false,
};

export interface LoadedApplicationAction extends Action {
    type: "LOADED_APPLICATION";
}

const loadedApplication = (): LoadedApplicationAction => {
    return {
        type: "LOADED_APPLICATION",
    };
};

export const applicationStateReducer = (state: IApplicationState = initialState, action: LoadedApplicationAction) => {
    switch (action.type) {
        case "LOADED_APPLICATION": {
            const applicationState = {
                initialLoaded: true,
            };
            return {...state, ...applicationState};
        }
        default:
            return state;
    }
};

export const applicationActionCreator = {
    loadedApplication
};
