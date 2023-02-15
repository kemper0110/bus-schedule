import React from 'react';
import {observer} from "mobx-react-lite";
import {useParams} from "react-router";
import RouteService from "../services/RouteService";
import SitMap from "../features/SitMap/SitMap";

const Route = observer(function Route() {
        const params = useParams();
        const route = RouteService.getById(Number(params.id))
        if (route === undefined)
            return <h1>Рейс не найден</h1>
        return (
            <div>
                <SitMap/>
            </div>
        );
    }
)

export default Route;