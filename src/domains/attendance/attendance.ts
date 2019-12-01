import {Document} from "../common/document";

export enum AttendanceKindEnum {
    GO_TO_WORK = 10,
    LEAVE_WORK = 20,
}

export class AttendanceKind {
    kind: AttendanceKindEnum;

    constructor(kind: AttendanceKindEnum) {
        this.kind = kind
    }

    public static toString(kind: AttendanceKindEnum) {
        switch (kind) {
            case AttendanceKindEnum.GO_TO_WORK:
                return '出勤';
            case AttendanceKindEnum.LEAVE_WORK:
                return '退勤';
            default:
                return '';
        }
    }
}


export interface Attendance extends Document {
    type: AttendanceKindEnum
    content?: string
}
