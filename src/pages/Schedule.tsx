import React from 'react';
import {observer} from "mobx-react-lite";
import {useSearchParams} from "react-router-dom";
import RouteForm, {RouteFormData} from "../features/RouteForm/RouteForm";
import RouteList from "../features/RouteList/RouteList";

const Schedule = observer(function Schedule() {
    const [params] = useSearchParams()

    return (
        <div>
            <RouteForm/>
            {/*<h1>{params.get("from")}</h1>*/}
            {/*<h1>{params.get("to")}</h1>*/}
            {/*<h1>{params.get("date")}</h1>*/}
            <RouteList/>
        </div>
    );
});

export default Schedule;


/*
басс даун


        Сущности
Перевозчик
Автобусы перевозчика
Автобусы
Рейс

 */