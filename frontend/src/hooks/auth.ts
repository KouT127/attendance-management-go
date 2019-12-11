import {actionCreator, AppState} from "../store";
import {useDispatch, useSelector} from "react-redux";
import {firebaseApp} from "../lib/firebase";
import {UserPayload} from "../redux/states/UserState";
import {useCallback} from "react";
import {User} from "../domains/user/User";
import axios from "axios";

const userSelector = (state: AppState) => state.user;

export const useAuth = () => {
  const dispatch = useDispatch();

  const getUser = async (token: string, user_id: string): Promise<User> => {
    const response = await axios.get("http://localhost:8080/v1/users/mine", {
      headers: { authorization: token }
    });
    const user = response.data.user;
    if (!user.email) {
      return User.initializeUser(user.id);
    }
    return new User(user.id, user.name, user.email, user.imageUrl, false);
  };

  const getToken = useCallback(async () => {
    const currentUser = firebaseApp.auth().currentUser;
    if (!currentUser) {
      return "";
    }
    return await currentUser.getIdToken(true);
  }, []);

  const subscribeAuth = useCallback(() => {
    firebaseApp.auth().onAuthStateChanged(async user => {
      if (!user || !user.uid || !user.email) {
        dispatch(actionCreator.applicationActionCreator.loadedApplication());
        return;
      }
      const token = await user.getIdToken();
      const authorizedUser = await getUser(token, user.uid);
      const payload: UserPayload = {
        initialLoaded: true,
        userState: {
          id: user.uid,
          username: authorizedUser.username,
          email: authorizedUser.email,
          imageUrl: authorizedUser.imageUrl,
          shouldEdit: authorizedUser.shouldEdit
        }
      };
      dispatch(actionCreator.userActionCreator.loadedUser(payload));
      dispatch(actionCreator.applicationActionCreator.loadedApplication());
    });
  }, []);

  return {
    subscribeAuth,
    getToken
  };
};

export const useUserSelector = () => {
  const user = useSelector(userSelector);
  const isAuthenticated = !!user.id;
  const shouldEdit = user.shouldEdit;

  return {
    user,
    isAuthenticated,
    shouldEdit
  };
};
