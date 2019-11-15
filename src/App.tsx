import React, { Suspense } from "react";
import {BrowserRouter as Router, Route, Switch} from "react-router-dom";

import {SignUp} from "./pages/users/SignUp";
import {DefaultLayout} from "./layouts/DefaultLayout";
import {AttendanceUser} from "./pages/attendance/AttendanceUser";
import {AttendanceScan} from "./pages/attendance/AttendanceScan";


const App: React.FC = () => {
    return (
        <Router>
            <Switch>
                <main>
                    <DefaultRoute/>
                    <UsersRoute/>
                </main>
            </Switch>
        </Router>
    );
};

const DefaultRoute = () => {
    return (
        <>
            <Route path="/attendance" exact>
                <DefaultLayout>
                    <AttendanceUser/>
                </DefaultLayout>
            </Route>
            <Route path="/attendance/scan" exact>
                <DefaultLayout>
                    <AttendanceScan/>
                </DefaultLayout>
            </Route>
        </>
    );
};

const UsersRoute = () => {
    return (
        <>
            <Route path="/users/new" exact>
                <DefaultLayout>
                    <SignUp/>
                </DefaultLayout>
            </Route>
        </>
    );
};


export default App;
