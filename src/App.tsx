import React from "react";
import {BrowserRouter as Router, Redirect, Route, Switch} from "react-router-dom";

import {SignIn} from "./pages/auth/SignIn";
import {AttendanceUser} from "./pages/attendance/AttendanceUser";
import {AttendanceScan} from "./pages/attendance/AttendanceScan";
import {useDispatch} from "react-redux";
import {actionCreator} from "./store";
import {useAuthUser} from "./hooks/auth";
import {Header} from "./components/header/Header";
import {InitialLoading} from "./pages/common/InitialLoading";
import {firebaseApp} from "./lib/firebase";


interface AuthParameters {
    children: any
}

const App: React.FC = () => {
    const dispatch = useDispatch();
    dispatch(actionCreator.userActionCreator.observeAuth());
    return (<Routes/>)
};

const Auth = (props: AuthParameters) => {
    const {isAuthenticated} = useAuthUser();
    return isAuthenticated ? props.children : <Redirect to={'/'}/>
};

const Routes = () => {
    return (
        <>
            <Header title='Time'/>
            <Router>
                <Switch>
                    <main className={'contents'}>
                        <AuthRoute/>
                        <Auth>
                            <Route path="/home" exact component={AttendanceUser}/>
                            <Route path="/attendance/scan" exact component={AttendanceScan}/>
                        </Auth>
                    </main>
                </Switch>
            </Router>
        </>
    );
};

const AuthRoute = () => {
    return (
        <>
            <Route path='/' exact component={InitialLoading}/>
            <Route path="/signin" exact component={SignIn}/>
        </>
    );
};


export default App;
