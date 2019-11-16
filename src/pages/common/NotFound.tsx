import React, {useEffect} from "react";
import {useLocation, useRouteMatch} from "react-router";

export const NotFound = () => {
    const routeMatch = useRouteMatch();
    const location = useLocation();
    useEffect(() => {
        console.log(routeMatch, location);
    });
    return (
        <div>
            <h1>NotFound</h1>
        </div>
    )
};
