import React, {FC, useEffect} from "react";
import {PacmanLoader} from "react-spinners";

import "./InitialLoading.sass"
import {useApplication} from "../../hooks/application";
import {useUserSelector} from "../../hooks/auth";
import {useHistory} from "react-router";

export const Splash: FC = () => {
    const {initialLoaded} = useApplication();
    const {isAuthenticated} = useUserSelector();
    const {push} = useHistory();
    useEffect(() => {
        if (!initialLoaded) {
            return;
        }
        if (isAuthenticated) {
            push('/home');
        } else {
            push('/signin');
        }
    }, [initialLoaded, isAuthenticated]);

    return (
        <div className='initial-loading__section'>
            <PacmanLoader
                sizeUnit={"px"}
                size={30}
                color={'black'}
                loading={true}
            />
        </div>
    );
};
