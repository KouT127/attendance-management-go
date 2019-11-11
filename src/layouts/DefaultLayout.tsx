import React, { ReactNode } from "react";
import { Header } from "../components/header/Header";

type Props = {
  children: ReactNode
};

export const DefaultLayout = (props: Props) => {
  return (
    <>
      <Header title='Time'/>
      {props.children}
    </>
  );
};
