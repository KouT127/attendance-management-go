import React from "react";
import {BrowserRouter as Router, Redirect, Route, Switch} from "react-router-dom";

import {SignIn} from "./pages/auth/SignIn";
import {AttendanceUser} from "./pages/attendance/AttendanceUser";
import {AttendanceScan} from "./pages/attendance/AttendanceScan";
import {useDispatch} from "react-redux";
import {actionCreator} from "./store";
import {useAuthUser} from "./hooks/auth";
import {Header} from "./components/header/Header";


interface AuthParameters {
    isAuthenticated: boolean;
    children: any
}

const App: React.FC = () => {
    const dispatch = useDispatch();
    dispatch(actionCreator.userActionCreator.observeAuth());
    return (<Routes/>)
};

const Auth = (props: AuthParameters) => {
    console.log(props.children);
    return props.isAuthenticated ? props.children : <Redirect to={'/signin'}/>
};

const Routes = () => {
    const {isAuthenticated} = useAuthUser();
    return (
        <>
            <Header title='Time'/>
            <Router>
                <Switch>
                    <main className={'contents'}>
                        <UsersRoute/>
                        <Auth isAuthenticated={isAuthenticated}>
                            <Route path="/attendance" exact component={AttendanceUser}/>
                            <Route path="/attendance/scan" exact component={AttendanceScan}/>
                        </Auth>
                    </main>
                </Switch>
            </Router>
        </>
    );
};

const UsersRoute = () => {
    return (
        <>
            <Route path="/signin" exact component={SignIn}/>
        </>
    );
};


export default App;
