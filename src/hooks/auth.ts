import {actionCreator, AppState} from "../store";
import {useDispatch, useSelector} from "react-redux";
import {firebaseApp} from "../lib/firebase";
import {UserPayload} from "../redux/states/UserState";
import {useCallback, useEffect} from "react";

const userSelector = (state: AppState) => state.user;

export const useAuth = () => {
    const dispatch = useDispatch();
    const observeAuth = useCallback(() => {
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
    }, []);
    return {
        observeAuth
    }
};

export const useUserSelector = () => {
    const user = useSelector(userSelector);
    const isAuthenticated = !!user.id;

    return {
        user,
        isAuthenticated
    }
};
