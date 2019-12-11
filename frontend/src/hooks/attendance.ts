import {createContext, useCallback, useState} from "react";
import {Attendance} from "../domains/attendance/Attendance";
import {useAuth} from "./auth";
import axios from "axios";

export const AttendanceContext = createContext({
    attendances: Array<Attendance>()
});

export const useAttendance = () => {
    const {getToken} = useAuth();
    const [attendances, setAttendances] = useState<Array<Attendance>>([]);


    const fetchAttendance = useCallback(async () => {
        const token = await getToken();
        const response = await axios.get("http://localhost:8080/v1/attendances", {
            headers: {
                authorization: token
            }
        });
        const data = response.data.attendances;
        const attendances = data.map((value: any) => {
            const attendance: Attendance = {
                ...value
            };
            return attendance;
        });
        setAttendances(attendances);
    }, []);

    return {
        fetchAttendance,
        attendances
    };
};