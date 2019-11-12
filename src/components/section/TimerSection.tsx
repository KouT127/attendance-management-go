import {RoundedButton} from "../button/RoundedButton";
import React, {useEffect, useState} from "react";
import * as Moment from "moment";
import "./TimerSection.sass"
import {db} from "../../lib/firebase";
import * as firebase from "firebase";
import {AttendanceType, IAttendance} from "../../domains/attendance/model";


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
    const [attendance, setAttendance] = useState<IAttendance>({
        type: AttendanceType.GO_TO_WORK,
        content: '',
        createdAt: undefined,
        updatedAt: undefined,
    });
    const onChangeInputText = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        setAttendance({
            ...attendance,
            [event.target.name]: event.target.value
        });
    };

    useEffect(() => {
        startTimer();
        return () => {
            console.log('clean up')
        }
    }, []);


    const addAttendance = () => {
        console.log(attendance);
        db.collection('users')
            .doc('a324al-sdflasdf')
            .collection('attendances')
            .add({
                ...attendance,
                createdAt: firebase.firestore.FieldValue.serverTimestamp(),
                updatedAt: firebase.firestore.FieldValue.serverTimestamp(),
            })
    };
    return (<section className='timer-section'>
            <h3 className='timer-section-date'>
                {currentDate}
            </h3>
            <h2 className='timer-section-timestamp'>
                {currentTime}
            </h2>
            <textarea
                name='content'
                className='timer-section-textarea'
                onChange={onChangeInputText}/>
            <RoundedButton
                title={"出勤する"}
                appearance={"black"}
                onClick={addAttendance}/>
        </section>
    )
};
