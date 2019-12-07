import React, {useState} from "react";
import QrReader from 'react-qr-reader'
import './AttendanceScan.sass'

enum FacingMode {
    User = 'user',
    Environment = 'environment',
};

export const AttendanceScan = () => {
    const [facingMode, setFacingMode] = useState<FacingMode>(FacingMode.Environment);
    const {scanData} = useScanner();
    const handleError = (err: string) => {
        console.error(err)
    };
    const toggleMode = (mode: FacingMode) => {
        if (FacingMode.Environment) {
            return FacingMode.User
        }
        return FacingMode.Environment
    };
    const handleClickToggleButton = () => {
        setFacingMode(toggleMode(facingMode));
        console.log(facingMode)
    };
    return (
        <div>
            <section className='attendance-scan-area'>
                <figure className='attendance-scan-area-camera'>
                    <QrReader
                        delay={300}
                        facingMode={facingMode}
                        onError={handleError}
                        onScan={scanData}
                        style={{width: '100%'}}
                    />
                </figure>
                <button onClick={handleClickToggleButton}>
                    切り替え
                </button>
            </section>
            <section className='attendance-infomation-area'>
                test
            </section>
        </div>
    )
};

const useScanner = () => {
    let timeOut: NodeJS.Timeout;
    let userId: string | null = null;

    const scanData = (data: string | null) => {
        if (data === null) {
            return;
        }
        userId = data;
        resetTimeOut();
        if (userId === data) {
            sendUserId(userId);
        }
    };
    const resetTimeOut = () => {
        clearTimeout(timeOut);
        timeOut = setTimeout(() => {
            console.log('reset');
            userId = ''
        }, 3000);
    };
    const sendUserId = (userId: string) => {
        //    send request
    };
    return {
        scanData,
    }
};
