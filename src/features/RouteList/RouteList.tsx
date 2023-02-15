import React from 'react';
import {observer} from "mobx-react-lite";
import RouteService, {Route} from "../../services/RouteService";
import {useSearchParams} from "react-router-dom";
import {Table} from "antd";
import type { ColumnsType } from 'antd/es/table';

const RouteList = observer(function RouteList() {
        const [params] = useSearchParams()
        const data = RouteService.getRoutes(params.get("from") || "", params.get("to") || "", params.get("date") || "");
        const columns: ColumnsType<Route> = [
            {
                title: 'Название',
                dataIndex: 'name',
            },
            {
                title: 'Время в пути',
                dataIndex: 'duration',
                sorter: (a, b) => a.duration.localeCompare(b.duration),
                // new Date(a.duration) < new Date(b.duration) ? 1 : -1
            },
            {
                title: 'Отправление',
                dataIndex: 'timestamp_from',
                sorter: (a, b) => a.timestamp_from.localeCompare(b.timestamp_from),
            },
            {
                title: 'Прибытие',
                dataIndex: 'timestamp_to',
                sorter: (a, b) => a.timestamp_to.localeCompare(b.timestamp_to),
            },
            {
                title: 'Цена',
                dataIndex: 'price',
                sorter: (a, b) => Number(a.price) - Number(b.price),
            }
        ];
        return <Table columns={columns} dataSource={data}/>;
    }
);
export default RouteList;