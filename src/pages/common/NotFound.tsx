import React, {useEffect} from "react";
import useRouter from "use-react-router";

export const NotFound = () => {
    const {match} = useRouter();
    useEffect(() => {
        console.log(match);
    }, [match]);
    return (
        <div>
            <h1>NotFound</h1>
        </div>
    )
};
