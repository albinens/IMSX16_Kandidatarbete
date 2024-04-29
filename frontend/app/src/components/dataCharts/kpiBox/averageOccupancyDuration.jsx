import React, { useEffect, useState } from "react";
import KpiBox from "./kpiBox";


//props : data: []
const AverageOccupancyDuration = (props) => {

  const [data, setData] = useState(props.data)
  const [value, setValue] = useState(0)

  const treatData = () => {
    let tempData = []
    let occupancySpans = []

    props.data.map((obj) => {
      tempData.push({
        name: obj.roomName,
        data: obj.data
      })
    })
    console.log('Temp data', tempData)
    tempData.map((room) => {
      let current = 0
      room.data.map((row) => {
        if(row.occupancy > 0){
          current += 5
        } else if (current > 0){
          occupancySpans.push(current)
        }
      })
    })
    setValue((occupancySpans.reduce((a, b) => a + b, 0) / occupancySpans.length).toFixed(0))
    console.log('Occupancy spans', occupancySpans)
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
      unit="minutes"
    />
  )
}

export default AverageOccupancyDuration;