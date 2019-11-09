import React from "react";
import QrReader from 'react-qr-reader'


export const TimesScan = () => {
    let timeOut: NodeJS.Timeout;
    let userId: string | null = null;

    const handleScan = (data: string | null) => {
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

    const handleError = (err: string) => {
        console.error(err)
    };
    return (
        <div>
            <QrReader
                delay={300}
                onError={handleError}
                onScan={handleScan}
                style={{width: '100%'}}
            />
        </div>
    )
};
