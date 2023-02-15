import {makeAutoObservable} from "mobx";

class PassengerStore {
    name: string = ""
    sex: null | "male" | "female" = null
    birthday: string = ""
    passport: string = ""
    sit: number | null = null
    constructor() {
        makeAutoObservable(this)
    }
}

export default PassengerStore;