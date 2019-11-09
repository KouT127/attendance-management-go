import React, {useEffect, useState} from "react";
import "./AttendanceUser.sass";
import {RoundedButton} from "../../components/button/RoundedButton";
import * as Moment from "moment";

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
            <AttendanceUserInformation/>
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
                <AttendanceUserItem/>
            </ol>
        </div>
    );
};

const AttendanceUserItem = () => {
    return (
        <li className='attendance-list-item'>
            <div className='attendance-list-item-left'>
                <h3 className='attendance-list-item-left-name'>name</h3>
                <p className='attendance-list-item-left-kind'>出勤</p>
            </div>
            <p className='attendance-list-item-right'>0:02:28</p>
        </li>
    );
};

const AttendanceUserInformation = () => {
    return (
        <div className='attendance-user'>
            <div className='attendance-user-information'>
                <figure className='attendance-user-information-icon'>
                    <img className='attendance-user-information-icon-image' src='http://via.placeholder.com/80x80'/>
                </figure>
                <section className='attendance-user-information-body'>
                    <h3 className='attendance-user-information-body-name'>
                        name
                    </h3>
                    <p className='attendance-user-information-body-identifier'>
                        ID: <span>1234678</span>
                    </p>
                </section>
            </div>
        </div>
    );
};
