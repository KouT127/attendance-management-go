import React from "react";
import {BrowserRouter as Router, Route, Switch} from "react-router-dom";

import {SignUp} from "./pages/users/SignUp";
import {DefaultLayout} from "./layouts/DefaultLayout";
import {TimesUser} from "./pages/times/TimesUser";
import {TimesScan} from "./pages/times/TimesScan";


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
            <Route path="/times" exact>
                <DefaultLayout>
                    <TimesUser/>
                </DefaultLayout>
            </Route>
            <Route path="/times/scan" exact>
                <DefaultLayout>
                    <TimesScan/>
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
