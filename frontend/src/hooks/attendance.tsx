import React, {createContext, useCallback, useContext, useEffect, useState} from "react";
import {Attendance, AttendanceKindEnum} from "../domains/attendance/Attendance";
import {useAuth} from "./auth";
import axios from "axios";

interface Props {
  children: React.ReactElement<any>;
}

interface AttendanceProviderProps {
  attendances: Array<Attendance>;
  latestKindType: AttendanceKindEnum;
  createAttendance: (attendance: Attendance) => Promise<void>;
  fetchAttendance: () => Promise<void>;
}

export const useAttendanceStore = () => useContext(AttendanceContext);

export const AttendanceContext = createContext<AttendanceProviderProps>({
  attendances: [],
  latestKindType: AttendanceKindEnum.GO_TO_WORK,
  createAttendance: (attendance: Attendance) => new Promise<void>(() => {}),
  fetchAttendance: () => new Promise<void>(() => {})
});

export const AttendanceProvider = (props: Props) => {
  const [attendances, setAttendances] = useState<Array<Attendance>>([]);
  const { getToken } = useAuth();
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

  const createAttendance = useCallback(
    async (attendance: Attendance) => {
      const token = await getToken();
      const response = await axios.post(
        `http://localhost:8080/v1/attendances`,
        {
          ...attendance
        },
        {
          headers: { authorization: token }
        }
      );
      const newAttendances = [response.data.attendance, ...attendances];
      setAttendances(newAttendances);
    },
    [attendances]
  );

  const fetchAttendance = useCallback(async () => {
    const token = await getToken();
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
  }, []);

  const value = {
    attendances,
    latestKindType,
    createAttendance,
    fetchAttendance
  };

  return (
    <AttendanceContext.Provider value={value}>
      {props.children}
    </AttendanceContext.Provider>
  );
};
