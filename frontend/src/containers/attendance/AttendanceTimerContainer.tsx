import React, {useEffect} from "react";
import {useTimer} from "../../hooks/timer";
import {Timer} from "../../components/common/Timer";

export const AttendanceTimerContainer: React.FC = () => {
    const {currentDate, currentTime, startTimer} = useTimer();

    useEffect(() => {
        startTimer();
    }, []);


    return (
        <Timer formatted_date={currentDate}
               formatted_time={currentTime}
        />
    )
};
