import {AppState} from "../store";
import {useSelector} from "react-redux";

const applicationSelector = (state: AppState) => state.application;
export const useApplication = () => {
    const application = useSelector(applicationSelector);
    const initialLoaded = application.initialLoaded;
    return {
        initialLoaded,
    }
};
