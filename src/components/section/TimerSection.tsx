import {RoundedButton} from "../button/RoundedButton";
import React, {useEffect, useState} from "react";
import * as Moment from "moment";
import "./TimerSection.sass"
import {db} from "../../lib/firebase";
import * as firebase from "firebase";

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

export const TimerSection = () => {
    const {currentDate, currentTime, startTimer} = useTimer();
    useEffect(() => {
        startTimer();
        return () => {
            console.log('clean up')
        }
    }, []);
    const addAttendance = () => {
        console.log(firebase.database.ServerValue.TIMESTAMP);
        db.collection('users')
            .doc('a324al-sdflasdf')
            .collection('attendance')
            .add({createdAt: firebase.database.ServerValue.TIMESTAMP})
    };
    return (<section className='timer-section'>
            <h3 className='timer-section-date'>
                {currentDate}
            </h3>
            <h2 className='timer-section-timestamp'>
                {currentTime}
            </h2>
            <textarea className='timer-section-textarea'/>
            <RoundedButton
                title={"出勤する"}
                appearance={"black"}
                onClick={addAttendance}/>
        </section>
    )
};
