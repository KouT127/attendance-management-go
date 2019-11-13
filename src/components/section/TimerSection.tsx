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
    const handleChangeInput = (event: React.ChangeEvent<HTMLInputElement>) => {
        const target = event.target;
        const value = target.type === 'checkbox' ? target.checked : target.value;
        setAttendance({
            ...attendance,
            [event.target.name]: event.target.value
        });
    };

    const handleChangeTextareaText = (event: React.ChangeEvent<HTMLTextAreaElement>) => {
        const target = event.target;
        setAttendance({
            ...attendance,
            [target.name]: target.value
        });
    };

    useEffect(() => {
        startTimer();
    }, []);

    const addAttendance = () => {
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
            <h3 className='timer-section__date'>
                {currentDate}
            </h3>
            <h2 className='timer-section__timestamp'>
                {currentTime}
            </h2>
            <textarea
                name='content'
                className='timer-section__textarea'
                onChange={handleChangeTextareaText}/>
            <RoundedButton
                title={"出勤する"}
                appearance={"black"}
                onClick={addAttendance}/>
        </section>
    )
};
