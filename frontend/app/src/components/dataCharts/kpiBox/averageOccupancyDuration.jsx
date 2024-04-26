import React, { useEffect, useState } from "react";
import KpiBox from "./kpiBox";


//props : data: []
const AverageOccupancyDuration = (props) => {

  const [data, setData] = useState(props.data)
  const [value, setValue] = useState(0)

  const treatData = () => {

  }


  useEffect(() => {
    setData(props.data)
    console.log('Data loaded AverageOccupancyDuration.jsx', data)
    treatData()
  }, [props.data])

  return(
    <KpiBox 
      title="Average Occupancy Duration"
      value={value} 
      unit="hh:mm"
    />
  )
}

export default AverageOccupancyDuration;