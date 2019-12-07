import React from "react";
import "./Header.sass";

type Props = {
  title: string;
}

export const Header = (props: Props) => {
  return (
    <header className='header'>
      <div className='header-container'>
        <a href='/' className='header-link'>
          <h1 className='header-title'>{props.title}</h1>
        </a>
      </div>
    </header>
  );
};
