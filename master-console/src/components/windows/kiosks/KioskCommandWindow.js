import {React, useState} from "react"
import PopUpWindow from '../common/PopUpWindow'
import {FormGroup, Input, Label} from 'reactstrap';
import {JWT_TOKEN_KEY} from '../../../Constants'
import axios from "axios";

const KioskCommandWindow = (props) => {
    const kioskId = props.kioskForCommand
    const [command, setCommand] = useState("")
    const [commandResult, setCommandResult] = useState("")

    const onSubmitAction = () => {
        let queryParams = {}
        queryParams['id'] = kioskId
        queryParams['command'] = command

        axios.post('/kiosksCommander', queryParams, {headers: {'Authentication': localStorage.getItem(JWT_TOKEN_KEY)}}).then((response) => {
            setCommandResult(response.data)
        }).catch(error => {
            if (error.response.status === 401) {
                localStorage.removeItem(JWT_TOKEN_KEY)
                window.location.reload();
            } else if (error.response.status === 403) {
                alert("action is forbidden")
            } else if (error.response.status === 400) {
                alert("Bad Request, " + error.response.data)
            } else if (error.response.status === 503) {
                alert("Kiosk is inactive")
            } else {
                throw error;
            }
        })
    }

    const onCommandChange = (event) => {
        setCommand(event.target.value)
    }

    const fieldStyle = {float: 'right'}

    return (
        <PopUpWindow {...props} onSubmit={onSubmitAction}>
            <Label>Kiosk: {kioskId}</Label>
            <FormGroup key='commandGroupKey'>
                <Label for="commandField">Command</Label>
                <Input type="text" name="command" id="commandField" value={command} onChange={onCommandChange}
                       style={fieldStyle} required/>
            </FormGroup>
            <FormGroup key='commandResultKey'>
                <Label for="commandResultField">Result</Label>
                <Input type="textarea" name="commandResult" id="commandResult" value={commandResult}
                       style={{width: '415px', height: '200px', float: 'right'}} enabled={false}/>
            </FormGroup>
        </PopUpWindow>
    )
}

export default KioskCommandWindow;