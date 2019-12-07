import {useEffect, useState} from "react";
import Moment from "moment";

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
        const moment = Moment();
        let date = moment.format("YYYY/MM/DD");

        let time = moment.format("HH:mm:ss");
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
};
