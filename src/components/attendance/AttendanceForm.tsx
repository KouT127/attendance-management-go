import React, {useCallback, useEffect, useState} from "react";
import {AttendanceKindEnum, IAttendance} from "../../domains/attendance/attendance";
import {db} from "../../lib/firebase";
import {RoundedButton} from "../common/RoundedButton";
import * as firebase from "firebase";
import {AttendanceTimerContainer} from "../../containers/attendance/AttendanceTimerContainer";

type Props = {
    documents: firebase.firestore.QueryDocumentSnapshot[]
}

export const AttendanceFormContainer = (props: Props) => {
    const [title, setTitle] = useState('');
    const [attendance, setAttendance] = useState<IAttendance>({
        type: AttendanceKindEnum.GO_TO_WORK,
        content: '',
        createdAt: undefined,
        updatedAt: undefined,
    });

    useEffect(() => {
        const latestAttendance = props.documents.length > 0 ? props.documents[0] : undefined;
        const latestKindType: AttendanceKindEnum = latestAttendance && latestAttendance.data().type;
        const kindType = latestKindType === AttendanceKindEnum.GO_TO_WORK ? AttendanceKindEnum.LEAVE_WORK : AttendanceKindEnum.GO_TO_WORK
        const buttonTitle = kindType === AttendanceKindEnum.GO_TO_WORK ? '出勤する' : '退勤する';
        setTitle(buttonTitle);
        setAttendance({
            type: kindType,
        })
    }, [props.documents]);

    const handleChangeTextareaText = useCallback((event: React.ChangeEvent<HTMLTextAreaElement>) => {
        const target = event.target;
        setAttendance({
            ...attendance,
            [target.name]: target.value
        });
    }, []);


    const handleClickButton = useCallback(() => {
        db.collection('users')
            .doc('a324al-sdflasdf')
            .collection('attendances')
            .add({
                ...attendance,
                createdAt: firebase.firestore.FieldValue.serverTimestamp(),
                updatedAt: firebase.firestore.FieldValue.serverTimestamp(),
            })
    }, []);

    return (
        AttendanceForm({
            buttonTitle: title,
            onClickButton: handleClickButton,
            onChangeTextArea: handleChangeTextareaText
        })
    )
};

type AttendanceFormProps = {
    buttonTitle: string
    onChangeTextArea: (event: React.ChangeEvent<HTMLTextAreaElement>) => void
    onClickButton: () => void
}

export const AttendanceForm = (props: AttendanceFormProps) => {

    return (
        <section className='timer-section'>
            <AttendanceTimerContainer/>
            <textarea
                name='content'
                className='timer-section__textarea'
                onChange={props.onChangeTextArea}/>
            <RoundedButton
                title={props.buttonTitle}
                appearance={"black"}
                onClick={props.onClickButton}/>
        </section>
    )
};
