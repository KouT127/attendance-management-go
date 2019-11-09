import React, { useEffect, useState } from "react";
import "./TimesUser.sass";
import { RoundedButton } from "../../components/button/RoundedButton";
import * as Moment from "moment";

function useTimer() {
  let timer: NodeJS.Timeout;
  const [currentDate, setDate] = useState("");
  const [currentTime, setTime] = useState("");

  const setCurrentTime = () => {
    // @ts-ignore
    let date = Moment().format("YYYY/MM/DD");
    // @ts-ignore
    let time = Moment().format("HH:mm:ss");
    setDate(date);
    setTime(time);
  };

  const startTimer = () => {
    timer = setInterval(setCurrentTime, 1);
  };

  return {
    currentDate,
    currentTime,
    startTimer
  };
}

export const TimesUser = () => {
  const { currentDate, currentTime, startTimer } = useTimer();
  useEffect(() => {
    startTimer();
  }, []);
  return (
    <div className='times'>
      <section className='times-section'>
        <h3 className='times-date'>
          {currentDate}
        </h3>
        <h2 className='times-timestamp'>
          {currentTime}
        </h2>
        <RoundedButton
          title={"出勤する"}
          appearance={"black"}/>
      </section>
      <ol className='times-list'>
        {TimesUserItem()}
      </ol>
    </div>
  );
};

const TimesUserItem = () => {
  return (
    <li className='times-list-item'>
      <div className='times-list-item-left'>
        <h3 className='times-list-item-left-name'>name</h3>
        <p className='times-list-item-left-kind'>出勤</p>
      </div>
      <p className='times-list-item-right'>0:02:28</p>
    </li>
  );
};
