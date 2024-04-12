import React from "react";
import './weekChart.css'

const Legend = ({ data, selectedItems, onChange }) => {

return(
    <div className="legendContainer-weekplot">
      {data.map((d) => (
        <div className="checkbox" style={{ color: d.color }} key={d.name}>
          {(
            <input
              type="checkbox"
              value={d.name}
              checked={selectedItems === undefined ? false : selectedItems.includes(d.name)}
              onChange={() => onChange(d.name)}
            />
          )}
          <label>{d.name}</label>
        </div>
      ))}
    </div>
  )
}




export default Legend;
