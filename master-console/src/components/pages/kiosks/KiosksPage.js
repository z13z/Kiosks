import React from 'react'
import './KiosksPage.css'
import RightPanel from "../common/RightPanel/RightPanel"
import SearchPanel from "../common/SearchPanel/SearchPanel"

function kiosksPage() {
    return (
        <RightPanel>
            <SearchPanel>
                <label>
                    id: <label>
                    <input type="number"/>
                </label>
                </label>
                <label>
                    name: <label>
                    <input type="text"/>
                </label>
                </label>
                <button>search</button>
            </SearchPanel>

            <ul className="KiosksPageDiv">
                <li>
                    kiosks
                </li>
            </ul>

        </RightPanel>

    )
}

export default kiosksPage