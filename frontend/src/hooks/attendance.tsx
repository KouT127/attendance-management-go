import React, {createContext, Dispatch, SetStateAction, useContext, useEffect, useRef, useState} from "react";
import {Attendance, AttendanceKindEnum} from "../domains/attendance/Attendance";
import axios from "axios";

interface Props {
    children: React.ReactElement<any>;
}

interface AttendanceProviderProps {
    attendances: Array<Attendance>;
    latestKindType: AttendanceKindEnum;
    setCreateEvent: Dispatch<SetStateAction<CreateParams | undefined>>;
    fetchAttendance: (token: string) => Promise<void>;
}

export const useAttendanceStore = () => useContext(AttendanceContext);

export const AttendanceContext = createContext<AttendanceProviderProps>({
    attendances: [],
    latestKindType: AttendanceKindEnum.GO_TO_WORK,
    setCreateEvent: value => {
    },
    fetchAttendance: token => new Promise<void>(() => {
    })
});

interface CreateParams {
    token: string;
    userId: string;
    remark: string;
    kind: AttendanceKindEnum;
}

export const AttendanceProvider = (props: Props) => {
    const isLoading = useRef(false);
    const [createEvent, setCreateEvent] = useState<CreateParams | undefined>();
    const [attendances, setAttendances] = useState<Array<Attendance>>([]);
    const [latestKindType, setLatestKindType] = useState<AttendanceKindEnum>(
        AttendanceKindEnum.GO_TO_WORK
    );

    useEffect(() => {
        if (attendances.length === 0) {
            return;
        }
        const latest = attendances[0];
        const kindType =
            latest.kind === AttendanceKindEnum.GO_TO_WORK
                ? AttendanceKindEnum.LEAVE_WORK
                : AttendanceKindEnum.GO_TO_WORK;
        setLatestKindType(kindType);
    }, [attendances]);

    useEffect(() => {
        if (!createEvent || isLoading.current) {
            return;
        }
        createAttendance(createEvent);
    }, [createEvent, isLoading]);

    const createAttendance = async (params: CreateParams) => {
        isLoading.current = true;
        const attendance = {
            remark: params.remark
        };

        const response = await axios.post(
            `http://localhost:8080/v1/attendances`,
            {
                ...attendance
            },
            {
                headers: {
                    authorization: params.token
                }
            }
        );

        const newAttendances = [response.data.attendance, ...attendances];
        setAttendances(newAttendances.slice(0, 5));
        setTimeout(() => {
            isLoading.current = false;
        }, 2000);
    };

    const fetchAttendance = async (token: string) => {
        const response = await axios.get("http://localhost:8080/v1/attendances", {
            headers: {
                authorization: token
            }
        });
        const data = response.data.attendances || [];
        const attendances = data.map((value: any) => {
            const attendance: Attendance = {
                ...value
            };
            return attendance;
        });
        setAttendances(attendances);
    };

    const value = {
        attendances,
        latestKindType,
        setCreateEvent,
        fetchAttendance
    };

    return (
        <AttendanceContext.Provider value={value}>
            {props.children}
        </AttendanceContext.Provider>
    );
};
