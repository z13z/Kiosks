import React from 'react'
import './RightPanel.css'

function rightPanel(props) {
    return (
        <div className="RightPanelDiv">
            {props.children}
        </div>
    )
}

export default rightPanel