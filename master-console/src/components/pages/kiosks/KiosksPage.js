import {React, useState} from 'react'
import './KiosksPage.css'
import RightPanel from "../common/RightPanel/RightPanel"
import SearchPanel from "../common/SearchPanel/SearchPanel"
import KiosksGrid from "./KiosksGrid";

const KiosksPage = () => {
    const [currentState, setCurrentState] = useState({
        kioskId: "",
        kioskName: ""
    });

    let newState = {...currentState}

    const updateKioskName = (event) => {
        newState.kioskName = event.target.value
    }

    const updateKioskId = (event) => {
        newState.kioskId = event.target.value
    }

    const updateState = () => {
        setCurrentState(newState)
    }

    return (
        <RightPanel>
            <SearchPanel>
                <label key="kioskIdField">
                    id: <label>
                    <input type="number" onChange={updateKioskId}/>
                </label>
                </label>
                <label key="kioskNameField">
                    name: <label>
                    <input type="text" onChange={updateKioskName}/>
                </label>
                </label>
                <button key="kioskSearchButton" onClick={updateState}>search</button>
            </SearchPanel>
            <KiosksGrid id={currentState.kioskId} name={currentState.kioskName}/>
        </RightPanel>

    )
}

export default KiosksPage