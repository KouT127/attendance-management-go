import {useEffect, useState} from "react";
import * as Moment from "moment";

export const useTimer = () => {
    let timer: NodeJS.Timeout;
    const [currentDate, setDate] = useState("");
    const [currentTime, setTime] = useState("");

    useEffect(() => {
        return () => {
            console.log('unsubscribe timer');
            clearInterval(timer);
        }
    }, []);

    const setCurrentTime = () => {
        // @ts-ignore
        let date = Moment().format("YYYY/MM/DD");
        // @ts-ignore
        let time = Moment().format("HH:mm:ss");
        setDate(date);
        setTime(time);
    };

    const startTimer = () => {
        timer = setInterval(setCurrentTime, 1);
    };

    return {
        currentDate,
        currentTime,
        startTimer
    };
}
