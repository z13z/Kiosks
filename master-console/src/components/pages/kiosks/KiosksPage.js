import {React, useState} from 'react'
import './KiosksPage.css'
import RightPanel from "../common/RightPanel/RightPanel"
import SearchPanel from "../common/SearchPanel/SearchPanel"
import KiosksGrid from "./KiosksGrid";
import KioskCommandWindow from "../../windows/kiosks/KioskCommandWindow";

const KiosksPage = () => {
    const [currentState, setCurrentState] = useState({
        kioskId: "",
        kioskName: "",
        showWindow: false,
        kioskForCommand: null
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

    const closeCommandWindow = () => {
        newState.showWindow = false
        newState.kioskForCommand = null
        updateState()
    }

    const sendCommandAction = (id) => {
        newState.showWindow = true
        newState.kioskForCommand = id
        updateState()
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
            <KiosksGrid id={currentState.kioskId} name={currentState.kioskName} sendCommandAction={sendCommandAction}/>
            {currentState.showWindow ? (
                <KioskCommandWindow onClose={closeCommandWindow} kioskForCommand={currentState.kioskForCommand}/>
            ) : null}
        </RightPanel>

    )
}

export default KiosksPage