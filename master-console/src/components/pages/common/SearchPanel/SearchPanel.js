import React from 'react'
import './SearchPanel.css'

function searchPanel(props) {
    return (
        <div className="SearchPanelDiv">

            {props.children.map((field, ) => {
                return (
                    <div className="SearchPanelComponent">
                        {field}
                    </div>
                )
            })}
        </div>
    )
}

export default searchPanel