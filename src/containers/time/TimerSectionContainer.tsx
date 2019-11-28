import React, {useEffect} from "react";
import {useTimer} from "../../hooks/timer";
import {TimerSection} from "../../components/section/TimerSection";

export const TimerSectionContainer: React.FC = () => {
    const {currentDate, currentTime, startTimer} = useTimer();

    useEffect(() => {
        startTimer();
    }, []);


    return (
        <TimerSection formatted_date={currentDate}
                      formatted_time={currentTime}
        />
    )
};
