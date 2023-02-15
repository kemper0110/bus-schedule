import React from 'react';
import {observer} from "mobx-react-lite";
import {Layout, Typography} from "antd";
const {Title} = Typography;
const {Header: AntHeader, Footer, Content: AntContent } = Layout;


const Header = observer(function Header() {
        return (
            <AntHeader>
                <Title style={{ color: 'whitesmoke'}}>Автовокзалы.сру</Title>
            </AntHeader>
        );
    }
);
export default Header;