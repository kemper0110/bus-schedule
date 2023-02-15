import {makeAutoObservable} from "mobx";

export default class OrderStore {
    
    constructor() {
        makeAutoObservable(this)
    }
}