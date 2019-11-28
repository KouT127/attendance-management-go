import React, {useEffect, useState} from "react";

import "./TimerSection.sass"

type Props = {
    formatted_date: string
    formatted_time: string
}

export const TimerSection: React.FC<Props> = (props: Props) => {
    return (
        <>
            <p className='timer-section__date'>
                {props.formatted_date}
            </p>
            <p className='timer-section__timestamp'>
                {props.formatted_time}
            </p>
        </>
    )
};
