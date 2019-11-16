import React, {FC, useEffect} from "react";
import {PacmanLoader} from "react-spinners";

import "./InitialLoading.sass"
import {useApplication} from "../../hooks/application";
import {useAuthUser} from "../../hooks/auth";
import useRouter from "use-react-router";

export const Splash: FC = () => {
    const {initialLoaded} = useApplication();
    const {isAuthenticated} = useAuthUser();
    const {history} = useRouter();
    useEffect(() => {
        if (!initialLoaded) {
            return;
        }
        if (isAuthenticated) {
            history.replace('/home');
        } else {
            history.replace('/signin');
        }
    }, [initialLoaded, isAuthenticated]);

    return (
        <div className='initial-loading__section'>
            <PacmanLoader
                sizeUnit={"px"}
                size={30}
                color={'#123abc'}
                loading={true}
            />
        </div>
    );
};
