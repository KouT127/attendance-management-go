import classNames from "classnames";
import React from "react";
import "./RoundedButton.sass";

type Appearance = "blue" | "navy" | "black"

type Props = {
  title: string;
  appearance: Appearance
  onClick?: () => void;
}

export const RoundedButton = (props: Props) => {
  const classes = classNames("rounded-button", {
    "blue-button": props.appearance === "blue",
    "navy-button": props.appearance === "navy",
    "black-button": props.appearance === "black"
  });
  return (
    <a
      className={classes}
      onClick={props.onClick}>
      {props.title}
    </a>
  );
};
