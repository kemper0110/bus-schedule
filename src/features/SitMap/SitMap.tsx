import React from 'react';
import {observer} from "mobx-react-lite";

enum State {
    Invalid = 0, Empty = 1, Busy = 2
}

type Cell = {
    state: State
    position?: number
}

const SitMap = observer(function SitMap() {
    const map = [
        [{position: 4, state: State.Empty}, {position: 8, state: State.Empty}, {position: 12, state: State.Empty}, {position: 16, state: State.Empty}, {state: State.Invalid}],
        [{position: 3, state: State.Empty}, {position: 7, state: State.Empty}, {position: 11, state: State.Empty}, {position: 15, state: State.Empty}, {position: 19, state: State.Busy}],
        [{state: State.Invalid}, {state: State.Invalid}, {state: State.Invalid}, {state: State.Invalid}, {state: State.Invalid}],
        [{position: 2, state: State.Empty}, {position: 6, state: State.Empty}, {position: 10, state: State.Empty}, {position: 14, state: State.Busy}, {position: 18, state: State.Busy}],
        [{position: 1, state: State.Empty}, {position: 5, state: State.Busy}, {position: 9, state: State.Empty}, {position: 13, state: State.Empty}, {position: 17, state: State.Busy}]
    ]
    const makeCell = (cell: Cell) => {
        const style = {
            // fontFamily: "'Museo Sans Cyrl',sans-serif",
            textAlign: 'center' as 'center',
            fontSize: 24,
            borderRadius: 4,
            margin: 3,
            width: 40,
            height: 40,
            border: "1px solid lightgray",
        }
        switch (cell.state) {
            case State.Invalid:
                return <button style={{ ...style,
                    color: "white",
                    background: "lightgray",
                }}/>
            case State.Busy:
                return <button style={{ ...style,
                    color: "white",
                    background: "gray",
                }}>{cell.position}</button>
            case State.Empty:
                return <button style={{ ...style,
                    color: 'black',
                    background: "white"
                }}>{cell.position}</button>
        }
    };
    return (
        <>
            <table>
                {/*<thead>*/}
                {/*<tr> {Array(map.length).fill(<th/>)} </tr>*/}
                {/*</thead>*/}
                <tbody>
                {
                    map.map((row, i) => (
                        <tr key={i}>
                            {row.map((cell, j) => (
                                <td key={j}>{makeCell(cell)}</td>
                            ))}
                        </tr>
                    ))
                }
                </tbody>
            </table>
            {/*<Row>*/}
            {/*    {makeCell(3)}*/}
            {/*    {makeCell(7)}*/}
            {/*    {makeCell(11)}*/}
            {/*    {makeCell(15)}*/}
            {/*</Row>*/}
            {/*<Row>*/}
            {/*    {makeCell(3)}*/}
            {/*    {makeCell(7)}*/}
            {/*    {makeCell(11)}*/}
            {/*    {makeCell(15)}*/}
            {/*</Row>*/}
            {/*<Row>*/}
            {/*    {makeCell(3)}*/}
            {/*    {makeCell(7)}*/}
            {/*    {makeCell(11)}*/}
            {/*    {makeCell(15)}*/}
            {/*</Row>*/}
            {/*<Row>*/}
            {/*    {makeCell(3)}*/}
            {/*    {makeCell(7)}*/}
            {/*    {makeCell(11)}*/}
            {/*    {makeCell(15)}*/}
            {/*</Row>*/}
        </>
    );
})

export default SitMap;