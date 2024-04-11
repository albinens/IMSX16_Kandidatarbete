import React from "react";
import './weekChart.css'

const Legend = ({ data, selectedItems, onChange }) => {

const checked = (name) => {
  if(selectedItems === undefined || selectedItems.length === 0) return false;

  return selectedItems.includes(name);
}

return(
    <div className="legendContainer-weekplot">
      {data.map((d) => (
        <div className="checkbox" style={{ color: d.color }} key={d.name}>
          {(
            <input
              type="checkbox"
              value={d.name}
              checked={checked(d.name)}
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
