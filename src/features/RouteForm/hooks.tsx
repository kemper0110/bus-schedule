import { useNavigate, useSearchParams } from "react-router-dom"
import { RouteFormData } from "./RouteForm"
import { useState } from 'react'
import { SchedulePath } from "../../pages/Routes"
import dayjs from "dayjs";


export const useFormState = () => {
    const [params] = useSearchParams()
    return useState<RouteFormData>({
        from: params.get("from") || "",
        to: params.get("to") || "",
        date: params.get("date") || dayjs().toString()
    })
}


export const useRouteFormSubmit = () => {
    const navigate = useNavigate();
    const onSubmit = (data: RouteFormData) => {
        console.log(data)
        const path = SchedulePath + '?' + new URLSearchParams(data).toString()
        navigate(path)
    }
    return onSubmit;
};