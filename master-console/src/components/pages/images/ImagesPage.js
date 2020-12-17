import {React, useState} from "react"
import "./ImagesPage.css"
import RightPanel from "../common/RightPanel/RightPanel"
import SearchPanel from "../common/SearchPanel/SearchPanel"
import ImagesGrid from "./ImagesGrid";

const ImagesPage = () => {
    const [currentState, setCurrentState] = useState({
        imageId: "",
        imageName: ""
    });

    let newState = {...currentState}

    const updateImageName = (event) => {
        newState.imageName = event.target.value
    }

    const updateImageId = (event) => {
        newState.imageId = event.target.value
    }

    const updateState = () => {
        setCurrentState(newState)
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
            </SearchPanel>
            <ImagesGrid imageId={currentState.imageId} imageName={currentState.imageName}/>
        </RightPanel>

    )
}

export default ImagesPage