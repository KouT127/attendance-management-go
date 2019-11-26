import React, {useCallback, useState} from "react"
import './EditUser.sass'
import {useUserDocuments} from "../../hooks/firestore";

type FormProps = {
    username: string
}

export const EditUser = () => {
    const [inputValue, setInputValue] = useState<FormProps>({username: ''});
    const {setUserDocument} = useUserDocuments();
    const handleChangeInput = useCallback((event: React.ChangeEvent<HTMLInputElement>) => {
        const target = event.target;
        const value = target.type === 'checkbox' ? target.checked : target.value;
        setInputValue({
            ...inputValue,
            [event.target.name]: value
        });
    }, [inputValue]);

    const handleClick = useCallback( async () => {
        await setUserDocument(inputValue.username)
    }, [inputValue]);

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
                    <input className='edit-user__text-input'
                           type={'text'}
                           disabled={true}
                           value={'example@example.com'}/>
                    <label>ユーザー名</label>
                    <input className='edit-user__text-input'
                           type={'text'}
                           name={'username'}
                           onChange={handleChangeInput}/>
                </div>
                <input className='edit-user__button'
                        type={'button'}
                       name='enter'
                       value='登録'
                       onClick={handleClick}/>
            </form>
        </div>
    )
};
