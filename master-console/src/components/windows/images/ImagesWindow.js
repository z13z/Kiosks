import {React, useState} from "react"
import PopUpWindow from '../common/PopUpWindow'
import {FormGroup, Input, Label} from 'reactstrap';
import {JWT_TOKEN_KEY} from '../../../Constants'
import axios from "axios";

const ImagesWindow = (props) => {
    const [imageName, setImageName] = useState("")
    const [imageScript, setImageScript] = useState("")

    const onSubmitAction = () => {
        let queryParams = {}
        queryParams['name'] = imageName
        queryParams['script'] = imageScript

        axios.put('/image', queryParams, {headers: {'Authentication': localStorage.getItem(JWT_TOKEN_KEY)}}).then(() => {
            props.successfullyUpdated()
        }).catch(error => {
            if (error.response.status === 401) {
                localStorage.removeItem(JWT_TOKEN_KEY)
                window.location.reload();
            } else if (error.response.status === 403) {
                alert("action is forbidden")
            } else if (error.response.status === 400) {
                alert("image can't be inserted in database. check if name is unique")
            } else {
                throw error;
            }
            props.onClose()
        })
    }

    const onImageNameChange = (event) => {
        setImageName(event.target.value)
    }
    const onImageScriptChange = (event) => {
        setImageScript(event.target.value)
    }

    const fieldStyle = {float: 'right'}

    return (
        <PopUpWindow {...props} onSubmit={onSubmitAction}>
            <FormGroup key='imageNameGroupKey'>
                <Label for="imageNameField">Name</Label>
                <Input type="text" name="imageName" id="imageNameField" value={imageName} onChange={onImageNameChange}
                       style={fieldStyle} required/>
            </FormGroup>
            <FormGroup key='imageScriptFieldKey'>
                <Label for="imageScriptField">Script</Label>
                <Input type="textarea" name="imageScript" id="imageScriptField" value={imageScript}
                       onChange={onImageScriptChange} style={{width: '415px', height: '200px', float: 'right'}}/>
            </FormGroup>
        </PopUpWindow>
    )
}

export default ImagesWindow;