import React, {useEffect} from "react";
import {BrowserRouter as Router, Redirect, Route, Switch} from "react-router-dom";

import {SignIn} from "./pages/auth/SignIn";
import {AttendanceUser} from "./pages/attendance/AttendanceUser";
import {AttendanceScan} from "./pages/attendance/AttendanceScan";
import {useDispatch} from "react-redux";
import {actionCreator} from "./store";
import {useAuthUser} from "./hooks/auth";
import {Header} from "./components/header/Header";
import {useApplication} from "./hooks/application";
import {PulseLoader} from "react-spinners";
import {Splash} from "./pages/common/Splash";


export type HeaderProps = {
    children: any
}

const App: React.FC = () => {
    const dispatch = useDispatch();
    dispatch(actionCreator.userActionCreator.observeAuth());
    return (<Routes/>)
};

const Auth: React.FC<HeaderProps> = (props) => {
    const {isAuthenticated} = useAuthUser();
    return isAuthenticated ? props.children : <Redirect to={'/signin'}/>
};
export const Loading = (props: HeaderProps) => {
    const {initialLoaded} = useApplication();
    const {isAuthenticated} = useAuthUser();
    useEffect(() => {
        if (!initialLoaded) {
            return;
        }
    }, [initialLoaded, isAuthenticated]);
    if (!initialLoaded) {
        return (
            <div className='initial-loading__section'>
                <PulseLoader
                    sizeUnit={"px"}
                    size={30}
                    color={'#123abc'}
                    loading={true}
                />
            </div>
        );
    }
    return props.children
};

const Routes = () => {
    return (
        <>
            <Header title='Time'/>
            <Router>
                <Switch>
                    <main className={'contents'}>
                        <Loading>
                            <AuthRoute/>
                            <Auth>
                                <Route path="/home" component={AttendanceUser}/>
                                <Route path="/scan" exact component={AttendanceScan}/>
                            </Auth>
                        </Loading>
                    </main>
                </Switch>
            </Router>
        </>
    );
};

const AuthRoute = () => {
    return (
        <>
            <Route path="/" exact component={Splash}/>
            <Route path="/signin" exact component={SignIn}/>
        </>
    );
};


export default App;
