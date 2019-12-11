import { useAuth, useUserSelector } from "./auth";
import { useDispatch } from "react-redux";
import { useCallback } from "react";
import axios from "axios";
import { actionCreator } from "../store";

export const useUserDetail = () => {
  const { user } = useUserSelector();
  const { getToken } = useAuth();
  const dispatch = useDispatch();

  const setUserData = useCallback(async (name: string) => {
    const token = await getToken();
    const response = await axios.put(
      `http://localhost:8080/v1/users/${user.id}`,
      {
        name: name,
        email: user.email,
        imageUrl: user.imageUrl
      },
      {
        headers: {
          authorization: token
        }
      }
    );
    const userData = response.data.user;
    dispatch(
      actionCreator.userActionCreator.loadedUser({
        initialLoaded: true,
        userState: {
          ...userData
        }
      })
    );
  }, []);

  return {
    setUserData
  };
};
