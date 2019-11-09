import { default as React } from "react";
import './ArtilesListItem.sass'

export const ArticlesListItem: React.FC = () => {
  return (
    <div className="list-item">
      <div className="list-item-content">
        <div className="list-item-content-left">
          <h3 className="list-item-content-left-title">test </h3>
          <p className="list-item-content-left-content">testdesc </p>
        </div>
        <div className="list-item-content-image">
          <img src="http://via.placeholder.com/120x120" alt=""/>
        </div>
      </div>
      <div className="list-item-footer">
        <div className="list-item-footer-inner">
          <img className="list-item-footer-image" src="http://via.placeholder.com/40x40" alt=""/>
            <div className="list-item-footer-inner-section">
              <p className="list-item-footer-inner-name">UserName</p>
              <p className="list-item-footer-inner-day">昨日</p>
            </div>
        </div>
      </div>
    </div>
  );
};
