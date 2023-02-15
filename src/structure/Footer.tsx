import { observer } from 'mobx-react-lite';
import React from 'react'
import {Col, Divider, Layout, Row} from "antd";

const {Footer: AntFooter} = Layout;

const Footer = observer(function Footer() {
    return (
        <AntFooter>
            <Divider/>
            <Row gutter={[8, 8]}>
                <Col span={12}>
                    <p>Поддержка</p>
                    <p>О компании</p>
                    <p>Условия использования</p>
                    <p>Возврат билета</p>
                </Col>
                <Col span={12}>
                    <p>Контроль качества</p>
                    <p>Виджет на сайт</p>
                    <p>Правила и документы</p>
                    <p>Транспортным компаниям</p>
                </Col>
            </Row>
            <Divider/>
            © 2008—2023, ООО «НеАвтовокзалы.сру»
        </AntFooter>
    )
});

export default Footer;