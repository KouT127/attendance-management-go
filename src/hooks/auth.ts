import {actionCreator, AppState} from "../store";
import {useDispatch, useSelector} from "react-redux";
import {firebaseApp} from "../lib/firebase";
import {UserPayload} from "../redux/states/UserState";
import {useCallback} from "react";
import {User} from "../domains/user/user";

const userSelector = (state: AppState) => state.user;

export const useAuth = () => {
    const dispatch = useDispatch();

    const getOrCreateUserDocument = async (user_id: string): Promise<User> => {
        const docReference = firebaseApp
            .firestore()
            .collection('users')
            .doc(user_id);
        const documentSnapshot = await docReference.get();
        const data = documentSnapshot.data();
        if (!data) {
            const user = User.initializeUser();
            await docReference.set(user.toJson());
            return user
        }
        return new User(data.username, data.email, data.imageUrl, false);
    };

    const subscribeAuth = useCallback(() => {
        firebaseApp.auth().onAuthStateChanged(async (user) => {
            if (!user || !user.uid || !user.email) {
                dispatch(actionCreator.applicationActionCreator.loadedApplication());
                return;
            }
            const userDocument = await getOrCreateUserDocument(user.uid);
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
        subscribeAuth
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
