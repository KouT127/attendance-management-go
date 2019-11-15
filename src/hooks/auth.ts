import {AppState} from "../store";
import {useSelector} from "react-redux";

const userSelector = (state: AppState) => state.user;
export const useAuthUser = () => {
    const user = useSelector(userSelector);
    const isAuthorized = !!user.id;
    return {
        user,
        isAuthenticated: isAuthorized
    }
};
