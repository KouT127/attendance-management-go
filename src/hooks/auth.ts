import {actionCreator, AppState} from "../store";
import {useDispatch, useSelector} from "react-redux";
import {firebaseApp} from "../lib/firebase";
import {UserPayload} from "../redux/states/UserState";
import {useCallback} from "react";
import {User} from "../domains/user/user";

const userSelector = (state: AppState) => state.user;

export const useAuth = () => {
    const dispatch = useDispatch();

    const getUserDocument = async (user_id: string): Promise<User> => {
        const documentSnapshot = await firebaseApp
            .firestore()
            .collection('users')
            .doc(user_id)
            .get();
        const data = documentSnapshot.data();
        if (!data) {
            return User.initializeUser();
        }
        return new User(data.username, data.email, data.imageUrl, false);
    };

    const observeAuth = useCallback(() => {
        firebaseApp.auth().onAuthStateChanged(async (user) => {
            if (!user || !user.uid || !user.email) {
                dispatch(actionCreator.applicationActionCreator.loadedApplication());
                return;
            }
            const userDocument = await getUserDocument(user.uid);
            const payload: UserPayload = {
                initialLoaded: true,
                userState: {
                    id: user.uid,
                    name: userDocument.username,
                    email: userDocument.email,
                    imageUrl: userDocument.imageUrl,
                    shouldEdit: userDocument.shouldEdit,
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
    const shouldEdit = user.shouldEdit;

    return {
        user,
        isAuthenticated,
        shouldEdit,
    }
};
