import React from "react";
import "./AttendanceUserInformationHeader.sass"
import {useUserSelector} from "../../hooks/auth";

export const AttendanceUserInformationHeader = () => {
    const {user} = useUserSelector();
    return (
        <div className='attendance-user'>
            <div className='attendance-user-information'>
                <figure className='attendance-user-information-icon'>
                    <img className='attendance-user-information-icon-image' src={user.imageUrl || 'http://via.placeholder.com/80x80'}
                         alt={''}/>
                </figure>
                <section className='attendance-user-information-body'>
                    <h3 className='attendance-user-information-body-name'>
                        {user.username}
                    </h3>
                    <p className='attendance-user-information-body-identifier'>
                        ID: <span>{user.id}</span>
                    </p>
                </section>
            </div>
        </div>
    );
};
