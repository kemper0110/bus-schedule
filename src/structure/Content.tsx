import React from 'react';
import {Routes} from "react-router-dom";
import {Route as RouterRoute} from "react-router";
import {MainPath, RoutePath, SchedulePath} from "../pages/Routes";
import Main from "../pages/Main";
import Schedule from "../pages/Schedule";
import Route from "../pages/Route";
import {observer} from "mobx-react-lite";
import {Layout} from "antd";
const {Content: AntContent } = Layout;

const Content = observer(function Content() {
    return (
        <AntContent>
            <Routes>
                <RouterRoute path={MainPath} element={<Main/>}/>
                <RouterRoute path={SchedulePath} element={<Schedule/>}/>
                <RouterRoute path={RoutePath + ":id"} element={<Route/>}/>
            </Routes>
        </AntContent>
    );
});

export default Content;