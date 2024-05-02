import React, { useEffect, useState } from "react";
import KpiBox from "./kpiBox";


//props : data: []
const AverageOccupancyRate = (props) => {

  const [data, setData] = useState(props.data)
  const [value, setValue] = useState(0)

  const treatData = () => {
    let tempData = []
    let totalOccupancy = 0
    let totalRecordings = 0

    props.data.map((obj) => {
      tempData.push({
        name: obj.roomName,
        data: obj.data
      })
    })
    console.log('Temp data', tempData)
    tempData.map((room) => {
      room.data.map((row) => {
        totalOccupancy += row.occupancy
        if(row.occupancy !== 0){
          totalRecordings++
        }
      })
    })

    setValue((totalOccupancy / totalRecordings).toFixed(2))
  }


  useEffect(() => {
    setData(props.data)
    console.log('Data loaded AverageOccupancyRate.jsx', data)
    if(data.length === 0){
      return
    }
    treatData()
  }, [props.data])

  return(
    <KpiBox 
      title="Average Room Occupancy Rate"
      value={value} 
      unit="visitors"
    />
  )
}

export default AverageOccupancyRate;