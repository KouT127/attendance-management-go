import React from "react";
// @ts-ignore
import { RoundedButton } from "../../components/button/RoundedButton";
import "./SignUp.sass";
import "../../base/base.sass";


export const SignUp = () => {
  return (
    <div>
      <div className='container'>
        <div className='container-header'>
          <h2 className='container-header-title'>
            新規ユーザー登録
          </h2>
        </div>
        <div className='container-section'>
          <div className='container-section-image'>
            <div>
              <img className='container-section-image-icon'
                   src='https://placehold.jp/150x150.png'/>
              <span className='container-section-image-title'>画像アップロード</span>
            </div>
          </div>
        </div>
        <div className='container-section'>
          <span>入力はあとで変更出来ます</span>
        </div>
        <div className='container-input-text'>
          <label className='container-input-label'>ニックネーム</label>
          <div>
            <input className='container-input-text-area' type='text'/>
          </div>
        </div>
        <div className='container-input-text'>
          <label className='container-input-label'>ユーザーID</label>
          <div>
            <input className='container-input-text-area' type='text'/>
          </div>
        </div>
        <div className='container-input-text'>
          <label className='container-input-label'>メールアドレス</label>
          <div className='text'>
            <input className='container-input-text-area' type='text'/>
          </div>
        </div>
        <div>
          <input type='checkbox'/>
          <span>
              <a href='/'>利用規約</a>
              と<a href='/'>プライバシーポリシー</a>に同意する
            </span>
        </div>
        <div className='container-input-button-section'>
          <RoundedButton
            title='登録する'
            appearance='navy'/>
        </div>
      </div>
    </div>
  );
};
