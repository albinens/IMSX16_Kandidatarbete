import React from 'react'
import "../legend.css"

const HorizontalLegend = (props) => {
  let green = "#8ED264";
  let yellow = "#F4EC32";
  let red = "#E5414B";
  return (
    <div className='horizontal-legend-container'>
        <div className='legend-circle' style={{backgroundColor: green}}></div>
        <p className='legend-text'>{props.green || "Available"} </p>
        <div className='legend-circle' style={{backgroundColor: red}}></div>
        <p className='legend-text'>{props.yellow || "Booked"}</p>
        <div className='legend-circle' style={{backgroundColor: yellow}}></div>
        <p className='legend-text'>{props.red || "Occupied"}</p>

    </div>
  )
}

export default HorizontalLegend