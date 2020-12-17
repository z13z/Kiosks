import React, {Component} from 'react'
import './ImagesPage.css'
import axios from "axios"

class ImagesGrid extends Component {
    state = {
        images: []
    }

    componentDidMount() {
        this.loadImages()
    }

    stateDidChange(props, state) {
        return JSON.stringify(this.state) !== JSON.stringify(state) || JSON.stringify(this.props) !== JSON.stringify(props)
    }

    shouldComponentUpdate(nextProps, nextState, nextContext) {
        return this.stateDidChange(nextProps, nextState)
    }

    componentDidUpdate(prevProps, prevState, snapshot) {
        if (this.stateDidChange(prevProps, prevState)) {
            this.loadImages();
        }
    }

    render() {
        return <div className="ImagesGridDiv">
            <ul>
                {this.state.images != null ? this.state.images.map(image => <li
                    key={image.id}>{image.name}</li>) : null}
            </ul>
        </div>
    }

    loadImages() {
        let queryParams = {params: {}}
        if (this.props.imageId != null && this.props.imageId !== "") {
            queryParams.params.id = this.props.imageId
        }
        if (this.props.imageName != null && this.props.imageName !== "") {
            queryParams.params.name = this.props.imageName
        }
        axios.get("/image", queryParams).then(response => {
            this.setState({
                image: response.data
            })
        })
    }
}

export default ImagesGrid