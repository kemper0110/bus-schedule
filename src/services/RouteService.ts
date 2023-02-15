export type Route = {
    id: number
    name: string
    duration: string
    timestamp_from: string
    timestamp_to: string
    price: string
}

export default class RouteService {
    static data: Route[] = [
        {
            id: 1,
            name: "Kameshki - Rostov",
            duration: "3:25",
            timestamp_from: "03:25 04/02",
            timestamp_to: "06:60 04/02",
            price: "799"
        },
        {
            id: 2,
            name: "Kamensk - Rostov",
            duration: "2:25",
            timestamp_from: "07:25 04/02",
            timestamp_to: "10:60 04/02",
            price: "899"
        },
        {
            id: 3,
            name: "Kamensk - Rostov",
            duration: "2:45",
            timestamp_from: "09:25 04/02",
            timestamp_to: "12:60 04/02",
            price: "499"
        }
    ]
    static getRoutes(from: string, to: string, date: string): Route[] {
        return RouteService.data
    }

    static getById(route_id: number): Route | undefined {
        return RouteService.data.find(({id}) => route_id === id)
    }
}