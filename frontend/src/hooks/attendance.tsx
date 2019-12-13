import React, {createContext, useCallback, useContext, useEffect, useState} from "react";
import {Attendance, AttendanceKindEnum} from "../domains/attendance/Attendance";
import axios from "axios";

interface Props {
  children: React.ReactElement<any>;
}

interface AttendanceProviderProps {
  attendances: Array<Attendance>;
  latestKindType: AttendanceKindEnum;
  createAttendance: (token: string, attendance: Attendance) => Promise<void>;
  fetchAttendance: (token: string) => Promise<void>;
}

export const useAttendanceStore = () => useContext(AttendanceContext);

export const AttendanceContext = createContext<AttendanceProviderProps>({
  attendances: [],
  latestKindType: AttendanceKindEnum.GO_TO_WORK,
  createAttendance: (token, attendance) => new Promise<void>(() => {}),
  fetchAttendance: token => new Promise<void>(() => {})
});

export const AttendanceProvider = (props: Props) => {
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

  const createAttendance = useCallback(
    async (token: string, attendance: Attendance) => {
      const response = await axios.post(
        `http://localhost:8080/v1/attendances`,
        {
          ...attendance
        },
        {
          headers: {
            authorization: token
          }
        }
      );

      const newAttendances = [response.data.attendance, ...attendances];
      setAttendances(newAttendances.slice(0, 5));
    },
    [attendances]
  );

  const fetchAttendance = useCallback(
    async (token: string) => {
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
    },
    [attendances]
  );

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
