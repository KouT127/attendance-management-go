import {useCallback, useState} from "react";
import {firebaseApp} from "../lib/firebase";
import {useUserSelector} from "./auth";
import axios from "axios";
import {useDispatch} from "react-redux";
import {actionCreator} from "../store";
import {Attendance} from "../domains/attendance/Attendance";

export const useAttendanceDocuments = () => {
    const {user} = useUserSelector();
    const [attendances, setAttendances] = useState<Array<Attendance>>([]);

    const fetchAttendance = useCallback(async () => {
        const currentUser = firebaseApp.auth().currentUser;
        if (!currentUser) {
            return
        }
        const token = await currentUser.getIdToken();
        const response = await axios.get('http://localhost:8080/v1/attendances', {headers: {'authorization': token}});
        const data = response.data.attendances;
        const attendances = data.map((value: any) => {
            const attendance: Attendance = {...value};
            return attendance
        });
        setAttendances(attendances);
    }, []);

    return {
        fetchAttendance,
        attendances,
    }
};

export const useUserDocuments = () => {
    const {user} = useUserSelector();
    const dispatch = useDispatch();

    const setUserData = useCallback(async (name: string) => {
        const currentUser = firebaseApp.auth().currentUser;
        if (!currentUser) {
            return
        }
        const token = await currentUser.getIdToken();
        const response = await axios.put(
            `http://localhost:8080/v1/users/${currentUser.uid}`,
            {name: name, email: user.email, imageUrl: user.imageUrl,},
            {headers: {'authorization': token}});
        const userData = response.data.user;
        dispatch(actionCreator.userActionCreator.loadedUser({initialLoaded: true, userState: {...userData,}}))
    }, []);

    return {
        setUserData
    }
};
