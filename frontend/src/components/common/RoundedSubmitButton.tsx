import React from "react";
import "./RoundedSubmitButton.sass"
import classNames from "classnames";

type Props = {
    title: string
    className?: string
}

export const RoundedSubmitButton = (props: Props) => {
    const classes = classNames('rounded-submit__button',
        props.className,
    );
    return (
        <input
            className={classes}
            type="submit"
            value={props.title}
        />
    )
};
