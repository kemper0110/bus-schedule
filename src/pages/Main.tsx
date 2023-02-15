import React from 'react';
import {observer} from "mobx-react-lite";
import RouteForm from "../features/RouteForm/RouteForm";
import {Card} from "antd";

const Main = observer(function Main() {
        return (
            <div style={{margin: "0 auto"}}>
                <Card>
                    <h1>Поиск расписаний автобусов</h1>
                    <h3>Найти билеты легко и быстро на любой автобусный маршрут
                        на территории России и ближнего зарубежья</h3>
                    <RouteForm/>
                </Card>
            </div>
        );
    }
)

export default Main;