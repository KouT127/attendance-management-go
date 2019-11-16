import React, {useEffect} from "react";
import useRouter from "use-react-router";

export const NotFound = () => {
    const {location, match, history} = useRouter();
    useEffect(() => {
        console.log(match, location, history);
    }, [match]);
    return (
        <div>
            <h1>NotFound</h1>
        </div>
    )
};
