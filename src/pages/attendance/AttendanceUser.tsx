import React, {useEffect, useState} from "react";
import "./AttendanceUser.sass";
import {RoundedButton} from "../../components/button/RoundedButton";
import * as Moment from "moment";
import {AttendanceUserListItem} from "../../components/list_item/AttendanceUserListItem";
import {AttendanceUserInformationHeader} from "../../components/header/AttendanceUserInfomationHeader";

function useTimer() {
    let timer: NodeJS.Timeout;
    const [currentDate, setDate] = useState("");
    const [currentTime, setTime] = useState("");

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

export const AttendanceUser = () => {
    const {currentDate, currentTime, startTimer} = useTimer();
    useEffect(() => {
        startTimer();
    }, []);
    return (
        <div className='attendance'>
            <AttendanceUserInformationHeader/>
            <section className='attendance-section'>
                <h3 className='attendance-date'>
                    {currentDate}
                </h3>
                <h2 className='attendance-timestamp'>
                    {currentTime}
                </h2>
                <RoundedButton
                    title={"出勤する"}
                    appearance={"black"}/>
            </section>
            <ol className='attendance-list'>
                <AttendanceUserListItem/>
            </ol>
        </div>
    );
};
