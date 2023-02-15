import React from 'react';
import './App.module.css';
import {BrowserRouter, Link, Navigate, Route as RouterRoute, Routes} from "react-router-dom";
import {App as AntApp, Card, Col, Divider, Layout, Row, Space} from 'antd';
import Header from "./structure/Header";
import Content from "./structure/Content";
import Footer from "./structure/Footer";



function App() {
    return (
        <BrowserRouter>
            <AntApp>
                <Layout>
                    <Header/>
                    <Content/>
                    <Footer/>
                </Layout>
            </AntApp>
        </BrowserRouter>
    );
}

export default App;
