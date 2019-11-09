import React from "react";
import "../../components/header/AttendanceUserInformationHeader.sass"

export const AttendanceUserInformationHeader = () => {
    return (
        <div className='attendance-user'>
            <div className='attendance-user-information'>
                <figure className='attendance-user-information-icon'>
                    <img className='attendance-user-information-icon-image' src='http://via.placeholder.com/80x80'/>
                </figure>
                <section className='attendance-user-information-body'>
                    <h3 className='attendance-user-information-body-name'>
                        name
                    </h3>
                    <p className='attendance-user-information-body-identifier'>
                        ID: <span>1234678</span>
                    </p>
                </section>
            </div>
        </div>
    );
};
