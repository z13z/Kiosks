import {React, useState} from "react"
import "./ImagesPage.css"
import RightPanel from "../common/RightPanel/RightPanel"
import SearchPanel from "../common/SearchPanel/SearchPanel"
import ImagesGrid from "./ImagesGrid";
import ImagesWindow from "../../windows/images/ImagesWindow";

const ImagesPage = () => {
    const [currentState, setCurrentState] = useState({
        imageId: "",
        imageName: "",
        showWindow: false,
        forceGridUpdate: false
    });

    let newState = {...currentState}

    const updateImageName = (event) => {
        newState.imageName = event.target.value
    }

    const updateState = () => {
        setCurrentState(newState)
    }

    const updateImageId = (event) => {
        newState.imageId = event.target.value
    }

    const addCreateImageWindow = () => {
        newState.showWindow = true
        updateState()
    }

    const closeImageWindow = () => {
        newState.showWindow = false
        updateState()
    }

    const successfullyUpdated = () => {
        newState.forceGridUpdate = !newState.forceGridUpdate
        closeImageWindow()
    }

    return (
        <RightPanel>
            <SearchPanel>
                <label key="imageIdField">
                    id: <label>
                    <input type="number" onChange={updateImageId}/>
                </label>
                </label>
                <label key="imageNameField">
                    name: <label>
                    <input type="text" onChange={updateImageName}/>
                </label>
                </label>
                <button key="imageSearchButton" onClick={updateState}>search</button>
                <button key="createImageButton" onClick={addCreateImageWindow}>create</button>
            </SearchPanel>
            <ImagesGrid id={currentState.imageId} name={currentState.imageName} forceGridUpdate={currentState.forceGridUpdate}/>
            {currentState.showWindow ? (
                    <ImagesWindow onClose={closeImageWindow} successfullyUpdated={successfullyUpdated}/>)
                : null}
        </RightPanel>

    )
}

export default ImagesPage