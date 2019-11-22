import React from "react"
import './EditUser.sass'

export const EditUser = () => {
    return (
        <div>
            <form>
                <section className='edit-user__header'>
                    <h1 className='edit-user__header-title'>ユーザー</h1>
                </section>
                <section className='edit-user__image-box'>
                    <figure>
                        <img className='edit-user__image'/>
                    </figure>
                </section>
                <div className='edit-user__form-section'>
                    <label>メールアドレス</label>
                    <input className='edit-user__text-input' type={'text'}/>
                    <label>ユーザー名</label>
                    <input className='edit-user__text-input' type={'text'}/>
                </div>
            </form>
        </div>
    )
};
